'use client'
import React, { useState, useEffect } from 'react';
import { useAuth } from "@clerk/nextjs";
import ProductCard from '@/components/Product';
import productsData from '@/public/products_dummy_data.json';
import styles from '@/app/page.module.css';

interface Product {
  id: number;
  shopId: number;
  categoryId: number;
  title: string;
  description: string;
  price: number;
  imageUrl: string;
}

const Page: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [cart, setCart] = useState<{product: Product, quantity: number}[]>([]);
  const { getToken } = useAuth();
  const shopId = 1;

  useEffect(() => {
    // For now, we'll use the dummy data instead of fetching from an API
    setProducts(productsData);
  }, []);

  const handleAddToCart = (product: Product, quantity: number) => {
    setCart(prevCart => {
      const existingItem = prevCart.find(item => item.product.id === product.id);
      if (existingItem) {
        return prevCart.map(item => 
          item.product.id === product.id 
            ? { ...item, quantity: item.quantity + quantity }
            : item
        );
      } else {
        return [...prevCart, { product, quantity }];
      }
    });
  };

  const getTotalPrice = () => {
    return cart.reduce((total, item) => total + item.product.price * item.quantity, 0);
  };

  return (
    <div className={styles.page}>
      <div className={styles.productList}>
        {products.map(product => (
          <ProductCard 
            key={product.id} 
            product={product} 
            onAddToCart={handleAddToCart} 
          />
        ))}
      </div>
      <div className={styles.checkoutSection}>
        <h2>Checkout</h2>
        {cart.map(item => (
          <div key={item.product.id} className={styles.cartItem}>
            <span>{item.product.title}</span>
            <span>x {item.quantity}</span>
            <span>${(item.product.price * item.quantity).toFixed(2)}</span>
          </div>
        ))}
        <p className={styles.total}>Total: ${getTotalPrice().toFixed(2)}</p>
        <button className={styles.checkoutButton}>Proceed to Checkout</button>
      </div>
    </div>
  );
};

export default Page;