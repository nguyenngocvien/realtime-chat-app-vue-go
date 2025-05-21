<template>
    <div class="chat-list h-full p-4">
        <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-800">Chats</h3>
            <Icon @click="showModal = true" name="plus" customClass="w-4 h-4 text-gray-500 hover:text-gray-700" />
        </div>
        <ul class="space-y-2">
            <li v-for="(chat, index) in chats" :key="index" @click="selectChat(chat)"
                :ref="el => chatRefs[chat.id] = el" :class="[
                    'p-3 rounded-lg hover:bg-gray-300 cursor-pointer transition-all duration-200',
                    selectedChatId === chat.id ? 'bg-gray-300 font-semibold' : ''
                ]">
                <div class="flex items-center gap-3">
                    <div
                        class="relative w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center overflow-hidden">
                        <template v-if="chat.type === 'group'">
                            <!-- Group Icon SVG -->
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-600" fill="none"
                                viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M17 20h5v-2a4 4 0 00-3-3.87M9 20H4v-2a4 4 0 013-3.87m9-3.13a4 4 0 100-8 4 4 0 000 8zm-9 0a4 4 0 100-8 4 4 0 000 8z" />
                            </svg>
                        </template>
                        <template v-else>
                            <div>
                                <img v-if="chat.avatar && !imageError" :src="chat.avatar" alt="Avatar"
                                    class="w-10 h-10 rounded-full" @error="imageError = true" />
                                <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                    stroke-width="1.5" stroke="currentColor" class="w-10 h-10 text-gray-400">
                                    <path stroke-linecap="round" stroke-linejoin="round"
                                        d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
                                </svg>
                            </div>

                        </template>
                    </div>
                    <div>
                        <p class="font-medium text-gray-800">{{ chat.name }}</p>
                        <p class="text-sm truncate"
                            :class="{ 'text-gray-800 font-semibold': chat.unreadCount > 0, 'text-gray-600': chat.unreadCount === 0 }">
                            {{
                                chat.lastMessage?.content }}</p>
                    </div>
                </div>
            </li>
        </ul>

        <!-- Modal create -->
        <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div class="bg-white rounded-lg shadow-lg p-6 w-80">
                <input v-model="newGroupName" type="text" placeholder="Name"
                    class="w-full p-2 border border-gray-300 rounded mb-8 focus:outline-none" />
                <div class="flex justify-end gap-2">
                    <button @click="showModal = false" class="p-2 rounded bg-gray-300 hover:bg-gray-400">Cancel</button>
                    <button @click="createGroup"
                        class="p-2 rounded bg-blue-500 text-white hover:bg-blue-600">Create</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, nextTick, ComponentPublicInstance } from 'vue'
import { createGroupChat, fetchChatsByUserId } from '@/api/chat'
import type { Chat } from '@/types'
import Icon from './base/Icon.vue'
import { useRouter } from 'vue-router';
import { parseJWT } from '@/utils';

const router = useRouter();

const props = defineProps<{
    selectedChat: Chat | null
}>()

const emit = defineEmits<{
    (e: 'select-chat', chat: Chat): void
}>()

const chats = ref<Chat[]>([])
const showModal = ref(false)
const newGroupName = ref('')
const selectedChatId = ref('')
const imageError = ref(false)

// chatRefs: map lưu ref phần tử chat item, dùng cho scrollIntoView
const chatRefs = reactive<Record<string, Element | ComponentPublicInstance | null>>({})

async function loadChats() {
    const token = localStorage.getItem('token');
    if (!token) {
        router.push('/login');
        return;
    }
    
    const userId = parseJWT(token).userId;
    try {
        chats.value = await fetchChatsByUserId(userId);
    } catch (error) {
        console.error('Failed to load chats:', error);
    }
}

function selectChat(chat: Chat) {
    selectedChatId.value = chat.id
    emit('select-chat', chat)
    router.push(`/chats/${chat.id}`);

    nextTick(() => {
        const el = chatRefs[chat.id]
        if (el instanceof HTMLElement) {
            el.scrollIntoView({ behavior: 'smooth', block: 'nearest' })
        }
    })

    if (chat.unreadCount > 0) {
        chat.unreadCount = 0
    }
}

async function createGroup() {
    if (newGroupName.value.trim() === '') {
        alert('Please enter group name')
        return
    }

    const token = localStorage.getItem('token');
    if (!token) {
        router.push('/login');
        return;
    }
    const userId = parseJWT(token).userId;

    try {
        const newChat = await createGroupChat(userId, newGroupName.value);
        chats.value.push(newChat);
        newGroupName.value = '';
        showModal.value = false;
    } catch (error) {
        console.error('Failed to create group:', error);
        alert('Failed to create group');
    }
}

onMounted(() => {
    loadChats()
})

</script>

<style scoped>
.chat-list {
    width: 150px;
}

@media (min-width: 768px) {
    .chat-list {
        width: 280px;
        /* Desktop */
    }
}
</style>