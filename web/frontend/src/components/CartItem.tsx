import React from 'react';
import { CartItem as CartItemType } from '../types/cart';

const CartItem: React.FC<{ item: CartItemType }> = ({ item }) => {
  return (
    <div>
      <h3>{item.product.name}</h3>
      <p>Quantity: {item.quantity}</p>
      <p>Price: ${item.product.price}</p>
    </div>
  );
};

export default CartItem;
