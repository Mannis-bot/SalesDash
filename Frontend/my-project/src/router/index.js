import { createRouter, createWebHistory } from 'vue-router';

// User views
import Login from '../views/User/Login.vue';
import Signup from '../views/User/Signup.vue';
import Home from '../views/User/Home.vue';
import Cart from '../views/User/Cart.vue';
import Checkout from '../views/User/Checkout.vue';
import Orders from '../views/User/Orders.vue';
import Receipt from '../views/User/Receipt.vue'

// Admin views
import AdminLogin from '../views/Admin/AdminLogin.vue';
import AdminDashboard from '../views/Admin/AdminDashboard.vue';
import AdminCustomers from '../views/Admin/AdminCustomers.vue';
import AdminOrders from '../views/Admin/AdminOrders.vue';

const routes = [
  { path: '/login', name: 'Login', component: Login, meta: { requiresLayout: false } },
  { path: '/signup', name: 'Signup', component: Signup, meta: { requiresLayout: false } },
  { path: '/home', name: 'Home', component: Home, meta: { requiresLayout: true } },
  { path: '/cart', name: 'Cart', component: Cart, meta: { requiresLayout: true } },
  { path: '/checkout', name: 'Checkout', component: Checkout, meta: { requiresLayout: true } },
  { path: '/orders', name: 'Orders', component: Orders, meta: { requiresLayout: true } },
  { path: '/receipt', name: 'Receipt', component: Receipt, meta: { requiresLayout: true } },

  { path: '/admin/login', name: 'AdminLogin', component: AdminLogin, meta: { requiresLayout: false } },
  { 
    path: '/admin/dashboard', 
    name: 'AdminDashboard', 
    component: AdminDashboard, 
    meta: { requiresAuth: true, requiresLayout: true }
  },
  { 
    path: '/admin/customers', 
    name: 'AdminCustomers', 
    component: AdminCustomers, 
    meta: { requiresAuth: true, requiresLayout: true }
  },
  { 
    path: '/admin/orders', 
    name: 'AdminOrders', 
    component: AdminOrders, 
    meta: { requiresAuth: true, requiresLayout: true }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});


router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    const isAuthenticated = localStorage.getItem('adminToken'); 
    if (isAuthenticated) {
      next(); 
    } else {
      next('/admin/login');
    }
  } else {
    next(); 
  }
});

export default router;
