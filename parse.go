package safe

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func Parse(filepath string) (BillBunches, error) {
	// open file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true

	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	bbs := BillBunches{}

	for _, line := range data {
		// first entry is labeled
		// denomination: count
		labelled := strings.Split(strings.ReplaceAll(line[0], " ", ""), ":")
		//labelled := strings.Split(line[0], ":")

		denomination := labelled[0]
		firstCountS := labelled[1]

		firstCount, err := strconv.Atoi(firstCountS)
		if err != nil {
			return nil, err
		}

		bill, err := StrToBill(denomination)
		if err != nil {
			return nil, err
		}

		bbs = append(bbs, BillBunch{
			Bill:  *bill,
			Count: uint(firstCount),
		})

		for _, billCountS := range line[1:] {
			billCount, err := strconv.Atoi(billCountS)
			if err != nil {
				return nil, err
			}

			bbs = append(bbs, BillBunch{
				Bill:  *bill,
				Count: uint(billCount),
			})
		}
	}

	return bbs, nil
}
