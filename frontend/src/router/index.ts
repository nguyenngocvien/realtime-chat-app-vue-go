import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import ChatView from '../views/ChatView.vue';
import Login from '@/components/Login.vue';
import Signup from '@/components/Signup.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup,
  },
  {
    path: '/',
    name: 'Chats',
    component: ChatView,
    meta: { requiresAuth: true },
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  if (to.meta.requiresAuth && !token) {
    next('/login');
  } else {
    next();
  }
});

export default router;