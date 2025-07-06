CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);

CREATE TABLE subscriptions(
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE publication(
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(120) NOT NULL,
    content VARCHAR(300),
    author_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
    likes INT DEFAULT 0,
    createdAt TIMESTAMP DEFAULT current_timestamp
) ENGINE=INNODB;

INSERT INTO publication(title, content, author_id) VALUES
("pub from user 1", "description...", 7),
("pub from user 2", "description...", 8),
("pub from user 3", "description...", 9),
("pub from user 4", "description...", 10);

select p.* from publication p inner join users u on u.id = p.author_id inner join subscriptions s on p.author_id = s.user_id where u.id = 10 or s.follower_id = 10;