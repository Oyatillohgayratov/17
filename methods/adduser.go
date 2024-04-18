package methods

import "homework/config"


func Add(users []config.User) []config.User {
	var user config.User
	user.Id = 6
	user.Name = "Djin-vu"
	user.Age = 22
	user.Position = "hunter"
	
	users = append(users, user)
	return users 
}


