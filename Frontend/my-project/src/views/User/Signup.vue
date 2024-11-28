<template>
  <div class="signup-page flex justify-center items-center min-h-screen bg-gray-50">
    <div class="signup-container bg-white p-[40px] rounded-lg shadow-xl w-[400px]">
      <h1 class="text-3xl font-bold text-center text-green-600 mb-[20px]">User Signup</h1>

      <form @submit.prevent="handleSignup">
        <div class="mb-[20px]">
          <label for="name" class="block text-lg text-gray-700">Full Name</label>
          <input
            type="text"
            id="name"
            v-model="name"
            placeholder="Enter your full name"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>

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

        <div class="mb-[20px]">
          <label for="confirmPassword" class="block text-lg text-gray-700">Confirm Password</label>
          <input
            type="password"
            id="confirmPassword"
            v-model="confirmPassword"
            placeholder="Confirm your password"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>

        <div v-if="isVerificationSent" class="mb-[20px]">
          <label for="verificationCode" class="block text-lg text-gray-700">Verification Code</label>
          <input
            type="text"
            id="verificationCode"
            v-model="verificationCode"
            placeholder="Enter verification code"
            class="w-full px-[15px] py-[10px] border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>

        <button
          v-if="!isVerificationSent"
          type="submit"
          class="w-full py-[12px] bg-green-600 text-white text-lg font-semibold rounded-lg hover:bg-green-700 transition duration-200"
        >
          Sign Up
        </button>

        <button
          v-if="isVerificationSent"
          @click="verifyCode"
          type="button"
          class="w-full py-[12px] bg-green-600 text-white text-lg font-semibold rounded-lg hover:bg-green-700 transition duration-200"
        >
          Verify Code
        </button>
      </form>

      <div class="mt-[15px] text-center text-gray-600">
        <p>Already have an account? 
          <router-link to="/login" class="text-green-600 font-semibold hover:underline">Log In</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      name: '',
      email: '',
      password: '',
      confirmPassword: '',
      verificationCode: '',
      isVerificationSent: false,
      sentVerificationCode: '',
    };
  },
  methods: {
    async handleSignup() {
      if (this.password !== this.confirmPassword) {
        alert('Passwords do not match!');
        return;
      }

      try {
        const response = await axios.post('http://localhost:8080/signup', {
          name: this.name,
          email: this.email,
          password: this.password,
        });

        console.log('Signup successful:', response);
        this.isVerificationSent = true;
        this.sentVerificationCode = response.data.verificationCode; 
        console.log('Verification code sent:', this.sentVerificationCode);
      } catch (error) {
        console.error('Signup error:', error);
      }
    },

    async verifyCode() {
      if (this.verificationCode !== this.sentVerificationCode) {
        alert('Invalid verification code!');
        return;
      }

      try {
        
        const response = await axios.post('http://localhost:8080/verify', {
          email: this.email,
          verification_code: this.verificationCode,
        });

        if (response.status === 200) {
          this.$router.push('/login');
        }
      } catch (error) {
        alert('Verification failed!');
        console.error('Verification error:', error);
      }
    },
  },
};
</script>





  
  <style scoped>
  /* Full page styling for center alignment */
  .signup-page {
    background-color: #f9fafb;
  }
  
  /* Container for the signup form */
  .signup-container {
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
  
  /* Login Link Styling */
  router-link {
    color: #38a169;
    text-decoration: none;
    font-weight: 600;
  }
  
  router-link:hover {
    text-decoration: underline;
  }
  </style>
  
  