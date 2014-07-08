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

type UserMap struct {
	Map   map[string]User
	robot *Robot
}

func NewUserMap(robot *Robot) *UserMap {
	return &UserMap{
		Map:   make(map[string]User, 0),
		robot: robot,
	}
}

func (um *UserMap) All() map[string]User {
	return um.Map
}

func (um *UserMap) Get(id string) (User, error) {
	user, ok := um.Map[id]
	if !ok {
		return User{}, fmt.Errorf("could not find user with id %s", id)
	}
	return user, nil
}

func (um *UserMap) Set(id string, user User) error {
	um.Map[id] = user

	if err := um.Save(); err != nil {
		return err
	}
	return nil
}

// Encode func
func (um *UserMap) Encode() ([]byte, error) {
	data, err := json.Marshal(um.Map)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// Decode func
func (um *UserMap) Decode() (map[string]User, error) {
	data, err := um.robot.Store.Get("users")
	if err != nil {
		return nil, err
	}

	users := map[string]User{}
	if err := json.Unmarshal(data, &users); err != nil {
		return users, err
	}

	return users, nil
}

// Load func
func (um *UserMap) Load() error {
	data, err := um.Decode()
	if err != nil {
		return err
	}

	um.Map = data
	return nil
}

// Save func
func (um *UserMap) Save() error {
	data, err := um.Encode()
	if err != nil {
		return err
	}

	return um.robot.Store.Set("users", data)
}
