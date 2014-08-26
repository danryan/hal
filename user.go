package hal

import (
	"encoding/json"
	"fmt"
	"sync"
)

// User is a chat participant
type User struct {
	ID      string
	Name    string
	Roles   []string
	Options map[string]interface{}
}

func (u *User) Get(k string) (interface{}, error) {
	v, ok := u.Options[k]
	if !ok {
		return nil, fmt.Errorf("%s not found in user options", k)
	}
	return v, nil
}

func NewUser() *User {
	return &User{
		Options: make(map[string]interface{}),
	}
}

// UserMap handles the known users
type UserMap struct {
	Map   map[string]User
	robot *Robot
	sync.Mutex
}

// NewUserMap returns an initialized UserMap
func NewUserMap(robot *Robot) *UserMap {
	return &UserMap{
		Map:   make(map[string]User, 0),
		robot: robot,
	}
}

// All returns the underlying map of all users
func (um *UserMap) All() []User {
	um.Lock()

	users := make([]User, len(um.Map))
	for _, user := range um.Map {
		users = append(users, user)
	}

	um.Unlock()
	return users
}

// Get looks up a user by id and returns a User object
func (um *UserMap) Get(id string) (User, error) {
	um.Lock()
	defer um.Unlock()

	user, ok := um.Map[id]
	if !ok {
		return User{}, fmt.Errorf("could not find user with id %s", id)
	}
	return user, nil
}

// GetByName looks up a user by name and returns a User object
func (um *UserMap) GetByName(name string) (User, error) {
	um.Lock()
	defer um.Unlock()

	for _, user := range um.Map {
		if user.Name == name {
			if user.Options == nil {
				user.Options = make(map[string]interface{})
			}
			return user, nil
		}
	}
	return User{Options: make(map[string]interface{})}, fmt.Errorf("could not find user with name %s", name)
}

// Set adds or updates a user in the UserMap and persists it to the store
func (um *UserMap) Set(id string, user User) error {
	um.Lock()

	// initialize user.Options if nothing's in there yet
	if user.Options == nil {
		user.Options = make(map[string]interface{})
	}
	um.Map[id] = user
	if err := um.Save(); err != nil {
		um.Unlock()
		return err
	}

	um.Unlock()
	return nil
}

// Encode marshals a UserMap to JSON
func (um *UserMap) Encode() ([]byte, error) {
	data, err := json.Marshal(um.Map)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// Decode unmarshals a JSON object into a map of strings to Users
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

// Load retrieves known users from the store and populates the UserMap
func (um *UserMap) Load() error {
	um.Lock()

	data, err := um.Decode()
	if err != nil {
		um.Unlock()
		return err
	}

	um.Map = data

	um.Unlock()
	return nil
}

// Save persists known users to the store
func (um *UserMap) Save() error {
	data, err := um.Encode()
	if err != nil {
		return err
	}

	return um.robot.Store.Set("users", data)
}
