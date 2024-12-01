package service

type Store interface {
}

type Service struct {
	store Store
}

func NewService(store Store) (*Service, error) {
	return &Service{store: store}, nil
}
