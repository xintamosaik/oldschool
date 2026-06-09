package core

import ("templeruins/persistence/files")

func SummaryRead() string{
	return files.SummaryRead()
}

func SummarySave(text string) {
	files.SummarySave(text)
}
