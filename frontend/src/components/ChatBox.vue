<template>
    <div class="chat-box shadow-lg p-4 bg-gray-400 w-full">
        <!-- Top Bar -->
        <div class="h-8 text-gray-700 mb-4">
            <h2 class="text-lg font-semibold">{{ selectedUser ? selectedUser.name : 'Select a chat' }}
            </h2>
        </div>
        <div ref="messagesContainer"
            class="messages h-[calc(100vh-16rem)] overflow-y-auto mb-6 rounded-lg scrollbar-none scrollbar-thumb-gray-300 scrollbar-track-gray-100">
            <div v-if="!selectedUser" class="text-center text-gray-500">
                Select a user to start chatting
            </div>
            <Message v-for="(message, index) in messages" :key="index" :text="message.text" :sender="message.sender"
                :timestamp="message.timestamp" :isOwnMessage="message.sender === currentUser.username" />
        </div>

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
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Message from './Message.vue'
import { Message as MessageType, User } from '../types'
import { nextTick } from 'vue'

export default defineComponent({
    name: 'ChatBox',
    components: { Message },
    props: {
        selectedUser: {
            type: Object,
            default: null
        }
    },
    data() {
        return {
            newMessage: '' as string,
            currentUser: { id: '1', username: 'You' } as User,
            chatData: {
                1: [
                    { text: 'Hey, how are you?', sender: 'John Doe', timestamp: new Date(), isMine: false },
                    { text: 'I’m good, thanks!', sender: 'You', timestamp: new Date(), isMine: true }
                ],
                2: [
                    { text: 'See you tomorrow!', sender: 'Jane Smith', timestamp: new Date(), isMine: false },
                    { text: 'Sure thing!', sender: 'You', timestamp: new Date(), isMine: true }
                ],
                3: [
                    { text: 'Anyone free tonight?', sender: 'Group Chat', timestamp: new Date(), isMine: false },
                    { text: 'I’m in!', sender: 'You', timestamp: new Date(), isMine: true }
                ]
            } as { [key: string]: MessageType[] }
        }
    },
    computed: {
        messages(): MessageType[] {
            if (!this.selectedUser) return [];
            return this.chatData[this.selectedUser.id] || [];
        }
    },
    methods: {
        sendMessage() {
            if (this.newMessage.trim() && this.selectedUser) {
                const newMsg: MessageType = {
                    text: this.newMessage,
                    sender: this.currentUser.username,
                    timestamp: new Date(),
                    isMine: true
                };
                if (!this.chatData[this.selectedUser.id]) {
                    this.chatData[this.selectedUser.id] = [];
                }
                this.chatData[this.selectedUser.id].push(newMsg);
                this.newMessage = '';
                // TODO: Integrate with WebSocket to send to backend

                this.scrollToBottom();
            }
        },

        scrollToBottom() {
            nextTick(() => {
                const messagesContainer = this.$refs.messagesContainer as HTMLElement;
                if (messagesContainer) {
                    messagesContainer.scrollTop = messagesContainer.scrollHeight;
                }
            });
        }
    }
})
</script>

<style scoped>
.chat-box {
    display: flex;
    flex-direction: column;
}

.messages {
    display: flex;
    flex-direction: column;
    gap: 1rem;
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
</style>