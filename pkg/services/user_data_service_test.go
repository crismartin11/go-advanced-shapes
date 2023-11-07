package services

import (
	"errors"
	"io"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Service", func() {
	Context("On User Data Service", userDataService)
})

func userDataService() {
	Context("On execute user data service", func() {

		body := io.NopCloser(strings.NewReader(`{"data": { "email": "cris@gmail.com" }}`))
		emptyBody := io.NopCloser(strings.NewReader(`{}`))

		Context("Works successfully", func() {

			h := MockHttpClient{}
			h.On("Do", mock.Anything).Return(&http.Response{Body: body, StatusCode: 200}, nil)

			uds := NewUserDataService(&h)
			_, err := uds.GetUserData("2")
			It("Should not return error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("Fail in http request", func() {

			h := MockHttpClient{}
			h.On("Do", mock.Anything).Return(&http.Response{Body: emptyBody, StatusCode: 400}, errors.New("some error"))

			uds := NewUserDataService(&h)
			_, err := uds.GetUserData("2")
			It("Should return error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
}
