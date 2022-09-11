package utils_test

import (
	"testing"

	"github.com/MRickers/Erbhoflauf/utils"
)

func TestDeserialize(t *testing.T) {
	type Test struct {
		A string
		B int
	}

	test, err := utils.Deserialize[Test]("{\"A\":\"Hallo\", \"B\":5}")

	if err != nil {
		t.Fatalf("Conversion failed: %s", err)
	}

	if test.A != "Hallo" {
		t.Fatalf("%s != Hallo", test.A)
	}

	if test.B != 5 {
		t.Fatalf("%d != 5", test.B)
	}
}

func TestSerialize(t *testing.T) {
	type Test struct {
		A string
		B int
	}

	var test = Test{
		A: "Hello There",
		B: 5000,
	}
	serialized, err := utils.Serialize[Test](test)

	if err != nil {
		t.Fatalf("Conversion failed: %s", err)
	}

	if serialized != "{\"A\":\"Hello There\",\"B\":5000}" {
		t.Fatalf("%s != {\"A\":\"Hello There\",\"B\":5000}", serialized)
	}
}
