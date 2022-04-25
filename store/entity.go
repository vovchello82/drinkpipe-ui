package store

type Entity interface {
	GetId() string
	ConvertToJson() ([]byte, error)
}
