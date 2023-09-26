// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type Users struct {
	UserID         int
	FullName       string
	PhoneNumber    int64
	HashedPassword string
	Salt           string
}

type AddNewUserInput struct {
	FullName       string
	PhoneNumber    int64
	HashedPassword string
	Salt           string
}

type SelectUserByPhoneNumberInput struct {
	PhoneNumber int64
}

type IncrementUserSuccessfulLoginInput struct {
	UserID int
}

type UpdateFullNameInput struct {
	UserID   int
	FullName string
}

type UpdatePhoneNumberInput struct {
	UserID      int
	PhoneNumber int64
}
