package api

type SubmittedFlag struct {
	Key string
}

type NewTask struct {
	Title   string
	Content string
	Points  int
	Key     string
}

type LoginCredentials struct {
	Email    string
	Password string
}

type CreatedUser struct {
	ID       uint
	Email    string
	Username string
	Points   int
	Admin    bool
}

type TokenResponse struct {
	Access  string
	Refresh string
}
