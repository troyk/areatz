package areatz

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var areacodeURL = "add your url here"

type AreaCode struct {
	AreaCode  string `json:"area_code"`
	GMTOffset int    `json:"gmt_offset"`
	DST       bool   `json:"dst"`
	State     string `json:"state"`
	Region    string `json:"region"`
}

// Time calculates time for AreaCode struct using UTC and the GMT offset
func (ac AreaCode) Time() time.Time {
	return time.Now().UTC().Add(time.Duration(ac.GMTOffset) * time.Hour)
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
		}
		// Uncomment lines below for visual of time and formatting
		// fmt.Println(ac.State, AreaCode.Time(*ac).Format("3:04PM"))
		// fmt.Println(ac.State, AreaCode.Time(*ac).Format("Mon Jan _2 15:04:05 2006"))
		codes = append(codes, ac)
	}

	return codes, err
}

func AreaCodesToJSON() ([]byte, error) {
	codes, err := GetAreaCodes()
	json_output := make([]byte, 0)

	for i := 0; i < len(codes); i++ {
		code, err := json.Marshal(codes[i])
		if err != nil {
			return nil, err
		}
		json_output = append(json_output, code...)
	}

	return json_output, err
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
