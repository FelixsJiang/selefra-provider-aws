



package apigateway



import (


	"testing"





	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"


	"github.com/golang/mock/gomock"


	"github.com/selefra/selefra-provider-aws/aws_client"




	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-sdk/table_schema_generator"




)







func buildApigatewayVpcLinks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayClient(ctrl)



	l := types.VpcLink{}
	err := faker.FakeObject(&l)


	if err != nil {


		t.Fatal(err)


	}


	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(


		&apigateway.GetVpcLinksOutput{


			Items: []types.VpcLink{l},




		}, nil)







	return aws_client.AwsServices{


		Apigateway: m,


	}


}





func TestVpcLinks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayVpcLinksGenerator{}), buildApigatewayVpcLinks, aws_client.TestOptions{})




}
