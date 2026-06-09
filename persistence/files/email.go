package files

func EmailSave(text string) {
	lineSave(text, "email.txt")
}

func EmailRead() string {
	return lineRead("email.txt")
}
