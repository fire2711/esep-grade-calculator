package esepunittests

type GradeCalculator struct {
	grades []Grade
	scheme string
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		scheme: "letter",
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	if gc.scheme == "pass or fail"{
		if numericalGrade >= 60 {
		return "Pass"
	}
	return "Fail"
	
	}

	switch {
	case numericalGrade >= 90:
		return "A"
	case numericalGrade >= 80:
		return "B"
	case numericalGrade >= 70:
		return "C"
	case numericalGrade >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAverage := computeWeightedAverage(gc.grades, Assignment)
	examAverage := computeWeightedAverage(gc.grades, Exam)
	essayAverage := computeWeightedAverage(gc.grades, Essay)

	weightedGrade := assignmentAverage*0.5 + examAverage*0.35 + essayAverage*0.15
	return int(weightedGrade)
}

func computeWeightedAverage(grades []Grade, gradeType GradeType) float64 {
	sum := 0
	count := 0
	for _, g := range grades {
		if g.Type == gradeType {
			sum += g.Grade
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
