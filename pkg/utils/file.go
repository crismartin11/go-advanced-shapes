package utils

import (
	"bytes"
	"context"
	"fmt"
	"go-advanced-shapes/pkg/constants"
	"go-advanced-shapes/pkg/models"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/rs/zerolog/log"
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

// GetFileReader. DEPRECATED method
func GetFileReader(shapes []models.IShape) (io.Reader, error) {
	var buffer bytes.Buffer
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	// Note: for just to see the order
	for _, sh := range shapes {
		log.Info().Msg(fmt.Sprintf("LIST. detail %s", sh.Detail()))
	}

	log.Info().Msg("CALCULATE. start")
	for i, sh := range shapes {

		wg.Add(1)
		go writeInBuffer(&buffer, sh, wg, m, i)
	}

	wg.Wait()
	reader := strings.NewReader(buffer.String())

	log.Info().Msg("CALCULATE. End")
	return reader, nil
}

// writeInBuffer. DEPRECATED method
func writeInBuffer(buffer *bytes.Buffer, sh models.IShape, wg *sync.WaitGroup, m *sync.RWMutex, i int) {
	m.Lock()
	buffer.WriteString(sh.Detail() + "\n")
	m.Unlock()

	log.Info().Msg(fmt.Sprintf("CALCULATE. iteracion %d", i))
	log.Info().Msg(fmt.Sprintf("buffer %v", buffer.String()))

	wg.Done()
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////

func GetFileReaderChannels(shapes []models.IShape) (io.Reader, error) {
	log.Info().Msg("CHANNELS. Start")

	var buffer bytes.Buffer
	channelDetails := make(chan string)

	go makeDetails(channelDetails, shapes)

	for detail := range channelDetails {
		log.Info().Msg(fmt.Sprintf("CHANNELS. Read Detail %s", detail))
		buffer.WriteString(detail + "\n")
	}

	log.Info().Msg(fmt.Sprintf("CHANNELS. End. %s", buffer.String()))
	reader := strings.NewReader(buffer.String())
	return reader, nil
}

func makeDetails(channelDetails chan string, shapes []models.IShape) {
	for i, sh := range shapes {
		log.Info().Msg(fmt.Sprintf("CHANNELS. Make Detail %s. Iteration %d", sh.Detail(), i))
		channelDetails <- sh.Detail()
		//time.Sleep(2 * time.Second)
	}
	close(channelDetails)
}
