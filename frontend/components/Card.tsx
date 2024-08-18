import React from 'react';
import styles from './Card.module.css';

interface CardProps {
  title: string;
  description: string;
  likes: number;
  comments: number;
  avatar: string;
}

const Card: React.FC<CardProps> = ({ title, description, likes, comments, avatar }) => {
  return (
    <div className={styles.card}>
      <div className={styles.content}>
        <h2 className={styles.title}>{title}</h2>
        <p className={styles.description}>{description}</p>
      </div>
      <div className={styles.footer}>
        <img src={avatar} alt="avatar" className={styles.avatar} />
        <span className={styles.likes}>❤️ {likes}</span>
        <span className={styles.comments}>💬 {comments}</span>
      </div>
    </div>
  );
};

export default Card;
