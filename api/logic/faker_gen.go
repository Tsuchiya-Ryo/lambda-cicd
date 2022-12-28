package logic

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/Tsuchiya-Ryo/lambda-cicd/api/model"

	"github.com/go-faker/faker/v4"
	"github.com/jszwec/csvutil"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const LAYOUT = "2006-01-02"

type FakeFields struct {
	Name string  `faker:"name"`
	ID   string  `faker:"uuid_digit"`
	Age  uint    `faker:"boundary_start=20, boundary_end=120"`
	Num  float32 `faker:"boundary_start=12.65, boundary_end=184.05"`
}

func GenerateData(from string, to string) ([]model.PutCSV, error) {
	fromDate, diffDays, err := getDiffDays(from, to)
	if err != nil {
		return nil, fmt.Errorf("\nFailed to get DiffDays: %w", err)
	}
	out := make([]model.PutCSV, diffDays+1)
	for i := range out {
		f := FakeFields{}
		_ = faker.FakeData(&f)
		out[i] = model.PutCSV{
			Date: fromDate.AddDate(0, 0, i).Format(LAYOUT),
			Name: f.Name,
			ID:   f.ID,
			Age:  f.Age,
			Num:  f.Num,
		}
	}
	return out, nil
}

func ToBody(input []model.PutCSV) (io.Reader, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(transform.NewWriter(&buf, unicode.UTF8.NewEncoder()))
	enc := csvutil.NewEncoder(w)
	for _, record := range input {
		err := enc.Encode(record)
		if err != nil {
			return nil, err
		}
	}
	w.Flush()
	return &buf, nil
}

func getDiffDays(from string, to string) (time.Time, int, error) {
	fromDate, err := time.Parse(LAYOUT, from)
	if err != nil {
		return time.Time{}, 0, err
	}
	toDate, err := time.Parse(LAYOUT, to)
	if err != nil {
		return time.Time{}, 0, err
	}
	if fromDate.After(toDate) {
		return time.Time{}, 0, fmt.Errorf("fromDate must be on or before the toDate")
	}

	diffDays := int(toDate.Sub(fromDate).Hours() / 24)
	if diffDays > 1000 {
		return time.Time{}, 0, fmt.Errorf("too long duration")
	}
	return fromDate, diffDays, nil
}
