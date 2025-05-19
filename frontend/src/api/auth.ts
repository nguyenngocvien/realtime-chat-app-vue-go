import axios from 'axios';

interface LoginResponse {
  token: string;
}

export async function login(username: string, password: string): Promise<string> {
  try {
    const response = await axios.post<LoginResponse>('/api/login', { username, password });
    const token = response.data.token;
    localStorage.setItem('token', token);
    return token;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
}