package hal

import (
	"fmt"
	"github.com/danryan/env"
	"strings"
)

// UserHasRole determines whether the Response's user has a given role
func UserHasRole(res *Response, role string) bool {
	user := res.Envelope.User
	for _, r := range user.Roles {
		if r == role {
			return true
		}
	}

	return false
}

// Auth type to group authentication methods
type Auth struct {
	robot  *Robot
	admins []string
}

type authConfig struct {
	Enabled bool   `env:"key=HAL_AUTH_ENABLED default=true"`
	Admins  string `env:"key=HAL_AUTH_ADMIN"`
}

// NewAuth returns a pointer to an initialized Auth
func NewAuth(r *Robot) *Auth {
	a := &Auth{robot: r}

	c := &authConfig{}
	env.MustProcess(c)

	if c.Enabled {
		if c.Admins != "" {
			a.admins = strings.Split(c.Admins, ",")
		}

		r.Handle(
			addUserRoleHandler,
			removeUserRoleHandler,
			listUserRolesHandler,
			listAdminsHandler,
		)
	}

	return a
}

// Admins returns a slice of admin Users
func (a *Auth) Admins() (admins []User) {
	for _, name := range a.admins {
		user, err := a.robot.Users.GetByName(name)
		if err != nil {
			continue
		}
		admins = append(admins, user)
	}

	return
}

// HasRole checks whether a user located by id has a given role(s)
func (a *Auth) HasRole(id string, roles ...string) bool {
	user, err := a.robot.Users.Get(id)
	if err != nil {
		return false
	}

	if len(user.Roles) == 0 {
		return false
	}

	for _, r := range roles {
		for _, b := range user.Roles {
			if b == r {
				return true
			}
		}
	}

	return false
}

// UsersWithRole returns a slice of Users that have a given role
func (a *Auth) UsersWithRole(role string) (users []User) {
	for _, user := range a.robot.Users.All() {
		if a.HasRole(user.ID, role) {
			users = append(users, user)
		}
	}
	return
}

// AddRole adds a role to a User
func (a *Auth) AddRole(user User, r string) error {
	if r == "admin" {
		return fmt.Errorf(`the "admin" role can only be defined by the HAL_AUTH_ADMIN environment variable`)
	}

	if a.HasRole(user.ID, r) {
		return fmt.Errorf("%s already has the %s role", user.Name, r)
	}

	user.Roles = append(user.Roles, r)
	a.robot.Users.Set(user.ID, user)

	return nil
}

// RemoveRole adds a role to a User
func (a *Auth) RemoveRole(user User, role string) error {
	if role == "admin" {
		return fmt.Errorf(`the "admin" role can only be defined by the HAL_AUTH_ADMIN environment variable`)
	}

	if !a.HasRole(user.ID, role) {
		return fmt.Errorf("%s already does not have the %s role", user.Name, role)
	}

	roles := make([]string, len(user.Roles)-1)

	for _, r := range user.Roles {
		if r != role {
			roles = append(roles, role)
		}
	}

	user.Roles = roles
	a.robot.Users.Set(user.ID, user)

	return nil
}

// IsAdmin checks whether a user is an admin
func (a *Auth) IsAdmin(user User) bool {
	for _, a := range a.admins {
		if a == user.Name {
			return true
		}
	}

	return false
}

var addUserRoleHandler = &Handler{
	Pattern: `(?i)@?(.+) (?:has)(?: the)? (["'\w: -_]+) (?:role)`,
	Method:  RESPOND,
	Run: func(res *Response) error {
		name := strings.TrimSpace(res.Match[1])
		role := strings.ToLower(res.Match[2])

		for _, s := range []string{"", "who", "what", "where", "when", "why"} {
			if s == name {
				return nil // don't match
			}
		}

		user, err := res.Robot.Users.GetByName(name)
		if err != nil {
			return res.Reply(err.Error())
		}

		if err := res.Robot.Auth.AddRole(user, role); err != nil {
			return res.Reply(err.Error())
		}

		return res.Reply(fmt.Sprintf("%s now has the %s role", name, role))
	},
}

var removeUserRoleHandler = &Handler{
	Pattern: `(?i)@?(.+) (?:does(?:n't| not) have)(?: the)? (["'\w: -_]+) (role)`,
	Method:  RESPOND,
	Run: func(res *Response) error {
		name := strings.TrimSpace(res.Match[1])
		role := strings.ToLower(res.Match[2])

		for _, s := range []string{"", "who", "what", "where", "when", "why"} {
			if s == name {
				return nil // don't match
			}
		}

		user, err := res.Robot.Users.GetByName(name)
		if err != nil {
			return res.Reply(err.Error())
		}

		if err := res.Robot.Auth.RemoveRole(user, role); err != nil {
			return res.Reply(err.Error())
		}

		return res.Reply(fmt.Sprintf("%s no longer has the %s role", name, role))
	},
}

var listUserRolesHandler = &Handler{
	Pattern: `(?i)(?:what roles? does) @?(.+) (?:have)\??`,
	Method:  RESPOND,
	Run: func(res *Response) error {
		name := res.Match[1]

		user, err := res.Robot.Users.GetByName(name)
		// return if we didn't find a user
		if err != nil {
			res.Reply(err.Error())
		}

		roles := user.Roles

		if res.Robot.Auth.IsAdmin(user) {
			roles = append(roles, "admin")
		}

		if len(roles) == 0 {
			return res.Reply(name + " has no roles")
		}

		return res.Reply(fmt.Sprintf("%s has the following roles: %s", name, strings.Join(roles, ", ")))
	},
}

var listAdminsHandler = &Handler{
	Pattern: `who (?:has)(?: the)? admin role\??`,
	Method:  RESPOND,
	Run: func(res *Response) error {
		admins := res.Robot.Auth.Admins()
		names := make([]string, len(admins))

		if len(names) == 0 {
			return res.Reply(`no users have the "admin" role`)
		}

		for i, u := range admins {
			names[i] = u.Name
		}

		return res.Reply(fmt.Sprintf(`the following users have the "admin" role: %s`, strings.Join(names, ", ")))
	},
}
