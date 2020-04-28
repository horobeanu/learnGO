package entity

type user struct {
	ID int // starts with capital letter => public
	// 0 default value for int
	password string // starts with low letter => private
	// empty string default value for strings
	FirstName string
}

// Bind a function to a type
// by using (someName bindingType) before function params
// (uc User) <- this transforms a function into a method
func (uc user) GetPassword() string {
	return uc.password
}

// Constructor - returns a pointer to an empty "user" type
func NewUser() *user {
	return &user{}
}

// private method
func (uc user) cantTouchThis() bool {
	return true
}
