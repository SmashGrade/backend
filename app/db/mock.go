package db

import (
	"time"

	"github.com/SmashGrade/backend/app/models"
	"github.com/google/uuid"
)

// Fill MockDB with prefill functions
func prefillMockDB(p Provider) {
	prefillFields(p)
	prefillConversions(p)
	prefillCourses(p)
	prefillCurriculum(p)
	prefillCurriculumtype(p)
	prefillEvaluationtype(p)
	prefillExam(p)
	prefillExamEvaluation(p)
	prefillExamtype(p)
	prefillFocus(p)
	prefillGradetype(p)
	prefillModules(p)
	prefillRole(p)
	prefillSelectedCourse(p)
	prefillState(p)
	prefillStudyStage(p)
	prefillUser(p)
}

// Field that will be added to the mock DB
func Field_1() models.Field {
	var field models.Field
	field.Description = "description Field 1"
	field.ID = 1
	return field
}

// Field that will be added to the mock DB
func Field_2() models.Field {
	var field models.Field
	field.Description = "description Field 2"
	field.ID = 2
	return field
}

// add all the fields to the fields table of the mockDB
func prefillFields(p Provider) {
	field_1 := Field_1()
	field_2 := Field_2()
	p.DB().Table("fields").Create(&field_1)
	p.DB().Table("fields").Create(&field_2)
}

// Conversion that will be added to the mock DB
func Conversion_1() models.Conversion {
	var conversion models.Conversion
	conversion.EEExamID = 1
	conversion.EESelectedCourseUserID = 1
	conversion.EESelectedCourseCourseVersion = 1
	conversion.EESelectedCourseClassStartyear = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	conversion.EEExamID = 1
	conversion.Value = 3.8
	conversion.ID = 1
	return conversion
}

// Conversion that will be added to the mock DB
func Conversion_2() models.Conversion {
	var conversion models.Conversion
	conversion.EEExamID = 2
	conversion.EESelectedCourseUserID = 2
	conversion.EESelectedCourseCourseVersion = 2
	conversion.EESelectedCourseClassStartyear = time.Date(2024, time.February, 2, 0, 0, 0, 0, time.UTC)
	conversion.EEExamID = 2
	conversion.Value = 5.5
	conversion.ID = 2
	return conversion
}

// add all the conversions to the conversions table of the mockDB
func prefillConversions(p Provider) {
	conversion_1 := Conversion_1()
	conversion_2 := Conversion_2()
	p.DB().Table("conversions").Create(&conversion_1)
	p.DB().Table("conversions").Create(&conversion_2)
}

// Course that will be added to the mock DB
func Course_1() models.Course {
	var course models.Course
	course.Description = "Course 1"
	course.Number = "NR01"
	course.Version = 1
	course.ID, _ = uuid.Parse("071e2e14-88d8-497a-aa20-c67539655686")
	return course
}

// Course that will be added to the mock DB
func Course_2_1() models.Course {
	var course models.Course
	course.Description = "Course 2"
	course.Number = "NR02"
	course.Version = 1
	course.ID, _ = uuid.Parse("6dd6ac96-9f30-408a-90b5-dd1cec955808")
	return course
}

// Course that will be added to the mock DB
func Course_2_2() models.Course {
	var course models.Course
	course.Description = "Course 2"
	course.Number = "NR02"
	course.Version = 2
	course.ID, _ = uuid.Parse("6dd6ac96-9f30-408a-90b5-dd1cec955808")
	return course
}

// add all the courses to the courses table of the mockDB
func prefillCourses(p Provider) {
	course_1 := Course_1()
	course_2_1 := Course_2_1()
	course_2_2 := Course_2_2()
	p.DB().Table("courses").Create(&course_1)
	p.DB().Table("courses").Create(&course_2_1)
	p.DB().Table("courses").Create(&course_2_2)
}

// Curriculum that will be added to the mock DB
func Curriculum_1() models.Curriculum {
	var curriculum models.Curriculum
	curriculum.Description = "Curriculum Description 1"
	curriculum.ID = 1
	return curriculum
}

// Curriculum that will be added to the mock DB
func Curriculum_2() models.Curriculum {
	var curriculum models.Curriculum
	curriculum.Description = "Curriculum Description 2"
	curriculum.ID = 2
	return curriculum
}

// add all the curriculum to the courses table of the mockDB
func prefillCurriculum(p Provider) {
	curriculum_1 := Curriculum_1()
	curriculum_2 := Curriculum_2()
	p.DB().Table("curriculums").Create(&curriculum_1)
	p.DB().Table("curriculums").Create(&curriculum_2)
}

// Curriculumtype that will be added to the mock DB
func Curriculumtype_1() models.Curriculumtype {
	var curriculum models.Curriculumtype
	curriculum.Description = "Curriculumtype Description 1"
	curriculum.ID = 1
	return curriculum
}

// Curriculumtype that will be added to the mock DB
func Curriculumtype_2() models.Curriculumtype {
	var curriculum models.Curriculumtype
	curriculum.Description = "Curriculumtype Description 2"
	curriculum.ID = 2
	return curriculum
}

// add all the curriculumtype to the courses table of the mockDB
func prefillCurriculumtype(p Provider) {
	curriculumtype_1 := Curriculumtype_1()
	curriculumtype_2 := Curriculumtype_2()
	p.DB().Table("curriculumtypes").Create(&curriculumtype_1)
	p.DB().Table("curriculumtypes").Create(&curriculumtype_2)
}

// Evaluationtype that will be added to the mock DB
func Evaluationtype_1() models.Evaluationtype {
	var evaluationtype models.Evaluationtype
	evaluationtype.Description = "Evaluationtype Description 1"
	evaluationtype.ID = 1
	return evaluationtype
}

// Evaluationtype that will be added to the mock DB
func Evaluationtype_2() models.Evaluationtype {
	var evaluationtype models.Evaluationtype
	evaluationtype.Description = "Evaluationtype Description 2"
	evaluationtype.ID = 2
	return evaluationtype
}

// add all the evaluationtype to the courses table of the mockDB
func prefillEvaluationtype(p Provider) {
	evaluationtype_1 := Evaluationtype_1()
	evaluationtype_2 := Evaluationtype_2()
	p.DB().Table("evaluationtypes").Create(&evaluationtype_1)
	p.DB().Table("evaluationtypes").Create(&evaluationtype_2)
}

// Exam that will be added to the mock DB
func Exam_1() models.Exam {
	var exam models.Exam
	exam.Description = "Exam Description 1"
	exam.ID = 1
	return exam
}

// Exam that will be added to the mock DB
func Exam_2() models.Exam {
	var exam models.Exam
	exam.Description = "Exam Description 2"
	exam.ID = 2
	return exam
}

// add all the exams to the exams table of the mockDB
func prefillExam(p Provider) {
	exam_1 := Exam_1()
	exam_2 := Exam_2()
	p.DB().Table("exams").Create(&exam_1)
	p.DB().Table("exams").Create(&exam_2)
}

// ExamEvaluation that will be added to the mock DB
func ExamEvaluation_1() models.ExamEvaluation {
	var examEvaluation models.ExamEvaluation
	examEvaluation.OriginalValue = "3.9"
	examEvaluation.ID = 1
	return examEvaluation
}

// ExamEvaluation that will be added to the mock DB
func ExamEvaluation_2() models.ExamEvaluation {
	var examEvaluation models.ExamEvaluation
	examEvaluation.OriginalValue = "4.2"
	examEvaluation.ID = 2
	return examEvaluation
}

// add all the ExamEvaluation to the exam_evaluations table of the mockDB
func prefillExamEvaluation(p Provider) {
	examEvaluation_1 := ExamEvaluation_1()
	examEvaluation_2 := ExamEvaluation_2()
	p.DB().Table("exam_evaluations").Create(&examEvaluation_1)
	p.DB().Table("exam_evaluations").Create(&examEvaluation_2)
}

// Examtype that will be added to the mock DB
func Examtype_1() models.Examtype {
	var examtype models.Examtype
	examtype.Description = "Examtype Description 1"
	examtype.ID = 1
	return examtype
}

// Examtype that will be added to the mock DB
func Examtype_2() models.Examtype {
	var examtype models.Examtype
	examtype.Description = "Examtype Description 2"
	examtype.ID = 2
	return examtype
}

// add all the examtypes to the examtypes table of the mockDB
func prefillExamtype(p Provider) {
	examtype_1 := Examtype_1()
	examtype_2 := Examtype_2()
	p.DB().Table("examtypes").Create(&examtype_1)
	p.DB().Table("examtypes").Create(&examtype_2)
}

// Focus that will be added to the mock DB
func Focus_1() models.Focus {
	var focus models.Focus
	focus.Description = "Focus Description 1"
	focus.ID = 1
	return focus
}

// Focus that will be added to the mock DB
func Focus_2() models.Focus {
	var focus models.Focus
	focus.Description = "Focus Description 2"
	focus.ID = 2
	return focus
}

// add all the focuss to the focuss table of the mockDB
func prefillFocus(p Provider) {
	focus_1 := Focus_1()
	focus_2 := Focus_2()
	p.DB().Table("focus").Create(&focus_1)
	p.DB().Table("focus").Create(&focus_2)
}

// Gradetype that will be added to the mock DB
func Gradetype_1() models.Gradetype {
	var gradetype models.Gradetype
	gradetype.Description = "Gradetype Description 1"
	gradetype.ID = 1
	return gradetype
}

// Gradetype that will be added to the mock DB
func Gradetype_2() models.Gradetype {
	var gradetype models.Gradetype
	gradetype.Description = "Gradetype Description 2"
	gradetype.ID = 2
	return gradetype
}

// add all the gradetypes to the gradetypes table of the mockDB
func prefillGradetype(p Provider) {
	gradetype_1 := Gradetype_1()
	gradetype_2 := Gradetype_2()
	p.DB().Table("gradetypes").Create(&gradetype_1)
	p.DB().Table("gradetypes").Create(&gradetype_2)
}

// Module that will be added to the mock DB
func Module_1() models.Module {
	var module models.Module
	module.Description = "Module 1"
	module.Number = "NR02"
	module.Version = 1
	module.ID, _ = uuid.Parse("0c256c99-4c64-4747-ac1a-d0362fc75319")
	return module
}

// Module that will be added to the mock DB
func Module_2_1() models.Module {
	var module models.Module
	module.Description = "Module 2"
	module.Number = "NR02"
	module.Version = 1
	module.ID, _ = uuid.Parse("4c14f1db-82e2-41f3-96cc-71e3b3d52d13")
	return module
}

// Module that will be added to the mock DB
func Module_2_2() models.Module {
	var module models.Module
	module.Description = "Module 2"
	module.Number = "NR02"
	module.Version = 2
	module.ID, _ = uuid.Parse("4c14f1db-82e2-41f3-96cc-71e3b3d52d13")
	return module
}

// add all the modules to the modules table of the mockDB
func prefillModules(p Provider) {
	module_1 := Module_1()
	module_2_1 := Module_2_1()
	module_2_2 := Module_2_2()
	p.DB().Table("modules").Create(&module_1)
	p.DB().Table("modules").Create(&module_2_1)
	p.DB().Table("modules").Create(&module_2_2)
}

// Role that will be added to the mock DB
func Role_1() models.Role {
	var role models.Role
	role.Description = "Role Description 1"
	role.ID = 1
	return role
}

// Role that will be added to the mock DB
func Role_2() models.Role {
	var role models.Role
	role.Description = "Role Description 2"
	role.ID = 2
	return role
}

// add all the roles to the roles table of the mockDB
func prefillRole(p Provider) {
	role_1 := Role_1()
	role_2 := Role_2()
	p.DB().Table("roles").Create(&role_1)
	p.DB().Table("roles").Create(&role_2)
}

// SelectedCourse that will be added to the mock DB
func SelectedCourse_1() models.SelectedCourse {
	var selectedCourse models.SelectedCourse
	selectedCourse.UserID = 1
	selectedCourse.CourseID = 1
	selectedCourse.CourseVersion = 1
	selectedCourse.ClassStartyear = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	selectedCourse.Dispensed = false
	return selectedCourse
}

// SelectedCourse that will be added to the mock DB
func SelectedCourse_2() models.SelectedCourse {
	var selectedCourse models.SelectedCourse
	selectedCourse.UserID = 2
	selectedCourse.CourseID = 2
	selectedCourse.CourseVersion = 2
	selectedCourse.ClassStartyear = time.Date(2024, time.February, 2, 0, 0, 0, 0, time.UTC)
	selectedCourse.Dispensed = false
	return selectedCourse
}

// add all the selectedCourses to the selectedCourses table of the mockDB
func prefillSelectedCourse(p Provider) {
	selectedCourse_1 := SelectedCourse_1()
	selectedCourse_2 := SelectedCourse_2()
	p.DB().Table("selected_courses").Create(&selectedCourse_1)
	p.DB().Table("selected_courses").Create(&selectedCourse_2)
}

// State that will be added to the mock DB
func State_1() models.State {
	var state models.State
	state.Description = "State Description 1"
	state.ID = 1
	return state
}

// State that will be added to the mock DB
func State_2() models.State {
	var state models.State
	state.Description = "State Description 2"
	state.ID = 2
	return state
}

// add all the states to the states table of the mockDB
func prefillState(p Provider) {
	state_1 := State_1()
	state_2 := State_2()
	p.DB().Table("states").Create(&state_1)
	p.DB().Table("states").Create(&state_2)
}

// StudyStage that will be added to the mock DB
func StudyStage_1() models.StudyStage {
	var studyStage models.StudyStage
	studyStage.Description = "StudyStage Description 1"
	studyStage.ID = 1
	return studyStage
}

// StudyStage that will be added to the mock DB
func StudyStage_2() models.StudyStage {
	var studyStage models.StudyStage
	studyStage.Description = "StudyStage Description 2"
	studyStage.ID = 2
	return studyStage
}

// add all the studyStages to the studyStages table of the mockDB
func prefillStudyStage(p Provider) {
	studyStage_1 := StudyStage_1()
	studyStage_2 := StudyStage_2()
	p.DB().Table("study_stages").Create(&studyStage_1)
	p.DB().Table("study_stages").Create(&studyStage_2)
}

// add all the users to the users table of the mockDB
func prefillUser(p Provider) {
	p.DB().Table("users").Create(&models.User{
		Name:  "Jakob Ferber",
		Email: "jakob.ferber@hftm.ch",
	})
	p.DB().Table("users").Create(&models.User{
		Name:  "Rafael Stauffer",
		Email: "rafael.stauffer@hftm.ch",
	})
}
