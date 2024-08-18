CREATE TABLE IF NOT EXISTS user_profile (
  model TEXT NOT NULL,
  user_i_d TEXT NOT NULL,
  username TEXT NOT NULL,
  bio TEXT NOT NULL,
  avatar_u_r_l TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS creator (
  model TEXT NOT NULL,
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS article (
  model TEXT NOT NULL,
  creator_i_d INTEGER NOT NULL,
  creator TEXT NOT NULL,
  title TEXT NOT NULL,
  content_u_r_l TEXT NOT NULL,
  likes INTEGER NOT NULL,
  views INTEGER NOT NULL,
  tags TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
  model TEXT NOT NULL,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS shop (
  model TEXT NOT NULL,
  creator_i_d INTEGER NOT NULL,
  creator TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS category (
  model TEXT NOT NULL,
  name TEXT NOT NULL,
  parent_i_d TEXT NOT NULL,
  parent TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS product (
  model TEXT NOT NULL,
  shop_i_d INTEGER NOT NULL,
  shop TEXT NOT NULL,
  category_i_d INTEGER NOT NULL,
  category TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  price REAL NOT NULL,
  image_u_r_l TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS comment (
  model TEXT NOT NULL,
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL,
  article_i_d INTEGER NOT NULL,
  article TEXT NOT NULL,
  content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS article_like (
  user_profile_i_d INTEGER NOT NULL,
  article_i_d INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS article_embedding (
  article_i_d INTEGER NOT NULL,
  article TEXT NOT NULL,
  embedding vector NOT NULL
);

CREATE TABLE IF NOT EXISTS product_embedding (
  product_i_d INTEGER NOT NULL,
  product TEXT NOT NULL,
  embedding vector NOT NULL
);

CREATE TABLE IF NOT EXISTS user_article_interaction (
  model TEXT NOT NULL,
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL,
  article_i_d INTEGER NOT NULL,
  article TEXT NOT NULL,
  interaction_type TEXT NOT NULL,
  duration INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS user_product_interaction (
  model TEXT NOT NULL,
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL,
  product_i_d INTEGER NOT NULL,
  product TEXT NOT NULL,
  interaction_type TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_article_recommendation (
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL,
  recommended_articles TEXT NOT NULL,
  last_updated TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_product_recommendation (
  user_profile_i_d INTEGER NOT NULL,
  user_profile TEXT NOT NULL,
  recommended_products TEXT NOT NULL,
  last_updated TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS kafka_event (
  model TEXT NOT NULL,
  event_type TEXT NOT NULL,
  payload TEXT NOT NULL,
  processed BOOLEAN NOT NULL
);

