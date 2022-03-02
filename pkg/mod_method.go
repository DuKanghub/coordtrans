package pkg

import (
	"fmt"
	"github.com/DuKanghub/coordtrans/utils/coordTransform"
	"strconv"
	"strings"
)

type Mod struct{}

func modCore(coords []string, f func(float64, float64) (float64, float64)) []string {
	var results []string
	for _, coord := range coords {
		if strings.Contains(coord, ",") {
			arr := strings.Split(coord, ",")
			if len(arr) != 2 {
				results = append(results, "")
				continue
			}
			lon, err := strconv.ParseFloat(arr[0], 64)
			if err != nil {
				results = append(results, "")
				continue
			}
			lat, err := strconv.ParseFloat(arr[1], 64)
			if err != nil {
				results = append(results, "")
				continue
			}
			x, y := f(lon, lat)
			results = append(results, fmt.Sprintf("%f,%f", x, y))
		}
	}
	return results
}

// WGS84toBD09 WGS84坐标转百度坐标
func (m *Mod) WGS84toBD09(coords []string) []string {
	return modCore(coords, coordTransform.WGS84toBD09)
}

// BD09toWGS84 百度坐标转WGS84坐标
func (m *Mod) BD09toWGS84(coords []string) []string {
	return modCore(coords, coordTransform.BD09toWGS84)
}

func (m *Mod) WGS84toGCJ02(coords []string) []string {
	return modCore(coords, coordTransform.WGS84toGCJ02)
}

func (m *Mod) GCJ02toWGS84(coords []string) []string {
	return modCore(coords, coordTransform.GCJ02toWGS84)
}

func (m *Mod) GCJ02toBD09(coords []string) []string {
	return modCore(coords, coordTransform.GCJ02toBD09)
}

func (m *Mod) BD09toGCJ02(coords []string) []string {
	return modCore(coords, coordTransform.BD09toGCJ02)
}
