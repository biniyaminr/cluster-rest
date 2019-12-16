package sep

//UserAccountRepository includes action for Account Managment
type UserAccountRepository interface {
	CheckUser(username string, password string) (User, error)
	AddUser(user User) error
}
