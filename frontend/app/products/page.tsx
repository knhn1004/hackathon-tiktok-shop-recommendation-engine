'use client'
import React, { useState, useEffect } from 'react';
import { useAuth } from "@clerk/nextjs";

interface Product {
  ID: number,
  ShopId: number,
  CategoryID: number,
  Title: string,
  Description: string,
  Price: number,
  ImageURL: string,
}

const Page: React.FC = () => {
  const [product, setProduct] = useState<Product | null>(null);
  const [quantity, setQuantity] = useState(1);
  const { getToken } = useAuth();
  const shopId = 1;
  
  useEffect(() => {
    const fetchProduct = async () => {
      try {
        const token = await getToken({ template: 'default' });
        const response = await fetch(`http://localhost:8080/api/shops/${shopId}/products`, {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
          }
        });
        const data = await response.json();
        setProduct(data[0]);
      } catch (error) {
        console.error('Error fetching product:', error);
      }
    };
    
    fetchProduct();
  }, []);
  
  const handleAddToCart = async () => {
    try {
        const token = await getToken({ template: 'default' });
        const response = await fetch(`http://localhost:8080/api/shops/${shop_id}/products`, {
            headers: {
                'Authorization': `Bearer ${token}`,
              'Content-Type': 'application/json',
            },
        });
        const data = await response.json();

    } catch (error) {
        console.error('Error adding to cart', error);
    }
  };

  if (!product) {
    return <div>Loading...</div>;
  }

  return (
    <div className="product-page">
      <h1>{product.Title}</h1>
      <img src={product.ImageURL} alt={product.Title} />
      <p>{product.Description}</p>
      <p>Price: ${product.Price}</p>
      <div>
        <label htmlFor="quantity">Quantity:</label>
        <input
          type="number"
          id="quantity"
          value={quantity}
          onChange={(e) => setQuantity(Math.max(1, parseInt(e.target.value)))}
          min="1"
        />
      </div>
      <button onClick={handleAddToCart}>Add to Cart</button>
    </div>
  );

};

export default Page; 