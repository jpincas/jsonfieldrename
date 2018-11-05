package jsonfieldrename

import (
	"strings"
	"testing"
)

type TestEntity struct {
	MySimpleKey string
	MyIntField  int
}

func simpleRenamer(in string) string {
	return "renamed"
}

func Test_All(t *testing.T) {
	entity := TestEntity{
		MySimpleKey: "test",
		MyIntField:  1,
	}

	cases := []struct{
		rename func(string) string
		expected string
	}{
		{simpleRenamer, `{"renamed":"test","renamed":1}`},
		{strings.ToUpper, `{"MYSIMPLEKEY":"test","MYINTFIELD":1}`},
	}

	for i, tc := range cases {
		result, err := Marshal(entity, tc.rename)
		if err != nil {
			t.Errorf("(%d) Error received from marshall", i)
			t.FailNow()
		}
		if string(result) != tc.expected {
			t.Errorf("(%d) Expected %s; Got %s", i, tc.expected, result)
		}	
	}
}
