package core

import ("templeruins/persistence/files")

func LinkeInRead() string{
	return files.LinkeInRead()
}

func LinkeInSave(text string) {
	files.LinkeInSave(text)
}
