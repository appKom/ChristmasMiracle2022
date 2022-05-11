package api

type SubmittedFlag struct {
	Key string `json:"key"`
}

type NewTask struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Points      int    `json:"points"`
	Key         string `json:"key"`
	ReleaseDate string `json:"release_date"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatedUser struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Points      int    `json:"points"`
	Admin       bool   `json:"admin"`
	SolvedTasks []uint `json:"solved_tasks"`
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshToken struct {
	Refresh string `json:"refresh_token"`
}
