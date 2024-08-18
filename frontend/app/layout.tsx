import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';
import {
	ClerkProvider,
	SignInButton,
	SignedIn,
	SignedOut,
	UserButton,
} from '@clerk/nextjs';
import styles from './page.module.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
	title: 'Tiktok Shop Products Recommender',
	description:
		'Conceptual Golang + Python Grpc shop products recommendation engine',
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<ClerkProvider
			publishableKey={process.env.NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY}
		>
			<html lang="en">
				<body className={inter.className}>
					<div className={styles.authContainer}>
						<SignedOut>
							<SignInButton />
						</SignedOut>
						<SignedIn>
							<UserButton />
						</SignedIn>
					</div>
					<main>
						<div className={styles.container}>
							{children}
						</div>
					</main>
				</body>
			</html>
		</ClerkProvider>
	);
}
