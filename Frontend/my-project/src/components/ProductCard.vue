<template>
  <div class="product-grid-container grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 p-4">
    <div v-for="product in products" :key="product.id" class="w-full">
      <section class="bg-white w-[230px] h-auto pb-[10px] rounded-md flex flex-col items-center pt-[40px]">
        <img :src="product.imagePath" alt="Product Image" width="150px" class="">
        <div class="productInfo w-[100%] mt-[30px] flex flex-col gap-[12px] px-[20px]">
          <div class="flex justify-between text-sm">
            <p class="text-[13px] text-green-400">{{ product.category }}</p>
            <p class="text-[13px] flex items-center gap-[5px]">
              <VueFeather type="star" size="12" class="text-yellow-500" /> 4.8
            </p>
          </div>
          
          <h4 class="font-semibold ">{{ product.name }}</h4>
          <p class="text-xs text-gray-500">500g</p>
          
          <div class="price flex justify-between items-center">
            <p class="text-[15px] font-semibold">{{ product.price }}</p>
            <button 
              @click="addToCart(product)" 
              class="flex items-center gap-[5px] bg-[#DDECE7] rounded-md py-[4px] px-[12px] text-sm text-[#1A8156] font-semibold">
              <VueFeather type="shopping-bag" size="18" /> Add
            </button>
          </div>
        </div>
      </section>
    </div>

    
    <div v-if="message" class="absolute top-[50px] left-1/2 transform -translate-x-1/2 bg-green-500 text-white py-2 px-4 rounded-md">
      {{ message }} added to Cart!
    </div>

    
    
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      products: [],
      message: "",  
      cart: [],    
    };
  },
  created() {
    this.fetchProducts();
    this.loadCartFromLocalStorage(); 
  },
  methods: {
    async fetchProducts() {
      try {
        const response = await axios.get("http://localhost:8080/products");
        this.products = response.data.map(product => ({
          ...product,
          imagePath: this.getImagePath(product.image),
        }));
      } catch (error) {
        console.error("There was an error fetching products:", error);
      }
    },
    getImagePath(imageName) {
      return new URL(`../assets/${imageName}`, import.meta.url).href;
    },
    loadCartFromLocalStorage() {
     
      const storedCart = JSON.parse(localStorage.getItem("cart"));
      if (storedCart) {
        this.cart = storedCart;
      }
    },
    addToCart(product) {
      
      this.cart.push(product);

     
      localStorage.setItem("cart", JSON.stringify(this.cart));

     
      console.log("Cart:", this.cart);

      
      this.message = product.name;

      
      setTimeout(() => {
        this.message = "";
      }, 3000);
    },
  },
};
</script>


<style lang="scss" scoped>

.cart-preview {
  position: fixed;
  bottom: 10px;
  right: 10px;
  background-color: white;
  padding: 10px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.cart-preview ul {
  list-style-type: none;
  padding: 0;
}

.cart-preview li {
  padding: 5px 0;
}
</style>
