package sep

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//UserAccountRepositoryImpl implements repository
type UserAccountRepositoryImpl struct {
	conn *sql.DB
}

//CheckUser check the username and password and returns User object with all the user details in it
func (usr *UserAccountRepositoryImpl) CheckUser(username string, password string) (User, error) {

	//returs a slice of encrypted password
	bs, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	ps = string(bs)

	row := usr.conn.QueryRow("SELECT * FROM all_users WHERE username = $1 AND password = $2 ", username, ps)

	u := User{}

	err := row.Scan(&u.Username, &u.Password, &u.FirstName, &u.LastName, &u.ProfilePicture, &u.Email)
	if err != nil {
		return u, err
	}

	return u, nil
}

//AddUser add user to the database
func (usr *UserAccountRepositoryImpl) AddUser(user User) error {

	bs, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	ps = string(bs)

	_, err := usr.conn.Exec("INSERT INTO all_users (username,email,password,profile_picture,firstname,lastname) values($1, $2, $3, $4, $5, $6)",
		user.Username, user.Email, ps, user.ProfilePicture, user.FirstName, user.LastName)
	if err != nil {
		return errors.New("Unable to add User")
	}

	return nil
}
