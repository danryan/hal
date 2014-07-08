package hal

import (
	"encoding/json"
	"fmt"
)

// User is a chat participant
type User struct {
	ID      string
	Name    string
	Roles   []string
	Options map[string]interface{}
}

type UserMap map[string]User

func (um *UserMap) Get(id string) (*User, error) {
	user, ok := (*um)[id]
	if !ok {
		return nil, fmt.Errorf("could not find user with id %s", id)
	}
	return &user, nil
}

func (um *UserMap) Set(id string, user *User) *User {
	(*um)[id] = *user
	return user
}

func (um *UserMap) Encode() ([]byte, error) {
	data, err := json.Marshal(um)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}
