package schemas

type Curriculum struct {
	Id             int64  `json:"id"`
	Focus          string `json:"focus"`
	Field          string `json:"field"`
	CurriculumType string `json:"curriculumType"`
	IsActive       bool   `json:"isActiv"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}

type CurriculumRes struct {
	Id             int64          `json:"id"`
	Focus          string         `json:"focus"`
	Field          string         `json:"field"`
	CurriculumType string         `json:"curriculumType"`
	IsActive       bool           `json:"isActiv"`
	StartDate      string         `json:"startDate"`
	EndDate        string         `json:"endDate"`
	Fieldmanager   []Fieldmanager `json:"fieldmanager"`
	Moules         []ModuleRes    `json:"modules"`
}

type CurriculumReq struct {
	Focus           string  `json:"focus"`
	Field           string  `json:"field"`
	CurriculumType  string  `json:"curriculumType"`
	IsActive        bool    `json:"isActiv"`
	StartDate       string  `json:"startDate"`
	EndDate         string  `json:"endDate"`
	FieldmanagerRef []int64 `json:"fieldmanagerRef"`
	MoulesRef       []int64 `json:"modulesRef"`
}

type CurriculumFilter struct {
	CurriculumTypes []CurriculumType `json:"curriculumTypes"`
	Fields          []Field          `json:"fields"`
	Focuses         []Focus          `json:"focuses"`
}

type CurriculumType struct {
	Id            int64  `json:"id"`
	Description   string `json:"description"`
	DurationYears int64  `json:"durationYears"`
}

type Focus struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	FieldRef    int64  `json:"fieldRef"`
}

type Field struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
}
