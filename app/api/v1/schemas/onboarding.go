package schemas

type OnboardingReq struct {
	CurriculumRef int `json:"curriculumRef"`
	StartYear     int `json:"startYear"`
}

type OnboardingFilter struct {
	CurriculumTypes []CurriculumType `json:"curriculumType"`
	Curriculums     []Curriculum     `json:"curriculums"`
}
