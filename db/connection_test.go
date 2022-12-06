package db_test

import (
	"go-docker/db"
	"testing"
)

func TestConnection(t *testing.T) {
	db.Connection("../.env")
}
