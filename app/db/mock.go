package db

import (
	"time"

	"github.com/SmashGrade/backend/app/models"
)

func prefillMockDB(p Provider) {
	p.DB().Table("roles").FirstOrCreate(&Role1)
	p.DB().Table("roles").FirstOrCreate(&Role2)

	p.DB().Table("users").FirstOrCreate(&User1)
	p.DB().Table("users").FirstOrCreate(&User2)
	p.DB().Table("users").FirstOrCreate(&User3)
	p.DB().Table("users").FirstOrCreate(&User4)
	p.DB().Table("users").FirstOrCreate(&User5)

	p.DB().Table("fields").FirstOrCreate(&Field1)
	p.DB().Table("fields").FirstOrCreate(&Field2)

	p.DB().Table("curriculumtypes").FirstOrCreate(&CurriculumTyp3)
	p.DB().Table("curriculumtypes").FirstOrCreate(&CurriculumTyp1)
	p.DB().Table("curriculumtypes").FirstOrCreate(&CurriculumTyp2)

	p.DB().Table("states").FirstOrCreate(&State1)
	p.DB().Table("states").FirstOrCreate(&State2)

	p.DB().Table("focus").FirstOrCreate(&Focus1)
	p.DB().Table("focus").FirstOrCreate(&Focus2)

	p.DB().Table("curriculums").FirstOrCreate(&Curriculum1)
	p.DB().Table("curriculums").FirstOrCreate(&Curriculum2)

	p.DB().Table("evaluationtypes").FirstOrCreate(&EvaluationType1)
	p.DB().Table("evaluationtypes").FirstOrCreate(&EvaluationType2)

	p.DB().Table("study_stages").FirstOrCreate(&StudyStage1)
	p.DB().Table("study_stages").FirstOrCreate(&StudyStage2)

	p.DB().Table("modules").FirstOrCreate(&Module1)
	p.DB().Table("modules").FirstOrCreate(&Module2)

	p.DB().Table("courses").FirstOrCreate(&Course1)
	p.DB().Table("courses").FirstOrCreate(&Course2)
	p.DB().Table("courses").FirstOrCreate(&Course3)

	p.DB().Table("selected_courses").FirstOrCreate(&SelectedCourse1)
	p.DB().Table("selected_courses").FirstOrCreate(&SelectedCourse2)

	p.DB().Table("examtypes").FirstOrCreate(&Examtype1)
	p.DB().Table("examtypes").FirstOrCreate(&Examtype2)

	p.DB().Table("exams").FirstOrCreate(&Exam1)
	p.DB().Table("exams").FirstOrCreate(&Exam2)

	p.DB().Table("gradetypes").FirstOrCreate(&GradeType1)
	p.DB().Table("gradetypes").FirstOrCreate(&GradeType2)

	p.DB().Table("exam_evaluations").FirstOrCreate(&ExamEvaluation1)
	p.DB().Table("exam_evaluations").FirstOrCreate(&ExamEvaluation2)
	p.DB().Table("exam_evaluations").FirstOrCreate(&ExamEvaluation3)

	p.DB().Table("conversions").FirstOrCreate(&Conversion1)
	p.DB().Table("conversions").FirstOrCreate(&Conversion2)
	p.DB().Table("conversions").FirstOrCreate(&Conversion3)
	p.DB().Table("conversions").FirstOrCreate(&Conversion4)
}

var Role1 = models.Role{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Description: "Student",
}
var Role2 = models.Role{
	Basemodel: models.Basemodel{
		ID: 2,
	},
	Description: "Dozent",
}

var User1 = models.User{
	Basemodel: models.Basemodel{
		ID: 1,
	},
	Name:  "Kurt Munter",
	Email: "kurt.munter@hftm.ch",
	Roles: []*models.Role{
		&Role2,
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
		&Role1,
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
		&Role1,
	},
}
var User4 = models.User{
	Basemodel: models.Basemodel{
		ID: 4,
	},
	Name:  "Bruno Borer",
	Email: "bruno.borer@hftm.ch",
	Roles: []*models.Role{
		&Role2,
	},
}
var User5 = models.User{
	Basemodel: models.Basemodel{
		ID: 5,
	},
	Name:  "Simeon Liniger",
	Email: "simeon.liniger@hftm.ch",
	Roles: []*models.Role{
		&Role2,
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
