<template>
    <div class="flex max-w-6xl rounded-xl bg-gray-400 w-full overflow-hidden">
        <div>
            <ChatList :selected-chat="selectedChat" @select-chat="selectUser"/>
        </div>
        <div class="flex-1">
            <ChatBox :selected-chat="selectedChat"/>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ChatList from '@/components/ChatList.vue';
import ChatBox from '@/components/ChatBox.vue'
import { Chat } from '@/types';

export default defineComponent({
    name: 'ChatView',
    components: { ChatList, ChatBox },
    data() {
        return {
            selectedChat: null as any
        }
    },
    methods: {
        selectUser(chat: Chat) {
            this.selectedChat = chat;
        }
    }
})
</script>

<style scoped></style>