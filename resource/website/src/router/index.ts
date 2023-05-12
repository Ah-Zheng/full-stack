import { createRouter, createWebHashHistory } from 'vue-router';
import Layout from '@/views/layout/Layout.vue';

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: '',
                component: () => import('@/views/nav/HomeView.vue')
            }
        ]
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

export default router;