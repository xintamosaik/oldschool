package core

import ("templeruins/persistence/files")

func ExperienceRead() []files.Job{
	return files.ExperienceRead()
}

func ExperienceSave(experience []files.Job) {
	files.ExperienceSave(experience)
}
