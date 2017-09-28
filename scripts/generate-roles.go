package main

import "gopkg.in/bblfsh/sdk.v1/uast"
import "fmt"

func main() {
	lastRole := uast.Visibility + 1
	fmt.Println(
		`////////////////////////////////////////////////////
// Automatically generated by "generate-roles.go" //
////////////////////////////////////////////////////

#include "roles.h"

#include <stddef.h>

static const char *id_to_roles[] = {`)

	for i := 0; i < int(lastRole); i++ {
		var name string
		if i == 0 {
			name = "roleInvalid"
		} else {
			roleID := uast.Role(i)
			name = "role" + roleID.String()
		}
		fmt.Printf("    \"%s\",\n", name)
	}
	fmt.Println("};")

	fmt.Println("#define TOTAL_ROLES", int(lastRole))

	fmt.Println(`
const char *RoleNameForId(uint16_t id) {
  if (id >= TOTAL_ROLES) {
    return NULL;
  }
  return id_to_roles[id];
}`)
}
