package model

import (
	"github.com/google/uuid"
	"github.com/nstoker/gorocktrack/app"
	"github.com/sirupsen/logrus"
)

// Users structure for normal user access
type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

// Admin returns true if the user is an administrator
func (u *User) Admin() (bool, error) {
	sql := `SELECT admin FROM users WHERE id=$1`
	isAdmin := false
	row := app.DB.QueryRow(sql, u.ID)

	err := row.Scan(&isAdmin)
	if err != nil {
		return false, err
	}

	return isAdmin, err
}

// Find finds a user based on the ID
// func(u*User)Find error {
// 	u=User{}
// 	return fmt.Errorf("not implemented")
// }

// ShowAllUsers TODO: Needs a check for admin authentication
func ShowAllUsers() ([]User, error) {
	rows, err := app.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		logrus.Errorf("ShowAllUsers failed %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			logrus.Errorf("showAllUsers scan error '%+v'", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FindByEmail(email string) (*User, error) {
	sql := `SELECT id, name, email FROM users WHERE email=$1;`
	u := &User{}

	row := app.DB.QueryRow(sql, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email)

	// Possibly not an ideomatic Golang expression, but :shrugs
	return u, err
}

// AddNewUser to the system
func AddNewUser(name, email, password string) (*User, error) {
	return AddAdminUser(name, email, password, false)
}

// AddAdminUser to the system
func AddAdminUser(name, email, password string, admin bool) (*User, error) {
	sql := `INSERT INTO users(name,email,password,admin) VALUES ($1, $2, $3, $4);`

	user := &User{}
	_, err := app.DB.Exec(sql, name, email, password, admin)
	if err != nil {
		logrus.Warnf("returning error %v", err)
		return user, err
	}

	user, err = FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
