<template>
    <div class="chat-list h-full p-4">
        <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-800">Chats</h3>
            <button @click="showModal = true" class="text-gray-500 hover:text-gray-700 p-1">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
            </button>
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
                        <template v-if="chat.isGroup">
                            <!-- Group Icon SVG -->
                            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-600" fill="none"
                                viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M17 20h5v-2a4 4 0 00-3-3.87M9 20H4v-2a4 4 0 013-3.87m9-3.13a4 4 0 100-8 4 4 0 000 8zm-9 0a4 4 0 100-8 4 4 0 000 8z" />
                            </svg>
                        </template>
                        <template v-else>
                            <img :src="chat.avatar" alt="Avatar" class="w-10 h-10 rounded-full" />
                        </template>
                    </div>
                    <div>
                        <p class="font-medium text-gray-800">{{ chat.name }}</p>
                        <p class="text-sm truncate"
                            :class="{ 'text-gray-800 font-semibold': chat.unreadCount > 0, 'text-gray-600': chat.unreadCount === 0 }">
                            {{
                                chat.lastMessage }}</p>
                    </div>
                </div>
            </li>
        </ul>


        <!-- Modal create -->
        <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div class="bg-white rounded-lg shadow-lg p-6 w-80">
                <input v-model="newGroupName" type="text" placeholder="Name"
                    class="w-full p-2 border border-gray-300 rounded mb-8 focus:outline-none"/>
                <div class="flex justify-end gap-2">
                    <button @click="showModal = false"
                        class="p-2 rounded bg-gray-300 hover:bg-gray-400">Cancel</button>
                    <button @click="createGroup"
                        class="p-2 rounded bg-blue-500 text-white hover:bg-blue-600">Create</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { ComponentPublicInstance, defineComponent } from 'vue'

export default defineComponent({
    name: 'ChatList',
    props: {
        selectedUser: {
            type: Object,
            default: null
        }
    },
    emits: ['select-user'],
    data() {
        return {
            chats: [
                { id: 1, name: 'John Doe', lastMessage: 'Hey, how are you?', avatar: '@/assets/images/avatar.png', unreadCount: 2 },
                { id: 2, name: 'Jane Smith', lastMessage: 'See you tomorrow!', avatar: '@/assets/images/avatar.png', unreadCount: 0 },
                { id: 3, name: 'Team Alpha', lastMessage: 'Anyone free tonight?', isGroup: true, unreadCount: 5 }
            ],
            showModal: false,
            newGroupName: '',
            selectedChatId: 0,
            chatRefs: {} as Record<number, Element | ComponentPublicInstance | null>
        }
    },
    methods: {
        selectChat(chat: any) {
            this.selectedChatId = chat.id;
            this.$emit('select-user', chat);

            // Auto scroll to item is selected
            this.$nextTick(() => {
                const el = this.chatRefs[chat.id];
                if (el && 'scrollIntoView' in el) {
                    el.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
                }
            });

            // Reset unread count khi click vào
            if (chat.unreadCount > 0) {
                chat.unreadCount = 0;
            }
        },
        createGroup() {
            if (this.newGroupName.trim() === '') {
                alert('Please enter group name');
                return;
            }

            // add new group to lists
            this.chats.push({
                name: this.newGroupName,
                lastMessage: 'Group has been created',
                isGroup: true,
                unreadCount: 0,
                id: this.chats.length + 1
            });

            // Reset & close modal
            this.newGroupName = '';
            this.showModal = false;
        }
    }
})
</script>

<style scoped>
.chat-list {
    width: 150px;
}

@media (min-width: 768px) {
    .chat-list {
        width: 300px;
        /* Desktop */
    }
}
</style>