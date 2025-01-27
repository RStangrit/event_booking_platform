package users

// to describe the business logic of working with users.
func IsEmailUnique(email string) (bool, error) {
	count, err := checkEmailPresence(email)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
