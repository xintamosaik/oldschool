package core

import ("templeruins/persistence/files")

func JobUpdate(updated files.Job) {
	files.JobUpdate(updated)
}

func JobRead(id string)files.Job{
	return files.JobRead(id)
}