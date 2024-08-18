'use client'
import React, { useState, useEffect } from 'react';
import Card from '@/components/Card';
import styles from '@/app/page.module.css';
import articles from '@/public/articles.json';
import { useUser } from "@clerk/nextjs";

const Page: React.FC = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [viewStartTime, setViewStartTime] = useState<number | null>(null);
  // const { user } = useUser();

  useEffect(() => {
    // Start tracking view time
    setViewStartTime(Date.now());

    return () => {
      if (viewStartTime) {
        const duration = Date.now() - viewStartTime; 
        console.log(duration);
        // Create interface for user_article_interactions
        // POST user_article_interactions
        // user_profile_id: user.id,
        // article_id: articles[currentIndex].id,
        // interaction_type: "view"
        // duration: duration,
      }
    };
  }, [currentIndex]);

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