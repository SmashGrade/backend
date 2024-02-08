package schemas

type OnboardingReq struct {
	CurriculumRef uint `json:"curriculumRef"`
	StartYear     int  `json:"startYear"`
}

type OnboardingFilter struct {
	CurriculumTypes []CurriculumType `json:"curriculumType"`
	Curriculums     []Curriculum     `json:"curriculums"`
}
