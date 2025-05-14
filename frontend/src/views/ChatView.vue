<template>
    <div class="flex max-w-6xl rounded-xl bg-gray-400 w-full overflow-hidden">
        <ChatList :selected-user="selectedUser" @select-user="selectUser"/>
        <ChatBox :selected-user="selectedUser"/>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ChatList from '@/components/ChatList.vue';
import ChatBox from '@/components/ChatBox.vue'

export default defineComponent({
    name: 'ChatView',
    components: { ChatList, ChatBox },
    data() {
        return {
            selectedUser: null as any
        }
    },
    methods: {
        selectUser(user: any) {
            this.selectedUser = user;
            console.log('Selected user:', user);
        }
    }
})
</script>

<style scoped></style>