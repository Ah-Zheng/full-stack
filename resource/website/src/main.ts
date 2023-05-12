import { createApp } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import router from './router'
import './style.css'
import App from './App.vue'

import 'normalize.css';
import '@/icons';

createApp(App)
    .component('fa-icon', FontAwesomeIcon)
    .use(router)
    .mount('#app')
