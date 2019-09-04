package main

import (
	"flag"
	"log"

	"github.com/pawelkowalak/dummy_service"
	"github.com/pawelkowalak/dummy_service/db/postgre"
	"github.com/pawelkowalak/dummy_service/db/sqlite"
)

var dbType = flag.String("db_type", "", "DB type")

func main() {
	flag.Parse()
	var s *dummy_service.Service
	if *dbType == "postgres" {
		s = dummy_service.New(postgre.New("postgres://url_to_db"))
	} else if *dbType == "sqlite" {
		s = dummy_service.New(sqlite.New("sqlite_local_db"))
	} else {
		log.Fatal("Unknown DB type")
	}
	s.DoWork()
}
