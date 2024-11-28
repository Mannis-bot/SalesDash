<template>
    <div class="bg-gray-100 min-h-screen p-6">
      <div v-if="order" class="bg-white rounded-lg shadow-lg p-6 max-w-xl mx-auto">
        <h1 class="text-2xl font-bold text-center mb-6">Receipt</h1>
  
        <div>
          <p class="text-sm font-medium">Order Number: <span class="font-bold">{{ order.orderNumber }}</span></p>
          <p class="text-sm font-medium">Comment: <span class="font-bold">{{ order.comment }}</span></p>
        </div>
  
        <table class="w-full mt-6 border-collapse">
          <thead>
            <tr>
              <th class="border-b py-2">Product</th>
              <th class="border-b py-2">Qty</th>
              <th class="border-b py-2">Unit Price (Ksh)</th>
              <th class="border-b py-2">Total (Ksh)</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in order.products" :key="item.name">
              <td class="py-2">{{ item.name }}</td>
              <td class="py-2 text-center">{{ item.quantity }}</td>
              <td class="py-2 text-right">{{ item.unitPrice }}</td>
              <td class="py-2 text-right">{{ item.total }}</td>
            </tr>
          </tbody>
        </table>
  
        <div class="text-right mt-4">
          <p class="font-semibold text-lg">Total: Ksh {{ order.total }}</p>
        </div>
      </div>
  
      <div v-if="showMessage" class="fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-blue-500 text-white p-4 rounded-lg shadow-lg">
        <p class="text-center">Redirecting to Home in <span>{{ countdown }}</span> seconds...</p>
      </div>
  
      <div v-else class="text-center text-gray-600">
        <p>No order data available.</p>
      </div>
    </div>
  </template>
  
  <script>
  import { ref, onMounted, computed } from "vue";
  import axios from "axios";
  import { useRouter } from "vue-router";  
  
  export default {
    setup() {
      const order = ref(null);
      const router = useRouter(); 
      const showMessage = ref(true);  
      const countdown = ref(60);  
  
      
      const userName = computed(() => localStorage.getItem("userName") || "Unknown");
  
      onMounted(async () => {
        const storedOrder = JSON.parse(localStorage.getItem("orderDetails"));
        if (storedOrder) {
          order.value = storedOrder;
  
         
          console.log("Order to be sent:", storedOrder);
  
       
          const orderToSend = {
            orderCode: storedOrder.orderNumber,
            userName: userName.value,  
            products: storedOrder.products.map(product => ({
              name: product.name,
              quantity: product.quantity,
              unitPrice: parseFloat(product.unitPrice),
              total: parseFloat(product.total) 
            })),
            total: parseFloat(storedOrder.total), 
            comment: storedOrder.comment || "No comment"
          };
  
          
          try {
            const response = await axios.post("http://localhost:8080/place-order", orderToSend, {
              headers: {
                "Content-Type": "application/json"
              }
            });
            console.log("Order saved successfully:", response.data);
          } catch (error) {
            console.error("Error saving order:", error.response ? error.response.data : error.message);
          }
        } else {
          
          order.value = {
            orderNumber: "N/A",
            products: [],
            total: 0,
            comment: "No comment",
          };
        }
  
      
        const interval = setInterval(() => {
          countdown.value--;
          if (countdown.value <= 0) {
            clearInterval(interval); 
            showMessage.value = false;  
            router.push("/home");  
          }
        }, 1000); 
      });
  
      return {
        order,
        showMessage,
        countdown,
        userName,
      };
    },
  };
  </script>
  
  
  
  
  

  
  <style scoped>
  /* Supermarket-style theme */
  body {
    background-color: #f4f4f4;
  }
  
  table {
    width: 100%;
  }
  
  th,
  td {
    padding: 8px;
    text-align: left;
  }
  
  th {
    background-color: #e5e7eb;
  }
  
  .shadow-lg {
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  </style>
  