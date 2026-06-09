package core

import ("templeruins/persistence/files")

func ExperienceRead() []files.Job{
	return files.ExperienceRead()
}
 
func ExperienceAdd(new files.Job) {
	files.ExperienceAdd(new)
}

func ExperienceUpdate(updated files.Job) {
	files.ExperienceUpdate(updated)
}