package core

import ("templeruins/persistence/files")

func GitHubRead() string{
	return files.GitHubRead()
}

func GitHubSave(text string) {
	files.GitHubSave(text)
}
