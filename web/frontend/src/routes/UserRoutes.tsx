import React from 'react';
import { Routes, Route } from 'react-router-dom';
import HomePage from '../user/HomePage';
import ProductPage from '../user/ProductPage';
import ProductDetailPage from '../user/ProductDetailPage';
import CartPage from '../user/CartPage';
import CheckoutPage from '../user/CheckoutPage';
import UserProfilePage from '../user/UserProfilePage';
import PrivateRoute from './PrivateRoute';

const UserRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/products" element={<ProductPage />} />
      <Route path="/products/:id" element={<ProductDetailPage />} />
      <Route path="/cart" element={<CartPage />} />
      <Route path="/checkout" element={<CheckoutPage />} />
      <Route path="/profile" element={<PrivateRoute><UserProfilePage /></PrivateRoute>} />
    </Routes>
  );
};

export default UserRoutes;
