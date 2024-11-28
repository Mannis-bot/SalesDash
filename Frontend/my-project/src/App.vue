<template>
  <div id="app" class="min-h-screen flex">
    <Sidebar v-if="showLayout" />
    
    <div class="flex-1 flex flex-col">
      <Navbar v-if="showLayout" />

      <main class="flex-1 p-4 bg-[#f5f6f7]">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue';
import Sidebar from './components/Sidebar.vue';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

export default {
  components: {
    Navbar,
    Sidebar,
  },
  setup() {
    const route = useRoute();
    
    // Determine if layout components should be shown based on meta.requiresLayout
    const showLayout = computed(() => route.meta.requiresLayout !== false);

    return {
      showLayout,
    };
  },
};
</script>

<style>
html,
body,
#app {
  @apply min-h-screen;
}
</style>


