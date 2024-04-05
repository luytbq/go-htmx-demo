package types

const AUTHENTICATED_USER_KEY = "authenticatedUser"

type LayoutProps struct {
	ShowNavigation bool
	Redirect       bool
	RedirectURL    string
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
	Email    string
	Password string
}

type SignupCredentials struct {
	Email    string
	Password string
}

type SignupError struct {
	EmailError           string
	PasswordError        string
	ConfirmPasswordError string
}
