package handler

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type MeatHandler struct{}

func (m *MeatHandler) GetSummary(c echo.Context) error {
	log.Print("start get beef summary")
	result := map[string]int{}
	rawData, fetchErr := fetchData()
	if fetchErr != nil {
		log.Fatal(fetchErr)
		return c.JSON(http.StatusBadRequest, result)
	}

	cleanData := cleanData(rawData)
	for _, v := range cleanData {
		if item, ok := result[v]; ok {
			item += 1
			result[v] = item
		} else {
			result[v] = 1
		}
	}
	response := Response{
		Beef: result,
	}
	log.Print("get summary success")
	return c.JSON(http.StatusOK, response)
}

func fetchData() (string, error) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseData), nil
}

func cleanData(rawData string) []string {
	rawData = strings.Replace(rawData, ".", "", -1)
	rawData = strings.Replace(rawData, ",", "", -1)
	rawData = strings.Replace(rawData, "  ", "", -1)
	result := strings.Fields(rawData)
	return result
}
