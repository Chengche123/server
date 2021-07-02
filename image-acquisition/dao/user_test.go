package dao

import (
	"image-acquisition-service/conf"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	repo := newUserRepository()

	defer repo.Close()
}

func TestUserRepository_AddUser(t *testing.T) {
	repo := newUserRepository()
	defer repo.Close()

	id, err := repo.AddUser("xianchengyue", "qwe123456")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("id:  ", id)
}

func newUserRepository() *UserRepository {
	repo, err := NewUserRepository(conf.DSN)
	if err != nil {
		panic(err)
	}

	return repo
}

func TestUserRepository_FindUser(t *testing.T) {
	repo := newUserRepository()
	defer repo.Close()

	id, err := repo.FindUser("xianchengyue", "qwe123456")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("id: ", id)
}
