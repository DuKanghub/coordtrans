package pkg

import "github.com/DuKanghub/coordtrans/utils/coordTransform"

type Mod struct {}


// WGS84toBD09 WGS84坐标转百度坐标
func (m *Mod) WGS84toBD09(lon, lat float64) (float64, float64) {
	return coordTransform.WGS84toBD09(lon, lat)
}

// BD09toWGS84 百度坐标转WGS84坐标
func (m *Mod) BD09toWGS84(lon, lat float64) (float64, float64) {
	return coordTransform.BD09toWGS84(lon, lat)
}

func (m *Mod) WGS84toGCJ02(lon, lat float64) (float64, float64) {
	return coordTransform.WGS84toGCJ02(lon, lat)
}

func (m *Mod) GCJ02toWGS84(lon, lat float64) (float64, float64) {
	return coordTransform.GCJ02toWGS84(lon, lat)
}

func (m *Mod) GCJ02toBD09(lon, lat float64) (float64, float64) {
	return coordTransform.GCJ02toBD09(lon, lat)
}

func (m *Mod) BD09toGCJ02(lon, lat float64) (float64, float64) {
	return coordTransform.BD09toGCJ02(lon, lat)
}