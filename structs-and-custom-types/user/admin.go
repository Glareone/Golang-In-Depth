package user

// Admin
// Embed User field into Admin
type Admin struct {
	email    string
	password string
	// add User field of type "User" anonymously way
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email,
		password,
		User{
			// I can access lowercase variables
			// firstName, lastName, birthDate from here
			// because Admin and User are located in one package
			// lowercase variables are available from other places if all of them are in one package
			FirstName: "Dummy Admin Name",
			LastName:  "Dummy Admin Surname",
			BirthDate: "1/1/1234",
		},
	}
}
