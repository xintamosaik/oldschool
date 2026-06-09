package files
 
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

 
 