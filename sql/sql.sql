CREATE DATABASE IF NOT EXISTS api;
USE api;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(55) NOT NULL,
    nick varchar(55) NOT NULL unique,
    email varchar(55) NOT NULL unique,
    password varchar(100) NOT NULL,
    createAt timestamp default current_timestamp()
);

CREATE TABLE followers(
    user_id int NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    primary key(user_id, follower_id)
);