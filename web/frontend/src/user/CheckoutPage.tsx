import React, { useState, useEffect } from 'react';
import { useAuth } from '../services/authService';
import cartService from '../services/cartService';
import { CartItem } from '../types/cart';
import { formatPrice } from '../utils/formatPrice';
import { validateEmail } from '../utils/validate';

const CheckoutPage: React.FC = () => {
  const { user } = useAuth();
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [email, setEmail] = useState(user?.email || '');
  const [address, setAddress] = useState('');
  const [paymentMethod, setPaymentMethod] = useState('Credit Card');
  const [errors, setErrors] = useState<{ email?: string; address?: string }>({});

  useEffect(() => {
    const fetchCartItems = async () => {
      try {
        const items = await cartService.getCartItems();
        setCartItems(items);
      } catch (error) {
        console.error('Failed to fetch cart items', error);
      }
    };

    fetchCartItems();
  }, []);

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
    if (!validateEmail(e.target.value)) {
      setErrors((prev) => ({ ...prev, email: 'Invalid email address' }));
    } else {
      setErrors((prev) => ({ ...prev, email: '' }));
    }
  };

  const handleAddressChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAddress(e.target.value);
    if (e.target.value.trim() === '') {
      setErrors((prev) => ({ ...prev, address: 'Address is required' }));
    } else {
      setErrors((prev) => ({ ...prev, address: '' }));
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Perform validation checks
    if (!validateEmail(email)) {
      setErrors((prev) => ({ ...prev, email: 'Invalid email address' }));
      return;
    }

    if (address.trim() === '') {
      setErrors((prev) => ({ ...prev, address: 'Address is required' }));
      return;
    }

    // Process checkout
    console.log('Processing checkout with the following details:');
    console.log('Email:', email);
    console.log('Address:', address);
    console.log('Payment Method:', paymentMethod);
    console.log('Cart Items:', cartItems);
  };

  const totalPrice = cartItems.reduce((acc, item) => acc + item.product.price * item.quantity, 0);

  return (
    <div className="checkout-page">
      <h1>Checkout</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={handleEmailChange}
            required
          />
          {errors.email && <span className="error">{errors.email}</span>}
        </div>
        <div>
          <label htmlFor="address">Address:</label>
          <input
            type="text"
            id="address"
            value={address}
            onChange={handleAddressChange}
            required
          />
          {errors.address && <span className="error">{errors.address}</span>}
        </div>
        <div>
          <label htmlFor="payment-method">Payment Method:</label>
          <select
            id="payment-method"
            value={paymentMethod}
            onChange={(e) => setPaymentMethod(e.target.value)}
          >
            <option value="Credit Card">Credit Card</option>
            <option value="PayPal">PayPal</option>
            <option value="Bank Transfer">Bank Transfer</option>
          </select>
        </div>
        <button type="submit">Place Order</button>
      </form>
      <h2>Order Summary</h2>
      <ul>
        {cartItems.map((item) => (
          <li key={item.product._id}>
            {item.product.name} - {item.quantity} x {formatPrice(item.product.price)}
          </li>
        ))}
      </ul>
      <p>Total: {formatPrice(totalPrice)}</p>
    </div>
  );
};

export default CheckoutPage;
