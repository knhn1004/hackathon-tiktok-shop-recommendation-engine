"use client";

import React, { useState, useEffect } from 'react';
import { useParams } from 'next/navigation';
import { useAuth } from '@clerk/nextjs';
import Link from 'next/link';
import styles from './ProductDetail.module.css';

interface Product {
  id: number;
  title: string;
  description: string;
  price: number;
  imageUrl: string;
  shop: {
    name: string;
  };
  shopId: number;
  category: {
    name: string;
  };
}

const ProductDetail: React.FC = () => {
  const { id } = useParams();
  const { getToken } = useAuth();
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const token = await getToken();
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/products/${id}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch product');
        }

        const data = await response.json();
        setProduct(data);
      } catch (err) {
        setError('Error fetching product details');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchProduct();
  }, [id, getToken]);

  if (loading) {
    return <div className={styles.loading}>Loading...</div>;
  }

  if (error) {
    return <div className={styles.error}>{error}</div>;
  }

  if (!product) {
    return <div className={styles.error}>Product not found</div>;
  }

  return (
    <div className={styles.container}>
      {product && (
        <div className={styles.breadcrumb}>
          <Link href="/">Home</Link> &gt; 
          <Link href={`/shop/${product.shopId}`}>Shop</Link> &gt; 
          <span>{product.title}</span>
        </div>
      )}
      <h1 className={styles.title}>{product.title}</h1>
      <div className={styles.content}>
        <div className={styles.imageContainer}>
          <img
            src="https://via.placeholder.com/400x300?text=Product+Image"
            alt={product.title}
            className={styles.image}
          />
        </div>
        <div className={styles.details}>
          <p className={styles.description}>{product.description}</p>
          <p className={styles.price}>Price: ${product.price.toFixed(2)}</p>
          <p className={styles.shop}>Shop: {product.shop.name}</p>
          <p className={styles.category}>Category: {product.category.name}</p>
          <button className={styles.buyButton}>Add to Cart</button>
        </div>
      </div>
    </div>
  );
};

export default ProductDetail;