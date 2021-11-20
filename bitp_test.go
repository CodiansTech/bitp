package bitp

import (
	"fmt"
	"testing"
)

type Permission int

const (
	CreatePost     Permission = 1 << iota //          1 = 00001
	DeletePost                            // 1 << 1 = 2 = 00010
	EditPost                              // 1 << 2 = 4 = 00100
	ChangeAvatar                          // 1 << 3 = 8 = 01000
	ChangePassword                        // 1 << 4 = 16 = 10000
)

func TestPermission(t *testing.T) {
	userPermissions := 24

	if HasPermission(userPermissions, int(CreatePost)) {
		fmt.Printf("User can create posts")
	} else {
		fmt.Println("User cannot create posts")
	}

	SetPermission(&userPermissions, int(CreatePost))
	if HasPermission(userPermissions, int(CreatePost)) {
		fmt.Printf("User can create posts")
	} else {
		fmt.Println("User cannot create posts")
	}

	RemovePermission(&userPermissions, int(CreatePost))
	if HasPermission(userPermissions, int(CreatePost)) {
		fmt.Printf("User can create posts")
	} else {
		fmt.Println("User cannot create posts")
	}

	fmt.Println(IntegerToBits(userPermissions)) // 11000
}
