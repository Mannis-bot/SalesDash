<template>
    <div class="bg-green-50 min-h-screen p-6">
      <h1 class="text-2xl font-bold text-green-700 mb-6">Checkout</h1>
  
      
      <div class="bg-white rounded-lg shadow-lg p-6">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="border-b">
              <th class="py-3 px-4">Product</th>
              <th class="py-3 px-4">Unit Price (Ksh)</th>
              <th class="py-3 px-4">Quantity</th>
              <th class="py-3 px-4">Total Price (Ksh)</th>
              <th class="py-3 px-4">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="item in cartItems"
              :key="item.name"
              class="border-b hover:bg-green-100"
            >
              <td class="py-3 px-4">{{ item.name }}</td>
              <td class="py-3 px-4">{{ item.unitPrice.toFixed(2) }}</td>
              <td class="py-3 px-4">
                {{ item.quantity }}
              </td>
              <td class="py-3 px-4">{{ (item.unitPrice * item.quantity).toFixed(2) }}</td>
              <td class="py-3 px-4">
                <button
                  @click="incrementItem(item)"
                  class="bg-green-500 text-white px-2 py-1 rounded hover:bg-green-600"
                >
                  +
                </button>
                <button
                  @click="decrementItem(item)"
                  class="bg-red-500 text-white px-2 py-1 rounded ml-2 hover:bg-red-600"
                >
                  -
                </button>
                <button
                  @click="removeItem(item)"
                  class="bg-gray-500 text-white px-2 py-1 rounded ml-2 hover:bg-gray-600"
                >
                  Remove
                </button>
              </td>
            </tr>
          </tbody>
        </table>
  
        <div class="mt-6 text-right">
          <h3 class="text-xl font-semibold text-green-700">
            Total: Ksh {{ totalCartPrice.toFixed(2) }}
          </h3>
        </div>
      </div>
  
      
      <div class="bg-white rounded-lg shadow-lg p-6 mt-6">
        <label for="comment" class="block text-lg font-semibold text-green-700 mb-2">
          Add a Comment/Message:
        </label>
        <textarea
          id="comment"
          v-model="comment"
          placeholder="e.g., Deliver by 9 PM"
          class="w-full p-3 border rounded-lg focus:outline-none focus:ring focus:ring-green-300"
          rows="4"
        ></textarea>
      </div>
  
   
      <div class="mt-6 text-right">
        <button
          @click="makeOrder"
          class="bg-green-500 text-white font-bold py-3 px-6 rounded-lg hover:bg-green-600"
        >
          Make Order
        </button>
      </div>
    </div>
  </template>
  
  <script>
  import { ref, computed, onMounted } from "vue";
  import { useRouter } from "vue-router";
  
  export default {
    setup() {
      const cartItems = ref([]);
      const comment = ref("");
      const router = useRouter();
  
      const updateCartItems = () => {
        const storedCart = JSON.parse(localStorage.getItem("cart")) || [];
        const items = [];
        storedCart.forEach((item) => {
          const existingItem = items.find((i) => i.name === item.name);
          if (existingItem) {
            existingItem.quantity += 1;
          } else {
            items.push({
              name: item.name,
              unitPrice: parseFloat(item.price),
              quantity: 1,
            });
          }
        });
        cartItems.value = items;
      };
  
      const incrementItem = (item) => {
        item.quantity++;
        updateLocalStorage();
      };
  
      const decrementItem = (item) => {
        if (item.quantity > 1) {
          item.quantity--;
        } else {
          removeItem(item);
        }
        updateLocalStorage();
      };
  
      const removeItem = (item) => {
        cartItems.value = cartItems.value.filter((i) => i.name !== item.name);
        updateLocalStorage();
      };
  
      const updateLocalStorage = () => {
        const updatedCart = [];
        cartItems.value.forEach((item) => {
          for (let i = 0; i < item.quantity; i++) {
            updatedCart.push({ name: item.name, price: item.unitPrice });
          }
        });
        localStorage.setItem("cart", JSON.stringify(updatedCart));
      };
  
      const totalCartPrice = computed(() => {
        return cartItems.value.reduce(
          (sum, item) => sum + item.unitPrice * item.quantity,
          0
        );
      });
  
      const makeOrder = () => {
  
  const orderCode = `ORD-${Date.now().toString(36).toUpperCase()}`;

  
  const orderDetails = {
    orderNumber: orderCode,
    products: cartItems.value.map((item) => ({
      name: item.name,
      unitPrice: item.unitPrice.toFixed(2),
      quantity: item.quantity,
      total: (item.unitPrice * item.quantity).toFixed(2),
    })),
    total: totalCartPrice.value.toFixed(2),
    comment: comment.value || "No comment",
  };

  
  localStorage.setItem("orderDetails", JSON.stringify(orderDetails));

  
  localStorage.removeItem("cart");

 
  cartItems.value = [];
  comment.value = "";

  
  alert("Thank you for your order!");

  
  router.push("/receipt");
};

  
      onMounted(() => {
        updateCartItems();
      });
  
      return {
        cartItems,
        incrementItem,
        decrementItem,
        removeItem,
        totalCartPrice,
        comment,
        makeOrder,
      };
    },
  };
</script>

  <style scoped>
  
  body {
    background-color: #e8f5e9;
  }
  
  table {
    width: 100%;
  }
  
  th,
  td {
    padding: 8px 12px;
  }
  
  .bg-green-50 {
    background-color: #f0fdf4;
  }
  </style>
  
  