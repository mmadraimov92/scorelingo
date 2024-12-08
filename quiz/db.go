package quiz

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
)

//go:embed estonian.csv
var estonian []byte

type db struct {
	words []word
}

type word struct {
	Estonian Estonian
	Russian  string
	English  string
}

type Estonian struct {
	FirstForm  string
	SecondForm string
	ThirdForm  string
}

func (e Estonian) allForms() string {
	return fmt.Sprintf("%s, %s, %s", e.FirstForm, e.SecondForm, e.ThirdForm)
}

func load() (*db, error) {
	r := csv.NewReader(bytes.NewReader(estonian))

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	words := make([]word, len(records))

	for i, line := range records {
		words[i] = word{
			Estonian: Estonian{
				FirstForm:  line[0],
				SecondForm: line[1],
				ThirdForm:  line[2],
			},
			Russian: line[3],
			English: line[4],
		}
	}

	return &db{words}, nil
}
