package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID       string     `json:"RouteId"`
	ClientID string     `json:"clientId"`
	Position [2]float64 `json:"position"`
	Finished bool       `json:"finished"`
}

func NewRoute() *Route {
  return &Route{}
}

var RouteNotFoundError error = errors.New("Route id not informed")
var PositionNotParsedError error = errors.New("Unable to parse position")

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return RouteNotFoundError
	}

	destinationFile, err := os.Open(fmt.Sprintf("destinations/%s.txt", r.ID))
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	scanner := bufio.NewScanner(destinationFile)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return PositionNotParsedError
		}
		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return PositionNotParsedError
		}
		r.Positions = append(r.Positions, Position{latitude, longitude})
	}

	return nil
}

func (r *Route) ExportJSONPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	route.ID = r.ID
	route.ClientID = r.ClientID
	for index, position := range r.Positions {
		route.Position = [2]float64{position.Latitude, position.Longitude}
		if total-1 == index {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}

	return result, nil
}
