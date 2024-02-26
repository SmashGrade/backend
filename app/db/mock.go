package db

import (
	"time"

	"github.com/SmashGrade/backend/app/models"
)

func prefillMockDB(p Provider) {
	role1 := models.Role{
		Description: "Student",
	}
	role2 := models.Role{
		Description: "Dozent",
	}
	p.DB().Table("roles").Create(&role1)
	p.DB().Table("roles").Create(&role2)

	user1 := models.User{
		Name:  "Kurt Munter",
		Email: "kurt.munter@hftm.ch",
		Roles: []*models.Role{
			&role2,
		},
	}
	user2 := models.User{
		Name:           "Jakob Ferber",
		Email:          "jakob.ferber@hftm.ch",
		ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
		Roles: []*models.Role{
			&role1,
		},
	}
	user3 := models.User{
		Name:           "Rafael Stauffer",
		Email:          "rafael.stauffer@hftm.ch",
		ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
		Roles: []*models.Role{
			&role1,
		},
	}
	user4 := models.User{
		Name:  "Bruno Borer",
		Email: "bruno.borer@hftm.ch",
		Roles: []*models.Role{
			&role2,
		},
	}
	user5 := models.User{
		Name:  "Simeon Liniger",
		Email: "simeon.liniger@hftm.ch",
		Roles: []*models.Role{
			&role2,
		},
	}
	p.DB().Table("users").Create(&user1)
	p.DB().Table("users").Create(&user2)
	p.DB().Table("users").Create(&user3)
	p.DB().Table("users").Create(&user4)
	p.DB().Table("users").Create(&user5)

	field1 := models.Field{
		Description: "Informatik",
		Users: []*models.User{
			&user1,
		},
	}
	field2 := models.Field{
		Description: "Elektrotechnik",
		Users: []*models.User{
			&user4,
		},
	}
	p.DB().Table("fields").Create(&field1)
	p.DB().Table("fields").Create(&field2)

	curriculumTyp1 := models.Curriculumtype{
		Description:   "Fix",
		DurationYears: 2,
	}
	curriculumTyp2 := models.Curriculumtype{
		Description:   "Vollzeit",
		DurationYears: 2,
	}
	curriculumTyp3 := models.Curriculumtype{
		Description:   "Berufsbegleitend",
		DurationYears: 3,
	}
	p.DB().Table("curriculumtypes").Create(&curriculumTyp1)
	p.DB().Table("curriculumtypes").Create(&curriculumTyp2)
	p.DB().Table("curriculumtypes").Create(&curriculumTyp3)

	state1 := models.State{
		Description: "Aktiv",
	}
	state2 := models.State{
		Description: "Inaktiv",
	}
	p.DB().Table("states").Create(&state1)
	p.DB().Table("states").Create(&state2)

	focus1 := models.Focus{
		Description: "Softwareentwicklung",
		Field:       field1,
	}
	focus2 := models.Focus{
		Description: "Wirtschaftsinformatik",
		Field:       field1,
	}
	p.DB().Table("focus").Create(&focus1)
	p.DB().Table("focus").Create(&focus2)

	curriculum1 := models.Curriculum{
		Description:    "Softwareentwicklung",
		EndValidity:    time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC),
		Focus:          focus1,
		Curriculumtype: curriculumTyp3,
		State:          state1,
		TerminatedBasemodel: models.TerminatedBasemodel{
			StartValidity: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
			ID:            1,
		},
	}
	curriculum2 := models.Curriculum{
		Description:    "Softwareentwicklung",
		Focus:          focus2,
		Curriculumtype: curriculumTyp2,
		State:          state1,
		EndValidity:    time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC),
		TerminatedBasemodel: models.TerminatedBasemodel{
			StartValidity: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
			ID:            2,
		},
	}
	p.DB().Table("curriculums").Create(&curriculum1)
	p.DB().Table("curriculums").Create(&curriculum2)

	evaluationType1 := models.Evaluationtype{ // TODO: was sind evaluationtypes
		Description: "F",
		Code:        "F",
	}
	evaluationType2 := models.Evaluationtype{
		Description: "D",
		Code:        "D",
	}
	p.DB().Table("evaluationtypes").Create(&evaluationType1)
	p.DB().Table("evaluationtypes").Create(&evaluationType2)

	studyStage1 := models.StudyStage{
		Description: "Grundstudium",
	}
	studyStage2 := models.StudyStage{
		Description: "Fachstudium",
	}
	p.DB().Table("study_stages").Create(&studyStage1)
	p.DB().Table("study_stages").Create(&studyStage2)

	module1 := models.Module{
		State:          state1,
		StudyStage:     studyStage1,
		EvaluationType: evaluationType1,
		Description:    "Schnittstellen-Technologien",
		Number:         "IN123",
		Curriculums: []*models.Curriculum{
			&curriculum1,
			&curriculum2,
		},
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 1,
			ID:      1,
		},
	}
	module2 := models.Module{
		State:          state1,
		StudyStage:     studyStage1,
		EvaluationType: evaluationType1,
		Description:    "Requirements Engineering",
		Number:         "AB123",
		Curriculums: []*models.Curriculum{
			&curriculum1,
			&curriculum2,
		},
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 1,
			ID:      2,
		},
	}
	p.DB().Table("modules").Create(&module1)
	p.DB().Table("modules").Create(&module2)

	course1 := models.Course{
		Description: "Datenbank Scripting",
		Number:      "IN311",
		Modules: []*models.Module{
			&module1,
		},
		TeachedBy: []*models.User{
			&user5,
		},
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 1,
			ID:      1,
		},
	}
	course2 := models.Course{
		Description: "Modellieren von Softwaresystemen",
		Number:      "IN000",
		Modules: []*models.Module{
			&module2,
		},
		TeachedBy: []*models.User{
			&user1,
		},
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 1,
			ID:      2,
		},
	}
	course3 := models.Course{
		Description: "Modellieren von Softwaresystemen",
		Number:      "IN231",
		Modules: []*models.Module{
			&module2,
		},
		TeachedBy: []*models.User{
			&user1,
		},
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 2,
			ID:      2,
		},
	}
	p.DB().Table("courses").Create(&course1)
	p.DB().Table("courses").Create(&course2)
	p.DB().Table("courses").Create(&course3)

	selectedCourse1 := models.SelectedCourse{
		UserID:         2,
		CourseID:       2,
		CourseVersion:  2,
		ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
		Dispensed:      false,
	}
	selectedCourse2 := models.SelectedCourse{
		UserID:         3,
		CourseID:       2,
		CourseVersion:  2,
		ClassStartyear: time.Date(2021, time.April, 1, 12, 0, 0, 0, time.UTC),
		Dispensed:      false,
	}
	p.DB().Table("selected_courses").Create(&selectedCourse1)
	p.DB().Table("selected_courses").Create(&selectedCourse2)

	examtype1 := models.Examtype{
		Description: "Schriftlich",
	}
	examtype2 := models.Examtype{
		Description: "Mündlich",
	}
	p.DB().Table("examtypes").Create(&examtype1)
	p.DB().Table("examtypes").Create(&examtype2)

	exam1 := models.Exam{
		Description: "Erste Prüfung",
		Weighting:   2,
		Examtype:    examtype1,
		Course:      course3,
	}
	exam2 := models.Exam{
		Description: "Präsentation",
		Weighting:   1,
		Examtype:    examtype2,
		Course:      course3,
	}
	p.DB().Table("exams").Create(&exam1)
	p.DB().Table("exams").Create(&exam2)

	gradeType1 := models.Gradetype{
		Description: "CH-Noten",
	}
	gradeType2 := models.Gradetype{
		Description: "Prozent",
	}
	p.DB().Table("gradetypes").Create(&gradeType1)
	p.DB().Table("gradetypes").Create(&gradeType2)

	examEvaluation1 := models.ExamEvaluation{
		RegisteredBy:     user5,
		SelectedCourse:   selectedCourse1,
		Exam:             exam1,
		OriginalValue:    "5.2",
		OrignialGradeTyp: gradeType1,
		EntryDate:        time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
	}
	examEvaluation2 := models.ExamEvaluation{
		RegisteredBy:     user5,
		SelectedCourse:   selectedCourse2,
		Exam:             exam1,
		OriginalValue:    "5.0",
		OrignialGradeTyp: gradeType1,
		EntryDate:        time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
	}
	p.DB().Table("exam_evaluations").Create(&examEvaluation1)
	p.DB().Table("exam_evaluations").Create(&examEvaluation2)

	conversion1 := models.Conversion{
		ID:             1,
		ExamEvaluation: examEvaluation1,
		Gradetype:      gradeType1,
		Value:          5.2,
	}
	conversion2 := models.Conversion{
		ID:             2,
		ExamEvaluation: examEvaluation1,
		Gradetype:      gradeType2,
		Value:          84,
	}
	conversion3 := models.Conversion{
		ID:             3,
		ExamEvaluation: examEvaluation2,
		Gradetype:      gradeType1,
		Value:          5.0,
	}
	conversion4 := models.Conversion{
		ID:             4,
		ExamEvaluation: examEvaluation2,
		Gradetype:      gradeType2,
		Value:          80,
	}
	p.DB().Table("conversions").Create(&conversion1)
	p.DB().Table("conversions").Create(&conversion2)
	p.DB().Table("conversions").Create(&conversion3)
	p.DB().Table("conversions").Create(&conversion4)
}
