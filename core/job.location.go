package core

import ("templeruins/persistence/files")

func JobLocationUpdate(id string, location string) {
	job := files.JobRead(id)
	job.Location = location
	files.JobUpdate(job)
}

func JobLocationRead(id string)string{
	return files.JobRead(id).Location
}
 