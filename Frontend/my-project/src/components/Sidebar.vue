<template>
  <div>
    <section class="w-[200px] min-h-screen border border-r pt-[10px] pl-[20px]">
      <h1 class="text-green-600 font-bold text-[25px] mb-[40px]">Grocery</h1>
      <div class="sideNav flex flex-col h-[400px]">
        <div class="upper flex flex-1 flex-col gap-[25px] w-[150px]">
         
          <router-link
            v-if="isUserLoggedIn"
            to="/home"
            active-class="active-link"
            class=" py-[6px] px-[8px] rounded-md  flex items-center gap-[10px]"
          >
            <VueFeather type="home" size="16" /> Home
          </router-link>

          

         
          <router-link
            v-if="isAdminLoggedIn"
            to="/admin/dashboard"
            active-class="active-link"
            class=" py-[6px] px-[8px] rounded-md  flex items-center gap-[10px]"
          >
            <VueFeather type="grid" size="16" /> Dashboard
          </router-link>

          

          <router-link
            v-if="isAdminLoggedIn"
            to="/admin/orders"
            active-class="active-link"
            class=" py-[6px] px-[8px] rounded-md  flex items-center gap-[10px]"
          >
            <VueFeather type="shopping-cart" size="16" /> Orders
          </router-link>
        </div>

        <div class="lower mt-auto">
          
          <button 
            v-if="isUserLoggedIn || isAdminLoggedIn"
            type="button" 
            @click="logout"
            class="bg-green-400 hover:bg-green-800 py-[6px] px-[16px] rounded-md text-white flex items-center gap-[10px]"
          >
            <VueFeather type="log-out" size="16" /> Logout
          </button>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
import { computed, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import VueFeather from 'vue-feather';

export default {
  components: {
    VueFeather,
  },
  setup() {
    const router = useRouter();

    const isUserLoggedIn = computed(() => {
      const loggedIn = localStorage.getItem('userLoggedIn');
      console.log('User Logged In:', loggedIn); 
      return loggedIn === 'true';  
    });

    const isAdminLoggedIn = computed(() => {
      const loggedIn = localStorage.getItem('adminLoggedIn');
      console.log('Admin Logged In:', loggedIn); 
      return loggedIn === 'true';  
    });


    watch([isUserLoggedIn, isAdminLoggedIn], () => {
      console.log('User logged in:', isUserLoggedIn.value);
      console.log('Admin logged in:', isAdminLoggedIn.value);
    });

   
    const logout = () => {
      if (isAdminLoggedIn.value) {
        localStorage.removeItem('adminLoggedIn');
        router.push('/admin/login'); 
      } else if (isUserLoggedIn.value) {
        localStorage.removeItem('userLoggedIn');
        router.push('/login'); 
      }
    };

    
    onMounted(() => {
      console.log('Sidebar mounted');
      console.log('isUserLoggedIn at mount:', isUserLoggedIn.value);
      console.log('isAdminLoggedIn at mount:', isAdminLoggedIn.value);
    });

    return {
      isUserLoggedIn,
      isAdminLoggedIn,
      logout,
    };
  },
};
</script>






<style  scoped>
.active-link {
  background-color: #ff5c33;
  color: white;
}

</style>