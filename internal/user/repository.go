package users

import "main/pkg/database"

// to interact with the database.
func (user *User) Save() error {
	query := "INSERT INTO users (name, email, password, role, created_at) VALUES (?, ?, ?, ?, ?)"
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.Role, "CURRENT_TIMESTAMP")

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.Id = int(userId)
	return err
}

func checkEmail(email string) (count int, err error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err = database.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, err
}
