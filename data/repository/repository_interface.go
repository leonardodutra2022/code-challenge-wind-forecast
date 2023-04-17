package repository

type IRepository interface {
	Create(interface{}) error
	GetAll() ([]interface{}, error)
	GetOne() (interface{}, error)
	Updates(obj interface{}) (interface{}, error)
	GetAlertByStatus(status bool) ([]interface{}, error)
}
