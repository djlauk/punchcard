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
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Project is a container for WorkLogEntry objects.
// I.e. every WorkLogEntry belongs to a Project.
type Project struct {
	Name        string
	Description string
	Reference   string `yaml:",omitempty"`
	Closed      bool   `yaml:",omitempty"`
}

// ProjectMap maps the Project.Name to the Project
type ProjectMap map[string]Project

// WorkLogEntry holds the details of a single entry in the work log.
type WorkLogEntry struct {
	Start     time.Time
	End       time.Time
	Project   string
	Message   string
	Reference string `yaml:",omitempty"`
}

// WorkLog is a collection of WorkLogEntry items
type WorkLog []WorkLogEntry

// PunchcardData is the overall type for data stored
type PunchcardData struct {
	Projects ProjectMap
	Entries  WorkLog
	Current  *WorkLogEntry `yaml:",omitempty"`
}

// AddProject adds a Project to the data
func (pcd *PunchcardData) AddProject(p *Project) error {
	if _, ok := pcd.Projects[p.Name]; ok {
		return fmt.Errorf("Project name already in use: %s", p.Name)
	}
	pcd.Projects[p.Name] = *p
	return nil
}

// RenameProject renames a Project and the entries
func (pcd *PunchcardData) RenameProject(oldName string, newName string) error {
	if oldName == newName {
		return fmt.Errorf("old and new name are the same")
	}
	project, ok := pcd.Projects[oldName]
	if !ok {
		return fmt.Errorf("Project name not found: %s", oldName)
	}
	if _, ok := pcd.Projects[newName]; ok {
		return fmt.Errorf("Project name already in use: %s", newName)
	}
	for i := range pcd.Entries {
		e := &pcd.Entries[i]
		if e.Project != oldName {
			continue
		}
		e.Project = newName
	}
	project.Name = newName
	delete(pcd.Projects, oldName)
	pcd.Projects[newName] = project
	return nil
}

// AddEntry adds a WorkLogEntry to the data
func (pcd *PunchcardData) AddEntry(entry *WorkLogEntry) error {
	if err := CheckEntry(entry, pcd.Projects); err != nil {
		return err
	}
	pcd.Entries = append(pcd.Entries, *entry)
	return nil
}

// SetCurrent adds a WorkLogEntry as the current
func (pcd *PunchcardData) SetCurrent(entry *WorkLogEntry) error {
	if err := CheckEntry(entry, pcd.Projects); err != nil {
		return err
	}
	if pcd.Current != nil {
		return fmt.Errorf("Already working on something")
	}
	pcd.Current = entry
	return nil
}

// FinishCurrent finishes the Current and adds it to Entries
func (pcd *PunchcardData) FinishCurrent(t time.Time) error {
	if pcd.Current == nil {
		return fmt.Errorf("Not working on anything")
	}
	pcd.Current.End = t
	if err := pcd.AddEntry(pcd.Current); err != nil {
		return err
	}

	pcd.Current = nil
	return nil
}

// GetLastEntry returns the last WorkLogEntry in the list
func (pcd *PunchcardData) GetLastEntry() *WorkLogEntry {
	if len(pcd.Entries) == 0 {
		return nil
	}
	return &pcd.Entries[len(pcd.Entries)-1]
}

// GetLastEntryForProject returns the last WorkLogEntry in the list for project
func (pcd *PunchcardData) GetLastEntryForProject(name string) *WorkLogEntry {
	for i := len(pcd.Entries) - 1; i >= 0; i-- {
		if pcd.Entries[i].Project == name {
			return &pcd.Entries[i]
		}
	}
	return nil
}

func newPunchcardData() *PunchcardData {
	x := PunchcardData{
		Projects: make(ProjectMap),
		Entries:  make(WorkLog, 0),
		Current:  nil,
	}
	return &x
}

// ReadData reads all the data from the data file
func ReadData(path string) (*PunchcardData, error) {
	punchcardData := newPunchcardData()
	file, err := os.Open(path)
	if err != nil {
		return punchcardData, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	err = d.Decode(punchcardData)
	return punchcardData, err
}

// WriteData reads all the data from the data file
func WriteData(path string, pcd *PunchcardData) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decode
	encoder := yaml.NewEncoder(file)

	// Start YAML decoding from file
	if err := encoder.Encode(pcd); err != nil {
		return err
	}

	return nil
}

// CheckEntry validates a WorkLogEntry
func CheckEntry(w *WorkLogEntry, projects ProjectMap) error {
	p, ok := projects[w.Project]
	if !ok {
		return fmt.Errorf("Project not found: %s", w.Project)
	}
	if p.Closed {
		return fmt.Errorf("Project is closed: %s", w.Project)
	}
	if w.Message == "" {
		return fmt.Errorf("Message must not be empty")
	}
	if !w.End.IsZero() && !w.Start.Before(w.End) {
		return fmt.Errorf("Start must be before End (or End is Zero)")
	}
	return nil
}
