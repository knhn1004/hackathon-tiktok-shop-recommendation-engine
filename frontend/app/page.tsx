'use client';
import React, { useState, useEffect, useRef, useCallback } from 'react';
import Link from 'next/link';
import Card from '@/components/Card';
import styles from '@/app/page.module.css';
import { ArrowLeft, ArrowRight } from 'lucide-react';
import { useAuth } from '@clerk/nextjs';

interface Article {
	id: number;
	title: string;
	content: string;
	likes: number[];
	creator: {
		userProfile: {
			avatarUrl: string;
		};
	};
	views: number;
	tags: string[] | null;
}

interface RecommendedProduct {
	id: number;
	title: string;
}

const Page: React.FC = () => {
	const [articles, setArticles] = useState<Article[]>([]);
	const [currentIndex, setCurrentIndex] = useState(0);
	const { isSignedIn, getToken } = useAuth();
	const viewStartTimeRef = useRef<number | null>(null);
	const isNavigatingRef = useRef(false);
	const [recommendedProducts, setRecommendedProducts] = useState<
		RecommendedProduct[]
	>([]);

	useEffect(() => {
		const fetchArticles = async () => {
			if (isSignedIn) {
				try {
					const token = await getToken({ template: 'default' });
					const response = await fetch(
						`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/articles/?page1`,
						{
							headers: {
								Authorization: `Bearer ${token}`,
								'Content-Type': 'application/json',
							},
						}
					);
					const data = await response.json();
					setArticles(data);
				} catch (error) {
					console.error('Error fetching articles:', error);
				}
			}
		};

		fetchArticles();
	}, [isSignedIn, getToken]);

	const sendInteraction = useCallback(
		async (duration: number) => {
			if (articles.length === 0) return;

			try {
				const token = await getToken({ template: 'default' });
				await fetch(
					`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/interactions/articles`,
					{
						method: 'POST',
						headers: {
							Authorization: `Bearer ${token}`,
							'Content-Type': 'application/json',
						},
						body: JSON.stringify({
							articleId: articles[currentIndex].id,
							interactionType: 'view',
							duration,
						}),
					}
				);
			} catch (error) {
				console.error('Error sending interaction:', error);
			}
		},
		[getToken, articles, currentIndex]
	);

	useEffect(() => {
		viewStartTimeRef.current = Date.now();

		return () => {
			if (viewStartTimeRef.current && isNavigatingRef.current) {
				const duration = Date.now() - viewStartTimeRef.current;
				sendInteraction(duration);
				isNavigatingRef.current = false;
			}
		};
	}, [currentIndex, sendInteraction]);

	const handleNavigation = (newIndex: number) => {
		isNavigatingRef.current = true;
		setCurrentIndex(newIndex);
	};

	const handleNext = () => {
		handleNavigation((currentIndex + 1) % articles.length);
	};

	const handleBack = () => {
		handleNavigation((currentIndex - 1 + articles.length) % articles.length);
	};

	const handleLike = async () => {
		try {
			const token = await getToken({ template: 'default' });
			const currentArticle = articles[currentIndex];
			const isLiked = currentArticle.likes.includes(1); // Assuming 1 represents the current user's like

			const response = await fetch(
				`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/articles/${currentArticle.id}/like`,
				{
					method: isLiked ? 'DELETE' : 'POST',
					headers: {
						Authorization: `Bearer ${token}`,
						'Content-Type': 'application/json',
					},
				}
			);
			if (isLiked) {
				const likeInteraction = await fetch(
					`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/interactions/articles`,
					{
						method: 'POST',
						headers: {
							Authorization: `Bearer ${token}`,
							'Content-Type': 'application/json',
						},
						body: JSON.stringify({
							productId: articles[currentIndex].id,
							interactionType: 'like',
						}),
					}
				);
			}

			if (response.ok) {
				setArticles(prevArticles => {
					const updatedArticles = prevArticles.map((article, index) => {
						if (index === currentIndex) {
							const updatedLikes = isLiked
								? article.likes.filter(like => like !== 1)
								: [...article.likes, 1];
							return { ...article, likes: updatedLikes };
						}
						return article;
					});
					return updatedArticles;
				});
			} else {
				throw new Error('Failed to update like status');
			}
		} catch (error) {
			console.error('Error liking article:', error);
		}
	};

	useEffect(() => {
		const fetchRecommendedProducts = async () => {
			if (isSignedIn && articles.length > 0) {
				try {
					const token = await getToken({ template: 'default' });
					const response = await fetch(
						`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/articles/${articles[currentIndex].id}/recommendations`,
						{
							headers: {
								Authorization: `Bearer ${token}`,
								'Content-Type': 'application/json',
							},
						}
					);
					const data = await response.json();
					setRecommendedProducts(data);
				} catch (error) {
					console.error('Error fetching recommended products:', error);
				}
			}
		};

		fetchRecommendedProducts();
	}, [isSignedIn, getToken, articles, currentIndex]);

	if (!isSignedIn) {
		return (
			<Card
				title="Sign in"
				description="Sign in to view articles"
				likes={0}
				comments={0}
				avatar=""
				handleLike={() => {}}
			/>
		);
	}

	if (articles.length === 0) {
		return <div>Loading...</div>;
	}

	const currentArticle = articles[currentIndex];

	return (
		<>
			<Card
				title={currentArticle.title}
				description={currentArticle.content}
				likes={currentArticle.likes.length}
				comments={currentArticle.views}
				avatar={currentArticle.creator.userProfile.avatarUrl}
				handleLike={handleLike}
				isLiked={currentArticle.likes.includes(1)}
				recommendedProducts={recommendedProducts}
			/>
			<div className={styles.buttons}>
				<button
					className={styles.arrowButton}
					onClick={handleBack}
					disabled={currentIndex === 0}
				>
					<ArrowLeft className={styles.arrowIcon} />
				</button>
				<button
					className={styles.arrowButton}
					onClick={handleNext}
					disabled={currentIndex === articles.length - 1}
				>
					<ArrowRight className={styles.arrowIcon} />
				</button>
			</div>
		</>
	);
};

export default Page;
