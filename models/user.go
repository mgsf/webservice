package models

// User to store user data
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers - returns the users stored in the models layer
func GetUsers() []*User {
	return users
}

//AddUser - Adds a user to the users collection
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
