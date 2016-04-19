package service

import (
	"testing"
)

func TestUserFlow(t *testing.T) {
	users, err := ListUsers()
	if err == nil {
		t.Errorf("no such file or directory error should be present but not")
	}
	if len(users) != 0 {
		t.Fatalf("0 users should be present in initial state")
	}

	err = AddUser(User{"awakia", "naoyoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = AddUser(User{"kawasy", "yoshi"})
	if err != nil {
		t.Errorf("%v", err)
	}

	err = AddUser(User{"awakia", "nao"})
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
	if (users[0] != User{"awakia", "nao"}) {
		t.Errorf("users[0] should be :%#v, but: %#v", User{"awakia", "nao"}, users[0])
	}
	if (users[1] != User{"kawasy", "yoshi"}) {
		t.Errorf("users[1] should be :%#v, but: %#v", User{"kawasy", "yoshi"}, users[1])
	}

	user, err := GetUser("awakia")
	if err != nil {
		t.Errorf("%v", err)
	}

	if (user != User{"awakia", "nao"}) {
		t.Errorf("user should be :%#v, but: %#v", User{"awakia", "nao"}, user)
	}
}
