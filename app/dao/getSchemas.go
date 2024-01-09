package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
	"gorm.io/gorm/clause"
)

func (db *Database) getFocus(focus *schemas.Focus, id uint) error {
	var eFocus entity.Focus

	result := db.Db.Preload("Field").First(&eFocus, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eFocus, &focus)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listFocus(focuses *[]schemas.Focus) error {
	var eFocuses []entity.Focus
	result := db.Db.Preload("Field").Find(&eFocuses)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eFocuses, &focuses)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) getField(field *schemas.Field, id uint) error {
	var efield entity.Field

	result := db.Db.First(&efield, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&efield, &field)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listFields(fields *[]schemas.Field) error {
	var eFields []entity.Field
	result := db.Db.Find(&eFields)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eFields, &fields)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) getModule(module *schemas.Module, id uint) error {
	var eModule entity.Module
	result := db.Db.Preload("State").First(&eModule, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eModule, &module)
	if err != nil {
		return nil
	}

	// TODO: What States are there?
	module.IsActiv = false
	if eModule.State.Description == "active" {
		module.IsActiv = true
	}

	return nil
}

func (db *Database) listModules(modules *[]schemas.Module) error {
	var eModules []entity.Module
	result := db.Db.Preload("State").Find(&eModules)
	if result.Error != nil {
		return result.Error
	}

	for _, eModule := range eModules {
		var module schemas.Module
		err := ParseEntityToSchema(&eModule, &module)
		if err != nil {
			return nil
		}

		// TODO: What States are there?
		module.IsActiv = false
		if eModule.State.Description == "active" {
			module.IsActiv = true
		}

		*modules = append(*modules, module)
	}

	return nil
}

// TODO:
func (db *Database) getModuleFilter(moduleFilter *schemas.ModuleFilter) error {

	return nil
}

func (db *Database) getStudyStage(studyStage *schemas.StudyStage, id uint) error {
	var eStudyStage entity.StudyStage

	result := db.Db.First(&eStudyStage)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eStudyStage, &studyStage)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listStudyStages(studyStage *[]schemas.StudyStage) error {
	var eStudyStages []entity.StudyStage

	result := db.Db.Find(&eStudyStages)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eStudyStages, &studyStage)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) getValuationCategory(valuationCategory *schemas.ValuationCategory, id uint) error {
	var eEvaluationType entity.Evaluationtype

	result := db.Db.First(&eEvaluationType, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eEvaluationType, &valuationCategory)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listValuationCategory(valuationCategorys *[]schemas.ValuationCategory) error {
	var eEvaluationTypes []entity.Evaluationtype

	result := db.Db.Find(&eEvaluationTypes)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eEvaluationTypes, &valuationCategorys)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) getCourseVersions(versions *[]uint, id uint) error {
	result := db.Db.Model(&entity.Course{}).Where("id = ?", id).Pluck("version", &versions)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *Database) getCourseRes(courseRes *schemas.CourseRes, id uint, version uint) error {
	var eCourse entity.Course
	var versions []uint

	err := db.getCourseVersions(&versions, id)
	if err != nil {
		return nil
	}
	courseRes.Versions = versions

	if version == 0 {
		version = findMax(versions)
	}

	result := db.Db.Where("id = ? AND version = ?", id, version).First(&eCourse)
	if result.Error != nil {
		return result.Error
	}

	err = ParseEntityToSchema(&eCourse, &courseRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listCourseRes(coursesRes *[]schemas.CoursesRes) error {
	var eCourses []entity.Course

	result := db.Db.Find(&eCourses)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eCourses, &coursesRes)
	if err != nil {
		return err
	}

	for _, courseRes := range *coursesRes {
		var versions []uint
		err = db.getCourseVersions(&versions, courseRes.Id)
		if err != nil {
			return nil
		}
		courseRes.Versions = versions
	}

	return nil
}

func (db *Database) getTeacher(teacher *schemas.Teacher, id uint) error {
	var eTeacher entity.User
	result := db.Db.First(&eTeacher, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eTeacher, &teacher)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listTeachers(teachers *[]schemas.Teacher) error {
	var eTeachers []entity.User
	var eUsers []entity.User
	result := db.Db.Preload("Roles").Find(eUsers)
	if result.Error != nil {
		return result.Error
	}

	for _, user := range eUsers {
		for _, role := range user.Roles {
			if role.Description == "teacher" {
				eTeachers = append(eTeachers, user)
			}
		}
	}

	err := ParseEntityToSchema(&eTeachers, &teachers)
	if err != nil {
		return nil
	}

	return nil
}

func (db *Database) getFieldManager(fieldmanager *schemas.Fieldmanager, id uint) error {
	var eFieldManager entity.User
	result := db.Db.First(&eFieldManager, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eFieldManager, &fieldmanager)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) listFieldManagers(fieldmanagers *[]schemas.Fieldmanager) error {
	var eFieldManagers []entity.User
	var eUsers []entity.User
	result := db.Db.Preload("Fields").Find(eUsers)
	if result.Error != nil {
		return result.Error
	}

	for _, user := range eUsers {
		if len(user.Fields) > 0 {
			eFieldManagers = append(eFieldManagers, user)
		}
	}

	err := ParseEntityToSchema(&eFieldManagers, &fieldmanagers)
	if err != nil {
		return nil
	}

	return nil
}

func extractRoleDescriptions(eUser entity.User) []string {
	var roles []string
	for _, role := range eUser.Roles {
		roles = append(roles, role.Description)
	}

	return roles
}

func (db *Database) getUser(user *schemas.User, id uint) error {
	var eUser entity.User
	result := db.Db.Preload("Roles").First(&eUser, id)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eUser, &user)
	if err != nil {
		return err
	}

	curriculumStartYear := eUser.ClassStartyear.Year()
	user.CurriculumStartYear = curriculumStartYear
	user.Roles = extractRoleDescriptions(eUser)

	return nil
}

func (db *Database) listUsers(users *[]schemas.User) error {
	var eUsers []entity.User
	result := db.Db.Preload("Roles").Find(&eUsers)
	if result.Error != nil {
		return result.Error
	}

	for _, eUser := range eUsers {
		var user schemas.User
		err := ParseEntityToSchema(&eUser, &user)
		if err != nil {
			return err
		}

		curriculumStartYear := eUser.ClassStartyear.Year()
		user.CurriculumStartYear = curriculumStartYear
		user.Roles = extractRoleDescriptions(eUser)
		*users = append(*users, user)
	}

	return nil
}

func (db *Database) getGradeTypes(gradeTypes *[]schemas.GradType, conversionId uint) error {
	var eConversions []entity.Conversion
	result := db.Db.Preload("Gradetype").First(&eConversions, conversionId)
	if result.Error != nil {
		return result.Error
	}

	for _, eConversion := range eConversions {
		var gradeType schemas.GradType
		gradeType.Description = eConversion.Gradetype.Description
		gradeType.Id = eConversion.GradetypeID
		gradeType.Grade = eConversion.Value
		*gradeTypes = append(*gradeTypes, gradeType)
	}

	return nil
}

func (db *Database) getConversionId(examEvaluationId uint) (error, uint) {
	var conversionId uint
	result := db.Db.Model(entity.Conversion{}).Select("id").Where("ee_exam_id = ?", examEvaluationId).First(&conversionId)
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, conversionId
}

func (db *Database) getGradeRes(gradRes *schemas.GradeRes, userId uint, courseVersion uint, courseId uint) error {
	var eExamEvaluation entity.ExamEvaluation
	result := db.Db.Where("selected_course_user_id = ? AND selected_course_course_version = ? AND selected_course_course_id = ?", userId, courseVersion, courseId).First(&eExamEvaluation)
	if result.Error != nil {
		return result.Error
	}

	// Get conversionId
	err, conversionId := db.getConversionId(eExamEvaluation.ID)
	if err != nil {
		return err
	}

	var gradeTypes []schemas.GradType
	err = db.getGradeTypes(&gradeTypes, conversionId)
	if err != nil {
		return err
	}

	gradRes.Id = eExamEvaluation.ID
	gradRes.Date = eExamEvaluation.EntryDate.String()
	gradRes.GradesPerType = gradeTypes

	return nil
}

func (db *Database) getExamRes(examRes *schemas.ExamRes, examId uint) error {
	var eExam entity.Exam
	result := db.Db.Preload("ExamType").First(&eExam, examId)
	if result.Error != nil {
		return result.Error
	}

	err := ParseEntityToSchema(&eExam, &examRes)
	if err != nil {
		return err
	}

	examRes.Type = eExam.Examtype.Description

	return nil
}

func (db *Database) ListExamRes(examsRes *[]schemas.ExamRes) error {
	var eExams []entity.Exam
	result := db.Db.Preload("ExamType").Find(&eExams)
	if result.Error != nil {
		return result.Error
	}

	for _, eExam := range eExams {
		var examRes schemas.ExamRes
		err := ParseEntityToSchema(&eExam, &examRes)
		if err != nil {
			return err
		}
		examRes.Type = eExam.Examtype.Description
		*examsRes = append(*examsRes, examRes)
	}

	return nil
}

func (db *Database) getCourseExamStudent(courseExamStudent *schemas.CourseExamStudent, examId uint, userId uint, courseVersion uint, courseId uint) error {
	var examRes schemas.ExamRes
	err := db.getExamRes(&examRes, examId)
	if err != nil {
		return err
	}

	var gradeRes schemas.GradeRes
	err = db.getGradeRes(&gradeRes, userId, courseVersion, courseId)
	if err != nil {
		return err
	}

	courseExamStudent.Id = examRes.Id
	courseExamStudent.Description = examRes.Description
	courseExamStudent.Weight = examRes.Weight
	courseExamStudent.Type = examRes.Type
	courseExamStudent.Grade = gradeRes

	return nil
}

func (db *Database) getAvgGradeTypes(gradeTypes *[]schemas.GradType, examId uint) error {
	// Search for Exam with id
	var eExam entity.Exam
	result := db.Db.Preload(clause.Associations).First(&eExam, examId)
	if result.Error != nil {
		return result.Error
	}

	// Use course_id and course_version from Exam to get each Exam_evalution
	var eExamEvaluations []entity.ExamEvaluation
	result = db.Db.Where("selected_course_course_id = ? AND selected_course_course_version = ?", eExam.Course.ID, eExam.Course.Version).Find(&eExamEvaluations)

	// Get all Conversion
	sumsMap := make(map[uint]float64)
	countMap := make(map[uint]uint)
	description := make(map[uint]string)

	for _, eExamEvaluation := range eExamEvaluations {
		err, conversionId := db.getConversionId(eExamEvaluation.ID)
		if err != nil {
			return err
		}

		// Calculate averge per gradeType_id
		var conversion entity.Conversion
		result = db.Db.Preload(clause.Associations).First(&conversion, conversionId)

		sumsMap[conversion.GradetypeID] += conversion.Value
		countMap[conversion.GradetypeID] += 1
		description[conversion.GradetypeID] = conversion.Gradetype.Description
	}

	for key, _ := range sumsMap {
		var gradeType schemas.GradType
		gradeType.Grade = sumsMap[key] / float64(countMap[key])
		gradeType.Id = key
		gradeType.Description = description[key]
		*gradeTypes = append(*gradeTypes, gradeType)
	}

	return nil
}

func (db *Database) getCourseExamTeacher(courseExamTeacher *schemas.CourseExamTeacher, examId uint) error {
	var examRes schemas.ExamRes
	err := db.getExamRes(&examRes, examId)
	if err != nil {
		return err
	}

	var gradeTypes []schemas.GradType
	err = db.getAvgGradeTypes(&gradeTypes, examId)
	if err != nil {
		return err
	}

	courseExamTeacher.Id = examRes.Id
	courseExamTeacher.Description = examRes.Description
	courseExamTeacher.Type = examRes.Type
	courseExamTeacher.Weight = examRes.Weight
	courseExamTeacher.AvgGrades = gradeTypes

	return nil
}

func (db *Database) getCourseFilter(courseFilter *schemas.CourseFilter) error {
	var modules []schemas.Module
	var teachers []schemas.Teacher

	err := db.listModules(&modules)
	if err != nil {
		return err
	}

	err = db.listTeachers(&teachers)
	if err != nil {
		return err
	}

	courseFilter.Modules = modules
	courseFilter.Teachers = teachers

	return nil
}

func (db *Database) listExams(exams *[]entity.Exam, courseId uint, courseVersion uint) error {
	result := db.Db.Where("course_id = ? AND course_version = ?").Find(&exams)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *Database) listCourses(courses *entity.Course, courseId uint, courseVersion uint) error {
	result := db.Db.Where("id = ? AND version = ?", courseId, courseVersion).First(&courses)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (db *Database) getCourseResTeacher(courseResTeacher *schemas.CourseResTeacher, courseId uint, version uint) error {
	var eCourse entity.Course
	var eExams []entity.Exam
	var courseExamTeachers []schemas.CourseExamTeacher

	err := db.listCourses(&eCourse, courseId, version)
	if err != nil {
		return err
	}

	err = db.listExams(&eExams, courseId, version)
	if err != nil {
		return nil
	}

	for _, exam := range eExams {
		var courseExamTeacher schemas.CourseExamTeacher

		err = db.getCourseExamTeacher(&courseExamTeacher, exam.ID)
		if err != nil {
			return err
		}

		courseExamTeachers = append(courseExamTeachers, courseExamTeacher)
	}

	courseResTeacher.Id = eCourse.ID
	courseResTeacher.Version = eCourse.Version
	courseResTeacher.Description = eCourse.Description
	courseResTeacher.Number = eCourse.Number
	courseResTeacher.Exams = courseExamTeachers

	return nil
}

func (db *Database) getCourseResStudent(courseResStudent *schemas.CourseResStudent, courseId uint, version uint, userId uint) error {
	var eCourse entity.Course
	var eExams []entity.Exam
	var courseExamStudents []schemas.CourseExamStudent

	err := db.listCourses(&eCourse, courseId, version)
	if err != nil {
		return err
	}

	err = db.listExams(&eExams, courseId, version)
	if err != nil {
		return err
	}

	for _, exam := range eExams {
		var courseExamStudent schemas.CourseExamStudent

		err = db.getCourseExamStudent(&courseExamStudent, exam.ID, userId, version, courseId)
		if err != nil {
			return nil
		}

		courseExamStudents = append(courseExamStudents, courseExamStudent)
	}

	courseResStudent.Id = eCourse.ID
	courseResStudent.Version = eCourse.Version
	courseResStudent.Description = eCourse.Description
	courseResStudent.Number = eCourse.Number
	courseResStudent.Exams = courseExamStudents

	return nil
}
