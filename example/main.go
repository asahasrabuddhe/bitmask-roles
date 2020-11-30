package main

import (
	role "bitmask-roles"
	"fmt"
)

func main() {
	r := role.NewRole(role.Admin)

	// Prints: Admin
	fmt.Println(r)

	r.AddRole(role.Analytic)
	// Prints: Analytic, Admin
	fmt.Println(r)

	// Prints: true true false
	fmt.Println(r.IsRole(role.Admin), r.IsRole(role.Analytic), r.IsRole(role.Curator))
}
