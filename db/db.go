package db

import (
	"encoding/json"
	"io/ioutil"
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
