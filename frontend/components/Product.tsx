import React from 'react';
import styles from '@/components/Product.module.css';
import Link from 'next/link';

interface Product {
	id: number;
	shopId: number;
	categoryId: number;
	title: string;
	description: string;
	price: number;
	imageUrl: string;
}

interface ProductCardProps {
	product: Product;
	onAddToCart: (product: Product, quantity: number) => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onAddToCart }) => {
	const [quantity, setQuantity] = React.useState(1);

	const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setQuantity(Math.max(1, parseInt(e.target.value)));
	};

	return (
		<div className={styles.productCard}>
			<Link href={`/products/${product.id}`}>
				<img
					src="https://via.placeholder.com/400x300?text=Product+Image"
					alt={product.title}
					className={styles.productImage}
				/>
			</Link>
			<div className={styles.productDetails}>
				<h2>{product.title}</h2>
				<p>{product.description}</p>
				<p className={styles.price}>${product.price.toFixed(2)}</p>
				<div className={styles.quantitySelector}>
					<label htmlFor={`quantity-${product.id}`}>Quantity:</label>
					<input
						type="number"
						id={`quantity-${product.id}`}
						value={quantity}
						onChange={handleQuantityChange}
						min="1"
					/>
				</div>
				<button
					onClick={() => onAddToCart(product, quantity)}
					className={styles.addToCartButton}
				>
					Add to Cart
				</button>
			</div>
		</div>
	);
};

export default ProductCard;
