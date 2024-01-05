package dao

import (
	"encoding/json"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (db *Database) ListExams(examsRes *[]schemas.ExamRes) error {
	var exams []entity.Exam
	db.Db.Find(&exams)

	jsonData, err := json.Marshal(exams)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonData), &examsRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) PostExam(exam *entity.Exam) {
	db.Db.Create(&exam)
}

func (db *Database) GetExam(exam *entity.Exam, id uint) {
	db.Db.First(&exam, id)
}

func (db *Database) EditExam(exam *entity.Exam, id uint) {
	exam.Basemodel.ID = id
	db.Db.Save(&exam)
}
