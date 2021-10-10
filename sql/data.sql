INSERT INTO users (name, nick, email, password)
VALUES
("User1", "user1", "user1@gmail.com", "$2a$10$.zJg9j0rGtEKo3VqIZKIbu8SZweHoRhF.59bN5u/Wx2.ylxaOWA76"),
("User2", "user2", "user2@gmail.com", "$2a$10$.zJg9j0rGtEKo3VqIZKIbu8SZweHoRhF.59bN5u/Wx2.ylxaOWA76"),
("User3", "user3", "user3@gmail.com", "$2a$10$.zJg9j0rGtEKo3VqIZKIbu8SZweHoRhF.59bN5u/Wx2.ylxaOWA76"),
("User4", "user4", "user4@gmail.com", "$2a$10$.zJg9j0rGtEKo3VqIZKIbu8SZweHoRhF.59bN5u/Wx2.ylxaOWA76")

INSERT INTO followers (user_id, follower_id)
VALUES
(1,2),
(3,2),
(1,3)