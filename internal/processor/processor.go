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
	log.Info().Msg("Create shape handler started (ProcessCreation).")

	// Get user data from api
	user, err := p.us.GetUserData(request.Id)
	if err != nil {
		log.Error().Msg("ProcessCreation.Error getting user data from API.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	} else if user.Data.Email == "" {
		log.Error().Msg("ProcessCreation.Error getting email user from data.")
		return models.NewResponseError(400, fmt.Sprintf("ProcessCreation. No se obtuvo el email del usuario con id %s", request.Id))
	}

	// Insert shape in table
	uuid, _ := uuid.NewUUID()
	err = p.d.CreateShape(uuid.String(), request.ShapeType, request.A, request.B, user.Data.Email)
	if err != nil {
		log.Error().Msg("ProcessCreation. Error creating shape in DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	log.Info().Msg("Create shape handler finished successfully (ProcessCreation).")
	return models.NewResponseOk("Creation process successful!")
}

func (p Processor) ProcessGeneration(ctx context.Context, request models.Request) (models.Response, error) {
	log.Info().Msg("ProcessGeneration. Generate shape list file handler started.")

	// Get shapes from table
	log.Info().Msg("ProcessGeneration. Getting list from dynamoDB.")
	listShapes, err := p.d.ListShapesByType(request.ShapeType)
	if err != nil {
		log.Error().Msg("ProcessGeneration. Error getting shape list from DynamoDB.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	// Generate Shape list
	log.Info().Msg("ProcessGeneration. Generating shapes.")
	var shapes = []models.IShape{}
	for _, item := range listShapes {
		elem, err := models.ShapeFactory(item.Id, item.ShapeType, item.A, item.B)
		if err != nil {
			log.Error().Msg("ProcessGeneration. Error making shape based on dynamo item.")
			return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
		}
		shapes = append(shapes, elem)
	}

	// Generate file
	log.Info().Msg("ProcessGeneration. Generating file.")
	fileReader, err := utils.GetFileReader(shapes)
	if err != nil {
		log.Error().Msg("ProcessGeneration. Error generating file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	// Upload File to s3
	log.Info().Msg("ProcessGeneration. Uploading file.")
	err = p.s3r.UploadFile(constants.BUCKET_NAME, utils.GetObjectKey(ctx, request), fileReader)
	if err != nil {
		log.Error().Msg("ProcessGeneration. Error uploading file.")
		return models.NewResponseError(400, fmt.Sprintf("ERROR: %s", err))
	}

	log.Info().Msg("ProcessGeneration. Generate shape list file handler finished successfully.")
	return models.NewResponseOk("Generation file process successful!")
}
