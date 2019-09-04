package postgre

type DB struct {
	url string
	conn float64
}

func New(url string) *DB {
	return &DB{url: url, conn: 1.5}
}

func (d *DB) Get(id string) (string, error) {
	return "postgre get", nil
}

func (d *DB) Delete(id string) error {
	return nil
}
