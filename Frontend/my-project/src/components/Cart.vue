<template>
    <div class="cart-container">
      <h2>Your Cart</h2>
      <ul>
        <li v-for="product in cartProducts" :key="product.id">
          <p>{{ product.name }}</p>
          <p>{{ product.price }}</p>
        </li>
      </ul>
      <div v-if="cartProducts.length === 0">Your cart is empty.</div>
    </div>
  </template>
  
  <script>
  import { mapState } from "vuex";
  
  export default {
    computed: {
      ...mapState({
        cartProducts: (state) => state.cartProducts,
      }),
    },
  };
  </script>
  
  <style scoped>
  .cart-container {
    padding: 20px;
  }
  </style>
  