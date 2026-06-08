package core

import ("templeruins/persistence/files")

func EmailRead() string{
	return files.EmailRead()
}

func EmailSave(text string) {
	files.EmailSave(text)
}
