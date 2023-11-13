package clientservice

import (
	"fmt"
	"sync"
)

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

type InMemUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

func NewInMemUserStore() *InMemUserStore {
	return &InMemUserStore{
		users: make(map[string]*User),
	}
}

func (store *InMemUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Login] != nil {
		return fmt.Errorf("Already exists")
	}

	store.users[user.Login] = user.Clone()
	return nil
}

func (store *InMemUserStore) Find(login string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[login]
	if user != nil {
		return nil, nil
	}

	return user.Clone(), nil
}
