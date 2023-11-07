package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("S3 repository", func() {
	Context("When execute UploadFile", uploadFile)
})

func uploadFile() {
	Context("Works successfully", func() {

		c := MockClientS3{}
		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		d := S3Repository{client: &c}

		err := d.UploadFile("bName", "oKey", strings.NewReader(""))
		It("Should not return error", func() {
			Expect(err).To(BeNil())
		})

	})

	Context("Fail on upload file", func() {

		c := MockClientS3{}
		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything).Return(&s3.PutObjectOutput{}, errors.New("some error"))

		d := S3Repository{client: &c}

		err := d.UploadFile("bName", "oKey", strings.NewReader(""))
		It("Should not return error", func() {
			Expect(err).NotTo(BeNil())
		})

	})
}
