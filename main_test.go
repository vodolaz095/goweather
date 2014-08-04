package goweather

import "testing"

func TestGetWeather(t *testing.T) {
	w, e := GetWeather("Moscow")
	if e != nil {
		t.Error("Error getting weather for Moscow", e)
	} else {
		if w.Name != "Moscow" {
			t.Error("We got weather not for Moscow, but for ", w.Name)
		}
	}
}
