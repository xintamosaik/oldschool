package files

func SubtitleSave(text string) {
	lineSave(text, "subtitle.txt")
}

func SubtitleRead() string {
	return lineRead("subtitle.txt")
}
