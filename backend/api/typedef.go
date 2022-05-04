package api

type SubmittedFlag struct {
	Key string
}

type NewTask struct {
	Title   string
	Content string
	Points int
	Key    string
}