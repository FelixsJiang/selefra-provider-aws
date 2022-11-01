package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildElasticacheParameterGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mockElasticache := mocks.NewMockElastiCache(ctrl)
	parameterGroupsOutput := elasticache.DescribeCacheParameterGroupsOutput{}
	err := faker.FakeObject(&parameterGroupsOutput)
	parameterGroupsOutput.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	parametersOutput := elasticache.DescribeCacheParametersOutput{}
	err = faker.FakeObject(&parametersOutput)
	parametersOutput.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&parameterGroupsOutput, nil)

	return aws_client.AwsServices{
		ElastiCache: mockElasticache,
	}
}

func TestElasticacheParameterGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticacheParameterGroupsGenerator{}), buildElasticacheParameterGroups, aws_client.TestOptions{})
}
