package model

// Writer ...
type Writer struct {
	ID       int        `db:"id"`
	Name     string     `db:"name"`
	Articles []*Article `db:"-"`
}
