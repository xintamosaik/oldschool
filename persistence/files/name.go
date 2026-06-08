package files

func NameSave(text string) {
	txtSave(text, "name.txt")
}

func NameRead() string {
	return txtRead("name.txt")
}
