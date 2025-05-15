<template>
    <div :class="['message', message?.isMine ? 'own-message' : 'other-message']">
        <!-- <p class="text-xs text-gray-500 mb-1 px-1">{{ sender }}</p> -->
        <div class="message-content p-3 rounded-lg transition duration-200 hover:shadow-sm bg-white">
            <p class="text-gray-700">{{ message?.content }}</p>
        </div>
        <small :class="['text-xs text-gray-500 mt-1 px-1', message?.isMine ? 'text-right' : 'text-left']">{{
            formattedTime }}</small>
    </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import type { Message } from '@/types'

const props = defineProps<{
    message: Message
}>()

const formattedTime = computed(() => {
    return props.message?.timestamp
        ? new Date(props.message.timestamp).toLocaleTimeString()
        : new Date().toLocaleTimeString()
})
</script>

<style scoped>
.message {
    display: flex;
    flex-direction: column;
    margin-bottom: 1rem;
    align-items: flex-start;
}

.own-message {
    align-items: flex-end;
}

.own-message .message-content {
    background-color: #e0f2fe;
}

.other-message .message-content {
    background-color: #fee2e2;
}

.message-content {
    max-width: 70%;
    border-radius: 0.75rem;
    padding: 0.75rem 1.25rem;
}
</style>