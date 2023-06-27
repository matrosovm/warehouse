package domain

// ProductTable ...
type ProductTable struct {
	ID   uint64    `db:"id"`
	Name string    `db:"name"`
	Size SizesType `db:"size"`
}

// SizesType ...
type SizesType struct {
	Width  uint64 `db:"width"`
	Height uint64 `db:"height"`
	Length uint64 `db:"length"`
}
