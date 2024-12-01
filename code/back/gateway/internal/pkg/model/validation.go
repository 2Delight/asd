package model

type ValidationData struct {
	IsValid bool
	Errors  []*Error
	Hints   []*Hint
}

type Error struct {
	Code    string
	Message string
}

type Hint struct {
	Message string
}
