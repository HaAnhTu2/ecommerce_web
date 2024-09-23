// src/components/ProductDetail.tsx
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getProductById } from '../services/api';
import { Product } from '../types/product';

const ProductDetailPage: React.FC= () => {
    const { id } = useParams<{ id: string }>();
    const [product, setProduct] = useState<Product | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchProduct = async () => {
            if (!id) {
                console.error('Product ID is undefined');
                setLoading(false);
                return;
            }

            try {
                const fetchedProduct = await getProductById(id);
                setProduct(fetchedProduct);
                setLoading(false);
            } catch (error) {
                console.error('Error fetching product:', error);
                setLoading(false);
            }
        };
        fetchProduct();
    }, [id]);

    if (loading) {
        return <div>Loading...</div>;
    }

    if (!product) {
        return <div>Product not found</div>;
    }

    return (
        <div className="card mb-3">
            <div className="row">
            <h4 className="card-title">{product.name}</h4>
                <div className='col-md-4'>
                    <div className="card-body">
                        <p className="card-text">Description: {product.description}</p>
                        <p className="card-text">Price: {product.price}</p>
                        <p className="card-text">Stock: {product.stock}</p>
                        <p className="card-text">Category: {product.category}</p>
                    </div>
                    <button type="submit" className="btn btn-outline-dark">buy</button>
                    <button type="submit" className="btn btn-outline-dark">add to cart</button>
                </div>
            </div>
        </div>
    );
};

export default ProductDetailPage;
