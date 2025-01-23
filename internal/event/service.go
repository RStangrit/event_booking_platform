package event

// to describe the business logic of working with users.
func IsTitleUnique(title string) (bool, error) {
	count, err := checkTitlePresence(title)

	if err != nil {
		return false, err
	}
	return count == 0, nil
}
