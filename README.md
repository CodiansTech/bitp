# Bitwise permissions simplified
A simple way to use bitwise operations in validating user permissions.

## Installation

```bash
go get github.com/ewoshi/bitp
```

## Usage

Please take a look at the testing file for a better understanding, usage of iota and enums are suggested for permissions struct.

```go
import "github.com/ewoshi/bitp"

userPermissions := 24
userRole := 1

# returns 'bool' false
bitp.HasPermission(userPermissions, int(userRole))

# adds userRole to userPermissions
bitp.SetPermission(&userPermissions, int(userRole))

# removes userRole from userPermission
bitp.RemovePermission(&userPermissions, int(userRole))
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.