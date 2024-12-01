package model

import (
	"time"
)

type Specification struct {
	Id        int64
	Name      string
	Content   string
	GitPath   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
