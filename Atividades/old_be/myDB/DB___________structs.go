package myDB

type Model interface{}

type Repository interface {
	Get() []Model
	Post() bool
	Put() bool
	Delete() string
}
