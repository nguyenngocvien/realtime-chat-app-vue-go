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
                <div class="flex items-center gap-4">
                    <input v-model="searchQuery" type="text" placeholder="Search..."
                        class="w-full p-2 rounded-lg bg-gray-400 text-gray-600 placeholder-gray-500 focus:outline-none focus:shadow-md focus:bg-gray-100"
                        @input="handleSearch" @focus="showDropdown = true" />
                </div>

                <!-- Dropdown -->
                <div v-if="showDropdown && (users.length > 0 || messages.length > 0)"
                    class="absolute top-16 left-4 right-4 bg-white border rounded-md shadow-lg z-10 max-h-96 overflow-y-auto">
                    <div v-if="users.length > 0" class="p-2">
                        <h4 class="text-sm font-semibold text-gray-700 mb-2">Users</h4>
                        <div v-for="user in users" :key="user.id" class="p-2 hover:bg-gray-100 cursor-pointer"
                            @click="selectUser(user)">
                            {{ user.username }}
                        </div>
                    </div>
                    <div v-if="messages.length > 0" class="p-2">
                        <h4 class="text-sm font-semibold text-gray-700 mb-2">Messages</h4>
                        <div v-for="message in messages" :key="message.id" class="p-2 hover:bg-gray-100 cursor-pointer"
                            @click="selectMessage(message)">
                            <p class="text-sm font-medium">{{ message.chatName }}</p>
                            <p class="text-sm text-gray-600 truncate">{{ message.content }}</p>
                        </div>
                    </div>
                    <button class="p-2 rounded bg-blue-500 text-white hover:bg-blue-600 transition duration-200"
                        @click="clearSearch">
                        Close
                    </button>
                </div>
            </div>

            <!-- User Avatar -->
            <div class="flex items-center">
                <img src="@/assets/images/avatar.png" alt="User Avatar" class="w-10 h-10 rounded-full" />
            </div>
        </div>
    </header>
</template>

<script lang="ts" setup>
import { searchMessages, searchUsers } from '@/api/chat';
import { SearchMessage, SearchUser } from '@/types';
import { watch } from 'fs';
import { defineComponent, onBeforeUnmount, onMounted, ref } from 'vue'

const searchQuery = ref('');
const users = ref<SearchUser[]>([]);
const messages = ref<SearchMessage[]>([]);
const showDropdown = ref(false);

const emit = defineEmits<{
    (e: 'toggle-sidebar'): void;
    (e: 'select-user', user: SearchUser): void;
    (e: 'select-message', message: SearchMessage): void;
}>();

function handleToggleSidebar() {
  emit('toggle-sidebar');
}

async function handleSearch() {
  if (searchQuery.value.trim() === '') {
    users.value = [];
    messages.value = [];
    return;
  }

  try {
    const [userResults, messageResults] = await Promise.all([
      searchUsers(searchQuery.value),
      searchMessages(searchQuery.value),
    ]);
    users.value = userResults;
    messages.value = messageResults;
  } catch (error) {
    console.error('Search failed:', error);
  }
}

function selectUser(user: SearchUser) {
  emit('select-user', user);
  clearSearch();
}

function selectMessage(message: SearchMessage) {
  emit('select-message', message);
  clearSearch();
}

function clearSearch() {
  searchQuery.value = '';
  users.value = [];
  messages.value = [];
  showDropdown.value = false;
}

// Close dropdown on outside click
function handleClickOutside(event: MouseEvent) {
  const target = event.target as HTMLElement;
  if (!target.closest('.bg-white.p-4.border-b.border-gray-200')) {
    showDropdown.value = false;
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside);
});

</script>

<style scoped>
/* Tailwind CSS handles styling */
</style>