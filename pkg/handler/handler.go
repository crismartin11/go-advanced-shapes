package handler

import (
	"context"
	"go-advanced-shapes/internal/processor"
	"go-advanced-shapes/pkg/models"
	"go-advanced-shapes/pkg/validator"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	p processor.IProcessor
	v validator.Validator
}

type IHandler interface {
	HandleApiRest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func New(p processor.IProcessor, v validator.Validator) IHandler {
	return Handler{
		p,
		v,
	}
}

var ginLambda *ginadapter.GinLambda

func (h Handler) HandleApiRest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sublogger := log.With().Str("component", "HandleApiRest").Logger()
	sublogger.Info().Msg("HandleApiRest started")
	r := gin.Default()

	// Read method
	r.POST("/read", func(c *gin.Context) {
		shapetype := c.Param("shapetype")
		sublogger.Info().Str("shapetype", shapetype).Msg("Http /read")

		var req models.Request
		err := c.ShouldBind(&req)
		if err != nil {
			sublogger.Error().Msg("Error in Bind.")
			c.JSON(http.StatusBadRequest, gin.H{
				"body":  "",
				"error": err.Error(),
			})
			return
		}

		err = h.v.ValidateRequest(req)
		if err != nil {
			sublogger.Error().Msg("Error in validation.")
			c.JSON(http.StatusBadRequest, gin.H{
				"body":  "",
				"error": err.Error(),
			})
			return
		}

		resp, err := h.p.ProcessGeneration(ctx, req)
		for k, v := range resp.Headers {
			c.Header(k, v)
		}
		c.JSON(resp.StatusCode, gin.H{
			"body":  resp.Body,
			"error": err.Error(),
		})
	})

	// Create method
	r.POST("/create", func(c *gin.Context) {
		shapetype := c.Param("shapetype")
		sublogger.Info().Str("shapetype", shapetype).Msg("Http /create")

		var req models.Request
		err := c.ShouldBind(&req)
		if err != nil {
			sublogger.Error().Msg("Error in bind.")
			c.JSON(http.StatusBadRequest, gin.H{
				"body":  "",
				"error": err.Error(),
			})
			return
		}

		err = h.v.ValidateRequest(req)
		if err != nil {
			sublogger.Error().Msg("Error in validation.")
			c.JSON(http.StatusBadRequest, gin.H{
				"body":  "",
				"error": err.Error(),
			})
			return
		}

		resp, err := h.p.ProcessCreation(req)
		for k, v := range resp.Headers {
			c.Header(k, v)
		}
		c.JSON(resp.StatusCode, gin.H{
			"body":  resp.Body,
			"error": err.Error(),
		})
	})

	ginLambda = ginadapter.New(r)

	return ginLambda.ProxyWithContext(ctx, req)
}
