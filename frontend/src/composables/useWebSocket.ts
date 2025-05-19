import { ref, onBeforeUnmount } from 'vue'
import { v4 as uuidv4 } from 'uuid'
import type { Message as MessageType, Chat, User } from '@/types'

export function useWebSocket(currentUser: User, selectedChat: Chat | null, onMessageReceived: (msg: MessageType) => void) {
    const ws = ref<WebSocket | null>(null)
    const reconnectTimeout = ref<ReturnType<typeof setTimeout> | null>(null)
    const isManuallyClosed = ref(false)

    function connectWebSocket() {
        const token = localStorage.getItem('token')
        ws.value = new WebSocket(`ws://localhost:8080/ws?token=${token}`)
        if (ws.value) {
            ws.value.close()
        }
        if (!selectedChat) return

        isManuallyClosed.value = false // reset flag

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
