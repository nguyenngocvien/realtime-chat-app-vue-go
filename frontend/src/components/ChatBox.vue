<template>
    <div class="chat-box shadow-lg bg-gray-400 w-full h-[calc(100vh-6rem)] flex">
        <div v-if="selectedChat" :class="['chat-content', { 'shrink': isShowDetail }]"
            class="flex-1 flex flex-col p-4 transition-width duration-300 ease-in-out">
            <!-- Top Bar -->
            <div class="h-8 text-gray-700 mb-4 px-2 flex items-center justify-between">
                <h2 class="text-lg font-semibold">{{ selectedChat ? selectedChat.name : 'Select a chat' }}
                </h2>
                <Icon @click="onIconClick" name="info" customClass="w-5 h-5 text-gray-600 hover:text-gray-800" />
            </div>
            <TransitionGroup name="message-fade" tag="div" ref="messagesContainer"
                class="messages flex-1 overflow-y-auto mb-6 rounded-lg scrollbar-none scrollbar-thumb-gray-300 scrollbar-track-gray-100">

                <Message v-for="(message, index) in messages" :key="index" :message="message" />

            </TransitionGroup>

            <!-- Input Area -->
            <div class="input-area flex gap-2 rounded-lg overflow-hidden bg-gray-300">
                <input v-model="newMessage" @keyup.enter="sendMessage" type="text" placeholder="Type a message..."
                    class="flex-1 p-3 focus:outline-none bg-transparent" />
                <button @click="sendMessage" class=" text-gray-500 p-3 transition duration-200">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                    </svg>
                </button>
            </div>
        </div>
        <div v-else class="flex-1 p-4 flex items-center justify-center w-full h-full">
            <div class="text-center text-gray-500">
                <h3>Welcome!</h3>
                Select a chat to start chatting
            </div>
        </div>

        <Transition name="slide-fade">
            <div v-if="isShowDetail && selectedChat" class="w-80">
                <ChatDetail :chat="selectedChat" />
            </div>
        </Transition>
    </div>
</template>

<script lang="ts" setup>
import { ref, watch, nextTick, defineProps, onUnmounted } from 'vue'
import { v4 as uuidv4 } from 'uuid'

import Message from './Message.vue'
import Icon from './base/Icon.vue'
import ChatDetail from './ChatDetail.vue'

import type { Chat, Message as MessageType, User } from '../types'

// Props
const props = defineProps<{
    selectedChat: Chat | null
}>()

// State
const isShowDetail = ref(false)
const newMessage = ref('')
const currentUser = ref<User>({ id: '1', username: 'You', name: 'You' })
const messages = ref<MessageType[]>([])
const { ws, connectWebSocket } = useWebSocket(
    currentUser.value,
    props.selectedChat,
    (msg: MessageType) => {
        messages.value.push(msg)
        scrollToBottom()
    }
)

// Methods
async function loadMessages() {
    if (!props.selectedChat) return

    try {
        messages.value = await fetchMessages(props.selectedChat.id);
        scrollToBottom()
    } catch (error) {
        console.error('Failed to fetch messages:', error)
    }
}

function sendMessage() {
    if (
        newMessage.value.trim() && 
        props.selectedChat && 
        ws.value &&
        ws.value.readyState === WebSocket.OPEN
    ) {
        const newMsg: MessageType = {
            id: uuidv4(),
            content: newMessage.value,
            senderId: currentUser.value.id,
            senderName: currentUser.value.name,
            timestamp: new Date().toISOString(),
            type: 'text',
            isMine: true,
        };
        ws.value.send(JSON.stringify({
            recipientId: props.selectedChat.id,
            text: newMessage.value
        }));

        messages.value.push(newMsg)
        newMessage.value = ''

        scrollToBottom()
    }
}

function onIconClick() {
    isShowDetail.value = !isShowDetail.value
}

function scrollToBottom() {
    nextTick(() => {
        const messagesContainer = refs.messagesContainer.value
        if (messagesContainer) {
            messagesContainer.scrollTop = messagesContainer.scrollHeight
        }
    })
}

// Watchers
watch(() => props.selectedChat, () => {
    loadMessages();
    connectWebSocket();
}, { immediate: true })

watch(messages, () => {
    scrollToBottom()
}, { deep: true })

// Refs (template refs)
import { ref as vueRef } from 'vue'
import { fetchMessages } from '@/api/chat'
import { useWebSocket } from '@/composables/useWebSocket'
const refs = {
    messagesContainer: vueRef<HTMLElement | null>(null)
}

// Cleanup
onUnmounted(() => {
    if (ws.value) {
        ws.value.close();
    }
});
</script>


<style scoped lang="css">
.chat-box {
    display: flex;
    flex-direction: row;
}

.messages {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    scroll-behavior: smooth;
    flex-grow: 1;
}

/* Custom scrollbar */
.scrollbar-none {
    scrollbar-width: none;
}

.scrollbar-thumb-gray-300::-webkit-scrollbar-thumb {
    background-color: #d1d5db;
    border-radius: 0.25rem;
}

.scrollbar-track-gray-100::-webkit-scrollbar-track {
    background-color: #f3f4f6;
}

.message-fade-enter-active,
.message-fade-leave-active {
    transition: all 0.3s ease;
}

.message-fade-enter-from {
    opacity: 0;
    transform: translateY(10px);
}

.message-fade-enter-to {
    opacity: 1;
    transform: translateY(0);
}

.message-fade-leave-from {
    opacity: 1;
    transform: translateY(0);
}

.message-fade-leave-to {
    opacity: 0;
    transform: translateY(10px);
}

.chat-content {
    transition: width 0.3s ease-in-out;
    width: 100%;
}

.chat-content.shrink {
    width: calc(100% - 20rem);
    /* Trừ đi chiều rộng của ChatDetail (w-80 = 20rem) + margin */
}


.slide-fade-enter-active {
    transition: all 0.3s ease;
}

.slide-fade-leave-active {
    transition: all 0.3s ease;
}

.slide-fade-enter-from {
    transform: translateX(100%);
    opacity: 0;
}

.slide-fade-enter-to {
    transform: translateX(0);
    opacity: 1;
}

.slide-fade-leave-from {
    transform: translateX(0);
    opacity: 1;
}

.slide-fade-leave-to {
    transform: translateX(100%);
    opacity: 0;
}
</style>