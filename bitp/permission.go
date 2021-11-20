package bitp

import "strconv"

// IntegerToBits gets an integer and returns it in stringified bits
func IntegerToBits(perms int) string {
	return strconv.FormatInt(int64(perms), 2)
}

// HasPermission gets user permissions arg (int) and permission(int) and runs a logical "AND" and "IS EQUAL" operator to validate if the user has permission
func HasPermission(perms int, perm int) bool {
	return int(perms)&perm == perm
}

// SetPermission gets user permissions pointer (int) and a permission(int) and runs a logical "OR" operator to update the user permission
func SetPermission(perms *int, perm int) {
	*perms = *perms | perm
}

// RemovePermission gets user permissions pointer (int) and a permission(int) and runs a logical "AND NOT" operator to remove a permission from the user permissions
func RemovePermission(perms *int, perm int) {
	*perms = *perms &^ perm
}
