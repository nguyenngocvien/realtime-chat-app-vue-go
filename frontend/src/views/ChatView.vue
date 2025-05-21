<template>
    <div class="h-screen flex flex-col bg-gray-500">
        <!-- Top Bar -->
        <div class="flex justify-center">
            <TopBar />
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
                        <ChatBox :selected-chat="selectedChat"/>
                    </div>
                </div>
            </main>
        </div>
    </div>

</template>

<script lang="ts" setup>
import { ref } from 'vue';
import ChatList from '../components/ChatList.vue';
import type { Chat } from '@/types';
import TopBar from '@/components/TopBar.vue';
import ChatBox from '@/components/ChatBox.vue';

const selectedChat = ref<Chat | null>(null);

function onSelectChat(chat: Chat) {
    selectedChat.value = chat;
}
</script>

<style scoped></style>