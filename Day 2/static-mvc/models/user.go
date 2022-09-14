package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var data = []User{
	{Id: 1, Username: "user1", Password: "pass1", Email: "email@example.com"},
	{Id: 2, Username: "user2", Password: "pass2", Email: "test@example.com"},
}

func GetAllUsers() []User {
	return data
}

func GetUserByID(id int) User {
	for _, user := range data {
		if user.Id == id {
			return user
		}
	}
	return User{}
}

func CreateUser(user User) []User {
	data = append(data, user)
	return data
}

func UpdateUser(id int, user User) []User {
	for i, u := range data {
		if u.Id == id {
			data[i] = user
		}
	}
	return data
}

func DeleteUser(id int) []User {
	var newData []User
	for _, user := range data {
		if user.Id != id {
			newData = append(newData, user)
		}
	}
	return newData
}
