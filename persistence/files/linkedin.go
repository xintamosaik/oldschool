package files

func LinkeInSave(text string) {
	lineSave(text, "linkedin.txt")
}

func LinkeInRead() string {
	return lineRead("linkedin.txt")
}
