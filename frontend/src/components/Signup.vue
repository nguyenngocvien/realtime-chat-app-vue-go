<template>
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-blue-200 to-gray-500">
        <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-sm">
            <h1 class="text-3xl font-bold text-center text-blue-600 mb-6">Sign Up</h1>
            <form @submit.prevent="handleSignup">
                <div class="mb-4">
                    <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
                    <input v-model="username" type="text" id="username"
                        class="text-sm mt-1 w-full p-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        placeholder="Enter your username" required />
                </div>
                <div class="mb-6">
                    <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                    <input v-model="password" type="password" id="password"
                        class="text-sm mt-1 w-full p-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        placeholder="Enter your password" required />
                </div>
                <div v-if="error" class="mb-4 text-red-500 text-sm text-center">{{ error }}</div>
                <button type="submit"
                    class="w-full bg-blue-600 text-white p-1 rounded-md hover:bg-blue-700 transition duration-300"
                    :disabled="loading">
                    {{ loading ? 'Signing up...' : 'Sign Up' }}
                </button>
            </form>
            <p class="mt-4 text-center text-sm text-gray-600">
                Already have an account? <router-link to="/login"
                    class="text-blue-600 hover:underline">Login</router-link>
            </p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { signup } from '../api/auth';

const router = useRouter();
const username = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

async function handleSignup() {
    error.value = '';
    loading.value = true;
    try {
        await signup(username.value, password.value);
        router.push('/');
    } catch (err: any) {
        error.value = err.message;
    } finally {
        loading.value = false;
    }
}
</script>

<style scoped>
</style>