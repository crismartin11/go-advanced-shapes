package service

import (
	"context"
	"fmt"

	"go-advanced-shapes/pkg/constants"
	"go-advanced-shapes/pkg/models"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/rs/zerolog/log"
)

// NOTE: this service exist just for apply the Adapter method.

type ServiceForAdapter struct {
	client *dynamodb.Client
}

type IServiceForAdapter interface {
	ListShapesByTypePartiQl(shapeType string) ([]models.Request, error)
	CreateShapePartiQl(id string, shapeType string, a float64, b float64, creator string) error // Note: I should pass an object as a param
}

func NewServiceForAdapter() IServiceForAdapter {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Configuration error, " + err.Error())
	}

	return ServiceForAdapter{client: dynamodb.NewFromConfig(cfg)}
}

func (s ServiceForAdapter) ListShapesByTypePartiQl(shapeType string) ([]models.Request, error) {
	sublogger := log.With().Str("component", "ListShapesByTypePartiQl").Logger()
	sublogger.Info().Msg("ListShapesByType service for adapter started")
	shapes := []models.Request{}

	var lastEvaluatedKey map[string]types.AttributeValue
	i := 0
	for {
		i++
		sublogger.Info().Str("topic", "Pagination").Msg(fmt.Sprintf("Page %d\n", i))

		// Query
		queryInput := &dynamodb.QueryInput{
			TableName:              aws.String(constants.SHAPE_TABLE_NAME),
			IndexName:              aws.String("tipo-index"),
			KeyConditionExpression: aws.String("tipo = :shapeType"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":shapeType": &types.AttributeValueMemberS{
					Value: shapeType,
				},
			},
			Limit: aws.Int32(2),
		}

		if lastEvaluatedKey != nil {
			queryInput.ExclusiveStartKey = lastEvaluatedKey
		}

		output, err := s.client.Query(context.TODO(), queryInput)
		if err != nil {
			return shapes, fmt.Errorf("ListShapesByType. Error al invocar API Query. %s", err)
		}

		// Append shapes
		shapesPage := []models.Request{}
		err = attributevalue.UnmarshalListOfMaps(output.Items, &shapesPage)
		if err != nil {
			return shapes, fmt.Errorf("ListShapesByType. Error al parsear items. %s", err)
		}

		if len(shapesPage) > 0 {
			shapes = append(shapes, shapesPage...)
		}

		// Evaluate loop
		if output.LastEvaluatedKey != nil {
			lastEvaluatedKey = output.LastEvaluatedKey
			sublogger.Info().Str("topic", "Pagination").Msg("LastEvaluatedKey is not nil so next page")
		} else {
			sublogger.Info().Str("topic", "Pagination").Msg("LastEvaluatedKey is nil so break")
			break
		}
	}

	return shapes, nil
}

func (s ServiceForAdapter) CreateShapePartiQl(id string, shapeType string, a float64, b float64, creator string) error {
	sublogger := log.With().Str("component", "CreateShapePartiQl").Logger()
	sublogger.Info().Msg("CreateShape service for adapter started")

	params, err := attributevalue.MarshalList([]interface{}{id, shapeType, a, b, creator})
	if err != nil {
		return fmt.Errorf("Create. Error in marshal params (%s). %s", id, err)
	}

	_, err = s.client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement:  aws.String(fmt.Sprintf("INSERT INTO %s VALUE {'id': ?, 'tipo': ?, 'a': ?, 'b': ?, 'creador': ?}", constants.SHAPE_TABLE_NAME)),
		Parameters: params,
	})
	if err != nil {
		return fmt.Errorf("Create. Error insertando figura (%s). %s", id, err)
	}

	return nil
}
