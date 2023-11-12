package schemas

type OnboardingReq struct {
	CurriculumRef int64 `json:"curriculumRef"`
	StartYear     int64 `json:"startYear"`
}

type OnboardingFilter struct {
	CurriculumTypes []CurriculumType `json:"curriculumType"`
	Curriculums     []Curriculum     `json:"curriculums"`
}
