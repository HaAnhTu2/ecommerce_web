// src/api/api.ts
import axios from 'axios';
import { User } from '../types/user';
import { Product } from '../types/product';

export const login = async (email: string, password: string): Promise<string> => {
    try {
        const response = await axios.post('api/login', {
            email,
            password,
        });
        const { token } = response.data;
        localStorage.setItem('token', token); // Save the token to localStorage
        return token;
        } catch (error) {
            throw new Error('Error logging in');
        }
};

export const logout = async (): Promise<void> => {
    try {
        await axios.delete('api/logout');
        localStorage.removeItem('token');
    } catch (error) {
        throw new Error('Error logging out');
    }
  };
export const getUsers = async (): Promise<User[]> => {
    const response = await axios.get('/api/user/get')//, {
    return response.data.users; 
  };

export const createUser = async (newUser: FormData): Promise<User> => {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
        const response = await axios.post('/api/user/create', newUser, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Authorization: `Bearer ${token}`,
            },
        });
        return response.data as User;
    } catch (error) {
        throw new Error('Error creating user');
    }
};

export const updateUser =async (id: string, user: Omit<User,'id'>):Promise<User> =>{
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
    const response = await axios.put(`/api/user/update/${id}`,user,{
        headers:{
            'Content-Type':'multipart/form-data',
            Authorization: `Bearer ${token}`,
        }
    });
    return response.data;
} catch (error) {
    throw new Error('Error updating user');
}
};

export const deleteUser = async (id: string):Promise<void> =>{
    const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
        const response = await axios.delete(`/api/user/delete/${id}`,{
            headers:{
                'Content-Type':'multipart/form-data',
                Authorization: `Bearer ${token}`,
            }
        });
        return response.data
};

export const getProducts = async (): Promise<Product[]> => {
    const response = await axios.get('/api/product/get')
    return response.data.products;
};

export const findNameProduct = async (name: string):Promise<Product> =>{
    const response = await axios.get(`/product/${name}`)
    return response.data.products;
};
export const getProductById = async (id: string): Promise<Product> => {
    const response = await axios.get(`/api/product/get/${id}`);
    return response.data.product;
};
export const createProduct = async (newProduct: FormData): Promise<Product> =>  {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
        const response = await axios.post('/api/product/create', newProduct, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Authorization: `Bearer ${token}`,
            },
        });
        return response.data as Product;
    } catch (error) {
        throw new Error('Error creating product');
    }
};

export const updateProduct =async (id: string, product: Omit<Product,'id'>):Promise<Product> =>{
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
    const response = await axios.put(`/api/product/update/${id}`,product,{
        headers:{
            'Content-Type':'multipart/form-data',
            Authorization: `Bearer ${token}`,
        }
    });
    return response.data;
} catch (error) {
    throw new Error('Error updating user');
}
};

export const deleteProduct = async (id: string):Promise<void> =>{
    const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('Token not found');
        }
        const response = await axios.delete(`/api/product/delete/${id}`,{
        headers:{
            'Content-Type':'multipart/form-data',
            Authorization: `Bearer ${token}`,
        }
    });
    return response.data
}