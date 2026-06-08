package files

func SubtitleSave(text string) {
	txtSave(text, "subtitle.txt")
}

func SubtitleRead() string {
	return txtRead("subtitle.txt")
}
