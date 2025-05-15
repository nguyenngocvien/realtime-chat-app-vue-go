export interface Message {
    id: string
    senderId: string
    senderName: string
    senderAvatar?: string
    recipientId?: string;
    isMine: boolean
    content: string
    type: 'text' | 'image' | 'file'
    fileUrl?: string
    fileName?: string
    fileSize?: number
    timestamp: string  // ISO format (UTC)
}

export interface User {
    id: string;
    username: string;
    name: string;
    email?: string;
    avatar?: string;
}

export interface Chat {
    id: string
    type: 'user' | 'group'
    name: string
    avatar?: string
    lastMessage?: Message
    unreadCount: number
    participants?: User[]
}