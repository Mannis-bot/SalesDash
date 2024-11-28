<template>
  <div class="orders-admin p-[20px] bg-gray-100 min-h-screen">
    <h2 class="text-2xl font-bold mb-[20px]">Orders</h2>

    <!-- Orders Table -->
    <div class="overflow-x-auto bg-white shadow-md rounded-lg p-[20px]">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-200 text-gray-700">
            <th class="py-[10px] px-[15px]">Code</th>
            <th class="py-[10px] px-[15px]">Customer Name</th>
            <th class="py-[10px] px-[15px]">Total Price</th>
            <th class="py-[10px] px-[15px]">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="order in orders" :key="order.code" class="border-b">
            <td class="py-[10px] px-[15px]">{{ order.code }}</td>
            <td class="py-[10px] px-[15px]">{{ order.customerName }}</td>
            <td class="py-[10px] px-[15px]">Ksh {{ order.totalPrice }}</td>
            <td class="py-[10px] px-[15px]">
              <button
                @click="viewOrder(order)"
                class="bg-blue-500 text-white px-[10px] py-[5px] rounded hover:bg-blue-600"
              >
                View More
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal for Viewing Order Details -->
    <div v-if="selectedOrder" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
      <div class="bg-white p-[20px] rounded-lg shadow-lg w-[350px] font-mono">
        <div class="text-center border-b pb-[10px]">
          <h3 class="text-2xl font-bold">Grocery Receipt</h3>
          <p class="text-sm text-gray-600">Thank you for shopping with us!</p>
        </div>

        <div class="mt-[10px]">
          <p><strong>Order Code:</strong> {{ selectedOrder.code }}</p>
          <p><strong>Customer Name:</strong> {{ selectedOrder.customerName }}</p>
        </div>

        <table class="w-full mt-[10px] border-collapse">
          <thead>
            <tr class="border-b">
              <th class="text-left py-[5px]">Product</th>
              <th class="text-right py-[5px]">Price</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in selectedOrder.products" :key="product.name" class="border-b">
              <td class="py-[5px]">{{ product.name }}</td>
              <td class="py-[5px] text-right">Ksh {{ product.price.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>

        <div class="mt-[10px] border-t pt-[10px]">
          <p class="flex justify-between font-bold">
            <span>Total:</span>
            <span>Ksh {{ selectedOrder.totalPrice.toFixed(2) }}</span>
          </p>
          <p v-if="selectedOrder.comment" class="text-sm text-gray-600 mt-[5px]">
            <strong>Comment:</strong> {{ selectedOrder.comment }}
          </p>
        </div>

        <div class="text-center mt-[15px] border-t pt-[10px]">
          <p class="text-xs text-gray-500">Visit again soon!</p>
        </div>

        <button
          @click="selectedOrder = null"
          class="mt-[10px] bg-blue-500 text-white px-[10px] py-[5px] w-full rounded hover:bg-blue-600"
        >
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      orders: [],
      processedOrderCodes: [],  // Track processed orders by their order_code
      selectedOrder: null,      // Define selectedOrder here to avoid the Vue warning
    };
  },

  methods: {
    // Method to set the selected order to view details
    viewOrder(order) {
      console.log(order);  // Log the selected order for debugging
      this.selectedOrder = order;
    },

    // Method to add new order to the orders list
    updateOrders(newOrder) {
      // Check if the order has already been processed using its order_code
      if (!this.processedOrderCodes.includes(newOrder.order_code)) {
        // Format the new order
        const formattedOrder = {
          code: newOrder.order_code,    // Map order_code to code
          customerName: newOrder.user_name || 'Unknown',  // Default to 'Unknown' if customerName is missing
          totalPrice: this.calculateTotalPrice(newOrder.unit_prices, newOrder.quantities), // Calculate total price dynamically
          products: this.formatProducts(newOrder.product_names, newOrder.unit_prices, newOrder.quantities),  // Format product details
          comment: newOrder.comment || '',  // Default to empty string if comment is missing
          created_at: newOrder.created_at,  // Include created_at
        };

        // Add the new order to the orders list
        this.orders.push(formattedOrder);

        // Mark the order as processed by adding its order_code to the processed list
        this.processedOrderCodes.push(newOrder.order_code);
      } else {
        console.log("Order already processed:", newOrder.order_code);  // Log if the order was already processed
      }
    },

    // Method to format product names into an array of product objects with prices and quantities
    formatProducts(productNames, unitPrices, quantities) {
      const productNamesArray = productNames.split(',');
      const unitPricesArray = unitPrices.split(',');
      const quantitiesArray = quantities.split(',');

      return productNamesArray.map((name, index) => ({
        name: name.trim(),
        price: this.formatPrice(parseFloat(unitPricesArray[index].trim())), // Format the unit price to Ksh
        quantity: parseInt(quantitiesArray[index].trim()), // Parse the quantity for each product
      }));
    },

    // Method to calculate the total price for an order based on unit prices and quantities
    calculateTotalPrice(unitPrices, quantities) {
      const unitPricesArray = unitPrices.split(',').map(price => parseFloat(price.trim()));
      const quantitiesArray = quantities.split(',').map(quantity => parseInt(quantity.trim()));

      let totalPrice = 0;

      // Calculate total price by multiplying unit prices with quantities
      for (let i = 0; i < unitPricesArray.length; i++) {
        totalPrice += unitPricesArray[i] * quantitiesArray[i];
      }

      return this.formatPrice(totalPrice); // Format the total price to Ksh
    },

    // Method to format prices into Ksh
    formatPrice(price) {
      return price.toFixed(2); // Keep price format consistent with 2 decimal places
    },
  },

  created() {
    // Create a WebSocket connection to the backend
    this.socket = new WebSocket("ws://localhost:8080/ws");  // Ensure the backend is running on this port
    
    // Listen for messages from the WebSocket server
    this.socket.onmessage = (event) => {
      const newOrder = JSON.parse(event.data);
      console.log("Received new order:", newOrder); // Log the new order
      this.updateOrders(newOrder);  // Add the new order to the orders list
    };

    // Handle WebSocket errors
    this.socket.onerror = (error) => {
      console.error("WebSocket Error: ", error);
    };
    
    // Handle WebSocket closure
    this.socket.onclose = () => {
      console.log("WebSocket connection closed");
    };
  },

  beforeDestroy() {
    // Close the WebSocket connection when the component is destroyed
    if (this.socket) {
      this.socket.close();
    }
  },
};
</script>

<style scoped>
/* Add your custom styles here */
</style>

  
  <style scoped>
  .orders-admin {
    background-color: #f8f9fa;
  }
  
  table {
    width: 100%;
    border-spacing: 0;
  }
  
  th,
  td {
    text-align: left;
  }
  </style>
  