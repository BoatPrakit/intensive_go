package nationalcard

import (
	"testing"
)

func TestValidateNationalCard(t *testing.T) {
	testcases := []struct {
		given    string
		expected string
	}{
		{
			given:    "123456789012",
			expected: "id digits incorrect",
		},
		{
			given:    "1234567890121",
			expected: "",
		},
		{
			given:    "1234567890122",
			expected: "id incorrect",
		},
		{
			given:    "1111111111111",
			expected: "id incorrect",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidateThaiID(testcase.given)
			if actual == nil && testcase.expected == "" {
				println("passed")
				return
			}

			if testcase.expected != actual.Error() {
				t.Errorf("given a id %s expected error is %v but actual error is %v", testcase.given, testcase.expected, actual)
			}
		})
	}
}
