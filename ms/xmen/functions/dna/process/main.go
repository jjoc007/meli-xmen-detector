package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	inj "github.com/jjoc007/meli-xmen-detector/xmen/functions/dna"
	"github.com/jjoc007/meli-xmen-detector/xmen/log"
	dnamodel "github.com/jjoc007/meli-xmen-detector/xmen/model/dna"
	"github.com/jjoc007/meli-xmen-detector/xmen/service/dna"
	"github.com/jjoc007/meli-xmen-detector/xmen/utils"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda process dna sequence")
	var dnaPayload *dnamodel.DNA
	body:= event.Body
	err := json.Unmarshal([]byte(body), &dnaPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on decoding body %v", dnaPayload)
		return utils.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	log.Logger.Debug().Msgf("dna %v , %+v", dnaPayload, cxt)

	err = inj.Instances["dnaService"].(servicedna.DNAService).Process(dnaPayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the dna")
		return utils.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	if !dnaPayload.IsMutant {
		return events.APIGatewayProxyResponse{StatusCode: 403}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
