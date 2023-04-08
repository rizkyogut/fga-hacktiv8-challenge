package main

func WaterCheckLevel(water int) (status string) {
	switch {
	case water <= 5:
		status = "aman"
	case water >= 6 && water <= 8:
		status = "siaga"
	case water > 8:
		status = "bahaya"
	}
	return
}

func WindCheckLevel(wind int) (status string) {
	switch {
	case wind <= 6:
		status = "aman"
	case wind >= 7 && wind <= 15:
		status = "siaga"
	case wind > 15:
		status = "bahaya"
		return
	}
	return
}
