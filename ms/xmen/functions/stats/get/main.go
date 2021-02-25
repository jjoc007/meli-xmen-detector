package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/meli-xmen-detector/xmen/functions/dna"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	servicestats "github.com/jjoc007/meli-xmen-detector/xmen/service/stats"
	"github.com/jjoc007/meli-xmen-detector/xmen/utils"
)

func LambdaHandler() (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda get stats")

	stats, err := inj.Instances["statsService"].(servicestats.StatsService).Get()
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the stats get")
		return utils.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	outStats, err := json.Marshal(stats)
	if err != nil {
		return utils.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Body:              string(outStats),
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
