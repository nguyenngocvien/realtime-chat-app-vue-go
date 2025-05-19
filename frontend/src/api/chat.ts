import { Chat, Message } from '@/types'
import axios from 'axios'

export async function fetchChatsByUserId(userId: string): Promise<Chat[]> {
    try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`/api/users/${userId}/chats`, {
            headers: {
                Authorization: `Bearer ${token}`
            }
        });
        return response.data.chats || [];
    } catch (error) {
        console.error('Failed to fetch chats, returning mock data:', error);
        return [];
    }
}

export async function fetchMessages(chatId: string): Promise<Message[]> {
    // try {
    //     const res = await axios.get(`/api/chats/${chatId}/messages`)
    //     return res.data.messages
    // } catch (error) {
    //     console.error('Load messages lỗi:', error)
    //     return []
    // }
    await new Promise(resolve => setTimeout(resolve, 300))

    return [
        {
            id: 'msg1',
            senderId: 'user1',
            senderName: 'Alice',
            senderAvatar: 'https://randomuser.me/api/portraits/women/1.jpg',
            isMine: false,
            content: 'Hello, how are you?',
            type: 'text',
            timestamp: new Date(Date.now() - 1000 * 60 * 10).toISOString() // 10 phút trước
        },
        {
            id: 'msg2',
            senderId: 'user2',
            senderName: 'You',
            isMine: true,
            content: 'I am fine, thanks! Here is the file you asked for.',
            type: 'file',
            fileUrl: 'https://example.com/files/sample.pdf',
            fileName: 'sample.pdf',
            fileSize: 234567,
            timestamp: new Date(Date.now() - 1000 * 60 * 5).toISOString() // 5 phút trước
        },
        {
            id: 'msg3',
            senderId: 'user1',
            senderName: 'Alice',
            senderAvatar: 'https://randomuser.me/api/portraits/women/1.jpg',
            isMine: false,
            content: 'hihihi',
            type: 'image',
            fileUrl: 'https://picsum.photos/id/237/200/300',
            timestamp: new Date(Date.now() - 1000 * 60 * 1).toISOString() // 1 phút trước
        }
    ]
}

export async function fetchChatMedia(chatId: string) {
    const res = await axios.get(`/api/chats/${chatId}/media`)
    return res.data.media
}