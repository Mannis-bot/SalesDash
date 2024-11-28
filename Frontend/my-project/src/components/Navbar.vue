<template>
  <div>
    <section class="flex justify-between pt-[13px] pl-[20px] mr-[30px] border-b pb-[10px]">
    
      <form action="">
        <input
          type="search"
          placeholder="Search..."
          class="border rounded-md px-[10px] py-[2px] outline-none focus:border-blue-500"
        />
      </form>

      
      <div v-if="!isAdminLoggedIn" class="relative">
        <VueFeather
          type="shopping-cart"
          size="20"
          @click="toggleCartDropdown"
        />
        <span
          v-if="cartItemCount > 0"
          class="absolute top-[-5px] right-[-5px] bg-red-500 text-white text-xs rounded-full w-[20px] h-[20px] flex items-center justify-center"
        >
          {{ cartItemCount }}
        </span>

       
        <div
          v-if="cartDropdownVisible"
          class="absolute top-[40px] right-0 mt-2 w-[300px] bg-white border border-blue-500 rounded-lg shadow-lg z-10 max-h-[400px] overflow-y-auto"
        >
          <div class="p-4 text-sm">
            

            <h5 class="font-semibold text-gray-700">Cart Items</h5>
            <ul>
              <li
                v-for="item in cartItems"
                :key="item.name"
                class="flex justify-between py-2"
              >
                <span>{{ item.quantity }} x {{ item.name }}</span>
                <span class="font-semibold">Ksh {{ item.totalPrice.toFixed(2) }}</span>
              </li>
            </ul>
            <div class="border-t pt-2">
              <p class="font-semibold text-lg">Total: Ksh {{ totalPrice.toFixed(2) }}</p>
            </div>

            
            <div class="mt-4">
              <button
                @click="goToCheckout"
                class="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600"
              >
                Go to Checkout
              </button>
            </div>
          </div>
        </div>
      </div>

     
      <div class="profile flex items-center gap-[10px] relative">
        <img
          src="../assets/profile.jpg"
          alt="ImageHere"
          width="30px"
          class="rounded-full"
        />
        <div class="flex flex-col text-sm">
          <h4 class="font-semibold">{{ userName }}</h4>
          <span>{{ isAdminLoggedIn ? 'Admin' : 'User' }}</span>
        </div>
        <VueFeather type="chevron-down" size="16" @click="toggleProfileDropdown" />

       
        <div
          v-if="profileDropdownVisible"
          class="absolute top-[40px] right-0 mt-2 w-[200px] bg-white border border-gray-300 rounded-lg shadow-lg z-10 p-4"
        >
          <p class="text-sm font-semibold">{{ userName }}</p>
          <p class="text-sm text-gray-600">{{ isAdminLoggedIn ? 'Admin' : 'User' }}</p>
        </div>
      </div>
    </section>
  </div>
</template>


<script>
import { computed, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import VueFeather from 'vue-feather';

export default {
  setup() {
    const router = useRouter();

    const cartItems = ref([]);
    const cartDropdownVisible = ref(false);
    const profileDropdownVisible = ref(false);

    const isAdminLoggedIn = computed(() => !!localStorage.getItem('adminToken'));
    const userName = computed(() => localStorage.getItem('userName') || 'Guest');

    const toggleCartDropdown = () => {
      cartDropdownVisible.value = !cartDropdownVisible.value;
      profileDropdownVisible.value = false; 
    };

    const toggleProfileDropdown = () => {
      profileDropdownVisible.value = !profileDropdownVisible.value;
      cartDropdownVisible.value = false; 
    };

    const updateCartItems = () => {
      const storedCart = JSON.parse(localStorage.getItem('cart')) || [];
      const items = [];
      storedCart.forEach((item) => {
        const existingItem = items.find((i) => i.name === item.name);
        if (existingItem) {
          existingItem.quantity += 1;
          existingItem.totalPrice += parseFloat(item.price);
        } else {
          items.push({
            name: item.name,
            quantity: 1,
            totalPrice: parseFloat(item.price),
          });
        }
      });
      cartItems.value = items;
    };

    const cartItemCount = computed(() =>
      cartItems.value.reduce((sum, item) => sum + item.quantity, 0)
    );

    const totalPrice = computed(() =>
      cartItems.value.reduce((sum, item) => sum + item.totalPrice, 0)
    );

    const goToCheckout = () => {
      router.push('/checkout');
    };

    const logout = () => {
      localStorage.clear();
      router.push('/login');
    };

    onMounted(() => {
      updateCartItems();
      setInterval(() => {
        updateCartItems();
      }, 2000);
    });

    return {
      isAdminLoggedIn,
      userName,
      cartDropdownVisible,
      profileDropdownVisible,
      toggleCartDropdown,
      toggleProfileDropdown,
      cartItems,
      cartItemCount,
      totalPrice,
      goToCheckout,
      logout,
    };
  },
};
</script>

<style scoped>

.profile .dropdown {
  position: relative;
  z-index: 999;
}


.profile .dropdown button {
  position: absolute;
  top: 10px;
  right: 10px;
  background: none;
  border: none;
  cursor: pointer;
  color: #aaa;
  transition: color 0.2s ease;
}

.profile .dropdown button:hover {
  color: #333;
}


.profile .dropdown {
  background-color: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.profile .dropdown .text-sm {
  padding: 12px;
}

.profile .dropdown button {
  padding: 8px;
}

.profile .dropdown .border-t {
  border-top: 1px solid #ddd;
}

.profile .dropdown .w-full {
  width: 100%;
}
</style>
