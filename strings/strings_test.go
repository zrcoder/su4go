package strings

import "testing"

func TestRemoveTailAfterFlag(t *testing.T) {
	tests := []struct {
		str    string
		flag   string
		expect string
	}{
		{"10.247.0.0/16", ".", "10.247.0."},
		{"10.247.0.0/16", "x", ""},
		{"", ".", ""},
	}
	for _, test := range tests {
		r := RemoveTailAfterLastFlag(test.str, test.flag)
		if r != test.expect {
			t.Error("Test RemoveTailAfterFlag faild.")
		} else {
			t.Log(r)
		}
	}
}

func TestSubStringBetween(t *testing.T) {
	tests := []struct {
		s      string
		s1     string
		s2     string
		expect string
	}{
		{"abc123SSS", "bc", "SS", "123"},
		{"abc123SSS", "bc", "444", ""},
		{"abc123SSS", "bc", "c12", ""},
		{"abc123SSS", "x", "SS", ""},
		{"", "a", "1", ""},
		{"fffabcXYZSSS===", "SSS", "abc", "XYZ"},
		{"fffabcXYZSSS===", "bc", "YZ", "X"},
		{"1223334444", "22", "333", ""},
		{"1223334444", "444", "333", ""},
	}
	for _, test := range tests {
		r := SubStringBetween(test.s, test.s1, test.s2)
		if r != test.expect {
			t.Error("Test SubStringBetween faild.")
		} else {
			t.Log(r)
		}
	}
}
