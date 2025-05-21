import axios, { AxiosError } from 'axios';

interface LoginResponse {
  token: string;
}

interface ErrorResponse {
  error: string;
}

export async function login(username: string, password: string): Promise<string> {
  try {
    const response = await axios.post<LoginResponse>('/api/login', { username, password }, {
      timeout: 5000, // Timeout 5s
    });
    const token = response.data.token;
    if (!token || typeof token !== 'string') {
      throw new Error('Invalid token received');
    }
    localStorage.setItem('token', token);
    return token;
  } catch (error) {
    const axiosError = error as AxiosError<ErrorResponse>;
    let errorMessage = 'Login failed';
    if (axiosError.response) {
      if (axiosError.response.status === 401) {
        errorMessage = 'Invalid credentials';
      } else if (axiosError.response.status === 400) {
        errorMessage = axiosError.response.data?.error || 'Invalid request';
      } else {
        errorMessage = 'Server error';
      }
    } else if (axiosError.request) {
      errorMessage = 'No response from server';
    }
    console.error('Login failed:', error);
    throw new Error(errorMessage);
  }
}

export async function signup(username: string, password: string): Promise<string> {
  try {
    const response = await axios.post<LoginResponse>('/api/signup', { username, password }, {
      timeout: 5000,
    });
    const token = response.data.token;
    if (!token || typeof token !== 'string') {
      throw new Error('Invalid token received');
    }
    localStorage.setItem('token', token);
    return token;
  } catch (error) {
    const axiosError = error as AxiosError<ErrorResponse>;
    let errorMessage = 'Signup failed';
    if (axiosError.response) {
      if (axiosError.response.status === 400) {
        errorMessage = axiosError.response.data?.error || 'Invalid request';
      } else if (axiosError.response.status === 409) {
        errorMessage = 'Username already exists';
      } else {
        errorMessage = 'Server error';
      }
    } else if (axiosError.request) {
      errorMessage = 'No response from server';
    }
    console.error('Signup failed:', error);
    throw new Error(errorMessage);
  }
}