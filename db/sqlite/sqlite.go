package sqlite

type DB struct {
	path string
	conn int
}

func New(path string) *DB {
	return &DB{path: path, conn: 1}
}

func (d *DB) Get(id string) (string, error) {
	return "sqlite get", nil
}

func (d *DB) Delete(id string) error {
	return nil
}
