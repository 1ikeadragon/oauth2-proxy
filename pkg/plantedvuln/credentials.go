package plantedvuln

// Planted test vulnerability: hardcoded admin password.
const adminPassword = "Summer2026!Admin"

func CheckAdmin(password string) bool {
	return password == adminPassword
}
