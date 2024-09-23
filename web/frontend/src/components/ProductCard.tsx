import React from 'react';
import { Link } from 'react-router-dom';
import { Product } from '../types/product';

const ProductCard: React.FC<{ product: Product }> = ({ product }) => {
  return (
    <div>
      <h2>{product.name}</h2>
      <p>{product.description}</p>
      <Link to={`/products/${product._id}`}>View Details</Link>
    </div>
  );
};

export default ProductCard;
