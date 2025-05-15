<template>
    <div v-if="chat" class="h-full bg-gray-400 rounded-lg shadow-lg max-w-md w-full p-6 relative">
        <div class="flex flex-col items-center space-y-4">
            <img v-if="chat.avatar" :src="chat.avatar" alt="User Avatar" class="w-24 h-24 rounded-full object-cover" />
            <h3 class="text-xl font-semibold">{{ chat.name }}</h3>

            <!-- User -->
            <template v-if="chat.type === 'user'">
                <p class="text-gray-600">{{ chat.name }}</p>
            </template>

            <!-- Group -->
            <template v-else-if="chat.type === 'group'">
                <div class="w-full">
                    <ul class="space-y-1">
                        <li v-for="user in chat.participants" :key="user.id" class="flex items-center space-x-2">
                            <img v-if="user.avatar" :src="user.avatar" alt="Avatar"
                                class="w-6 h-6 rounded-full object-cover" />
                            <span>{{ user.name }}</span>
                        </li>
                    </ul>
                </div> 
            </template>
        </div>

        <!-- Images -->
        <div v-if="images && images.length" class="mt-6">
            <h4 class="text-lg font-semibold mb-2">Hình ảnh</h4>
            <div class="grid grid-cols-3 gap-2">
                <img v-for="(img, index) in images" :key="index" :src="img"
                    class="w-full h-24 object-cover rounded-md" />
            </div>
        </div>

        <!-- Files -->
        <div v-if="files && files.length" class="mt-6">
            <h4 class="text-lg font-semibold mb-2">Files</h4>
            <ul class="space-y-2">
                <li v-for="(file, index) in files" :key="index" class="flex items-center space-x-2">
                    <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 16v1a2 2 0 002 2h12a2 2 0 002-2v-1M4 12h16M4 8h16M4 4h16" />
                    </svg>
                    <a :href="file.url" target="_blank" class="text-blue-600 hover:underline truncate">{{ file.name
                    }}</a>
                </li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts">
import { Chat, User } from '@/types'
import axios from 'axios'
import { defineComponent, PropType, ref, watch } from 'vue'

export default defineComponent({
    name: 'ChatDetail',
    props: {
        chat: {
            type: Object as PropType<Chat>,
            required: true
        }
    },
    emits: ['close'],
    setup(props, { emit }) {
        const images = ref<string[]>([])
        const files = ref<{ name: string, url: string }[]>([])

        const fetchMedia = async (chatId: string) => {
            try {
                const res = await axios.get(`/api/chats/${chatId}/media`)
                const media = res.data.media || []

                images.value = media
                    .filter((item: any) => item.type === 'image')
                    .map((item: any) => item.fileUrl)

                files.value = media
                    .filter((item: any) => item.type === 'file')
                    .map((item: any) => ({ name: item.fileName, url: item.fileUrl }))
            } catch (err) {
                console.error('Failed to fetch media:', err)
            }
        }

        // Watch chat change (optional, if chat can change dynamically)
        watch(() => props.chat.id, (newId) => {
            if (newId) fetchMedia(newId)
        }, { immediate: true })

        const close = () => emit('close')

        return {
            images,
            files,
            close
        }
    }
})
</script>