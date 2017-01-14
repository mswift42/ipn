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
func (pdb *programmeDB) findCategory(cat string) (*tv.Category, error) {
	for _, i := range pdb.Categories {
		if i.Name == cat {
			return i, nil
		}
	}
	return nil, errors.New("Can not find Category with Name: " + cat)
}

// ListCategory returns a String of all Programmes in an Iplayer
// Category, with a line seperator appended after every Programme.
func (pdb *programmeDB) ListCategory(cat string) string {
	var buffer bytes.Buffer
	category, err := pdb.findCategory(cat)
	if err != nil {
		return fmt.Sprintln(err)
	}
	for _, i := range category.Programmes {
		buffer.WriteString(i.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// TODO handle case mismatch.
func (pdb *programmeDB) FindTitle(cat string) string {
	var buffer bytes.Buffer
	for _, i := range pdb.Categories {
		for _, j := range i.Programmes {
			if strings.Contains(j.String(), cat) {
				buffer.WriteString(j.String() + "\n")
			}
		}
	}
	if buffer.String() == "" {
		return "No Matches found.\n"
	}
	return buffer.String()
}
