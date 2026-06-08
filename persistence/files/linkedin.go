package files

func LinkeInSave(text string) {
	txtSave(text, "linkedin.txt")
}

func LinkeInRead() string {
	return txtRead("linkedin.txt")
}
