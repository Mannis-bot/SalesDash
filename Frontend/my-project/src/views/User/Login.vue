<template>
  <div class="login-page flex justify-center items-center min-h-screen bg-gray-50">
    <div class="login-container bg-white p-[40px] rounded-lg shadow-xl w-[400px]">
      <h1 class="text-3xl font-bold text-center text-green-600 mb-[20px]">Login</h1>
  
      <form @submit.prevent="handleLogin">
        <div class="mb-[20px]">
          <label for="email" class="block text-lg text-gray-700">Email</label>
          <input
            type="email"
            id="email"
            v-model="email"
            placeholder="Enter your email"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>
  
        <div class="mb-[20px]">
          <label for="password" class="block text-lg text-gray-700">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Enter your password"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>
  
        <button
          type="submit"
          class="w-full py-[12px] bg-green-600 text-white text-lg font-semibold rounded-lg hover:bg-green-700 transition duration-200"
        >
          Log In
        </button>
      </form>
  
      <div v-if="errorMessage" class="mt-4 text-center text-red-600">
        <p>{{ errorMessage }}</p>
      </div>

      <div class="mt-[15px] text-center text-gray-600">
        <p>Don't have an account? 
          <router-link to="/signup" class="text-green-600 font-semibold hover:underline">Sign Up</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: '',
      errorMessage: '',
    };
  },
  methods: {
    async handleLogin() {
      if (!this.email || !this.password) {
        this.errorMessage = 'Please enter both email and password.';
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
          }),
        });

        const data = await response.json();

        if (response.ok) {
          if (data.isLoggedIn) {
            
            console.log('Login successful:', data); 

            
            localStorage.setItem('userLoggedIn', 'true');
            localStorage.setItem('userName', data.name || 'Guest'); 

            alert(`Logged in as: ${localStorage.getItem('userName')}`); 
            
            console.log('Stored name in localStorage:', localStorage.getItem('userName'));

            
            this.$router.push('/home');
          } else {
            this.errorMessage = 'Login failed: User not authorized or verified.';
          }
        } else {
          this.errorMessage = `Login failed: ${data.message}`;
        }
      } catch (error) {
        this.errorMessage = 'Login failed. Server error or no response.';
      }
    },
  },
};
</script>

  <style scoped>

  .login-page {
    background-color: #f9fafb;
  }
  
 
  .login-container {
    width: 100%;
    max-width: 400px;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  }
  
  
  label {
    font-size: 16px;
    font-weight: 600;
    color: #4a5568; 
  }
  
 
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
    border-color: #38a169; 
    box-shadow: 0 0 0 2px rgba(56, 161, 105, 0.5); 
  }

  button {
    width: 100%;
    background-color: #38a169; 
    color: white;
    padding: 12px;
    font-size: 16px;
    font-weight: 600;
    border-radius: 8px;
    transition: background-color 0.2s;
  }
  
  button:hover {
    background-color: #2f7a4b; 
  }
  
 
  router-link {
    color: #38a169;
    text-decoration: none;
    font-weight: 600;
  }
  
  router-link:hover {
    text-decoration: underline;
  }
  </style>
  