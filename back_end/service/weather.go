package service

import (
	"back_end/model"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Weather struct {
}

type Weather_Response struct {
	Data []struct {
		Weather struct {
			Icon        string `json:"icon"`
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"data"`
}

func (w *Weather) Weather_get(lat string, lon string) error {
	baseUrl := "https://api.weatherbit.io/v2.0/current?"
	url := baseUrl + "key=" + os.Getenv("KEY") + "&lat=" + lat + "&lon=" + lon

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var weather Weather_Response
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return err
	}

	weather_ := model.Weather{}
	weather_.Code = weather.Data[0].Weather.Code

	if err := model.DB.Model(&model.Weather{}).Create(&weather_).Error; err != nil {
		return err
	}
	return nil
}

func (w *Weather) Weather_broadcast() (error, int) {
	weather := model.Weather{}
	if err := model.DB.Model(&model.Weather{}).Last(&weather).Error; err != nil {
		return err, 0
	}
	return nil, weather.Code
}
