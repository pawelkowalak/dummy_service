package dummy_service

type Service struct {
	db DB
}

type DB interface {
	Get(string) (string, error)
	Delete(string) error
}

func New(db DB) *Service {
	return &Service{db: db}
}

func (s *Service) DoWork() {
	// Do some work here.
}