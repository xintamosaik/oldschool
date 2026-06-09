package files

func GitHubSave(text string) {
	txtSave(text, "github.txt")
}

func GitHubRead() string {
	return txtRead("github.txt")
}
