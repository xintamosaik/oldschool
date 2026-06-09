package core

import ("templeruins/persistence/files")

func CommercialImpactRead() string{
	return files.CommercialImpactRead()
}

func CommercialImpactSave(text string) {
	files.CommercialImpactSave(text)
}
