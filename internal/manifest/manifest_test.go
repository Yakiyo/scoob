package manifest

import (
	"fmt"
	"testing"

	json "github.com/json-iterator/go"
)

func TestParse(t *testing.T) {
	mf := Manifest{}
	err := json.UnmarshalFromString(testManifest, &mf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", mf)
}

var testManifest = `{
	"version": "1.10.14",
	"description": "Apache Ant is a Java library and command-line tool for compiling, assembling, testing and running Java and non-Java applications.",
	"homepage": "https://ant.apache.org/",
	"license": "Apache-2.0",
	"suggest": {
		"JDK": "java/openjdk"
	},
	"url": "https://www.apache.org/dist/ant/binaries/apache-ant-1.10.14-bin.zip",
	"hash": "sha512:6e85cf45726ee88c916976aba07488b79da84be1a31b5c5441a65c28bb4436b5a5a373402c78ac6a3827bd261fb924124bb9534e52d6429162eaf57b9737124c",
	"extract_dir": "apache-ant-1.10.14",
	"pre_install": [
		"# Record built-in libs in builtin_libs.json",
		"$builtin_libs = Get-ChildItem \"$dir\\lib\" | Select -ExpandProperty Name",
		"$builtin_libs | ConvertTo-Json | Out-File \"$dir\\lib\\builtin_libs.json\" -Encoding Ascii",
		"# Copy user libs to $dir",
		"if (Test-Path \"$persist_dir\\lib\") {",
		"    Get-ChildItem \"$persist_dir\\lib\" | Select -ExpandProperty Name | ForEach-Object {",
		"        if (!(Test-Path \"$dir\\lib\\$_\")) { Copy-Item \"$persist_dir\\lib\\$_\" \"$dir\\lib\" }",
		"    }",
		"}"
	],
	"env_add_path": "bin",
	"env_set": {
		"ANT_HOME": "$dir"
	},
	"uninstaller": {
		"script": [
			"# Only persist libs that were added by the user, but not built-in ones",
			"ensure \"$persist_dir\\lib\" | Out-Null",
			"$builtin_libs = Get-Content \"$dir\\lib\\builtin_libs.json\" | ConvertFrom-Json",
			"Get-ChildItem \"$dir\\lib\" -Exclude builtin_libs.json | Select -ExpandProperty Name | ForEach-Object {",
			"    if (!($builtin_libs -contains \"$_\")) { Copy-Item \"$dir\\lib\\$_\" \"$persist_dir\\lib\" }",
			"}"
		]
	},
	"checkver": {
		"url": "https://ant.apache.org/bindownload.cgi",
		"regex": "Currently, Apache Ant (?:[\\d.]+ and )?([\\d.]+) (?:is|are) the best"
	},
	"autoupdate": {
		"url": "https://www.apache.org/dist/ant/binaries/apache-ant-$version-bin.zip",
		"hash": {
			"url": "$url.sha512"
		},
		"extract_dir": "apache-ant-$version"
	}
}`
