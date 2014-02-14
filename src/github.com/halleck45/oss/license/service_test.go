package license

import(
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)


func TestScenario (t *testing.T) {
	filename := os.TempDir() + "/oss-unit-license.json"
	service := Service{Filename: filename}
	service.Init()

	// file is created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Init should create json file")
	}

	// file can be loaded
	err := service.Load()
	assert.Nil(t, err, "json should be loaded")

	// I can add Asset
	license := License{Identifier: "MIT", Name: "MIT License"}
	asset := Asset{License: license, File: "my-file.txt"}
	service.Add(asset)

	// Added Asset is found
	asset, err = service.Get("my-file.txt")
	assert.Nil(t, err, "Asset should be added to Manifest")

	// clean up
	os.Remove(filename)
}
