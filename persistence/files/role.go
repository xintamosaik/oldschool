package files

func TitleSave(text string) {
	lineSave(text, "title.txt")
}

func TitleRead() string {
	return lineRead("title.txt")
}
