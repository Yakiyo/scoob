package manifest

import (
	"fmt"
	"testing"

	json "github.com/json-iterator/go"
)

func TestParse(t *testing.T) {
	value := `{
		"version": "1.0.0",
		"url": "https://github.com/Yakiyo/scoob/releases/latest/download/blah-blah.zip",
		"bin": ["scoob.exe"]
	}`
	mf := Manifest{}
	err := json.UnmarshalFromString(value, &mf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", mf)
}
