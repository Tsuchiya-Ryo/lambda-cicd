package logic_test

import (
	"math/rand"
	"testing"

	"github.com/Tsuchiya-Ryo/lambda-cicd/api/logic"
	"github.com/Tsuchiya-Ryo/lambda-cicd/api/model"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

type args struct {
	From string
	To   string
}

func TestGenerateDataValids(t *testing.T) {
	type cases struct {
		Name    string
		ArgCase args
		Want    []model.PutCSV
	}
	testCases := []cases{
		{
			Name:    "single",
			ArgCase: args{From: "2022-01-01", To: "2022-01-01"},
			Want: []model.PutCSV{
				{Date: "2022-01-01", Age: 0x46, Num: 20.160484},
			},
		},
		{
			Name:    "double",
			ArgCase: args{From: "2022-01-01", To: "2022-01-02"},
			Want: []model.PutCSV{
				{Date: "2022-01-01", Age: 0x46, Num: 20.160484},
				{Date: "2022-01-02", Age: 0x30, Num: 123.4389},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			faker.SetRandomSource(rand.NewSource(42))
			got, err := logic.GenerateData(tc.ArgCase.From, tc.ArgCase.To)
			for i := 0; i < len(got); i++ {
				assert.Equal(t, got[i].Date, tc.Want[i].Date)
				assert.Equal(t, got[i].Age, tc.Want[i].Age)
				assert.Equal(t, got[i].Num, tc.Want[i].Num)
			}
			assert.NoError(t, err)
		})
	}
}

func TestGenerateDataInvalids(t *testing.T) {
	type cases struct {
		Name    string
		ArgCase args
		WantMsg string
	}
	testCases := []cases{
		{
			Name:    "cannotParse",
			ArgCase: args{From: "2022-1-1", To: "2022-01-01"},
			WantMsg: "cannot parse",
		},
		{
			Name:    "tooLong",
			ArgCase: args{From: "2022-01-01", To: "2100-01-01"},
			WantMsg: "too long duration",
		},
		{
			Name:    "inverse",
			ArgCase: args{From: "2022-12-31", To: "2022-01-01"},
			WantMsg: "fromDate must be on or before the toDate",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			faker.SetRandomSource(rand.NewSource(42))
			got, err := logic.GenerateData(tc.ArgCase.From, tc.ArgCase.To)
			assert.Nil(t, got)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.WantMsg)
		})
	}
}
