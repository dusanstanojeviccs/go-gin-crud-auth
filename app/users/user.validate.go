package users

import (
	"go-gin-crud-auth/utils"
)

func validate(validationMessages *[]*utils.ValidationMessage, user *User) error {
	if len(user.Name) < 3 {
		*validationMessages = append(*validationMessages, &utils.ValidationMessage{
			Field:   "name",
			Message: "Please enter a valid Name",
		})
	}

	existingUser, err := UserRepository.findByEmail(user.Email)

	if err != nil {
		return err
	}

	if existingUser != nil && existingUser.Id != user.Id {
		*validationMessages = append(*validationMessages, &utils.ValidationMessage{
			Field:   "email",
			Message: "Please enter a unique Email address",
		})
	}

	return nil
}
