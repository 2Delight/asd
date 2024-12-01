package model

type PushResult struct {
	CommitHash string
	IsSuccess  bool
}

type MLDevResult struct {
	Artifacts []string
	IsSuccess bool
}
