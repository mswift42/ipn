package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/mswift42/ipn/tv"
)

// ProgrammeDB stores all queried categories.
type programmeDB struct {
	Categories []*tv.Category `json:"categories"`
	Saved      time.Time      `json:"saved"`
}

func newProgrammeDB(cats []*tv.Category, saved time.Time) *programmeDB {
	return &programmeDB{Categories: cats, Saved: saved}
}

func LoadProgrammeDbFromJSON(jsonfilename string) (*programmeDB, error) {
	file, err := ioutil.ReadFile(jsonfilename)
	if err != nil {
		return nil, err
	}
	var pdb programmeDB
	json.Unmarshal(file, &pdb)
	return &pdb, nil
}

func (pdb *programmeDB) toJSON() ([]byte, error) {
	marshalled, err := json.MarshalIndent(pdb, "", "\t")
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (pdb *programmeDB) Save(filename string) error {
	pdb.Saved = time.Now()
	pdb.index()
	json, err := pdb.toJSON()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, json, 0644)
}

func (pdb *programmeDB) index() {
	index := 0
	for _, i := range pdb.Categories {
		for _, j := range i.Programmes {
			j.Index = index
			index++
		}
	}
}

// ListCategory returns a String of all Programmes in an Iplayer
// Category, with a line separator appended after every Programme.
func (pdb *programmeDB) ListCategory(category string) string {
	var buffer bytes.Buffer
	cat, err := pdb.findCategory(category)
	if err != nil {
		return fmt.Sprintln(err)
	}
	for _, i := range cat.Programmes {
		buffer.WriteString(i.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (pdb *programmeDB) findCategory(category string) (*tv.Category, error) {
	for _, i := range pdb.Categories {
		if i.Name == category {
			return i, nil
		}
	}
	return nil, errors.New("Can not find Category with Name: " + category)
}

// ListAvailableCategories returns a line separated string of
// all Categories saved in the Programme DB.
func (pdb *programmeDB) ListAvailableCategories() string {
	var buffer bytes.Buffer
	for _, i := range pdb.Categories {
		buffer.WriteString(i.Name + "\n")
	}
	return buffer.String()
}

// FindTitle returns a line separated String of all Programmes
// containing a given string in its Title. Case distinction is ignored.
func (pdb *programmeDB) FindTitle(title string) string {
	var buffer bytes.Buffer
	for _, i := range pdb.Categories {
		for _, j := range i.Programmes {
			if strings.Contains(strings.ToLower(j.String()),
				strings.ToLower(title)) {
				buffer.WriteString(j.String() + "\n")
			}
		}
	}
	if buffer.String() == "" {
		return "No Matches found.\n"
	}
	return buffer.String()
}
