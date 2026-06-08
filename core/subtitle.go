package core

import ("templeruins/persistence/files")

func SubtitleRead() string{
	return files.SubtitleRead()
}

func SubtitleSave(text string) {
	files.SubtitleSave(text)
}
