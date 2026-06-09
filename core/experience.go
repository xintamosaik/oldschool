package core

import ("templeruins/persistence/files")

func ExperienceRead() files.Experience{
	return files.ExperienceRead()
}

func ExperienceSave(experience files.Experience) {
	files.ExperienceSave(experience)
}
