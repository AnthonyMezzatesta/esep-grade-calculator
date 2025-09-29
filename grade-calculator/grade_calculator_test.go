package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 55, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}

}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	// Check that each GradeType's name returns the correct string per the grade type (Assignment, Exam, Essay)
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment.String() to return 'assignment'; got '%s' instead", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected Exam.String() to return 'exam'; got '%s' instead", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected Essay.String() to return 'essay'; got '%s' instead", Essay.String())
	}
}

func TestEmptyGrades(t *testing.T) {
	// Empty grade should result in a zero or F
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()
	
	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassFail_Pass(t *testing.T) {
	expected_value := "Pass"

	// Create a grade calculator using the new Pass/Fail function
	gradeCalculator := GradeCalculatorPassFail()

	gradeCalculator.AddGrade("assignment 1", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay 1", 78, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassFail_Fail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := GradeCalculatorPassFail()

	gradeCalculator.AddGrade("assignment 1", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay 1", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassFail_Seventy(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := GradeCalculatorPassFail()

	// Should pass with exactly seventy percent
	gradeCalculator.AddGrade("assignment 1", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay 1", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassFail_SixtyNine(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := GradeCalculatorPassFail()

	// Should fail with exactly sixty nine percent 
	gradeCalculator.AddGrade("assignment 1", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 69, Exam)
	gradeCalculator.AddGrade("essay 1", 69, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestLetterGrades(t *testing.T) {
	// Verify that the default constructor still produces letter grades
	expected_value := "C"

	gradeCalculator := NewGradeCalculator() // Default constructor

	gradeCalculator.AddGrade("assignment 1", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay 1", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected default constructor to return letter grade '%s'; got '%s' instead", expected_value, actual_value)
	}
}