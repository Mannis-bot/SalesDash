<template>
    <div class="orders-page bg-gray-100 p-[20px]">
      <h1 class="text-3xl font-bold mb-[20px]">Your Previous Orders</h1>
  
      <div v-if="orders.length === 0" class="text-center text-gray-500">
        <p>No orders placed yet.</p>
      </div>
  
      <div class="orders-grid grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-[20px]">
        <div v-for="order in orders" :key="order.code">
          <OrderTable :order="order" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import OrderTable from "../../components/OrderTable.vue"; 
  export default {
    components: {
      OrderTable,
    },
    data() {
      return {
        orders: [
          
          {
            code: "ORD123",
            customerName: "John Doe",
            totalPrice: 120.5,
            products: [
              { name: "Apple", price: 2.5 },
              { name: "Banana", price: 1.5 },
              { name: "Orange", price: 3.0 },
            ],
          },
          {
            code: "ORD124",
            customerName: "John Doe",
            totalPrice: 45.0,
            products: [
              { name: "Milk", price: 1.5 },
              { name: "Eggs", price: 2.5 },
              { name: "Bread", price: 3.0 },
            ],
          },
          
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  .orders-page {
    padding: 20px;
    background-color: #f9fafb;
  }
  
  .orders-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 20px;
  }
  </style>
  