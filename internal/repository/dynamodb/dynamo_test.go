package repository

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("DynamoDB repository", func() {
	Context("When execute ListShapesByType", listShapesByType)
	Context("When execute CreateShape", createShape)
})

func listShapesByType() {
	Context("Works successfully", func() {

		itemsResponse := []map[string]types.AttributeValue{
			{
				"id":      &types.AttributeValueMemberS{Value: "1"},
				"tipo":    &types.AttributeValueMemberS{Value: "RECTANGLE"},
				"a":       &types.AttributeValueMemberN{Value: "10.2"},
				"b":       &types.AttributeValueMemberN{Value: "15.4"},
				"creador": &types.AttributeValueMemberS{Value: "cris@gmail.com"},
			},
			{
				"id":      &types.AttributeValueMemberS{Value: "2"},
				"tipo":    &types.AttributeValueMemberS{Value: "ELLIPSE"},
				"a":       &types.AttributeValueMemberN{Value: "13.2"},
				"b":       &types.AttributeValueMemberN{Value: "13.4"},
				"creador": &types.AttributeValueMemberS{Value: "cr@gmail.com"},
			},
			{
				"id":      &types.AttributeValueMemberS{Value: "3"},
				"tipo":    &types.AttributeValueMemberS{Value: "CIRCLE"},
				"a":       &types.AttributeValueMemberN{Value: "15.2"},
				"b":       &types.AttributeValueMemberN{Value: "16.4"},
				"creador": &types.AttributeValueMemberS{Value: "mr@gmail.com"},
			},
		}

		c := MockClient{}
		c.On("Query", context.TODO(), mock.Anything, mock.Anything).Return(&dynamodb.QueryOutput{
			Items: itemsResponse,
		}, nil)

		d := DynamoDB{client: &c}

		res, err := d.ListShapesByType("RECTANGLE")
		It("Should return 3 elements", func() {
			Expect(len(res)).To(Equal(3))
		})
		It("Should not return error", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("Fail on list shapes query", func() {
		c := MockClient{}
		c.On("Query", context.TODO(), mock.Anything, mock.Anything).Return(&dynamodb.QueryOutput{}, errors.New("some error"))

		d := DynamoDB{client: &c}

		_, err := d.ListShapesByType("RECTANGLE")
		It("Should return error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
}

func createShape() {
	Context("Works successfully", func() {
		c := MockClient{}
		c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{}, nil)

		d := DynamoDB{client: &c}

		err := d.CreateShape("2", "RECTANGLE", 12.0, 13.0, "cr@gmail.com")
		It("Should not return error", func() {
			Expect(err).To(BeNil())
		})
	})

	Context("Fail on create shape", func() {
		c := MockClient{}
		c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&dynamodb.PutItemOutput{}, errors.New("some error"))

		d := DynamoDB{client: &c}

		err := d.CreateShape("2", "RECTANGLE", 12.0, 13.0, "cr@gmail.com")
		It("Should return error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
}
