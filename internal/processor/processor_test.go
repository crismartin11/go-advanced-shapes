package processor

import (
	"context"
	"errors"
	"go-advanced-shapes/pkg/models"
	"go-advanced-shapes/pkg/services"

	d "go-advanced-shapes/internal/repository/dynamodb"
	s "go-advanced-shapes/internal/repository/s3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Service", func() {
	Context("On ProcessCreation", processCreatio)
})

func processCreatio() {
	Context("On execute process creation", func() {

		//body := io.NopCloser(strings.NewReader(`{"data": { "email": "cris@gmail.com" }}`))
		//emptyBody := io.NopCloser(strings.NewReader(`{}`))
		//models.Request{Id: "1", ShapeType: "RECTANGLE", A: 15.2, B: 15.4}
		var userData models.UserDataResponse
		userData.Data.Email = "cr@gmail.com"

		Context("Works successfully", func() {
			dy := d.MockDynamoDB{}
			dy.On("CreateShape", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

			sr := s.MockS3Repository{}

			uds := services.MockUserDataService{}
			uds.On("GetUserData", mock.Anything).Return(userData, nil)

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessCreation(models.Request{})
			It("Should not return error", func() {
				Expect(err.Error()).To(Equal(""))
			})
		})

		Context("Fail on Query", func() {
			dy := d.MockDynamoDB{}
			dy.On("CreateShape", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some error"))

			sr := s.MockS3Repository{}

			uds := services.MockUserDataService{}
			uds.On("GetUserData", mock.Anything).Return(userData, nil)

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessCreation(models.Request{})
			It("Should return error", func() {
				Expect(err.Error()).NotTo(Equal(""))
				Expect(err.Error()).To(Equal("ERROR: some error"))
			})
		})

		Context("Fail on Get User data", func() {
			dy := d.MockDynamoDB{}
			dy.On("CreateShape", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

			sr := s.MockS3Repository{}

			uds := services.MockUserDataService{}
			uds.On("GetUserData", mock.Anything).Return(userData, errors.New("some error data"))

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessCreation(models.Request{})
			It("Should return error", func() {
				Expect(err.Error()).NotTo(Equal(""))
				Expect(err.Error()).To(Equal("ERROR: some error data"))
			})
		})

	})

	Context("On execute process generation", func() {

		list := []models.Request{
			{
				Id:        "1",
				ShapeType: "RECTANGLE",
				A:         15.4,
				B:         12.9,
				Creator:   "cr@gmail.com",
			},
		}

		Context("Works successfully", func() {
			dy := d.MockDynamoDB{}
			dy.On("ListShapesByType", mock.Anything).Return(list, nil)

			sr := s.MockS3Repository{}
			sr.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)

			uds := services.MockUserDataService{}

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessGeneration(context.TODO(), models.Request{})
			It("Should not return error", func() {
				Expect(err.Error()).To(Equal(""))
			})
		})

		Context("Fail on List shapes query", func() {
			dy := d.MockDynamoDB{}
			dy.On("ListShapesByType", mock.Anything).Return([]models.Request{}, errors.New("some error"))

			sr := s.MockS3Repository{}
			sr.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)

			uds := services.MockUserDataService{}

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessGeneration(context.TODO(), models.Request{})
			It("Should return error", func() {
				Expect(err.Error()).NotTo(Equal(""))
				Expect(err.Error()).To(Equal("ERROR: some error"))
			})
		})

		Context("Fail on upload file", func() {
			dy := d.MockDynamoDB{}
			dy.On("ListShapesByType", mock.Anything).Return(list, nil)

			sr := s.MockS3Repository{}
			sr.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some error"))

			uds := services.MockUserDataService{}

			p := New(&dy, &sr, &uds)
			_, err := p.ProcessGeneration(context.TODO(), models.Request{})
			It("Should return error", func() {
				Expect(err.Error()).NotTo(Equal(""))
				Expect(err.Error()).To(Equal("ERROR: some error"))
			})
		})
	})
}