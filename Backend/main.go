package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"database/sql"
	"math/rand"
	"time"
	"context"
	"strings"
	"strconv"


	"github.com/rs/cors"
	"github.com/gorilla/websocket"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mailgun/mailgun-go/v4"
	
)


var db *sql.DB


type Order struct {
	UserName   string        `json:"userName"`
	Products   []OrderProduct `json:"products"`
	Total      float64       `json:"total"`
	Comment    string        `json:"comment"`
	OrderCode  string        `json:"orderCode"`
}
type OrderProduct struct {
	Name     string  `json:"name"`
	UnitPrice float64 `json:"unitPrice"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}



type Product struct{
	ID       int    `json:"id"`
    Category string `json:"category"`
    Name     string `json:"name"`
    Price    string `json:"price"`
    Image    string `json:"image"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func init(){
	var err error
	db, err := sql.Open("mysql", "root:Marykiki@tcp(127.0.0.1:3306)/grocery_store")
	if err != nil{
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("Error pinging the databse", err)
	}
	fmt.Println("Successfully connected to the database!")
}

func initializeDB() {
    var err error
    db, err = sql.Open("mysql", "root:Marykiki@tcp(localhost:3306)/grocery_store")
    if err != nil {
        log.Fatal("Error opening database: ", err)
    }
    err = db.Ping() // Ping the database to check if the connection is alive
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    fmt.Println("Successfully connected to the database!")
}


func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Received POST request for admin login")

	// Check if database is initialized
	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	// Parse form data
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Declare variables to hold email and password from the database
	var dbEmail, dbPassword string

	// Query the database for the admin user by email
	err := db.QueryRow("SELECT email, password FROM users WHERE email = ?", email).Scan(&dbEmail, &dbPassword)
	if err != nil {
		// If no rows are found, return an error
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
		// Log any other errors
		log.Fatal(err)
	}

	// Debugging: print the values being compared
	fmt.Println("Email from Postman:", email)
	fmt.Println("Password from Postman:", password)
	fmt.Println("Email from DB:", dbEmail)
	fmt.Println("Password from DB:", dbPassword)

	// Check if the provided email and password match
	if email == dbEmail && password == dbPassword {
		fmt.Fprintf(w, "Welcome Admin!")
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}
}


func fetchProducts(w http.ResponseWriter, r *http.Request) {
    
    rows, err := db.Query("SELECT id, category, name, price, image FROM products")
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []Product

    
    for rows.Next() {
        var product Product
        err := rows.Scan(&product.ID, &product.Category, &product.Name, &product.Price, &product.Image)
        if err != nil {
            http.Error(w, "Error scanning product", http.StatusInternalServerError)
            return
        }
        products = append(products, product)
    }

    
    if err := rows.Err(); err != nil {
        http.Error(w, "Error during row iteration", http.StatusInternalServerError)
        return
    }

    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(products); err != nil {
        http.Error(w, "Error encoding products", http.StatusInternalServerError)
        return
    }
}



func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Send the verification email via Mailgun
func sendVerificationEmail(email, verificationCode string) error {
	// Mailgun setup
	domain := "sandboxba5d00dfe36c458890d2daad59e20d6d.mailgun.org" // Your Mailgun domain
	apiKey := "0dab8a7942a0645ac642a5ab8c4e2c47-c02fd0ba-8a911cd8"               // Your Mailgun API key
	mg := mailgun.NewMailgun(domain, apiKey)

	// Create the email message
	message := mg.NewMessage(
		"kinyanjui816@gmail.com", // Replace with your verified email
		"Email Verification Code",
		fmt.Sprintf("Your verification code is: %s", verificationCode),
		email,
	)

	ctx := context.Background()

// Send the email
_, _, err := mg.Send(ctx, message)
if err != nil {
    log.Println("Error sending email:", err)
    return err
}
	log.Println("Verification email sent successfully to", email)
	return nil
}

// Signup API - User registration
func signup(w http.ResponseWriter, r *http.Request) {
    var user struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Parse the incoming JSON request body
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Check if the email already exists in the database
    var existingUser struct {
        ID int
    }
    err = db.QueryRow("SELECT id FROM regular_users WHERE email = ?", user.Email).Scan(&existingUser.ID)
    if err == nil {
        http.Error(w, "Email already registered", http.StatusBadRequest)
        return
    }

    // Generate a verification code
    verificationCode := generateVerificationCode()

    // Insert the user into the database
    _, err = db.Exec("INSERT INTO regular_users (name, email, password, verification_code, is_verified) VALUES (?, ?, ?, ?, ?)",
        user.Name, user.Email, user.Password, verificationCode, false)
    if err != nil {
        http.Error(w, "Error registering user", http.StatusInternalServerError)
        return
    }

    // Send the verification email
    if err := sendVerificationEmail(user.Email, verificationCode); err != nil {
        http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
        return
    }

    // Respond to the client with the verification code
    response := map[string]interface{}{
        "message": "Signup successful! Check your email for verification code.",
        "verificationCode": verificationCode, // Return the verification code here
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}


// Verify API - Verify the user's email using the verification code
func verifyEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email            string `json:"email"`
		VerificationCode string `json:"verification_code"`
	}
	// Parse the incoming JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find the user by email
	var user struct {
		ID                int
		VerificationCode  string
		IsVerified        bool
	}
	err = db.QueryRow("SELECT id, verification_code, is_verified FROM regular_users WHERE email = ?", req.Email).Scan(&user.ID, &user.VerificationCode, &user.IsVerified)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the verification code matches
	if user.VerificationCode != req.VerificationCode {
		http.Error(w, "Invalid verification code", http.StatusBadRequest)
		return
	}

	// Mark the user as verified
	_, err = db.Exec("UPDATE regular_users SET is_verified = ? WHERE email = ?", true, req.Email)
	if err != nil {
		http.Error(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Email verified successfully!"}`))
}

// Login API - User login
func login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the incoming request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	var user struct {
		ID        int
		Name      string
		Password  string
		IsVerified bool
	}

	// Fetch the user details including the name
	err = db.QueryRow("SELECT id, name, password, is_verified FROM regular_users WHERE email = ?", req.Email).Scan(&user.ID, &user.Name, &user.Password, &user.IsVerified)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	// Check if the password matches
	if user.Password != req.Password {
		http.Error(w, `{"error": "Invalid password"}`, http.StatusUnauthorized)
		return
	}

	// Check if the email is verified
	if !user.IsVerified {
		http.Error(w, `{"error": "Email not verified"}`, http.StatusForbidden)
		return
	}

	// Send the response with the user's name, login status, and other details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "Login successful",
		"isLoggedIn": true,
		"name": "` + user.Name + `",
		"email": "` + req.Email + `"
	}`))
}



// Save order details into the database
func saveOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Received order:", order) 
	// Prepare product details as comma-separated strings
	var productNames []string
	var quantities []string
	var unitPrices []string
	var total float64

	// Collect product details
	for _, product := range order.Products {
		productNames = append(productNames, product.Name)
		quantities = append(quantities, fmt.Sprintf("%d", product.Quantity))
		unitPrices = append(unitPrices, fmt.Sprintf("%.2f", product.UnitPrice))
		total += product.Total
	}

	// Convert product details to comma-separated strings
	productNamesStr := strings.Join(productNames, ",")
	quantitiesStr := strings.Join(quantities, ",")
	unitPricesStr := strings.Join(unitPrices, ",")

	// Insert the order into the database
	query := "INSERT INTO orders (order_code, user_name, product_names, quantities, unit_prices, total, comment, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, order.OrderCode, order.UserName, productNamesStr, quantitiesStr, unitPricesStr, total, order.Comment, time.Now())
	if err != nil {
		http.Error(w, "Error inserting order into database", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Order placed successfully!")
}


// Get order details by order code
func getOrder(w http.ResponseWriter, r *http.Request) {
	// Get the order code from the query parameters
	orderCode := r.URL.Query().Get("orderCode")
	if orderCode == "" {
		http.Error(w, "Order code is required", http.StatusBadRequest)
		return
	}

	// Query the database to get the order details
	query := "SELECT user_name, product_names, quantities, unit_prices, total, comment FROM orders WHERE order_code = ?"
	var order Order
	var productNamesStr, quantitiesStr, unitPricesStr string
	err := db.QueryRow(query, orderCode).Scan(&order.UserName, &productNamesStr, &quantitiesStr, &unitPricesStr, &order.Total, &order.Comment)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching order details", http.StatusInternalServerError)
		}
		return
	}

	// Convert product details from comma-separated strings to slices
	productNames := strings.Split(productNamesStr, ",")
	quantities := strings.Split(quantitiesStr, ",")
	unitPrices := strings.Split(unitPricesStr, ",")

	// Prepare the products slice
	var products []OrderProduct
	for i := range productNames {
		quantity, _ := strconv.Atoi(quantities[i]) // Use strconv.Atoi
		unitPrice, _ := strconv.ParseFloat(unitPrices[i], 64) // Use strconv.ParseFloat
		total := unitPrice * float64(quantity)
	
		products = append(products, OrderProduct{
			Name:      productNames[i],
			UnitPrice: unitPrice,
			Quantity:  quantity,
			Total:     total,
		})
	}

	// Set the products in the order struct
	order.Products = products
	order.OrderCode = orderCode

	// Respond with the order details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}



func fetchOrders(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT id, order_code, user_name, product_names, quantities, unit_prices, total, comment, created_at FROM orders ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []map[string]interface{}
	for rows.Next() {
		var order map[string]interface{}
		var id int
		var order_code, user_name, product_names, quantities, unit_prices, comment string
		var total float64
		var created_at string

		err := rows.Scan(&id, &order_code, &user_name, &product_names, &quantities, &unit_prices, &total, &comment, &created_at)
		if err != nil {
			return nil, err
		}

		order = map[string]interface{}{
			"id":           id,
			"order_code":   order_code,
			"user_name":    user_name,
			"product_names": product_names,
			"quantities":   quantities,
			"unit_prices":  unit_prices,
			"total":        total,
			"comment":      comment,
			"created_at":   created_at,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("New client connected")

	// Handle incoming messages (if any)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			log.Println("Client disconnected")
			break
		}
	}
}

// Broadcast new orders to all connected WebSocket clients
func broadcastOrders(orderData string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(orderData))
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}

// Fetch orders from the database and send them to WebSocket clients
func fetchAndBroadcastOrders() {
    for {
        // Fetch orders
        orders, err := fetchOrders(db)
        if err != nil {
            log.Println("Error fetching orders:", err)
            time.Sleep(5 * time.Second) // Retry after some time
            continue
        }

        // Send each order as a JSON message
        for _, order := range orders {
            orderJSON, err := json.Marshal(order)
            if err != nil {
                log.Println("Error marshalling order:", err)
                continue
            }
            broadcastOrders(string(orderJSON))
        }

        // Wait before checking for new orders again
        time.Sleep(10 * time.Second)
    }
}



func main() {
	// Initialize the database
	initializeDB()

	http.HandleFunc("/ws", handleConnections)

	go fetchAndBroadcastOrders()

	// Define routes using http.HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Grocery Store Backend!")
	})

	http.HandleFunc("/admin/login", adminLogin)
	http.HandleFunc("/products", fetchProducts)
	http.HandleFunc("/signup", signup)          // POST for user signup
	http.HandleFunc("/verify", verifyEmail)     // POST for email verification
	http.HandleFunc("/login", login)     
	http.HandleFunc("/place-order", saveOrder)  
	http.HandleFunc("/get-order", getOrder)
     // POST for user login

	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5174"}, // Allow frontend on port 5174
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allow OPTIONS preflight
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Apply CORS middleware to all handlers
	handler := c.Handler(http.DefaultServeMux)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}