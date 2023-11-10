package webhook

import "testing"

func TestParseCmd(t *testing.T) {
	cases := map[string]string{
		"a":                   "",
		".apex   u":           CmdUser,
		".apex  u uu123":      CmdUser,
		".apex m":             CmdMap,
		".apex s ":            CmdShop,
		" .apex c":            CmdCraft,
		".apex t":             CmdTime,
		".apex h":             CmdHelp,
		".apex d":             CmdDist,
		".apex p":             CmdPick,
		" 。apex  u":           CmdUser,
		" 。apex  u user123  ": CmdUser,
		"。apex m ":            CmdMap,
		"。apex  s":            CmdShop,
		"。apex c ":            CmdCraft,
		"。apex t ":            CmdTime,
		"。apex h":             CmdHelp,
		" 。apex d ":           CmdDist,
		"。apex p ":            CmdPick,
		".apex":               CmdHelp,
		"。apex":               CmdHelp,
		".apex aaaa":          CmdHelp,
		".apex z":             CmdHelp,
		".apex    ":           CmdHelp,
	}
	for text, expected := range cases {
		actual := ParseCmd(text)
		if actual != expected {
			t.Errorf("ParseCmd(%s) == %s, expected %s", text, actual, expected)
		}
	}
}
