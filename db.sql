/*
    first execute the following statements if running for the first time:
        create database rss_manager_db
        \c rss_manager_db
*/

CREATE TABLE all_users(
    username VARCHAR(255) PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE user_information(
    username varchar(255) primary key references all_users(username) on update cascade on delete cascade,
    first_name varchar(255),
    last_name varchar(255),
    description text,
    profile_picture varchar(255)
);

CREATE TABLE rss_feeds (
    feed_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL,
    user_count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE subscribed_feeds(
    username VARCHAR(255) REFERENCES all_users(username) ON UPDATE CASCADE ON DELETE CASCADE,
    feed_id INTEGER REFERENCES rss_feeds(feed_id),
    category VARCHAR(255),
    CONSTRAINT subscribed_feed_pky PRIMARY KEY (username, feed_id)
);

CREATE TABLE favourite_feeds (
    username VARCHAR(255) REFERENCES all_users(username) ON UPDATE CASCADE ON DELETE CASCADE,
    feed_id INTEGER REFERENCES rss_feeds(feed_id),
    CONSTRAINT favourite_feed_pky PRIMARY KEY (username, feed_id)
);

CREATE TABLE articles (
    article_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    source_rss INTEGER REFERENCES rss_feeds(feed_id),
    title VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL
);

CREATE TABLE saved_articles(
    username VARCHAR(255) REFERENCES all_users(username) ON UPDATE CASCADE ON DELETE CASCADE,
    article_id INTEGER REFERENCES articles(article_id),
    date DATE,
    CONSTRAINT saved_article_pky PRIMARY KEY(username, article_id)
);

CREATE TABLE recommended_articles(
    article_id INTEGER REFERENCES articles(article_id),
    sender_username VARCHAR(255) REFERENCES all_users(username),
    receiver_username VARCHAR(255) REFERENCES all_users(username),
    message TEXT,
    CONSTRAINT recommended_article_pky PRIMARY KEY(article_id, sender_username, receiver_username)
);

CREATE TABLE public_entries(
    entry_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    article_id INTEGER REFERENCES articles(article_id),
    shared_by VARCHAR(255) REFERENCES all_users(username),
    shared_date date,
    likes_count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE likes(
    username VARCHAR(255) REFERENCES all_users(username) ON UPDATE CASCADE ON DELETE CASCADE,
    entry_id INTEGER REFERENCES public_entries(entry_id),
    liked_date DATE
);

CREATE TABLE comments(
    comment_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) REFERENCES all_users(username) ON UPDATE CASCADE ON DELETE CASCADE,
    entry_id INTEGER REFERENCES public_entries(entry_id),
    comment TEXT,
    commented_date DATE
);


