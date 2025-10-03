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

func TestGetGradeABoundary(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

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

func TestGetGradeBBoundary(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 74, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 77, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeCBoundary(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 63, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeDBoundary(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
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
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeFBoundary(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 59, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeTypes(t *testing.T) {
	expected_assignment_string := "assignment"
	expected_exam_string := "exam"
	expected_essay_string := "essay"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("team github assignment", 100, Assignment)
	gradeCalculator.AddGrade("essay on open source contributions", 100, Essay)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("waterfall vs agile assignment", 100, Assignment)
	gradeCalculator.AddGrade("team scrum assignment", 100, Assignment)
	gradeCalculator.AddGrade("essay on sprint retrospectives", 100, Essay)
	gradeCalculator.AddGrade("exam 2", 100, Exam)
	gradeCalculator.AddGrade("algorithms assignment", 100, Assignment)
	gradeCalculator.AddGrade("ai building assignment", 100, Assignment)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)
	gradeCalculator.AddGrade("exam 3", 100, Exam)

	for _, grade := range gradeCalculator.assignments {
		if grade.Type.String() != expected_assignment_string {
			t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_assignment_string, grade.Type.String())
		}
	}

	for _, grade := range gradeCalculator.exams {
		if grade.Type.String() != expected_exam_string {
			t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_exam_string, grade.Type.String())
		}
	}

	for _, grade := range gradeCalculator.essays {
		if grade.Type.String() != expected_essay_string {
			t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_essay_string, grade.Type.String())
		}
	}
}
