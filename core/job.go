package core

import "log"

import ("templeruins/persistence/files")

func JobUpdate(updated files.Job) {
	files.JobUpdate(updated)
	log.Println(updated.ID)
}

func JobRead(id string)files.Job{
	return files.JobRead(id)
}