package schemas

type Curriculum struct {
	Id             uint   `json:"id"`
	Focus          string `json:"focus"`
	Field          string `json:"field"`
	CurriculumType string `json:"curriculumType"`
	IsActive       bool   `json:"isActive"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}

type CurriculumRes struct {
	Id             uint           `json:"id"`
	Focus          string         `json:"focus"`
	Field          string         `json:"field"`
	CurriculumType string         `json:"curriculumType"`
	IsActive       bool           `json:"isActive"`
	StartDate      string         `json:"startDate"`
	EndDate        string         `json:"endDate"`
	Description    string         `json:"description"`
	Fieldmanager   []Fieldmanager `json:"fieldmanager"`
	Moules         []ModuleRes    `json:"modules"`
}

type CurriculumReq struct {
	Focus           string `json:"focus"`
	Field           string `json:"field"`
	CurriculumType  string `json:"curriculumType"`
	IsActive        bool   `json:"isActive"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	FieldmanagerRef []uint `json:"fieldmanagerRef"`
	MoulesRef       []uint `json:"modulesRef"`
}

type CurriculumFilter struct {
	CurriculumTypes []CurriculumType `json:"curriculumTypes"`
	Fields          []Field          `json:"fields"`
	Focuses         []Focus          `json:"focuses"`
}

type CurriculumType struct {
	Id            uint   `json:"id"`
	Description   string `json:"description"`
	DurationYears int    `json:"durationYears"`
}

type Focus struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	FieldRef    uint   `json:"fieldRef"`
}

type Field struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
}
