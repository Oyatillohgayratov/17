package methods

import "homework/config"

func Age(users []config.User) []config.User {
	for i := 0; i < len(users); i++ {
        if users[i].Id == 3{
			users[i].Age = 50
		}
    }
    return users
}