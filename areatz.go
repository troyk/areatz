package areatz

// Keep as package
// Time will be a method based on the info received
// JSON output (function to convert to JSON)

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var areacodeURL = "add your url here"

type AreaCode struct {
	AreaCode  string `json:"area_code"`
	GMTOffset int    `json:"gmt_offset"`
	DST       bool   `json:"dst"`
	State     string `json:"state"`
	Region    string `json:"region"`
	Time      string `json:"time"`
}

func GetAreaCodes() ([]*AreaCode, error) {
	doc, err := goquery.NewDocument(areacodeURL)
	if err != nil {
		return nil, err
	}
	sel := doc.Find(`table.gvGrid`)
	if sel.Size() < 1 {
		return nil, errors.New("table.gvGrid not found")
	}
	codes := make([]*AreaCode, 0)
	rows := sel.Find("tr")
	fmt.Println("rowsize", rows.Size())

	for i := 0; i < rows.Size(); i++ {
		if i == 0 {
			continue // skip headers
		}
		tr := rows.Eq(i)
		ac := &AreaCode{
			AreaCode:  tr.Find("td").First().Text(),
			GMTOffset: stringToInt(tr.Find("td.tz").Text()),
			DST:       stringToBool(tr.Find("td.dst").Text()),
			State:     tr.Find("td.time").Next().Next().Text(),
			Region:    tr.Find("td").Last().Text(),
			Time:      tr.Find("td.time").Text(),
		}
		codes = append(codes, ac)
	}

	// FOR TESTING ONLY; DELETE WHEN DONE
	for i := 0; i < len(codes); i++ {
		fmt.Println(
			"AreaCode:", codes[i].AreaCode,
			"GMTOffset:", codes[i].GMTOffset,
			"DST:", codes[i].DST,
			"State:", codes[i].State,
			"Region:", codes[i].Region,
			"Time:", codes[i].Time,
		)
	}
	// END OF TESTING BLOCK

	return codes, err
}

func stringToInt(val string) int {
	x, _ := strconv.Atoi(val)
	return x
}

func stringToBool(val string) bool {
	if val == "Y" {
		return true
	}
	return false
}
