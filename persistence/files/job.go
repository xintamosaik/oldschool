package files

type Job struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	IsCurrent bool   `json:"isCurrent"`

	StartYear  int `json:"startYear"`
	StartMonth int `json:"startMonth"` // e.g., "Jan", "April"

	// Pointers allow these to be null in JSON if isCurrent is true
	EndYear  int `json:"endYear"`
	EndMonth int `json:"endMonth"`

	Highlights []string `json:"highlights"`
}

func JobAdd(new Job){

}

func JobUpdate(updated Job){
	i := jobFind(updated.ID)
	if (i < 0) {
		return
	}
	experience := ExperienceRead()
	experience[i] = updated
}

func jobFind(id string)int {
	i := 0
	for _, job := range ExperienceRead() {
		if job.ID == id {

			return i
		}
		i = i + 1
	}

	return -1
}

func JobRead(id string) Job{
	for _, job := range ExperienceRead() {
		if job.ID == id {
			return job
		}
	}
	return Job{} // Fallback
}
 