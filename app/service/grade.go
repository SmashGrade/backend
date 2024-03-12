package service

import (
	"strconv"

	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/models"
)

type GradeService struct {
}

func NewGradeService() *GradeService {
	return &GradeService{}
}

// Calculates the average grade of a student
// Based on the given exam evaluations
func (s GradeService) CalculateAverage(exams []models.ExamEvaluation) (float64, error) {
	// Get the number of exams from the list
	numExams := len(exams)
	// Result variable
	result := 0.0
	// Weights of the exams
	var weights uint = 0

	// If there are no exams, we return 0
	if numExams == 0 {
		return 0, nil
	}

	for _, exam := range exams {
		switch exam.OrignialGradeTypeID {
		case config.GRADETYPE_NONE:
			// Deduct the number of exams by 1, because this exam does not count towards the average
			numExams--
			// In the case of a non graded exam, we skip it, because it does not count towards the average
			continue
		case config.GRADETYPE_NOTE:
			// If this grade has a weighting, we add it to the total weighting
			// Default is 1
			weights += exam.Exam.Weighting
			// Grades ranging from 1 to 6, where 1 is the worst and 6 is the best
			grade, err := strconv.ParseFloat(exam.OriginalValue, 64)
			if err != nil {
				return 0, err
			}
			result += grade * float64(exam.Exam.Weighting)
		case config.GRADETYPE_PERCENT:
			// If this grade has a weighting, we add it to the total weighting
			// Default is 1
			weights += exam.Exam.Weighting
			// We know that the original value is a percentage
			percentage, err := strconv.ParseFloat(exam.OriginalValue, 64)
			if err != nil {
				return 0, err
			}
			// Grades ranging from 0 to 100, where 0 is the worst and 100 is the best
			result += s.PercentageToGrade(percentage) * float64(exam.Exam.Weighting)
		}
	}

	// Calculate the average with weights
	result = result / float64(weights)

	// Return the result, with 0 exams the result will be 0
	return result, nil

}

// Converts a percentage to a grade from 1 to 6
// 1 being the worst and 6 being the best
func (s GradeService) PercentageToGrade(percentage float64) float64 {
	// Utilize the percentage to grade conversion formula for swiss schools
	// This formula is used to convert the percentage to a grade from 1 to 6 in 5% steps
	// This gives a very accurate representation of the grade based on the percentage
	return percentage/100*5 + 1
}
