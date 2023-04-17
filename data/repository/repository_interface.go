package repository

type IRepository interface {
	Create(interface{}) error
	GetAll() ([]interface{}, error)
}
