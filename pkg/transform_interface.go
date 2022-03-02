package pkg

type TransFormer interface {
	WGS84toBD09(lon, lat float64) (float64, float64)
	BD09toWGS84(lon, lat float64) (float64, float64)
	WGS84toGCJ02(lon, lat float64) (float64, float64)
	GCJ02toWGS84(lon, lat float64) (float64, float64)
	GCJ02toBD09(lon, lat float64) (float64, float64)
	BD09toGCJ02(lon, lat float64) (float64, float64)
}
// WGS84坐标系：即地球坐标系，国际上通用的坐标系。GPS标准坐标
// GCJ02坐标系：即火星坐标系，WGS84坐标系经加密后的坐标系。Google Maps，高德，腾讯地图在用。
// BD09坐标系：即百度坐标系，GCJ02坐标系经加密后的坐标系。
// longitude:经度, latitude:纬度

type Account struct {
	Method string
	AK  string
}

func NewTransFormer(t Account)  TransFormer {
	switch t.Method {
	case "mod":
		return &Mod{}
	case "bd":
		return &BdApi{ak: t.AK}
	default:
		return &Mod{}
	}
}