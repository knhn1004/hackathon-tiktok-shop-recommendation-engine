'use client';
import React, { useState, useEffect } from 'react';
import { useParams } from 'next/navigation';
import { useAuth } from '@clerk/nextjs';
import ProductCard from '@/components/Product';
import Link from 'next/link';
import styles from './ShopPage.module.css';

interface Product {
	id: number;
	title: string;
	description: string;
	price: number;
	imageUrl: string;
	shopId: number;
	categoryId: number;
}

const ShopPage: React.FC = () => {
	const { id } = useParams();
	const { getToken } = useAuth();
	const [products, setProducts] = useState<Product[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const fetchProducts = async () => {
			try {
				const token = await getToken({ template: 'default' });
				const response = await fetch(
					`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/shops/${id}/products`,
					{
						headers: {
							Authorization: `Bearer ${token}`,
						},
					}
				);

				if (!response.ok) {
					throw new Error('Failed to fetch products');
				}

				const data = await response.json();
				setProducts(data);
			} catch (err) {
				setError('Error fetching products');
				console.error(err);
			} finally {
				setLoading(false);
			}
		};

		fetchProducts();
	}, [id, getToken]);

	const handleAddToCart = (product: Product, quantity: number) => {
		// Implement add to cart functionality here
		console.log(`Added ${quantity} of ${product.title} to cart`);
	};

	if (loading) {
		return <div className={styles.loading}>Loading...</div>;
	}

	if (error) {
		return <div className={styles.error}>{error}</div>;
	}

	return (
		<div className={styles.container}>
			<div className={styles.breadcrumb}>
				<Link href="/">Home</Link> &gt;
				<Link href={`/shop/${id}`}>Shop {id}</Link>
			</div>
			<h1 className={styles.title}>Shop {id}</h1>
			<div className={styles.productGrid}>
				{products.map(product => (
					<ProductCard
						key={product.id}
						product={product}
						onAddToCart={handleAddToCart}
					/>
				))}
			</div>
		</div>
	);
};

export default ShopPage;
