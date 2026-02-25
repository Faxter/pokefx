package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  white\tspace   ",
			expected: []string{"white", "space"},
		},
		{
			input:    "Some dIfFeRenT CASES",
			expected: []string{"some", "different", "cases"},
		},
		{
			input:    "Battle-Armor",
			expected: []string{"battle-armor"},
		},
	}

	for _, testCase := range cases {
		actual := CleanInput(testCase.input)
		if len(testCase.expected) != len(actual) {
			t.Errorf("expected %d words, but got %d", len(testCase.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := testCase.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %s, got: %s", expectedWord, word)
			}
		}
	}
}
