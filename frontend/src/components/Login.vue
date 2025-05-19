<template>
    <div>
        <h2>Login</h2>
        <form @submit.prevent="handleLogin">
            <input v-model="username" placeholder="Username" required />
            <input v-model="password" type="password" placeholder="Password" required />
            <button type="submit">Login</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../api/auth'

const username = ref<string>('')
const password = ref<string>('')

const router = useRouter()

const handleLogin = async () => {
    try {
        await login(username.value, password.value)
        router.push('/chat') // Navigate to the chat page
    } catch (error) {
        alert('Login failed')
    }
}
</script>