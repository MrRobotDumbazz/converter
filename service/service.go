package service

var AllUnits = map[string]float64{
	"nm":      0.000000001,
	"mcm":     0.000001,
	"mm":      0.001,
	"cm":      0.01,
	"dm":      0.1,
	"m":       1,
	"km":      1000,
	"bit":     1,
	"byte":    8,
	"KiB":     8 * 1024,
	"MiB":     8 * 1024 * 1024,
	"GiB":     8 * 1024 * 1024 * 1024,
	"TiB":     8 * 1024 * 1024 * 1024 * 1024,
	"PiB":     8 * 1024 * 1024 * 1024 * 1024 * 1024,
	"bit/s":   1,
	"byte/s":  8,
	"Kbit/s":  1e3,
	"Mbit/s":  1e6,
	"Gbit/s":  1e9,
	"ns":      0.000000001,
	"mcs":     0.000001,
	"ms":      0.001,
	"sec":     1,
	"min":     60,
	"hour":    3600,
	"day":     86400,
	"week":    604800,
	"month":   2592000,
	"year":    31536000,
	"centure": 3153600000,
	"mcg":     0.000000001,
	"mg":      0.000001,
	"g":       0.001,
	"kg":      1,
	"ton":     1000,
}

func Convert(value float64, from, to string) float64 {
	return value * AllUnits[from] / AllUnits[to]
}
