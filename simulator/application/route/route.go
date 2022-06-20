package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat float64
	Lng float64
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
