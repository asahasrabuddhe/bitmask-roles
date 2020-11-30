package role

import "strings"

var rolesMap = map[string]role{
	"Guest":    Guest,
	"Admin":    Admin,
	"Analytic": Analytic,
	"Curator":  Curator,
	"User":     User,
}

func (r *role) String() string {
	var roles []string

	for key, value := range rolesMap {
		if r.IsRole(value) {
			roles = append(roles, key)
		}
	}

	return strings.Join(roles, ", ")
}
