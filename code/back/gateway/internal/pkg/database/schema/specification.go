package schema

import "time"

type Specification struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	GitPath   string    `db:"git_path"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
