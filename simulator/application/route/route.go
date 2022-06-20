package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("invalid route id")
	}
	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		lng, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		route.Positions = append(route.Positions, Position{
			Lat: lat, Lng: lng,
		})
	}
	return nil
}

func (route *Route) ExportJSONPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string
	total := len(route.Positions)
	for key, value := range route.Positions {
		partialRoute.ID = route.ID
		partialRoute.ClientID = route.ClientID
		partialRoute.Position = []float64{value.Lat, value.Lng}
		partialRoute.Finished = false
		if total-1 == key {
			partialRoute.Finished = true
		}
		jsonRoute, err := json.Marshal(partialRoute)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
