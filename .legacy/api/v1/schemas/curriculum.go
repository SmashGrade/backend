package schemas

// the minimal ressource to get in the api without any connections
type Curriculum struct {
	Id             uint   `json:"id"`
	Focus          string `json:"focus"`
	Field          string `json:"field"`
	CurriculumType string `json:"curriculumType"`
	IsActive       bool   `json:"isActive"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}

// the ressource to return will all connections in the api
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
	Modules        []ModuleRes    `json:"modules"`
}

// the ressource used for a post or update and has only minimal required connections via id
type CurriculumReq struct {
	Focus           string `json:"-"`
	Field           string `json:"-"`
	CurriculumType  string `json:"-"`
	IsActive        bool   `json:"isActive"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	FieldmanagerRef []uint `json:"fieldmanagerRef"`
	ModulesRef      []uint `json:"modulesRef"`
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
