"use client";
import React, { useState, useEffect, useRef } from "react";
import Card from "@/components/Card";
import styles from "@/app/page.module.css";
import { ArrowLeft, ArrowRight } from "lucide-react";
import { useAuth } from "@clerk/nextjs";

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

const Page: React.FC = () => {
  const [articles, setArticles] = useState<Article[]>([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const { isSignedIn, getToken } = useAuth();
  const viewStartTimeRef = useRef<number | null>(null);

  useEffect(() => {
    // Fetch articles from the API with authorization
    const fetchArticles = async () => {
      if (isSignedIn) {
        try {
          const token = await getToken({ template: "default" });
          const response = await fetch(
            "http://127.0.0.1:8080/api/articles/?page1",
            {
              headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
              },
            }
          );
          const data = await response.json();
          setArticles(data);
        } catch (error) {
          console.error("Error fetching articles:", error);
        }
      }
    };

    fetchArticles();
  }, [isSignedIn, getToken]);

  useEffect(() => {
    const sendInteraction = async (duration: number) => {
      try {
        const token = await getToken({ template: "default" });
        const response = await fetch(
          "http://127.0.0.1:8080/api/interactions/articles",
          {
            method: "POST",
            headers: {
              Authorization: `Bearer ${token}`,
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              articleId: articles[currentIndex].id,
              interactionType: "view",
              duration: duration,
            }),
          }
        );
        console.log(articles[currentIndex]);
        if (!response.ok) {
          throw new Error("Failed to send interaction");
        }
      } catch (error) {
        console.error("Error sending interaction:", error);
      }
    };

    // Start tracking view time
    viewStartTimeRef.current = Date.now();

    // Clean up function
    return () => {
      if (viewStartTimeRef.current && articles.length > 0) {
        const duration = Date.now() - viewStartTimeRef.current;
        sendInteraction(duration);
      }
    };
  }, [currentIndex, getToken, articles]);

  const handleNext = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % articles.length);
  };

  const handleBack = () => {
    setCurrentIndex(
      (prevIndex) => (prevIndex - 1 + articles.length) % articles.length
    );
  };

  // const handleLike = async () => {
  //   try {
  //     const token = await getToken({ template: "default" });
  //     await fetch(
  //       `http://127.0.0.1:8080/api/articles/${articles[currentIndex].id}/like`,
  //       {
  //         method: "POST",
  //         headers: {
  //           Authorization: `Bearer ${token}`,
  //           "Content-Type": "application/json",
  //         },
  //       }
  //     );
  //     setArticles((prevArticles) => {
  //       const updatedArticles = [...prevArticles];
  //       updatedArticles[currentIndex].likes.push(1);
  //       return updatedArticles;
  //     });
  //   } catch (error) {
  //     console.error("Error liking article:", error);
  //   }
  // };

  // const handleUnlike = async () => {
  //   try {
  //     const token = await getToken({ template: "default" });
  //     await fetch(
  //       `http://127.0.0.1:8080/api/articles/${articles[currentIndex].id}/like`,
  //       {
  //         method: "DELETE",
  //         headers: {
  //           Authorization: `Bearer ${token}`,
  //           "Content-Type": "application/json",
  //         },
  //       }
  //     );
  //     // Update the likes locally
  //     setArticles((prevArticles) => {
  //       const updatedArticles = [...prevArticles];
  //       updatedArticles[currentIndex].likes.pop(); // Assuming the last like is removed
  //       return updatedArticles;
  //     });
  //   } catch (error) {
  //     console.error("Error unliking article:", error);
  //   }
  // };

  if (!isSignedIn) {
    return (
      <div className={styles.container}>
        <Card
          title="Sign in"
          description="Sign in to view articles"
          likes={0}
          comments={0}
          avatar=""
          onLike={() => {}}
          onUnlike={() => {}}
        />
      </div>
    );
  }

  if (articles.length === 0) {
    return <div>Loading...</div>;
  }

  const currentArticle = articles[currentIndex];

  return (
    <div className={styles.container}>
      <Card
        title={currentArticle.title}
        description={currentArticle.content}
        likes={currentArticle.likes.length}
        comments={currentArticle.views}
        avatar={currentArticle.creator.userProfile.avatarUrl}
        // onLike={handleLike}
        // onUnlike={handleUnlike}
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