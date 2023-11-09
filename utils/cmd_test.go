package utils

import "testing"

func TestMatchPlayerName(t *testing.T) {
	cases := map[string]string{
		".apex u":              "",
		".apex u apex":         "apex",
		".apex u apex 123":     "apex 123",
		".apex    u   apex   ": "apex",
	}

	for text, expected := range cases {
		actual := MatchPlayerName(text)
		if actual != expected {
			t.Errorf("MatchPlayerName(%s) == %s, expected %s", text, actual, expected)
		}
	}
}

func TestMatchMapName(t *testing.T) {
	cases := map[string]bool{
		".apex m":         true,
		"。apex m":         true,
		".apex  m":        true,
		"。apex    m   ":   true,
		".apex ma":        false,
		".apex m 1":       false,
		"   .apex   m   ": true,
	}

	for text, expected := range cases {
		actual := MatchIsMap(text)
		if actual != expected {
			t.Errorf("MatchMapName(%s) == %v, expected %v", text, actual, expected)
		}
	}
}
