package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayv2ApisGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApisGenerator{}

func (x *TableAwsApigatewayv2ApisGenerator) GetTableName() string {
	return "aws_apigatewayv2_apis"
}

func (x *TableAwsApigatewayv2ApisGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApisGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApisGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"arn",
		},
	}
}

func (x *TableAwsApigatewayv2ApisGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config apigatewayv2.GetApisInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetApis(ctx, &config)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayv2ApisGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApisGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disable_execute_api_endpoint").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DisableExecuteApiEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cors_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CorsConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warnings").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Warnings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_gateway_managed").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ApiGatewayManaged")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disable_schema_validation").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DisableSchemaValidation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_info").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ImportInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					return arn.ARN{
						Partition:	cl.Partition,
						Service:	string("apigateway"),
						Region:		cl.Region,
						AccountID:	"",
						Resource:	fmt.Sprintf("/apis/%s", aws.ToString(result.(types.Api).ApiId)),
					}.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProtocolType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_key_selection_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApiKeySelectionExpression")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_selection_expression").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RouteSelectionExpression")).Build(),
	}
}

func (x *TableAwsApigatewayv2ApisGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiStagesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiAuthorizersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiDeploymentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiIntegrationsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiModelsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiRoutesGenerator{}),
	}
}
