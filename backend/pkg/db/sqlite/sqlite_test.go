package db_test

import (
	"database/sql"
	"os"
	"path"
	db "socialnetwork/pkg/db/sqlite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitialiseDatabase(t *testing.T) {
	tempPath := t.TempDir()
	cases := []struct {
		in string
	}
}

func TestCreateDatabase(t *testing.T) {
	tempPath := t.TempDir()
	testDBName := "testDB"
	db.CreateDatabase(tempPath, testDBName)
	filepath := path.Join(tempPath, testDBName)
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			t.Fatalf("file does not exist: %s", err)
		}
		file.Close()
	}
	db, err := sql.Open("sqlite3", filepath+"/.db?_foreign_keys=on")
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	db.Close()
}
