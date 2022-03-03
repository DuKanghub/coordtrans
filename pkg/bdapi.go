package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const bdApiUrl = "https://api.map.baidu.com/geoconv/v1/?coords=%s&from=%d&to=%d&ak=%s"

type BdApi struct {
	ak string
}

// 发起百度API请求
func apiRequest(coords []string, from, to int, ak string) []string {
	var results []string
	coordsMerged := strings.Join(coords, ";")
	url := fmt.Sprintf(bdApiUrl, coordsMerged, from, to, ak)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return results
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return results
	}
	fmt.Println(string(body))
	var result struct {
		Status int
		Result []struct {
			X float64
			Y float64
		}
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return results
	}
	if result.Status != 0 {
		return results
	}
	for i, r := range result.Result {
		results = append(results, fmt.Sprintf("%s,%f,%f", coords[i], r.X, r.Y))
	}
	return results
}

func request(coords []string, from, to int, ak string) []string {
	var results []string
	for i := 0; i < len(coords); i += 100 {
		end := i + 100
		if end > len(coords) {
			end = len(coords)
		}
		results = append(results, apiRequest(coords[i:end], from, to, ak)...)
	}
	return results
}

func (b *BdApi) WGS84toBD09(coords []string) []string {
	return request(coords, 1, 5, b.ak)
}

func (b *BdApi) BD09toWGS84(coords []string) []string {
	return request(coords, 5, 1, b.ak)
}

func (b *BdApi) GCJ02toBD09(coords []string) []string {
	return request(coords, 3, 5, b.ak)
}

func (b *BdApi) BD09toGCJ02(coords []string) []string {
	return request(coords, 5, 3, b.ak)
}

func (b *BdApi) WGS84toGCJ02(coords []string) []string {
	return request(coords, 1, 3, b.ak)
}

func (b *BdApi) GCJ02toWGS84(coords []string) []string {
	return request(coords, 3, 1, b.ak)
}
