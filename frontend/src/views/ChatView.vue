<template>
    <div class="h-screen flex flex-col bg-gray-500">
        <!-- Top Bar -->
        <div class="flex justify-center">
            <TopBar @select-user="handleSelectUser" @select-message="handleSelectMessage" />
        </div>

        <!-- Main Layout -->
        <div class="flex-1 justify-center">
            <!-- Main Content -->
            <main class="flex px-4 w-full justify-center">
                <div class="flex max-w-6xl rounded-xl bg-gray-400 w-full overflow-hidden">
                    <div>
                        <ChatList :selected-chat="selectedChat" @select-chat="onSelectChat" />
                    </div>
                    <div class="flex-1">
                        <ChatBox :selected-chat="selectedChat" />
                    </div>
                </div>
            </main>
        </div>
    </div>

</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import ChatList from '../components/ChatList.vue';
import type { Chat, SearchMessage, SearchUser } from '@/types';
import TopBar from '@/components/TopBar.vue';
import ChatBox from '@/components/ChatBox.vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const selectedChat = ref<Chat | null>(null);
const selectedMessageId = ref<string>('');

function onSelectChat(chat: Chat) {
    selectedChat.value = chat;
}

function handleSelectUser(user: SearchUser) {
    // Gửi sự kiện đến ChatList.vue
    // Logic tạo chat mới được xử lý trong ChatList.vue
}

function handleSelectMessage(message: SearchMessage) {
    selectedMessageId.value = message.id;
    router.push(`/chats/${message.chatId}`);
}

onMounted(() => {
    const token = localStorage.getItem('token');
    if (!token) {
        router.push('/login');
        return;
    }
})

</script>

<style scoped></style>