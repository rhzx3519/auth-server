package user

import (
	"database/sql"
	"fmt"
	"github.com/rhzx3519/auth-server/domain"
	"github.com/rhzx3519/auth-server/persistance/mysql"
)

func FindUser(email, password string) (domain.User, error) {
	row := mysql.DB.QueryRow("SELECT * FROM users WHERE email = ? and password = ?", email, password)
	// Loop through rows, using Scan to assign column data to struct fields.
	var user domain.User
	if err := row.Scan(&user.ID, &user.No, &user.Email, &user.Nickname,
		&user.Firstname, &user.Lastname, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("No such user with email: %s, password: %s", email, password)
		}
		return user, fmt.Errorf("getUserbyEmailAndPassword error %s: %v", email, err)
	}
	return user, nil
}
