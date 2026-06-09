package files

func CommercialImpactSave(text string) {
	lineSave(text, "commercial-impact.md")
}

func CommercialImpactRead() string {
	return lineRead("commercial-impact.md")
}
