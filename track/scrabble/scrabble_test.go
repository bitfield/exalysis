package scrabble

import (
	"fmt"
	"testing"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/testhelper"
	"github.com/exercism/exalysis/track/scrabble/tpl"
	"github.com/stretchr/testify/assert"
)

var suggestTests = []struct {
	path       string
	suggestion gtpl.Template
	expected   bool
}{
	{path: "./solutions/2", suggestion: tpl.MoveMap, expected: true},
	{path: "./solutions/3", suggestion: tpl.MoveMap, expected: false},
	{path: "./solutions/4", suggestion: tpl.MoveMap, expected: true},
	{path: "./solutions/2", suggestion: tpl.MapRune, expected: true},
	{path: "./solutions/2", suggestion: tpl.TrySwitch, expected: true},
	{path: "./solutions/5", suggestion: tpl.FlattenMap, expected: true},
	{path: "./solutions/5", suggestion: tpl.TrySwitch, expected: false},
	{path: "./solutions/5", suggestion: tpl.TypeConversion, expected: false},
	{path: "./solutions/6", suggestion: tpl.Regex, expected: true},
	{path: "./solutions/6", suggestion: tpl.Challenge, expected: true},
	{path: "./solutions/7", suggestion: tpl.FlattenMap, expected: false},
	{path: "./solutions/8", suggestion: tpl.SliceRuneConv, expected: true},
	{path: "./solutions/8", suggestion: tpl.TrySwitch, expected: false},
	{path: "./solutions/9", suggestion: tpl.FlattenMap, expected: true},
}

func Test_Suggest(t *testing.T) {
	for _, test := range suggestTests {
		_, pkg, err := testhelper.LoadExample(test.path, "scrabble")
		if err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		Suggest(pkg, r)

		assert.Equal(t, test.expected, r.HasSuggestion(test.suggestion),
			fmt.Sprintf("test failed: %+v", test))
	}
}
