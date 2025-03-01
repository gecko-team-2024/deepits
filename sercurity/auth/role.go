package auth

var roles = map[string][]string{
	"admin": {"SET", "GET", "DELETE"},
	"user":  {"SET", "GET"},
}

var userRoles = map[string]string{} // map luu tru ten nguoi dung va vai tro

// gan vai tro cho nguoi dung
func AssignRole(username, role string) {
	userRoles[username] = role
}

// kiem tra quyen han cua nguoi dung
func CanExecute(username, command string) bool {
	role, exists := userRoles[username]
	if !exists {
		return false
	}
	permissions, exists := roles[role]
	if !exists {
		return false
	}
	for _, perm := range permissions {
		if perm == command {
			return true
		}
	}
	return false
}
