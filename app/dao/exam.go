package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (db *Database) ListExams(examsRes *[]schemas.ExamRes) error {
	var eExams []entity.Exam
	result := db.Db.Find(&eExams)
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

// creates examn from req, returns entity if successful
func (db *Database) CreateExam(examReq *schemas.ExamReq) (*entity.Exam, error) {
	exam := &entity.Exam{}

	err := ParseSchemaToEntity(examReq, exam)
	if err != nil {
		return nil, err
	}

	err = db.GetCourseEntity(&exam.Course, exam.CourseID, exam.CourseVersion)
	if err != nil {
		return nil, err
	}

	err = db.Db.Create(exam).Error
	if err != nil {
		return nil, err
	}

	return exam, nil
}

func (db *Database) GetExam(examRes *schemas.ExamRes, id uint) error {
	var eExam entity.Exam

	result := db.Db.First(&eExam, id)
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

func (db *Database) EditExam(exam *entity.Exam, id uint) {
	exam.Basemodel.ID = id
	db.Db.Save(&exam)
}
