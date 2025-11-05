package main

type User struct {
	name string
}

func (u User) String() string {
	return u.name
}

type Admin struct {
	*User
	Perms map[string]bool
}

func (a Admin) String () string {
	if a.User != nil {
		return a.User.String()
	}

	return "<nil>"
}

