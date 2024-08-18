'use client'
import React, { useState, useEffect } from 'react';
import Card from '@/components/Card';
import styles from '@/app/page.module.css';
import articles from '@/public/articles.json';
import { useUser } from "@clerk/nextjs";
import { ArrowLeft, ArrowRight } from 'lucide-react';
import { useAuth } from '@clerk/nextjs';

interface Article {
  ID: number;
  Title: string;
  Content: string;
  Likes: number[];
  Creator: {
    UserProfile: {
      AvatarURL: string;
    };
  };
  Views: number;
  Tags: string[] | null;
}

const Page: React.FC = () => {
  const [articles, setArticles] = useState<Article[]>([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [viewStartTime, setViewStartTime] = useState<number | null>(null);
  const { getToken } = useAuth();

  useEffect(() => {
    // Fetch articles from the API with authorization
    const fetchArticles = async () => {
      try {
        const token = await getToken({ template: 'default' });
        const response = await fetch('http://127.0.0.1:8080/api/articles/?page1', {
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        });
        const data = await response.json();
        setArticles(data);
      } catch (error) {
        console.error('Error fetching articles:', error);
      }
    };

    fetchArticles();
  }, []);

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

  if (articles.length === 0) {
    return <div>Loading...</div>;
  }

  const currentArticle = articles[currentIndex];

  return (
    <div className={styles.container}>
      <Card 
        title={currentArticle.Title} 
        description={currentArticle.Content} 
        likes={currentArticle.Likes.length} 
        comments={currentArticle.Views}
        avatar={currentArticle.Creator.UserProfile.AvatarURL} 
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
    </div>
  );
};

export default Page;

// misc
// const Page: React.FC = () => {
//   const [currentIndex, setCurrentIndex] = useState(0);
//   const [viewStartTime, setViewStartTime] = useState<number | null>(null);
//   // const { user } = useUser();

//   useEffect(() => {
//     // Start tracking view time
//     setViewStartTime(Date.now());

//     return () => {
//       if (viewStartTime) {
//         const duration = Date.now() - viewStartTime; 
//         console.log(duration);
//         // Create interface for user_article_interactions
//         // POST user_article_interactions
//         // user_profile_id: user.id,
//         // article_id: articles[currentIndex].id,
//         // interaction_type: "view"
//         // duration: duration,
//       }
//     };
//   }, [currentIndex]);

//   const handleNext = () => {
//     setCurrentIndex((prevIndex) => (prevIndex + 1) % articles.length);
//   };

//   const handleBack = () => {
//     setCurrentIndex((prevIndex) => (prevIndex - 1 + articles.length) % articles.length);
//   };

//   return (
//     <div className={styles.container}>
//       <Card 
//         title={articles[currentIndex].title} 
//         description={articles[currentIndex].description} 
//         likes={articles[currentIndex].likes} 
//         comments={articles[currentIndex].comments} 
//         avatar={articles[currentIndex].avatar} 
//       />
//       <div className={styles.buttons}>
//         <button onClick={handleBack} disabled={currentIndex === 0}>Back</button>
//         <button onClick={handleNext} disabled={currentIndex === articles.length - 1}>Next</button>
//       </div>
//     </div>
//   );
// };


// export default Page;