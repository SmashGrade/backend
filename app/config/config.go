package config

import (
	"fmt"
)

// APIConfig is used to define the configuration of the API
type APIConfig struct {
	Host                string                     `yaml:"host"`                // Host is the address of the server
	Port                int                        `yaml:"port"`                // Port is the port of the server
	AutoMigrate         bool                       `yaml:"autoMigrate"`         // AutoMigrate is a flag to determine if the database should be migrated automatically
	DBConnectionStr     string                     `yaml:"dbConnectionStr"`     // DBConnectionStr is the connection string for the database
	ExamTypes           []string                   `yaml:"examTypes"`           // ExamTypes is the list of exam types
	ExamEvaluationTypes []ExamEvaluationTypeConfig `yaml:"examEvaluationTypes"` // EvalTypes is the list of evaluation types
	GradeTypes          []string                   `yaml:"gradeTypes"`          // GradeTypes is the list of grade types
	States              []string                   `yaml:"states"`              // States is the list of states
	CurriculumTypes     []string                   `yaml:"curriculumTypes"`     // CurriculumTypes is the list of curriculum types
}

type ExamEvaluationTypeConfig struct {
	code        string `yaml:"code"`        // code is the code of the evaluation type
	description string `yaml:"description"` // description is the description of the evaluation type
}

// Returns a new configuration with default values
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		Host:            "0.0.0.0",
		Port:            9000,
		AutoMigrate:     true,
		DBConnectionStr: "sqlite://data.db",
		ExamTypes:       []string{"Mündliche oder schriftliche Prüfung ", "Präsentationen", "Lernbericht", "schriftliche Arbeit", "Lernjournal"},
		GradeTypes:      []string{"Kein Eintrag", "Note (1-6)", "Prozentwert (0-100)"},
		ExamEvaluationTypes: []ExamEvaluationTypeConfig{
			{code: "F", description: "Modul bestanden, wenn jeder Kurs eine genügende Bewertung aufweist. (Art. 29)"},
			{code: "M", description: "Modul bestanden, wenn der Durchschnitt aller Kurse genügend und nicht mehr als ein Kurs im Modul ungenügend ist. (Art. 30)"},
			{code: "D", description: "Modul bestanden, wenn der Durchschnitt der Kurse genügend ist (mehr als 60%). (Art. 31)"},
			{code: "E", description: "Modul bestanden, wenn alle Kurse erfüllt sind. (Art. 32)"},
		},
		States:          []string{"Aktiv", "Inaktiv"},
		CurriculumTypes: []string{"Vollzeit", "Teilzeit"},
	}
}

// Returns the server address as a string
func (c *APIConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
