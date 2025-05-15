import { ref, onBeforeUnmount } from 'vue'
import { v4 as uuidv4 } from 'uuid'
import type { Message as MessageType, Chat, User } from '@/types'

const WS_BASE_URL = import.meta.env.VITE_WS_URL

export function useWebSocket(currentUser: User, selectedChat: Chat | null, onMessageReceived: (msg: MessageType) => void) {
    const ws = ref<WebSocket | null>(null)
    const reconnectTimeout = ref<ReturnType<typeof setTimeout> | null>(null)
    const isManuallyClosed = ref(false)

    function connectWebSocket() {
        if (ws.value) {
            ws.value.close()
        }
        if (!selectedChat) return

        isManuallyClosed.value = false // reset flag

        const url = `${WS_BASE_URL}?userId=${currentUser.id}`
        console.log('Connecting to WebSocket:', url)
        
        ws.value = new WebSocket(url)

        ws.value.onopen = () => {
            console.log('WebSocket connected')
        }

        ws.value.onmessage = (event) => {
            const msg: MessageType = JSON.parse(event.data)
            if (msg.senderId === selectedChat.id || msg.recipientId === selectedChat.id) {
                onMessageReceived({
                    id: msg.id?.toString() || uuidv4(),
                    content: msg.content,
                    senderId: msg.senderId,
                    senderName: msg.senderName || selectedChat.name || 'Unknown',
                    timestamp: new Date(msg.timestamp).toISOString(),
                    type: 'text',
                    isMine: msg.senderId === currentUser.id
                })
            }
        }

        ws.value.onclose = () => {
            console.log('WebSocket disconnected')
            ws.value = null
            if (!isManuallyClosed.value) {
                attemptReconnect()
            }
        }

        ws.value.onerror = (error) => {
            console.error('WebSocket error:', error)
            ws.value?.close()
        }
    }

    function attemptReconnect() {
        if (reconnectTimeout.value) {
            clearTimeout(reconnectTimeout.value)
        }
        console.log('🔄 Attempting to reconnect in 5 seconds...')
        reconnectTimeout.value = setTimeout(() => {
            connectWebSocket()
        }, 5000)
    }

    function disconnectWebSocket() {
        isManuallyClosed.value = true
        if (reconnectTimeout.value) {
            clearTimeout(reconnectTimeout.value)
            reconnectTimeout.value = null
        }

        if (ws.value) {
            ws.value.close()
            ws.value = null
        }
    }

    onBeforeUnmount(() => {
        disconnectWebSocket()
    })

    return {
        ws,
        connectWebSocket,
        disconnectWebSocket
    }
}
