package types

const AUTHENTICATED_USER_KEY = "authenticatedUser"

type LayoutProps struct {
	ShowNavigation bool
}

type AuthenticatedUser struct {
	IsLoggedIn bool
	Email      string
	UserName   string
}

type LoginError struct {
	EmailError    string
	PasswordError string
}

type LoginCredentials struct {
	Email string
}

type SignupCredentials struct {
	Email string
}

type SignupError struct {
	EmailError           string
	PasswordError        string
	ConfirmPasswordError string
}
