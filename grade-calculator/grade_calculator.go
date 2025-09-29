package esepunittests

type GradingMode int // Grading mode type, used in pass/fail implementation

const (
	LetterGrade GradingMode = iota
	PassFail
)

type GradeCalculator struct {
	grades []Grade
	mode GradingMode
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
		mode: LetterGrade, // Set mode to LetterGrade
	}
}

// Function to return if grade reached pass/fail threshold (seventy percent)
func GradeCalculatorPassFail() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode: PassFail, // Set mode to PassFail
	}
}

// Changed function name as it is only used to get the letter grade now
func (gc *GradeCalculator) GetLetterGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	// Letter grade mode
	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

// New function to get pass/fail grade
func (gc *GradeCalculator) GetPassFailGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	if numericalGrade >= 70 { // Seventy percent is the pass threshold according to assignment instructions
		return "Pass"
	}
	return "Fail"
}

// New GetFinalGrade function to determine which grading mode to return
func (gc *GradeCalculator) GetFinalGrade() string {

	if gc.mode == PassFail { // Return pass/fail  if in pass/fail mode
		return gc.GetPassFailGrade()
	}
	 // By default return letter grade
	return gc.GetLetterGrade()
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	// Filter grades by the type
	assignments := filterGradesByType(gc.grades, Assignment)
	exams := filterGradesByType(gc.grades, Exam)
	essays := filterGradesByType(gc.grades, Essay)
	
	// Compute averages for each type
	assignment_average := computeAverage(assignments)
	exam_average := computeAverage(exams)
	essay_average := computeAverage(essays)

	// Apply the weights for each category
	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

// Helper function to filter grades by type
func filterGradesByType(grades []Grade, gradeType GradeType) []Grade {
	var res []Grade
	// Loop through grades and append if it matches the type
	for _, grade := range grades {
		if grade.Type == gradeType {
			res = append(res, grade)
		}
	}
	return res
}

// Fixed computeAverage function
func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}

	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
