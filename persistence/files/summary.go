package files

func SummarySave(text string) {
	fileSave(text, "summary.md")
}

func SummaryRead() string {
	return fileRead("summary.md")
}
