package core

import ("templeruins/persistence/files")

func JobUpdate(new files.Job) {
	files.JobUpdate(new)
}

func JobRead(id string)files.Job{
	return files.JobRead(id)
}