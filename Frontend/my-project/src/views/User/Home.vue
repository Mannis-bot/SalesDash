<template>
    <div>
      <section class="mr-[30px]">
        <h1 class="text-center text-[25px] font-bold mb-[40px]">Featured <span class="text-[#1A8156]">Products</span></h1>
        <div>
    <ProductCard />
    
  </div>
      </section>
    </div>
</template>

<script>
import ProductCard from '../../components/ProductCard.vue'
    export default {
        components:{
            ProductCard,
        }
        
    }
</script>

<style lang="scss" scoped>

</style>