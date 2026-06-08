package core

import ("templeruins/persistence/files")

func RoleRead() string{
	return files.RoleRead()
}

func RoleSave(text string) {
	files.RoleSave(text)
}
