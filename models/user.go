package models

type User struct {
	Id       int
	UserName string
}

// variable blocks
var (
	users      []*User
	nextUserId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(newUser User) (User, error) {
	// assign the Id based on the nextUserId
	newUser.Id = nextUserId

	// increment the next user ID
	nextUserId++

	// append the address of the new user to the users slice
	users = append(users, &newUser)

	// return the result with nil error
	return newUser, nil
}
