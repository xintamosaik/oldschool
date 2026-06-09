package files

func RoleSave(text string) {
	lineSave(text, "role.txt")
}

func RoleRead() string {
	return lineRead("role.txt")
}
