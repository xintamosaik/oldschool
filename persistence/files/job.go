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
	EndYear  *int `json:"endYear,omitempty"`
	EndMonth *int `json:"endMonth,omitempty"`

	Highlights []string `json:"highlights"`
}

func JobAdd(new Job){

}

func JobUpdate(updated Job){

}

func JobRead(id string) Job{
	jobs := ExperienceRead()
	return jobs[0]
}
 