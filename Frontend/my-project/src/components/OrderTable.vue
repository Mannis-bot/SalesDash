<template>
  <div class="order-container mb-[20px] bg-white shadow-lg rounded-lg border-t-[5px] border-t-[#1A8156] w-[300px] mx-auto">
    <!-- Order Header (Receipt Style) -->
    <div class="text-center p-[15px] border-b-[2px] border-gray-200">
      <h3 class="text-xl font-bold">Order #{{ order.code }}</h3>
      <p class="text-sm text-gray-600">Thank you for shopping with us!</p>
    </div>

    <!-- Order Details -->
    <div class="px-[15px] py-[10px]">
      <p class="text-sm"><strong>Customer Name:</strong> {{ order.customerName }}</p>
      <p class="text-sm"><strong>Total Price:</strong> ${{ order.totalPrice.toFixed(2) }}</p>
    </div>

    <!-- Products List -->
    <table class="w-full border-collapse mb-[15px]">
      <thead>
        <tr>
          <th class="text-left py-[5px] text-sm">Product</th>
          <th class="text-right py-[5px] text-sm">Price</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in order.products" :key="product.name" class="border-b">
          <td class="py-[5px] text-sm">{{ product.name }}</td>
          <td class="py-[5px] text-right text-sm">${{ product.price.toFixed(2) }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Order Comment (if available) -->
    <div v-if="order.comment" class="px-[15px]">
      <p class="text-sm"><strong>Comment:</strong> {{ order.comment }}</p>
    </div>

    <!-- Footer with padding and nice rounding for a receipt-like finish -->
    <div class="footer p-[10px] bg-gray-100 text-center rounded-b-lg">
      <p class="text-sm text-gray-500">Thank you for your order!</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ["order"],
};
</script>

<style scoped>
.order-container {
  /* Ensure padding and spacing make it look like a receipt */
  padding: 10px;
  width: 300px;
  font-family: 'Courier New', Courier, monospace; /* Receipt-like font */
}

.table th, .table td {
  padding: 5px;
  font-size: 14px;
}

.footer {
  background-color: #f9f9f9;
  border-top: 2px solid #1A8156;
  font-size: 12px;
  text-align: center;
}
</style>

  