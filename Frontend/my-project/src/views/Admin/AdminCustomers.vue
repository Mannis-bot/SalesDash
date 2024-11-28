<template>
    <div class="customers-page p-[20px]">
      <h1 class="text-3xl font-bold text-green-600 mb-[20px]">Customers</h1>
  
      <!-- Display a message when no customers are available -->
      <div v-if="customers.length === 0" class="text-center text-gray-500">
        <p>No customers available.</p>
      </div>
  
      <!-- Display customers in a grid layout -->
      <div v-if="customers.length > 0" class="customers-grid grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-[20px]">
        <div v-for="customer in customers" :key="customer.id" class="customer-card p-[15px] bg-white shadow-lg rounded-lg border-l-[5px] border-l-green-500">
          <div class="text-center mb-[10px]">
            <h4 class="text-xl font-bold text-green-600">{{ customer.name }}</h4>
            <p class="text-sm text-gray-600">{{ customer.email }}</p>
          </div>
          <div class="text-center text-gray-700">
            <p class="text-sm">Joined: {{ customer.joinDate }}</p>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        customers: [
           {
            id: 1,
            name: "John Doe",
            email: "johndoe@example.com",
            joinDate: "2024-01-15",
           },
           {
             id: 2,
             name: "Jane Smith",
             email: "janesmith@example.com",
             joinDate: "2024-02-10",
           },
          // Add more customers here or fetch from an API
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  .customers-page {
    padding: 20px;
  }
  
  .customers-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
  }
  
  .customer-card {
    border: 1px solid #e0e0e0;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
  }
  
  .customer-card:hover {
    transform: translateY(-10px);
    box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.1);
  }
  
  .customer-card h4 {
    color: #38a169; /* Green for the name */
  }
  
  .customer-card p {
    color: #4a5568; /* Dark gray for other text */
  }
  </style>
  
  