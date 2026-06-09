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
// Experience represents the top-level JSON wrapper structure
type Experience struct {
	Jobs []Job `json:"jobs"`
}

func experienceSave(experience []Job) {
	jsonWrite("experience.json", experience)
}

func ExperienceRead() []Job {
	return jsonRead("experience.json")
}

 
func ExperienceAdd(new Job) {
	current := ExperienceRead()
	// implement
	experienceSave(current)
}

func ExperienceUpdate(updated Job) {
	current := ExperienceRead()
	// implement
	experienceSave(current)
}