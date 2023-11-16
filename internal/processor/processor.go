package processor

import (
	"context"
	"fmt"

	d "go-advanced-shapes/internal/repository/dynamodb"
	s "go-advanced-shapes/internal/repository/s3"
	"go-advanced-shapes/pkg/constants"
	"go-advanced-shapes/pkg/models"
	"go-advanced-shapes/pkg/services"
	"go-advanced-shapes/pkg/utils"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Processor struct {
	d   d.IDynamoDB
	s3r s.IS3Repository
	us  services.IUserDataService
}

type IProcessor interface {
	ProcessCreation(request models.Request) (models.Response, error)
	ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error)
}

func New(d d.IDynamoDB, s3r s.IS3Repository, us services.IUserDataService) IProcessor {
	return Processor{
		d,
		s3r,
		us,
	}
}

func (p Processor) ProcessCreation(request models.Request) (models.Response, error) {
	sublogger := log.With().Str("component", "ProcessCreation").Logger()
	sublogger.Info().Msg("Create shape handler started.")

	// Get user data from api
	user, err := p.us.GetUserData(request.Id)
	if err != nil {
		sublogger.Error().Msg("Error getting user data from API.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else if user.Data.Email == "" {
		sublogger.Error().Msg("Error getting email user from data.")
		return models.NewResponseError(400, fmt.Sprintf("ProcessCreation. No se obtuvo el email del usuario con id %s", request.Id))
	}

	// Insert shape in table
	uuid, _ := uuid.NewUUID()
	err = p.d.CreateShape(uuid.String(), request.ShapeType, request.A, request.B, user.Data.Email)
	if err != nil {
		sublogger.Error().Msg("Error creating shape in DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	sublogger.Info().Msg("Create shape handler finished successfully.")
	return models.NewResponseOk("Creation process successful!")
}

func (p Processor) ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error) {
	sublogger := log.With().Str("component", "ProcessGeneration").Logger()
	sublogger.Info().Msg("Generate shape list file handler started.")

	// Get shapes from table
	sublogger.Info().Msg("Getting list from dynamoDB.")
	listShapes, err := p.d.ListShapesByType(request.ShapeType)
	if err != nil {
		sublogger.Error().Msg("Error getting shape list from DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	// Generate Shape list
	sublogger.Info().Msg("Generating shapes.")
	var shapes = []models.IShape{}
	for _, item := range listShapes {
		elem, err := models.ShapeFactory(item.Id, item.ShapeType, item.A, item.B)
		if err != nil {
			sublogger.Error().Msg("Error making shape based on dynamo item.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
		shapes = append(shapes, elem)
	}

	if len(shapes) == 0 {
		sublogger.Info().Msg("Generate shape list file handler finished with empty list.")
		return models.NewResponseOk("Generation file process successful! Empty list (file wasn't generated)")
	}

	// Generate file
	sublogger.Info().Msg("Generating file.")
	//fileReader, err := utils.GetFileReader(shapes)
	fileReader, err := utils.GetFileReaderChannels(shapes)
	if err != nil {
		sublogger.Error().Msg("Error generating file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	// Upload File to s3
	sublogger.Info().Msg("Uploading file.")
	err = p.s3r.UploadFile(constants.BUCKET_NAME, utils.GetObjectKey(ctx, request), fileReader)
	if err != nil {
		sublogger.Error().Msg("Error uploading file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	sublogger.Info().Msg("Generate shape list file handler finished successfully.")
	return models.NewResponseOk("Generation file process successful!")
}
