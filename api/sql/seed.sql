INSERT INTO users (name, nick, email, password) VALUES
("User 1", "user_1", "user@gmail.com", "$2a$12$N./1cjv5DdGdLpLJXKmnvurDNV90NKZcZVtbyTVN2/3PRJ4HWSzzi"),
("User 2", "user_2", "user2@gmail.com", "$2a$12$N./1cjv5DdGdLpLJXKmnvurDNV90NKZcZVtbyTVN2/3PRJ4HWSzzi"),
("User 3", "user_3", "user3@gmail.com", "$2a$12$N./1cjv5DdGdLpLJXKmnvurDNV90NKZcZVtbyTVN2/3PRJ4HWSzzi"),
("User 4", "user_4", "user4@gmail.com", "$2a$12$N./1cjv5DdGdLpLJXKmnvurDNV90NKZcZVtbyTVN2/3PRJ4HWSzzi");

INSERT INTO subscriptions(user_id, follower_id)
VALUES
(5, 6),
(8, 10),
(10, 9);