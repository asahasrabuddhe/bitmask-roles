//go:generate stringer-bitmask role.go role
package role

type role uint64

type Role interface {
	IsRole(role role) bool
	AddRole(role role)
	RemoveRole(role role)
	String() string
}

func (r *role) IsRole(role role) bool { return *r&role != 0 }
func (r *role) AddRole(role role)     { *r |= role }
func (r *role) RemoveRole(role role)  { *r &= role }

const (
	Guest role = 1 << iota
	Admin
	Analytic
	Curator
	User
)

func NewRole(role role) Role {
	return &role
}

func NewRoleFromInt64(r int64) Role {
	rle := role(r)
	return &rle
}
