export interface Message {
    text: string;
    sender: string;
    timestamp: Date;
    isMine?: boolean;
}

export interface User {
    id: string;
    username: string;
}