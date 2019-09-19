package util

import (
	"context"
	"errors"

	"github.com/OrderSystem_WeiZhang/config"
	"googlemaps.github.io/maps"
)

func GetDistance(origin []string, destination []string) (distance string, err error) {

	var client *maps.Client
	if config.GoogleMapApiKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(config.GoogleMapApiKey))
	} else if config.GoogleMapClientId != "" || config.GoogleMapClientSignature != "" {
		client, err = maps.NewClient(maps.WithClientIDAndSignature(config.GoogleMapClientId, config.GoogleMapClientSignature))
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      pepareCoordinateParameter(origin),
		Destinations: pepareCoordinateParameter(destination),
	}

	resp, err := client.DistanceMatrix(context.Background(), r)
	if resp.Rows[0].Elements[0].Status != "OK" {
		err = errors.New("cannot find distance")
	}
	distance = resp.Rows[0].Elements[0].Distance.HumanReadable
	return

}

func pepareCoordinateParameter(input []string) (resp []string) {

	var s string
	s = input[0] + "," + input[1]

	resp = append(resp, s)
	return
}
