<template>
  <div class="login-page flex justify-center items-center min-h-screen bg-gray-50">
    <div class="login-container bg-white p-[40px] rounded-lg shadow-xl w-[400px]">
      <h1 class="text-3xl font-bold text-center text-green-600 mb-[20px]">Admin Login</h1>

      <form @submit.prevent="submitLogin">
        <!-- Email Field -->
        <div class="mb-[20px]">
          <label for="email" class="block text-lg text-gray-700">Email</label>
          <input
            type="email"
            id="email"
            v-model="email"
            required
            placeholder="Enter your email"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>

        <!-- Password Field -->
        <div class="mb-[20px]">
          <label for="password" class="block text-lg text-gray-700">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            required
            placeholder="Enter your password"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>

        <!-- Login Button -->
        <button
          type="submit"
          class="w-full py-[12px] bg-green-600 text-white text-lg font-semibold rounded-lg hover:bg-green-700 transition duration-200"
        >
          Login
        </button>
      </form>
    </div>
  </div>
</template>
<script>
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      email: '',
      password: '',
    };
  },
  setup() {
    const router = useRouter();
    return { router };
  },
  methods: {
    async submitLogin() {
      try {
        const response = await fetch('http://localhost:8080/admin/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
          body: new URLSearchParams({
            email: this.email,
            password: this.password,
          }),
        });

        if (response.ok) {
          const data = await response.text();
          alert(data); // "Welcome Admin!"

          // Save an authentication token (simulated here)
          localStorage.setItem('adminLoggedIn', 'true'); // Simplified; replace with actual token.
          console.log('Admin token saved:', localStorage.getItem('adminLoggedIn'));

          // Redirect to admin dashboard
          this.router.push('/admin/dashboard');
        } else {
          const errorMessage = await response.text();
          alert(errorMessage); // Invalid credentials
        }
      } catch (error) {
        console.error('Error during login:', error);
        alert('Something went wrong. Please try again.');
      }
    },
  },
};
</script>

  
  <style scoped>
  /* Full page styling for center alignment */
  .login-page {
    background-color: #f9fafb;
  }
  
  /* Container for the login form */
  .login-container {
    width: 100%;
    max-width: 400px;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  }
  
  /* Form field labels */
  label {
    font-size: 16px;
    font-weight: 600;
    color: #4a5568; /* Dark gray color */
  }
  
  /* Form fields */
  input {
    width: 100%;
    padding: 10px 15px;
    border-radius: 8px;
    border: 1px solid #ccc;
    outline: none;
    font-size: 14px;
    transition: border-color 0.3s;
  }
  
  input:focus {
    border-color: #38a169; /* Green for focus state */
    box-shadow: 0 0 0 2px rgba(56, 161, 105, 0.5); /* Green shadow */
  }
  
  /* Button styling */
  button {
    width: 100%;
    background-color: #38a169; /* Green */
    color: white;
    padding: 12px;
    font-size: 16px;
    font-weight: 600;
    border-radius: 8px;
    transition: background-color 0.2s;
  }
  
  button:hover {
    background-color: #2f7a4b; /* Darker green */
  }
  </style>
  