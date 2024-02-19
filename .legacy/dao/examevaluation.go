package dao

import (
	"fmt"
	"time"

	"github.com/SmashGrade/backend/legacy/api/v1/schemas"
	"github.com/SmashGrade/backend/legacy/entity"
)

// creates exam evaluation from req, returns entity if successful
func (db *Database) CreateExamEvaluation(examId, examVersion uint, startYear time.Time, examReq *schemas.ExamReqStudent) (*entity.ExamEvaluation, error) {
	var exam entity.Exam

	// get the exam to link
	// err := ParseSchemaToEntity(examReq, exam)
	err := db.Db.First(&exam, examId).Error
	if err != nil {
		return nil, err
	}

	// get the student to link
	student, err := db.GetUserEntityById(examReq.StudentRef)
	if err != nil {
		return nil, err
	}

	// identify the selected course to link to
	var course *entity.SelectedCourse
	for _, selCourse := range student.SelectedCourses {
		// we need a matching course start year, courseID and courseVersion
		if selCourse.ClassStartyear == startYear && selCourse.CourseID == exam.CourseID && selCourse.CourseVersion == exam.CourseVersion {
			course = &selCourse
			break
		}
	}
	if course == nil {
		return nil, fmt.Errorf("no course with id '%v' version '%v' and start year '%v' found", exam.CourseID, exam.CourseVersion, startYear)
	}

	// cast data from schema to entity
	var examEvaluation *entity.ExamEvaluation
	err = ParseEntityToSchema(examReq, examEvaluation)
	if err != nil {
		return nil, err
	}

	examEvaluation.Exam = exam
	examEvaluation.SelectedCourse = *course

	err = db.Db.Create(examEvaluation).Error
	if err != nil {
		return nil, err
	}

	return examEvaluation, nil
}
