package model

import "github.com/nstoker/gorocktrack/config"

// Users structure for normal user access
type User struct {
	User_ID string
	Name    string
	Email   string
}

// ShowAllUsers TODO: Needs a check for admin authentication
func ShowAllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.User_ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
