import React from 'react';
import styles from './Card.module.css';
import Link from 'next/link';

interface RecommendedProduct {
	id: number;
	title: string;
}

interface CardProps {
	title: string;
	description: string;
	likes: number;
	comments: number;
	avatar: string;
	handleLike: () => void;
	isLiked: boolean;
	recommendedProducts: RecommendedProduct[];
}

const Card: React.FC<CardProps> = ({
	title,
	description,
	likes,
	comments,
	avatar,
	handleLike,
	isLiked,
	recommendedProducts,
}) => {
	return (
		<div className={styles.card}>
			<div className={styles.mainContent}>
				<div className={styles.content}>
					<h2 className={styles.title}>{title}</h2>
					<p className={styles.description}>{description}</p>
				</div>
				<div className={styles.footer}>
					<img
						src={avatar ? avatar : 'https://via.placeholder.com/150'}
						alt="avatar"
						className={styles.avatar}
					/>
					<button onClick={handleLike} className={styles.likeButton}>
						<span className={styles.likes}>
							{isLiked ? '❤️' : '🤍'} {likes}
						</span>
					</button>
					<span className={styles.comments}>💬 {comments}</span>
				</div>
			</div>
			<div className={styles.recommendedProductsWrapper}>
				<span className={styles.recommendedProductsTitle}>
					Recommended Products
				</span>
				<div className={styles.recommendedProducts}>
					{recommendedProducts.map(product => (
						<Link href={`/products/${product.id}`} key={product.id}>
							<span className={styles.productChip}>{product.title}</span>
						</Link>
					))}
				</div>
			</div>
		</div>
	);
};

export default Card;
