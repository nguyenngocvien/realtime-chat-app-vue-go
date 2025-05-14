<template>
    <div :class="['message', isOwnMessage ? 'own-message' : 'other-message']">
        <!-- <p class="text-xs text-gray-500 mb-1 px-1">{{ sender }}</p> -->
        <div class="message-content p-3 rounded-lg transition duration-200 hover:shadow-sm bg-white">
            <p class="text-gray-700">{{ text }}</p>
        </div>
        <small  :class="['text-xs text-gray-500 mt-1 px-1', isOwnMessage ? 'text-right' : 'text-left']">{{ formattedTime }}</small>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'Message',
    props: {
        text: {
            type: String,
            required: true
        },
        sender: {
            type: String,
            required: true
        },
        isOwnMessage: {
            type: Boolean,
            required: true
        },
        timestamp: {
            type: Date,
            default: () => new Date()
        }
    },
    computed: {
        formattedTime(): string {
            return this.timestamp
                ? new Date(this.timestamp).toLocaleTimeString()
                : new Date().toLocaleTimeString()
        }
    }
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