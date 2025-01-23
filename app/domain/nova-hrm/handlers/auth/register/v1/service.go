package register

func createUser(
	username string,
	encPassword string,
	email string,
) (string, error) {
	id, err := createUserWithPassword(
		username,
		encPassword,
		email,
	)

	return id, err
}
