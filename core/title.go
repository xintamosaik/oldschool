package core

import ("templeruins/persistence/files")

func TitleRead() string{
	return files.TitleRead()
}

func TitleSave(text string) {
	files.TitleSave(text)
}
