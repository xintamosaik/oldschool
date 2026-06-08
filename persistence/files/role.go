package files

func RoleSave(text string) {
	txtSave(text, "role.txt")
}

func RoleRead() string {
	return txtRead("role.txt")
}
