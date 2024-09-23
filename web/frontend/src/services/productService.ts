import api from './api';
import { Product } from '../types/product';

const getAllProducts = async (): Promise<Product[]> => {
  const response = await api.get('/products');
  return response.data;
};

const getProductById = async (id: string): Promise<Product> => {
  const response = await api.get(`/products/${id}`);
  return response.data;
};

export default { getAllProducts, getProductById };
