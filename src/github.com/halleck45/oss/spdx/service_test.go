package spdx

import(
	"testing"
	"os"
)


func TestUpdate (t *testing.T) {
	filename := os.TempDir() + "/oss-unit-spdx.json"
	service := Service{LicenseFilename: filename}
	service.Update()

	// file is created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("SPDX Update should create json file with list of all licenses")
	}

	// valid contains valid json
	licenses := service.All()
	if 0 == len(licenses) {
		t.Error("SPDX licenses list should not be empty")
	}

	// clean up
	os.Remove(filename)

}
