/*
Copyright © 2020 Daniel J. Lauk <daniel.lauk@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package data

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// Test fixtures
var expectedProjects = map[string]Project{
	"admin": Project{
		Name:        "admin",
		Description: "Administrative work",
		Reference:   "",
		Closed:      false,
	},
	"skynet": Project{
		Name:        "skynet",
		Description: "Ultimately powerful AI, what could go wrong?!",
		Reference:   "https://www.example.com/",
		Closed:      true,
	},
}

var expectedWorkLogEntries = []WorkLogEntry{
	WorkLogEntry{
		Project:   "skynet",
		Message:   "Design emergency shutdown and override procedures",
		Reference: "task#1",
		Start:     time.Date(1997, 8, 3, 9, 30, 0, 0, time.Local),
		End:       time.Date(1997, 8, 3, 11, 30, 0, 0, time.Local),
	},
	WorkLogEntry{
		Project: "admin",
		Message: "Write memo that we postpone it as it's too hard. What could go wrong anyway?!",
		Start:   time.Date(1997, 8, 3, 11, 30, 0, 0, time.Local),
		End:     time.Date(1997, 8, 3, 11, 45, 0, 0, time.Local),
	},
	WorkLogEntry{
		Project: "skynet",
		Message: "Start up the whole thing. What could possibly go wrong after all?",
		Start:   time.Date(1997, 8, 4, 8, 30, 0, 0, time.Local),
		End:     time.Date(1997, 8, 4, 12, 0, 0, 0, time.Local),
	},
}

var testData = PunchcardData{
	Projects: expectedProjects,
	Entries:  expectedWorkLogEntries,
}

var testYaml = `# test yaml
projects:
  admin:
    name: admin
    description: Administrative work
  skynet:
    name: skynet
    description: Ultimately powerful AI, what could go wrong?!
    reference: https://www.example.com/
    closed: true
entries:
  - start: 1997-08-03T09:30:00+02:00
    end: 1997-08-03T11:30:00+02:00
    project: skynet
    message: Design emergency shutdown and override procedures
    reference: task#1
  - start: 1997-08-03T11:30:00+02:00
    end: 1997-08-03T11:45:00+02:00
    project: admin
    message: Write memo that we postpone it as it's too hard. What could go wrong anyway?!
  - start: 1997-08-04T08:30:00+02:00
    end: 1997-08-04T12:00:00+02:00
    project: skynet
    message: Start up the whole thing. What could possibly go wrong after all?
`

func TestReadData(t *testing.T) {
	file, err := ioutil.TempFile("", "punchcard.data.*.yaml")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())

	if _, err := file.WriteString(testYaml); err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}

	pcd, err := ReadData(file.Name())
	if err != nil {
		t.Error(err)
	}
	if len(pcd.Projects) != len(expectedProjects) {
		t.Errorf("len(pcd.Projects) incorrect, got %v, want %v", len(pcd.Projects), len(expectedProjects))
	}
	if len(pcd.Entries) != len(expectedWorkLogEntries) {
		t.Errorf("len(pcd.Entries) incorrect, got %v, want %v", len(pcd.Entries), len(expectedWorkLogEntries))
	}
	for i := 0; i < len(pcd.Entries); i++ {
		if pcd.Entries[i] != expectedWorkLogEntries[i] {
			t.Errorf("pcd.Entries[%d] incorrect, got %v, want %v", i, pcd.Entries[i], expectedWorkLogEntries[i])
		}
	}
	for key := range pcd.Projects {
		if pcd.Projects[key] != expectedProjects[key] {
			t.Errorf("pcd.Projects[%s] incorrect, got %v, want %v", key, pcd.Projects[key], expectedProjects[key])
		}
	}
}

func TestReadDataReturnsEmptyData(t *testing.T) {
	pcd, err := ReadData("/nosuchfile.projects.yaml")
	if err == nil {
		t.Error("Expected error")
	}
	if !os.IsNotExist(err) {
		t.Errorf("err is incorrect, got %v, want %v", err, "NotExist error")
	}
	if len(pcd.Entries) != 0 {
		t.Errorf("len(pcd.Entries) incorrect, got %v, want %v", len(pcd.Entries), 0)
	}
	if len(pcd.Projects) != 0 {
		t.Errorf("len(pcd.Projects) incorrect, got %v, want %v", len(pcd.Projects), 0)
	}
}

func TestWriteDataWorks(t *testing.T) {
	testFile, err := ioutil.TempFile("", "punchcard.data.*.yaml")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(testFile.Name())

	WriteData(testFile.Name(), &testData)
	file, err := os.Open(testFile.Name())
	if err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}
}

func TestWriteDataReadDataRoundTrip(t *testing.T) {
	testFile, err := ioutil.TempFile("", "punchcard.data.*.yaml")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(testFile.Name())

	if err = WriteData(testFile.Name(), &testData); err != nil {
		t.Error(err)
	}
	pcd, err := ReadData(testFile.Name())
	if err != nil {
		t.Error(err)
	}
	for key := range pcd.Projects {
		prj := pcd.Projects[key]
		expected := expectedProjects[key]
		if prj != expected {
			t.Errorf("pcd.Projects[%s] incorrect, got %v, want %v", key, prj, expected)
		}
	}
	for i := 0; i < len(pcd.Entries); i++ {
		e := pcd.Entries[i]
		expected := expectedWorkLogEntries[i]
		if e != expected {
			t.Errorf("pcd.Entries[%d] incorrect, got %v, want %v", i, e, expected)
		}
	}
}

func TestCheckEntryWorksWithGoodData(t *testing.T) {
	w := expectedWorkLogEntries[1]
	if err := CheckEntry(&w, expectedProjects); err != nil {
		t.Error(err)
	}
}

func TestCheckEntryFailsProjectNotFound(t *testing.T) {
	w := expectedWorkLogEntries[0]
	w.Project = "no-such-project"
	if err := CheckEntry(&w, expectedProjects); err == nil {
		t.Error("Expected error")
	}
}
func TestCheckEntryFailsProjectClosed(t *testing.T) {
	w := expectedWorkLogEntries[0]
	if err := CheckEntry(&w, expectedProjects); err == nil {
		t.Error("Expected error")
	}
}

func TestCheckEntryFailsEmptyMessage(t *testing.T) {
	w := expectedWorkLogEntries[0]
	w.Message = ""
	if err := CheckEntry(&w, expectedProjects); err == nil {
		t.Error("Expected error")
	}
}

func TestCheckEntryFailsStartAfterEnd(t *testing.T) {
	w := expectedWorkLogEntries[0]
	w.End = time.Now().Add(-5 * time.Minute)
	w.Start = time.Now()
	if err := CheckEntry(&w, expectedProjects); err == nil {
		t.Error("Expected error")
	}
}

func TestCheckEntryFailsStartEqualsEnd(t *testing.T) {
	w := expectedWorkLogEntries[0]
	w.End = time.Now()
	w.Start = w.End
	if err := CheckEntry(&w, expectedProjects); err == nil {
		t.Error("Expected error")
	}
}
