package conf

import "testing"

func TestCreateFile(t *testing.T) {
	if err := createFile(); err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	config := NewConfig()
	if err := readFile(config); err != nil {
		t.Fatal(err)
	}
	t.Log(config.String())
}

func TestWriteFile(t *testing.T) {
	config := NewDefaultConfig()
	config.SetMaxUsers(30000)
	config.SetMaxRoomsPerAdmin(30)
	t.Log(config.String())
	if err := writeFile(config); err != nil {
		t.Fatal(err)
	}
}