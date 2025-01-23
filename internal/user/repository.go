package users

import (
	"main/pkg/database"
	"time"
)

// to interact with the database.
func (user *UserRequest) Save() error {
	query := "INSERT INTO users (name, email, password, role, created_at) VALUES (?, ?, ?, ?, ?)"
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.Role, currentTime)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.Id = int(userId)
	return err
}

func getAll() ([]UserResponse, error) {
	var users []UserResponse
	query := "SELECT id, name, email FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func checkEmailPresence(email string) (count int, err error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err = database.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, err
}
