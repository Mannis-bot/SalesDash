<template>
    <div class="dashboard p-[20px] bg-gray-100 min-h-screen">
      <!-- Metrics Section -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-[20px] mb-[30px]">
        <div class="bg-white shadow-md p-[15px] rounded-lg">
          <h3 class="text-gray-500">Total Sales</h3>
          <p class="text-2xl font-bold text-green-600">$12,345</p>
        </div>
        <div class="bg-white shadow-md p-[15px] rounded-lg">
          <h3 class="text-gray-500">Total Orders</h3>
          <p class="text-2xl font-bold text-blue-600">320</p>
        </div>
        <div class="bg-white shadow-md p-[15px] rounded-lg">
          <h3 class="text-gray-500">Total Customers</h3>
          <p class="text-2xl font-bold text-yellow-600">150</p>
        </div>
        <div class="bg-white shadow-md p-[15px] rounded-lg">
          <h3 class="text-gray-500">Revenue</h3>
          <p class="text-2xl font-bold text-red-600">$8,500</p>
        </div>
      </div>
  
      <!-- Charts Section -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-[20px]">
        <div class="bg-white shadow-md p-[20px] rounded-lg">
          <h3 class="text-lg font-bold mb-[15px]">Sales Trend</h3>
          <apexchart type="bar" :options="barOptions" :series="barSeries" />
        </div>
        <div class="bg-white shadow-md p-[20px] rounded-lg">
          <h3 class="text-lg font-bold mb-[15px]">Category Distribution</h3>
          <apexchart type="pie" :options="pieOptions" :series="pieSeries" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import VueApexCharts from "vue3-apexcharts";
  
  export default {
    components: {
      apexchart: VueApexCharts,
    },
    data() {
      return {
        // Bar Chart
        barOptions: {
          chart: {
            id: "sales-bar",
          },
          xaxis: {
            categories: ["January", "February", "March", "April"],
          },
        },
        barSeries: [
          {
            name: "Sales",
            data: [1200, 2400, 3000, 1800],
          },
        ],
  
        // Pie Chart for Groceries
        pieOptions: {
          labels: ["Fruits", "Vegetables", "Dairy", "Beverages"],
        },
        pieSeries: [40, 30, 20, 10], // Placeholder data
      };
    },
  };
  </script>
  
  <style scoped>
  .dashboard {
    background-color: #f8f9fa;
  }
  </style>
  