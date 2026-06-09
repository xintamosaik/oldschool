package files

func GitHubSave(text string) {
	lineSave(text, "github.txt")
}

func GitHubRead() string {
	return lineRead("github.txt")
}
