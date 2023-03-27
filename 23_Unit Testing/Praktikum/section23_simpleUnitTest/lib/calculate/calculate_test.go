package calculate

import "testing"

func TestAddition(t *testing.T) {
	// result := Addition(5, 5)
	// if result != 10 {
	// 	t.Error("Expected 10, got ", result)
	// }

	if Addition(5, 5) != 10 {
		t.Error("Expected 5 (+) 5 to equal 10")
	}

	if Addition(-5, -5) != -10 {
		t.Error("Expected -5 (+) -5 to equal -10")
	}

	if Substraction(-5, 5) != -10 {
		t.Error("Expected -5 (-) 5 to equal 0")
	}

	if Multiplication(-5, 5) != -25 {
		t.Error("Expected -5 (*) 5 to equal -25")
	}

	if Division(-5, 5) != -1 {
		t.Error("Expected -5 (/) 5 to equal -1")
	}

}
