package areatz

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
			DST: stringToBool(tr.Find("td.dst").Text()),
		}
		codes = append(codes, ac)
	}
	fmt.Println(codes[0])
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

func main() {
	fmt.Printf(("Test string\n"))
}
