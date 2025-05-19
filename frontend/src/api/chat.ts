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
    try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`http://localhost:8080/api/chats/${chatId}/messages?limit=50`,
            {
                headers: { Authorization: `Bearer ${token}` },
            });
        return response.data.messages.reverse();
    } catch (error) {
        console.error('Load messages lỗi:', error)
        return []
    }
}

export async function fetchChatMedia(chatId: string) {
    const res = await axios.get(`/api/chats/${chatId}/media`)
    return res.data.media
}