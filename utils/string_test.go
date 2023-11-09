package utils

import (
	"testing"
)

func TestMd5(t *testing.T) {
	cases := map[string]string{
		"hello": "5d41402abc4b2a76b9719d911017c592",
		"world": "7d793037a0760186574b0282f2f435e7",
	}
	for k, v := range cases {
		if Md5(k) != v {
			t.Errorf("md5(%s) should be %s", k, v)
		}
	}
}
