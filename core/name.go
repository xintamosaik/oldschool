package core

import ("templeruins/persistence/files")

func NameRead() string{
	return files.NameRead()
}

func NameSave(text string) {
	files.NameSave(text)
}
