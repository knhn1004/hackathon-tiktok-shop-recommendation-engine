'use client'
import React, { useState } from 'react';
import Card from '../components/Card';
import styles from './page.module.css';
import articles from '../public/articles.json';

const Page: React.FC = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const handleNext = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % articles.length);
  };

  const handleBack = () => {
    setCurrentIndex((prevIndex) => (prevIndex - 1 + articles.length) % articles.length);
  };

  return (
    <div className={styles.container}>
      <Card 
        title={articles[currentIndex].title} 
        description={articles[currentIndex].description} 
      />
      <div className={styles.buttons}>
        <button onClick={handleBack} disabled={currentIndex === 0}>Back</button>
        <button onClick={handleNext} disabled={currentIndex === articles.length - 1}>Next</button>
      </div>
    </div>
  );
};

export default Page;
