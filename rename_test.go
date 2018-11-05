package jsonfieldrename

import "testing"

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

	expectedResult := `{"renamed":"test","renamed":1}`

	result, err := Marshal(entity, simpleRenamer)
	if err != nil {
		t.Error("Error received from marshall")
		t.FailNow()
	}

	if string(result) != expectedResult {
		t.Errorf("Expected %s; Got %s", expectedResult, result)
	}

}
