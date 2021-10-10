package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// UserRepo struct to create a repository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo - create a new user's repository
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (userRepo UserRepo) Create(user models.User) (uint64, error) {
	statement, error := userRepo.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	ID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return uint64(ID), nil
}

func (userRepo UserRepo) FindAll() ([]models.User, error) {
	rows, error := userRepo.db.Query("select * from users")
	if error != nil {
		return []models.User{}, error
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		error := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreateAt)
		if error != nil {
			return []models.User{}, error
		}
		users = append(users, user)
	}
	return users, nil
}

// Find - find users by name or nick
func (UserRepo UserRepo) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%
	rows, error := UserRepo.db.Query(
		"select id, name, nick, email, createAt from users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User

		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

func (UserRepo UserRepo) FindById(ID uint64) (models.User, error) {
	rows, err := UserRepo.db.Query(
		"SELECT id, name, nick, email, createAt FROM users WHERE id = ? ",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (UserRepo UserRepo) Update(ID uint64, data models.User) error {
	statement, err := UserRepo.db.Prepare(
		"UPDATE users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(data.Name, data.Nick, data.Email, ID); err != nil {
		return err
	}
	return nil
}

func (UserRepo UserRepo) Delete(ID uint64) error {
	statement, err := UserRepo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (UserRepo UserRepo) FindByEmail(email string) (models.User, error) {
	row, err := UserRepo.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()
	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// FollowUser - create a new row in followers table
func (UserRepo UserRepo) FollowUser(follower_id uint64, user_id uint64) error {
	statement, err := UserRepo.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user_id, follower_id); err != nil {
		return err
	}
	return nil
}

// UnFollowUser - remove a row in followers table
func (UserRepo UserRepo) UnFollowUser(follower_id uint64, user_id uint64) error {
	statement, err := UserRepo.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user_id, follower_id); err != nil {
		return err
	}
	return nil
}

// GetFollowers - Get all followers from an user
func (UserRepo UserRepo) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := UserRepo.db.Query(`
	   select u.id, u.name, u.nick, u.email, u.createAt
	   FROM users u INNER JOIN followers f ON (f.follower_id = u.id)
	   WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return []models.User{}, nil
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return []models.User{}, nil
		}
		users = append(users, user)
	}
	return users, nil
}

// GetFollowing - Get all users followed by user
func (UserRepo UserRepo) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := UserRepo.db.Query(`
	   select u.id, u.name, u.nick, u.email, u.createAt
	   FROM users u INNER JOIN followers f ON (f.user_id = u.id)
	   WHERE f.follower_id = ?
	`, userID)
	if err != nil {
		return []models.User{}, nil
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return []models.User{}, nil
		}
		users = append(users, user)
	}
	return users, nil
}
