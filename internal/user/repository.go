package users

import (
	"errors"
	"main/pkg/database"
	"main/pkg/util"
)

var dbAccessMode = util.GetEnvVariable("DB_ACCESS_MODE")

func (user *UserRequest) Save() error {
	switch dbAccessMode {
	case "gorm":
		unique, _ := IsEmailUnique(user.Email)
		if !unique {
			return errors.New("user email already exists")
		}
		newUser := user.ToUser()
		result := database.GORM.Create(&newUser)
		return result.Error
	case "raw_sql":
		unique, _ := IsEmailUnique(user.Email)
		if !unique {
			return errors.New("user email already exists")
		}

		query := "INSERT INTO users (name, email, password, role, created_at) VALUES (?, ?, ?, ?, ?)"
		statement, err := database.DB.Prepare(query)
		if err != nil {
			return err
		}

		defer statement.Close()

		currentTime := util.GetCurrentTime()

		result, err := statement.Exec(user.Name, user.Email, user.Password, user.Role, currentTime)

		if err != nil {
			return err
		}

		userId, err := result.LastInsertId()
		user.Id = int(userId)
		return err
	default:
		return errors.ErrUnsupported
	}
}

func getAllUsers() (any, error) {
	switch dbAccessMode {
	case "gorm":
		var users []UserResponse
		result := database.GORM.Select("id", "name", "email", "role", "created_at").Find(&users)
		return result, nil
	case "raw_sql":
		var users []UserResponse
		query := "SELECT id, name, email, role, created_at FROM users"
		rows, err := database.DB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var user UserResponse
			err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
		return users, nil
	default:
		return nil, errors.ErrUnsupported
	}

}

func getOneUser(userId int64) (*UserResponse, error) {
	query := `SELECT id, name, email, role, created_at FROM users WHERE id =?`
	row := database.DB.QueryRow(query, userId)

	var user UserResponse
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (user UserResponse) Update() error {
	query := `
	UPDATE users
	SET name = ?, email = ?, updated_at = ?
	WHERE id = ?
	`
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	currentTime := util.GetCurrentTime()

	_, err = statement.Exec(user.Name, user.Email, currentTime, user.Id)
	return err
}

func (user UserResponse) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	statement, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Id)
	return err
}

func checkEmailPresence(email string) (count int, err error) {
	switch dbAccessMode {
	case "gorm":
		result := database.GORM.Where("email = ?", email).Find(&email)
		if result.Error != nil {
			return 0, err
		}
		return count, err
	case "raw_sql":
		query := "SELECT COUNT(*) FROM users WHERE email = ?"
		err = database.DB.QueryRow(query, email).Scan(&count)
		if err != nil {
			return 0, err
		}
		return count, err
	default:
		return 0, errors.New("invalid db access mode")
	}
}
