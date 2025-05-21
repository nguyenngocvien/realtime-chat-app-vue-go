<template>
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-orange-200 to-gray-500">
        <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-sm">
            <h1 class="text-3xl font-bold text-center text-orange-400 mb-6">Login</h1>
            <form @submit.prevent="handleLogin">
                <div class="mb-4">
                    <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
                    <input v-model="username" type="text" id="username"
                        class="text-sm mt-1 w-full p-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-yellow-400"
                        placeholder="Enter your username" required />
                </div>
                <div class="mb-6">
                    <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                    <input v-model="password" type="password" id="password"
                        class="text-sm mt-1 w-full p-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-yellow-400"
                        placeholder="Enter your password" required />
                </div>
                <div v-if="error" class="mb-4 text-red-500 text-sm text-center">{{ error }}</div>
                <button type="submit"
                    class="w-full bg-orange-400 text-white p-1 rounded-md hover:bg-orange-500 transition duration-300"
                    :disabled="loading">
                    {{ loading ? 'Logging in...' : 'Login' }}
                </button>
            </form>
            <p class="mt-4 text-center text-sm text-gray-600">
                Don't have an account? <router-link to="/signup" class="text-orange-600 hover:underline">Sign up</router-link>
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../api/auth'

const username = ref<string>('')
const password = ref<string>('')
const error = ref<string>('')
const loading = ref<boolean>(false)

const router = useRouter()

const handleLogin = async () => {
    try {
        await login(username.value, password.value)
        router.push('/') // Navigate to the chat page
    } catch (e: any) {
        error.value = e.message || 'Login failed';
    } finally {
        loading.value = false
    }
}
</script>