package service

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/wantedly/slack-mention-converter/models"
)

func randomString() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func setup() {
	tmpdir := fmt.Sprintf("tmp/%v", randomString())
	os.MkdirAll(filepath.Join("..", tmpdir), 0755)
	UserMapCachePath = tmpdir + "/user_map.csv"
}

func teardown() {
	os.RemoveAll(filepath.Join("..", "tmp"))
}

func TestUserFlow(t *testing.T) {
	setup()
	users, err := ListUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 0 {
		t.Fatalf("0 users should be present in initial state")
	}

	err = AddUser(&models.User{"awakia", "naoyoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = AddUser(&models.User{"kawasy", "yoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = AddUser(&models.User{"awakia", "nao"})
	if err != nil {
		t.Errorf("%v", err)
	}

	users, err = ListUsers()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 2 {
		t.Fatalf("2 users should be present after adding users but %v users", len(users))
	}
	if (users[0] != &models.User{"awakia", "nao"}) {
		t.Errorf("users[0] should be :%#v, but: %#v", &models.User{"awakia", "nao"}, users[0])
	}
	if (users[1] != &models.User{"kawasy", "yoshi"}) {
		t.Errorf("users[1] should be :%#v, but: %#v", &models.User{"kawasy", "yoshi"}, users[1])
	}

	user, err := GetUser("awakia")
	if err != nil {
		t.Errorf("%v", err)
	}

	if (user != &models.User{"awakia", "nao"}) {
		t.Errorf("user should be :%#v, but: %#v", &models.User{"awakia", "nao"}, user)
	}
	teardown()
}
