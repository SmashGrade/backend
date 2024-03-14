package db

import (
	"time"

	"github.com/SmashGrade/backend/app/models"
)

type Mock struct {
	Table  string
	Entity any
}

func prefillMockDB(p Provider) error {
	mocks := []Mock{
		{
			Table:  "roles",
			Entity: &RoleStudent,
		},
		{
			Table:  "roles",
			Entity: &RoleDozent,
		},
		{
			Table:  "roles",
			Entity: &RoleKursadmin,
		},
		{
			Table:  "roles",
			Entity: &RoleFieldmanager,
		},
		{
			Table:  "users",
			Entity: &User1,
		},
		{
			Table:  "users",
			Entity: &User2,
		},
		{
			Table:  "users",
			Entity: &User3,
		},
		{
			Table:  "users",
			Entity: &User4,
		},
		{
			Table:  "users",
			Entity: &User5,
		},
		{
			Table:  "fields",
			Entity: &Field1,
		},
		{
			Table:  "fields",
			Entity: &Field2,
		},
		{
			Table:  "curriculumtypes",
			Entity: &CurriculumTyp1,
		},
		{
			Table:  "curriculumtypes",
			Entity: &CurriculumTyp2,
		},
		{
			Table:  "curriculumtypes",
			Entity: &CurriculumTyp3,
		},
		{
			Table:  "states",
			Entity: &State1,
		},
		{
			Table:  "states",
			Entity: &State2,
		},
		{
			Table:  "focus",
			Entity: &Focus1,
		},
		{
			Table:  "focus",
			Entity: &Focus2,
		},
		{
			Table:  "curriculums",
			Entity: &Curriculum1,
		},
		{
			Table:  "curriculums",
			Entity: &Curriculum2,
		},
		{
			Table:  "evaluationtypes",
			Entity: &EvaluationType1,
		},
		{
			Table:  "evaluationtypes",
			Entity: &EvaluationType2,
		},
		{
			Table:  "study_stages",
			Entity: &StudyStage1,
		},
		{
			Table:  "study_stages",
			Entity: &StudyStage2,
		},
		{
			Table:  "modules",
			Entity: &Module1,
		},
		{
			Table:  "modules",
			Entity: &Module2,
		},
		{
			Table:  "courses",
			Entity: &Course1,
		},
		{
			Table:  "courses",
			Entity: &Course2,
		},
		{
			Table:  "courses",
			Entity: &Course3,
		},
		{
			Table:  "selected_courses",
			Entity: &SelectedCourse1,
		},
		{
			Table:  "selected_courses",
			Entity: &SelectedCourse2,
		},
		{
			Table:  "examtypes",
			Entity: &Examtype1,
		},
		{
			Table:  "examtypes",
			Entity: &Examtype2,
		},
		{
			Table:  "exams",
			Entity: &Exam1,
		},
		{
			Table:  "exams",
			Entity: &Exam2,
		},
		{
			Table:  "gradetypes",
			Entity: &GradeType1,
		},
		{
			Table:  "gradetypes",
			Entity: &GradeType2,
		},
		{
			Table:  "exam_evaluations",
			Entity: &ExamEvaluation1,
		},
		{
			Table:  "exam_evaluations",
			Entity: &ExamEvaluation2,
		},
		{
			Table:  "exam_evaluations",
			Entity: &ExamEvaluation3,
		},
		{
			Table:  "conversions",
			Entity: &Conversion1,
		},
		{
			Table:  "conversions",
			Entity: &Conversion2,
		},
		{
			Table:  "conversions",
			Entity: &Conversion3,
		},
		{
			Table:  "conversions",
			Entity: &Conversion4,
		},
	}

	for _, mock := range mocks {
		res := p.DB().Table(mock.Table).FirstOrCreate(mock.Entity)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

var RoleStudent = models.Role{
	Basemodel: models.Basemodel{
		ID: 4,
	},
	Description: "Student",
	Claim:       "Student",
}
var RoleDozent = models.Role{
	Basemodel: models.Basemodel{
		ID: 3,
	},
	Description: "Dozent",
	Claim:       "Dozent",
}
var RoleFieldmanager = models.Role{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Fachbereichsleiter",
	Claim:       "Fachbereichsleiter",
}
var RoleKursadmin = models.Role{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Kursadministrator",
	Claim:       "Kursadministrator",
}

var User1 = models.User{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Name:  "Kurt Munter",
	Email: "kurt.munter@hftm.ch",
	Roles: []*models.Role{
		&RoleDozent, &RoleKursadmin, &RoleFieldmanager,
	},
}
var User2 = models.User{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Name:           "Jakob Ferber",
	Email:          "jakob.ferber@hftm.ch",
	ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
	Roles: []*models.Role{
		&RoleStudent,
	},
}
var User3 = models.User{
	Basemodel: models.Basemodel{
		ID: 3,
	},
	Name:           "Rafael Stauffer",
	Email:          "rafael.stauffer@hftm.ch",
	ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
	Roles: []*models.Role{
		&RoleStudent,
	},
}
var User4 = models.User{
	Basemodel: models.Basemodel{
		ID: 4,
	},
	Name:  "Bruno Borer",
	Email: "bruno.borer@hftm.ch",
	Roles: []*models.Role{
		&RoleDozent,
	},
}
var User5 = models.User{
	Basemodel: models.Basemodel{
		ID: 5,
	},
	Name:  "Simeon Liniger",
	Email: "simeon.liniger@hftm.ch",
	Roles: []*models.Role{
		&RoleDozent,
	},
}

var Field1 = models.Field{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Informatik",
	Users: []*models.User{
		&User1,
	},
}
var Field2 = models.Field{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Elektrotechnik",
	Users: []*models.User{
		&User4,
	},
}

var CurriculumTyp1 = models.Curriculumtype{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description:   "Fix",
	DurationYears: 2,
}
var CurriculumTyp2 = models.Curriculumtype{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description:   "Vollzeit",
	DurationYears: 2,
}
var CurriculumTyp3 = models.Curriculumtype{
	Basemodel: models.Basemodel{
		ID: 3,
	},
	Description:   "Berufsbegleitend",
	DurationYears: 3,
}

var State1 = models.State{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Aktiv",
}
var State2 = models.State{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Inaktiv",
}

var Focus1 = models.Focus{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Softwareentwicklung",
	Field:       Field1,
}
var Focus2 = models.Focus{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Wirtschaftsinformatik",
	Field:       Field1,
}

var Curriculum1 = models.Curriculum{
	Description:    "Softwareentwicklung",
	EndValidity:    time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC),
	Focus:          Focus1,
	Curriculumtype: CurriculumTyp3,
	State:          State1,
	TerminatedBasemodel: models.TerminatedBasemodel{
		StartValidity: time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
		ID:            1,
	},
}
var Curriculum2 = models.Curriculum{
	Description:    "Softwareentwicklung",
	Focus:          Focus2,
	Curriculumtype: CurriculumTyp2,
	State:          State1,
	EndValidity:    time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC),
	TerminatedBasemodel: models.TerminatedBasemodel{
		StartValidity: time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
		ID:            2,
	},
}

var EvaluationType1 = models.Evaluationtype{ // TODO: was sind evaluationtypes
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "F",
	Code:        "F",
}
var EvaluationType2 = models.Evaluationtype{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "D",
	Code:        "D",
}

var StudyStage1 = models.StudyStage{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Grundstudium",
}
var StudyStage2 = models.StudyStage{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Fachstudium",
}

var Module1 = models.Module{
	State:          State1,
	StudyStage:     StudyStage1,
	EvaluationType: EvaluationType1,
	Description:    "Schnittstellen-Technologien",
	Number:         "IN123",
	Curriculums: []*models.Curriculum{
		&Curriculum1,
		&Curriculum2,
	},
	VersionedBasemodel: models.VersionedBasemodel{
		Version: 1,
		ID:      1,
	},
}
var Module2 = models.Module{
	State:          State1,
	StudyStage:     StudyStage1,
	EvaluationType: EvaluationType1,
	Description:    "Requirements Engineering",
	Number:         "AB123",
	Curriculums: []*models.Curriculum{
		&Curriculum1,
		&Curriculum2,
	},
	VersionedBasemodel: models.VersionedBasemodel{
		Version: 1,
		ID:      2,
	},
}

var Course1 = models.Course{
	Description: "Datenbank Scripting",
	Number:      "IN311",
	Modules: []*models.Module{
		&Module1,
	},
	TeachedBy: []*models.User{
		&User5,
	},
	VersionedBasemodel: models.VersionedBasemodel{
		Version: 1,
		ID:      1,
	},
}
var Course2 = models.Course{
	Description: "Modellieren von Softwaresystemen",
	Number:      "IN000",
	Modules: []*models.Module{
		&Module2,
	},
	TeachedBy: []*models.User{
		&User1,
	},
	VersionedBasemodel: models.VersionedBasemodel{
		Version: 1,
		ID:      2,
	},
}
var Course3 = models.Course{
	Description: "Modellieren von Softwaresystemen",
	Number:      "IN231",
	Modules: []*models.Module{
		&Module2,
	},
	TeachedBy: []*models.User{
		&User1,
	},
	VersionedBasemodel: models.VersionedBasemodel{
		Version: 2,
		ID:      2,
	},
}

var SelectedCourse1 = models.SelectedCourse{
	UserID:         2,
	CourseID:       2,
	CourseVersion:  2,
	ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
	Dispensed:      false,
}
var SelectedCourse2 = models.SelectedCourse{
	UserID:         3,
	CourseID:       2,
	CourseVersion:  2,
	ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
	Dispensed:      false,
}

var Examtype1 = models.Examtype{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Schriftlich",
}
var Examtype2 = models.Examtype{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Mündlich",
}

var Exam1 = models.Exam{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Erste Prüfung",
	Weighting:   2,
	Examtype:    Examtype1,
	Course:      Course3,
}
var Exam2 = models.Exam{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Präsentation",
	Weighting:   1,
	Examtype:    Examtype2,
	Course:      Course3,
}

var GradeType1 = models.Gradetype{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "CH-Noten",
}
var GradeType2 = models.Gradetype{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Prozent",
}

var ExamEvaluation1 = models.ExamEvaluation{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	RegisteredBy:     User5,
	SelectedCourse:   SelectedCourse1,
	Exam:             Exam1,
	OriginalValue:    "5.2",
	OrignialGradeTyp: GradeType1,
	EntryDate:        time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
}
var ExamEvaluation2 = models.ExamEvaluation{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	RegisteredBy:     User5,
	SelectedCourse:   SelectedCourse2,
	Exam:             Exam1,
	OriginalValue:    "5.0",
	OrignialGradeTyp: GradeType1,
	EntryDate:        time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
}
var ExamEvaluation3 = models.ExamEvaluation{
	Basemodel: models.Basemodel{
		ID: 3,
	},
	RegisteredBy:   User5,
	SelectedCourse: SelectedCourse2,
	Exam:           Exam2,
	OriginalValue:  "4.8",
	EntryDate:      time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
}

var Conversion1 = models.Conversion{
	ID:             1,
	ExamEvaluation: ExamEvaluation1,
	Gradetype:      GradeType1,
	Value:          5.2,
}
var Conversion2 = models.Conversion{
	ID:             2,
	ExamEvaluation: ExamEvaluation1,
	Gradetype:      GradeType2,
	Value:          84,
}
var Conversion3 = models.Conversion{
	ID:             3,
	ExamEvaluation: ExamEvaluation2,
	Gradetype:      GradeType1,
	Value:          5.0,
}
var Conversion4 = models.Conversion{
	ID:             4,
	ExamEvaluation: ExamEvaluation2,
	Gradetype:      GradeType2,
	Value:          80,
}
