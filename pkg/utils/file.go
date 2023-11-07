package utils

import (
	"bytes"
	"context"
	"fmt"
	"go-advanced-shapes/pkg/constants"
	"go-advanced-shapes/pkg/models"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

func GetObjectKey(ctx context.Context, request models.Request) string {
	lc, ok := lambdacontext.FromContext(ctx)
	awsRequestID := "unknown"
	if ok {
		awsRequestID = lc.AwsRequestID
	}
	fileName := request.ShapeType + "-" + awsRequestID + "-" + time.Now().Format(constants.DATE_FORMAT) + ".txt"
	return constants.S3_DIRECTORY + "/" + fileName
}

func GetFileReader(shapes []models.IShape) (io.Reader, error) {
	var buffer bytes.Buffer
	for _, sh := range shapes {
		_, err := buffer.WriteString(sh.Detail() + "\n")
		if err != nil {
			return nil, fmt.Errorf("%s", err)
		}
	}
	reader := strings.NewReader(buffer.String())

	return reader, nil
}
