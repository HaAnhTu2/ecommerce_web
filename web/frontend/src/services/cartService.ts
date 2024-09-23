import api from './api';
import { CartItem } from '../types/cart';

const getCartItems = async (): Promise<CartItem[]> => {
  const response = await api.get('/cart');
  return response.data;
};

export default { getCartItems };
