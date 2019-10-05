package stringutil

import "testing"

func TestToUpper(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello, world", "HELLO, WORLD"},
		{"HeLlO, WorLd", "HELLO, WORLD"},
		{"", ""},
		{"some_env_var", "SOME_ENV_VAR"},
	}
	for _, c := range cases {
		got := ToUpper(c.in)
		if got != c.want {
			t.Errorf("ToUpper(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestToLower(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello, world", "hello, world"},
		{"HeLlO, WorLd", "hello, world"},
		{"", ""},
		{"THISISNOTANENVVAR", "thisisnotanenvvar"},
	}
	for _, c := range cases {
		got := ToLower(c.in)
		if got != c.want {
			t.Errorf("ToLower(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestCompareThese(t *testing.T) {
	cases := []struct {
		left, right, want string
	}{
		{"left", "four", "left and four are equal in length of characters.\n"},
		{"thirteen", "twenty", "thirteen has more characters than twenty.\n"},
		{"THISISNOT", "thisisnotanenvvar", "THISISNOT has fewer characters than thisisnotanenvvar.\n"},
	}
	for _, c := range cases {
		got := CompareThese(c.left, c.right)
		if got != c.want {
			t.Errorf("CompareThese(%s, %s) == %s, want %s", c.left, c.right, got, c.want)
		}
	}
}

func TestConcatThese(t *testing.T) {
	cases := []struct {
		first, second, want string
	}{
		{"lemon", "cat", "lemoncat"},
		{"red ", "square", "red square"},
		{" Some Thing", "bad", " Some Thingbad"},
	}
	for _, c := range cases {
		got := ConcatThese(c.first, c.second)
		if got != c.want {
			t.Errorf("ConcatThese(%s, %s) == %s, want %s", c.first, c.second, got, c.want)
		}
	}
}
