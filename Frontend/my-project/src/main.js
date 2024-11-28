import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import VueFeather from 'vue-feather';
import './style.css'

const app = createApp(App);


app.use(router).component('VueFeather', VueFeather);

app.mount('#app');



