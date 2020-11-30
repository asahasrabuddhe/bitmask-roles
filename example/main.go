package main

import (
	role "bitmask-roles"
	"fmt"
)

func main() {
	r := role.NewRole(role.Admin)

	fmt.Println(r)

	r.AddRole(role.Analytic)

	fmt.Println(r)

	fmt.Println(r.IsRole(role.Admin), r.IsRole(role.Analytic), r.IsRole(role.Curator))
}
