import { Chat, Message } from '@/types'
import axios, { AxiosError } from 'axios'

interface ErrorResponse {
    error?: string;
}

export async function fetchChatsByUserId(userId: string): Promise<Chat[]> {
    try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`/api/users/${userId}/chats`, {
            headers: { Authorization: `Bearer ${token}` },
            timeout: 5000,
        });
        return response.data.chats;
    } catch (error) {
        const axiosError = error as AxiosError<ErrorResponse>;
        throw new Error(axiosError.response?.data?.error || 'Failed to fetch chats');
    }
}

export async function createGroupChat(userId: string, groupName: string): Promise<Chat> {
    try {
        const token = localStorage.getItem('token');
        const response = await axios.post(
            '/api/chats',
            { name: groupName, type: 'group', participants: [userId] },
            {
                headers: { Authorization: `Bearer ${token}` },
                timeout: 5000,
            }
        );
        return response.data.chat;
    } catch (error) {
        const axiosError = error as AxiosError<ErrorResponse>;
        throw new Error(axiosError.response?.data?.error || 'Failed to create group chat');
    }
}

export async function fetchMessages(chatId: string): Promise<Message[]> {
    try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`/api/chats/${chatId}/messages?limit=50`,
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