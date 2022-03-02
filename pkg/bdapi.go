package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const bdApiUrl = "https://api.map.baidu.com/geoconv/v1/?coords=%s&from=%d&to=%d&ak=%s"

type BdApi struct {
	ak string
}

// 发起百度API请求
func apiRequest(lon, lat float64, from,to int, ak string) (float64, float64) {
	url := fmt.Sprintf(bdApiUrl, fmt.Sprintf("%f,%f", lon, lat), from, to, ak)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0
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
		return 0, 0
	}
	if result.Status != 0 {
		return 0, 0
	}
	return result.Result[0].X, result.Result[0].Y
}

func (b *BdApi) WGS84toBD09(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 1, 5, b.ak)
}

func (b *BdApi) BD09toWGS84(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 5, 1, b.ak)
}


func (b *BdApi) GCJ02toBD09(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 3, 5, b.ak)
}

func (b *BdApi) BD09toGCJ02(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 5, 3, b.ak)
}

func (b *BdApi) WGS84toGCJ02(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 1, 3, b.ak)
}

func (b *BdApi) GCJ02toWGS84(lon, lat float64) (float64, float64) {
	return apiRequest(lon, lat, 3, 1, b.ak)
}