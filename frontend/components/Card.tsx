import React from 'react';
import styles from './Card.module.css';

interface CardProps {
  title: string;
  description: string;
  likes: number;
  comments: number;
  avatar: string;
  handleLike: () => void;
  isLiked: boolean; // New prop
}

const Card: React.FC<CardProps> = ({ title, description, likes, comments, avatar, handleLike, isLiked }) => {
  return (
    <div className={styles.card}>
      <div className={styles.content}>
        <h2 className={styles.title}>{title}</h2>
        <p className={styles.description}>{description}</p>
      </div>
      <div className={styles.footer}>
        <img src={avatar} alt="avatar" className={styles.avatar} />
        <button onClick={handleLike} className={styles.likeButton}>
          <span className={styles.likes}>
            {isLiked ? '❤️' : '🤍'} {likes}
          </span>
        </button>
        <span className={styles.comments}>💬 {comments}</span>
      </div>
    </div>
  );
};

export default Card;