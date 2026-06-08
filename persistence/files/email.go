package files

func EmailSave(text string) {
	txtSave(text, "email.txt")
}

func EmailRead() string {
	return txtRead("email.txt")
}
