package users

import (
	"fmt"

	"github.com/markgardner10/bookstore_users-api/utils/errors"
	"github.com/markgardner10/bookstore_users-api/utils/date_utils"
)

var (
	// mock the database for now using a map
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if result, ok := usersDB[user.ID]; ok {
		user.ID = result.ID
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email
		user.DateCreated = result.DateCreated
		return nil
	}
	return errors.NewNotFoundError(fmt.Sprintf("user %d nto found", user.ID))
}

func (user *User) Save() *errors.RestErr {
	if _, ok := usersDB[user.ID]; !ok {
		for userInDB := range usersDB {
			fmt.Println(usersDB[userInDB].Email, user.Email)
			if usersDB[userInDB].Email == user.Email {
				return errors.NewBadRequestError(fmt.Sprintf("email address %s already registered", user.Email))
			}
		}
		
		user.DateCreated = date_utils.GetNowString()
		usersDB[user.ID] = user
		return nil
	}
	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
}
