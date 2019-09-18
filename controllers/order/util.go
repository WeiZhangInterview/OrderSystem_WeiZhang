package order

import (
	"errors"
	"strconv"

	"github.com/OrderSystem_WeiZhang/models"
)

func validateCoordinate(coordinate []string) (err error) {
	lat, lon := -200.00, -200.00
	if len(coordinate) != 2 {
		err = errors.New("NOT VALID INPUT")
		return
	}
	lat, err = strconv.ParseFloat(coordinate[0], 64)
	lon, err = strconv.ParseFloat(coordinate[1], 64)

	//validate input coordiante
	if err != nil || lat > 180 || lat < -180 {
		err = errors.New("not valid input")
		return
	}
	if lon > 90 || lon < -90 {
		err = errors.New("not valid input")
		return
	}
	return
}

func validateCoordinateInput(o model.OrderRequest) (err error) {
	err = validateCoordinate(o.Origin)
	if err != nil {
		return
	}
	err = validateCoordinate(o.Destination)
	if err != nil {
		return
	}
	return
}
