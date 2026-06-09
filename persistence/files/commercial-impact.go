package files

func CommercialImpactSave(text string) {
	fileSave(text, "commercial-impact.md")
}

func CommercialImpactRead() string {
	return fileRead("commercial-impact.md")
}
