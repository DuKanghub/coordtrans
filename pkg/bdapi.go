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
func apiRequest(coords string, from, to int, ak string) []string {
	var results []string
	url := fmt.Sprintf(bdApiUrl, coords, from, to, ak)
	resp, err := http.Get(url)
	if err != nil {
		return results
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return results
	}
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
	for _, r := range result.Result {
		results = append(results, fmt.Sprintf("%f,%f", r.X, r.Y))
	}
	return results
}

func (b *BdApi) WGS84toBD09(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 1, 5, b.ak)
}

func (b *BdApi) BD09toWGS84(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 5, 1, b.ak)
}

func (b *BdApi) GCJ02toBD09(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 3, 5, b.ak)
}

func (b *BdApi) BD09toGCJ02(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 5, 3, b.ak)
}

func (b *BdApi) WGS84toGCJ02(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 1, 3, b.ak)
}

func (b *BdApi) GCJ02toWGS84(coords []string) []string {
	str := strings.Join(coords, ";")
	return apiRequest(str, 3, 1, b.ak)
}
