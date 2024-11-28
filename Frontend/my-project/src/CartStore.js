
import { reactive } from 'vue';

export const cartStore = reactive({
  items: [], // This will hold all the products in the cart
  
  addToCart(product) {
    this.items.push(product);
  },

  getCartCount() {
    return this.items.length;
  },

  clearCart() {
    this.items = [];
  }
});
