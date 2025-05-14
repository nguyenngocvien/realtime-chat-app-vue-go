<template>
    <header class="text-white p-4 w-full top-0 z-10 max-w-6xl">
        <div class="mx-auto flex items-center justify-between">
            <!-- Hamburger Button -->
            <button @click="handleToggleSidebar" class="focus:outline-none rounded-full p-2 bg-gray-400">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                </svg>
            </button>

            <!-- Search Bar -->
            <div class="flex-1 mx-4 max-w-md">
                <input type="text" placeholder="Search..."
                    class="w-full p-2 rounded-lg bg-gray-400 text-gray-600 placeholder-gray-500 focus:outline-none focus:shadow-md focus:bg-gray-100" />
            </div>

            <!-- User Avatar -->
            <div class="flex items-center">
                <img src="@/assets/images/avatar.png" alt="User Avatar" class="w-10 h-10 rounded-full" />
            </div>
        </div>
    </header>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'TopBar',
    emits: ['toggle-sidebar'],
    methods: {
    handleToggleSidebar() {
      console.log('Hamburger clicked, emitting toggle-sidebar');
      this.$emit('toggle-sidebar');
    }
  }
})
</script>

<style scoped>
/* Tailwind CSS handles styling */
</style>