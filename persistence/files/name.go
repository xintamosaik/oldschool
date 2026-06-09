package files

func NameSave(text string) {
	lineSave(text, "name.txt")
}

func NameRead() string {
	return lineRead("name.txt")
}
