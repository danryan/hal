package hal

// Brain struct
type Brain struct {
	Users map[string]User
	store Store
}

// NewBrain
func NewBrain(robot *Robot) (*Brain, error) {
	return &Brain{
		Users: map[string]User{},
		store: robot.Store,
	}, nil
}

// Run starts up the brain
func (b *Brain) Run() error {
	return nil
}

// Stop shuts down the brain
func (b *Brain) Stop() error {
	return nil
}

// TODO(dryan): implement this!
func (b *Brain) getUsers() ([]User, error) {
	_, err := b.store.Get("hal:users")

	users := []User{}
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Get wraps the underlying store's Get function
func (b *Brain) Get(key string) ([]byte, error) {
	return b.store.Get(key)
}

// Set wraps the underlying store's Set function
func (b *Brain) Set(key string, data []byte) error {
	return b.store.Set(key, data)
}

// Delete wraps the underlying store's Delete function
func (b *Brain) Delete(key string) error {
	return b.store.Delete(key)
}
