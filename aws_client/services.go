package aws_client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)
import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)
import (
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
)
import (
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
)
import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

type AwsServices struct {
	Account                   AccountClient
	Accessanalyzer            AccessanalyzerClient
	Acm                       AcmClient
	Amp                       AmpClient
	Apigateway                ApigatewayClient
	Apigatewayv2              Apigatewayv2Client
	Applicationautoscaling    ApplicationautoscalingClient
	Apprunner                 ApprunnerClient
	Appsync                   AppsyncClient
	Athena                    AthenaClient
	Autoscaling               AutoscalingClient
	Backup                    BackupClient
	Cloudformation            CloudformationClient
	Cloudfront                CloudfrontClient
	Cloudhsmv2                Cloudhsmv2Client
	Cloudtrail                CloudtrailClient
	Cloudwatch                CloudwatchClient
	Cloudwatchlogs            CloudwatchlogsClient
	Codebuild                 CodebuildClient
	Codepipeline              CodepipelineClient
	Cognitoidentity           CognitoidentityClient
	Cognitoidentityprovider   CognitoidentityproviderClient
	Configservice             ConfigserviceClient
	Databasemigrationservice  DatabasemigrationserviceClient
	Dax                       DaxClient
	Directconnect             DirectconnectClient
	Docdb                     DocdbClient
	Dynamodb                  DynamodbClient
	Ec2                       Ec2Client
	Ecr                       EcrClient
	Ecrpublic                 EcrpublicClient
	Ecs                       EcsClient
	Efs                       EfsClient
	Eks                       EksClient
	Elasticache               ElasticacheClient
	Elasticbeanstalk          ElasticbeanstalkClient
	Elasticloadbalancing      ElasticloadbalancingClient
	Elasticloadbalancingv2    Elasticloadbalancingv2Client
	Elasticsearchservice      ElasticsearchserviceClient
	Kafka                     KafkaClient
	Emr                       EmrClient
	Eventbridge               EventbridgeClient
	Firehose                  FirehoseClient
	Frauddetector             FrauddetectorClient
	Fsx                       FsxClient
	Glacier                   GlacierClient
	Glue                      GlueClient
	Guardduty                 GuarddutyClient
	Identitystore             IdentitystoreClient
	Inspector                 InspectorClient
	Inspector2                Inspector2Client
	Iot                       IotClient
	Kinesis                   KinesisClient
	Kms                       KmsClient
	Lambda                    LambdaClient
	Savingsplans              SavingsplansClient
	Lightsail                 LightsailClient
	Mq                        MqClient
	Mwaa                      MwaaClient
	Neptune                   NeptuneClient
	Organizations             OrganizationsClient
	Qldb                      QldbClient
	Rds                       RdsClient
	Securityhub               SecurityhubClient
	Redshift                  RedshiftClient
	Resourcegroups            ResourcegroupsClient
	Route53                   Route53Client
	Route53domains            Route53domainsClient
	S3                        S3Client
	S3control                 S3controlClient
	Sagemaker                 SagemakerClient
	Secretsmanager            SecretsmanagerClient
	Servicecatalog            ServicecatalogClient
	Servicecatalogappregistry ServicecatalogappregistryClient
	Servicequotas             ServicequotasClient
	Quicksight                QuicksightClient
	Sesv2                     Sesv2Client
	Shield                    ShieldClient
	Sns                       SnsClient
	Sqs                       SqsClient
	Timestreamwrite           TimestreamwriteClient
	IAM                       IAMClient
	Sfn                       SfnClient
	Ssm                       SsmClient
	Ssoadmin                  SsoadminClient
	Appstream                 AppstreamClient
	Transfer                  TransferClient
	Waf                       WafClient
	Wafregional               WafregionalClient
	Wafv2                     Wafv2Client
	Elastictranscoder         ElastictranscoderClient
	Workspaces                WorkspacesClient
	Xray                      XrayClient
	Scheduler                 SchedulerClient
	Ses                       SesClient
	S3manager                 S3managerClient
	RAM                       RAMClient
}

func newAwsService(config aws.Config) *AwsServices {
	return &AwsServices{
		Elastictranscoder:         elastictranscoder.NewFromConfig(config),
		Account:                   account.NewFromConfig(config),
		Accessanalyzer:            accessanalyzer.NewFromConfig(config),
		Acm:                       acm.NewFromConfig(config),
		Apigateway:                apigateway.NewFromConfig(config),
		Amp:                       amp.NewFromConfig(config),
		Apigatewayv2:              apigatewayv2.NewFromConfig(config),
		Applicationautoscaling:    applicationautoscaling.NewFromConfig(config),
		Apprunner:                 apprunner.NewFromConfig(config),
		Appsync:                   appsync.NewFromConfig(config),
		Appstream:                 appstream.NewFromConfig(config),
		Athena:                    athena.NewFromConfig(config),
		Autoscaling:               autoscaling.NewFromConfig(config),
		Backup:                    backup.NewFromConfig(config),
		Cloudformation:            cloudformation.NewFromConfig(config),
		Cloudfront:                cloudfront.NewFromConfig(config),
		Timestreamwrite:           timestreamwrite.NewFromConfig(config),
		Cloudhsmv2:                cloudhsmv2.NewFromConfig(config),
		Savingsplans:              savingsplans.NewFromConfig(config),
		Cloudtrail:                cloudtrail.NewFromConfig(config),
		Cloudwatch:                cloudwatch.NewFromConfig(config),
		Cloudwatchlogs:            cloudwatchlogs.NewFromConfig(config),
		Codebuild:                 codebuild.NewFromConfig(config),
		Scheduler:                 scheduler.NewFromConfig(config),
		Codepipeline:              codepipeline.NewFromConfig(config),
		Sfn:                       sfn.NewFromConfig(config),
		Cognitoidentity:           cognitoidentity.NewFromConfig(config),
		Cognitoidentityprovider:   cognitoidentityprovider.NewFromConfig(config),
		Configservice:             configservice.NewFromConfig(config),
		Databasemigrationservice:  databasemigrationservice.NewFromConfig(config),
		Dax:                       dax.NewFromConfig(config),
		Directconnect:             directconnect.NewFromConfig(config),
		Docdb:                     docdb.NewFromConfig(config),
		Dynamodb:                  dynamodb.NewFromConfig(config),
		Securityhub:               securityhub.NewFromConfig(config),
		Ec2:                       ec2.NewFromConfig(config),
		Ecr:                       ecr.NewFromConfig(config),
		Ecrpublic:                 ecrpublic.NewFromConfig(config),
		Ecs:                       ecs.NewFromConfig(config),
		Kafka:                     kafka.NewFromConfig(config),
		Efs:                       efs.NewFromConfig(config),
		Eks:                       eks.NewFromConfig(config),
		Elasticache:               elasticache.NewFromConfig(config),
		Elasticbeanstalk:          elasticbeanstalk.NewFromConfig(config),
		Elasticloadbalancing:      elasticloadbalancing.NewFromConfig(config),
		Elasticloadbalancingv2:    elasticloadbalancingv2.NewFromConfig(config),
		Elasticsearchservice:      elasticsearchservice.NewFromConfig(config),
		Emr:                       emr.NewFromConfig(config),
		Eventbridge:               eventbridge.NewFromConfig(config),
		Firehose:                  firehose.NewFromConfig(config),
		Frauddetector:             frauddetector.NewFromConfig(config),
		Ses:                       ses.NewFromConfig(config),
		Fsx:                       fsx.NewFromConfig(config),
		Glacier:                   glacier.NewFromConfig(config),
		Glue:                      glue.NewFromConfig(config),
		Guardduty:                 guardduty.NewFromConfig(config),
		IAM:                       iam.NewFromConfig(config),
		Identitystore:             identitystore.NewFromConfig(config),
		Inspector:                 inspector.NewFromConfig(config),
		Inspector2:                inspector2.NewFromConfig(config),
		Iot:                       iot.NewFromConfig(config),
		Kinesis:                   kinesis.NewFromConfig(config),
		Kms:                       kms.NewFromConfig(config),
		Lambda:                    lambda.NewFromConfig(config),
		Lightsail:                 lightsail.NewFromConfig(config),
		Mq:                        mq.NewFromConfig(config),
		Mwaa:                      mwaa.NewFromConfig(config),
		RAM:                       ram.NewFromConfig(config),
		Neptune:                   neptune.NewFromConfig(config),
		Organizations:             organizations.NewFromConfig(config),
		Qldb:                      qldb.NewFromConfig(config),
		Rds:                       rds.NewFromConfig(config),
		Redshift:                  redshift.NewFromConfig(config),
		Resourcegroups:            resourcegroups.NewFromConfig(config),
		Route53:                   route53.NewFromConfig(config),
		Route53domains:            route53domains.NewFromConfig(config),
		S3:                        s3.NewFromConfig(config),
		S3control:                 s3control.NewFromConfig(config),
		Sagemaker:                 sagemaker.NewFromConfig(config),
		Secretsmanager:            secretsmanager.NewFromConfig(config),
		Servicecatalog:            servicecatalog.NewFromConfig(config),
		Servicecatalogappregistry: servicecatalogappregistry.NewFromConfig(config),
		Servicequotas:             servicequotas.NewFromConfig(config),
		Sesv2:                     sesv2.NewFromConfig(config),
		Shield:                    shield.NewFromConfig(config),
		Sns:                       sns.NewFromConfig(config),
		Sqs:                       sqs.NewFromConfig(config),
		Ssm:                       ssm.NewFromConfig(config),
		Ssoadmin:                  ssoadmin.NewFromConfig(config),
		Transfer:                  transfer.NewFromConfig(config),
		Waf:                       waf.NewFromConfig(config),
		Wafregional:               wafregional.NewFromConfig(config),
		Wafv2:                     wafv2.NewFromConfig(config),
		Workspaces:                workspaces.NewFromConfig(config),
		Xray:                      xray.NewFromConfig(config),
		S3manager:                 newS3ManagerFromConfig(config),
		Quicksight:                quicksight.NewFromConfig(config),
	}
}

type AwsServicesManager struct {
	AwsServicesManagerMap map[string]map[string]*AwsServices

	baseAwsConfig *aws.Config
}

func NewAwsServiceCache(baseAwsConfig *aws.Config) *AwsServicesManager {
	return &AwsServicesManager{
		baseAwsConfig:         baseAwsConfig,
		AwsServicesManagerMap: make(map[string]map[string]*AwsServices),
	}
}

func (x *AwsServicesManager) CacheCount() int {
	return len(x.AwsServicesManagerMap)
}

func (x *AwsServicesManager) initAwsServices(partition, region string) {
	if x.AwsServicesManagerMap[partition] == nil {
		x.AwsServicesManagerMap[partition] = make(map[string]*AwsServices)
	}
	awsConfig := x.baseAwsConfig.Copy()
	awsConfig.Region = region
	x.AwsServicesManagerMap[partition][region] = newAwsService(awsConfig)
}

func (x *AwsServicesManager) Get(partition, region string) *AwsServices {
	return x.AwsServicesManagerMap[partition][region]
}

type S3Manager struct {
	*s3.Client
}

func newS3ManagerFromConfig(cfg aws.Config) S3Manager {
	return S3Manager{
		s3.NewFromConfig(cfg),
	}
}

func (s3Manager S3Manager) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
	return manager.GetBucketRegion(ctx, s3Manager, bucket, optFns...)
}

type AccessanalyzerClient interface {
	GetAccessPreview(context.Context, *accessanalyzer.GetAccessPreviewInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAccessPreviewOutput, error)
	GetAnalyzedResource(context.Context, *accessanalyzer.GetAnalyzedResourceInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzedResourceOutput, error)
	GetAnalyzer(context.Context, *accessanalyzer.GetAnalyzerInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzerOutput, error)
	GetArchiveRule(context.Context, *accessanalyzer.GetArchiveRuleInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetArchiveRuleOutput, error)
	GetFinding(context.Context, *accessanalyzer.GetFindingInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingOutput, error)
	GetGeneratedPolicy(context.Context, *accessanalyzer.GetGeneratedPolicyInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetGeneratedPolicyOutput, error)
	ListAccessPreviewFindings(context.Context, *accessanalyzer.ListAccessPreviewFindingsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewFindingsOutput, error)
	ListAccessPreviews(context.Context, *accessanalyzer.ListAccessPreviewsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewsOutput, error)
	ListAnalyzedResources(context.Context, *accessanalyzer.ListAnalyzedResourcesInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
	ListAnalyzers(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error)
	ListArchiveRules(context.Context, *accessanalyzer.ListArchiveRulesInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListArchiveRulesOutput, error)
	ListFindings(context.Context, *accessanalyzer.ListFindingsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsOutput, error)
	ListPolicyGenerations(context.Context, *accessanalyzer.ListPolicyGenerationsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListPolicyGenerationsOutput, error)
	ListTagsForResource(context.Context, *accessanalyzer.ListTagsForResourceInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListTagsForResourceOutput, error)
}

type ApigatewayClient interface {
	GetAccount(context.Context, *apigateway.GetAccountInput, ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error)
	GetApiKey(context.Context, *apigateway.GetApiKeyInput, ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error)
	GetApiKeys(context.Context, *apigateway.GetApiKeysInput, ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	GetAuthorizer(context.Context, *apigateway.GetAuthorizerInput, ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error)
	GetAuthorizers(context.Context, *apigateway.GetAuthorizersInput, ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	GetBasePathMapping(context.Context, *apigateway.GetBasePathMappingInput, ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error)
	GetBasePathMappings(context.Context, *apigateway.GetBasePathMappingsInput, ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	GetClientCertificate(context.Context, *apigateway.GetClientCertificateInput, ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error)
	GetClientCertificates(context.Context, *apigateway.GetClientCertificatesInput, ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	GetDeployment(context.Context, *apigateway.GetDeploymentInput, ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error)
	GetDeployments(context.Context, *apigateway.GetDeploymentsInput, ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationPart(context.Context, *apigateway.GetDocumentationPartInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error)
	GetDocumentationParts(context.Context, *apigateway.GetDocumentationPartsInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersion(context.Context, *apigateway.GetDocumentationVersionInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error)
	GetDocumentationVersions(context.Context, *apigateway.GetDocumentationVersionsInput, ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDomainName(context.Context, *apigateway.GetDomainNameInput, ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error)
	GetDomainNames(context.Context, *apigateway.GetDomainNamesInput, ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	GetExport(context.Context, *apigateway.GetExportInput, ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error)
	GetGatewayResponse(context.Context, *apigateway.GetGatewayResponseInput, ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error)
	GetGatewayResponses(context.Context, *apigateway.GetGatewayResponsesInput, ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	GetIntegration(context.Context, *apigateway.GetIntegrationInput, ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error)
	GetIntegrationResponse(context.Context, *apigateway.GetIntegrationResponseInput, ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error)
	GetMethod(context.Context, *apigateway.GetMethodInput, ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error)
	GetMethodResponse(context.Context, *apigateway.GetMethodResponseInput, ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error)
	GetModel(context.Context, *apigateway.GetModelInput, ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error)
	GetModelTemplate(context.Context, *apigateway.GetModelTemplateInput, ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	GetModels(context.Context, *apigateway.GetModelsInput, ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	GetRequestValidator(context.Context, *apigateway.GetRequestValidatorInput, ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error)
	GetRequestValidators(context.Context, *apigateway.GetRequestValidatorsInput, ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	GetResource(context.Context, *apigateway.GetResourceInput, ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error)
	GetResources(context.Context, *apigateway.GetResourcesInput, ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetRestApi(context.Context, *apigateway.GetRestApiInput, ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error)
	GetRestApis(context.Context, *apigateway.GetRestApisInput, ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	GetSdk(context.Context, *apigateway.GetSdkInput, ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error)
	GetSdkType(context.Context, *apigateway.GetSdkTypeInput, ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error)
	GetSdkTypes(context.Context, *apigateway.GetSdkTypesInput, ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error)
	GetStage(context.Context, *apigateway.GetStageInput, ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error)
	GetStages(context.Context, *apigateway.GetStagesInput, ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetTags(context.Context, *apigateway.GetTagsInput, ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error)
	GetUsage(context.Context, *apigateway.GetUsageInput, ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error)
	GetUsagePlan(context.Context, *apigateway.GetUsagePlanInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error)
	GetUsagePlanKey(context.Context, *apigateway.GetUsagePlanKeyInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error)
	GetUsagePlanKeys(context.Context, *apigateway.GetUsagePlanKeysInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlans(context.Context, *apigateway.GetUsagePlansInput, ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	GetVpcLink(context.Context, *apigateway.GetVpcLinkInput, ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error)
	GetVpcLinks(context.Context, *apigateway.GetVpcLinksInput, ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
}

type Apigatewayv2Client interface {
	GetApi(context.Context, *apigatewayv2.GetApiInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiOutput, error)
	GetApiMapping(context.Context, *apigatewayv2.GetApiMappingInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappings(context.Context, *apigatewayv2.GetApiMappingsInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApis(context.Context, *apigatewayv2.GetApisInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error)
	GetAuthorizer(context.Context, *apigatewayv2.GetAuthorizerInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizers(context.Context, *apigatewayv2.GetAuthorizersInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error)
	GetDeployment(context.Context, *apigatewayv2.GetDeploymentInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeployments(context.Context, *apigatewayv2.GetDeploymentsInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDomainName(context.Context, *apigatewayv2.GetDomainNameInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNames(context.Context, *apigatewayv2.GetDomainNamesInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error)
	GetIntegration(context.Context, *apigatewayv2.GetIntegrationInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationResponse(context.Context, *apigatewayv2.GetIntegrationResponseInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponses(context.Context, *apigatewayv2.GetIntegrationResponsesInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrations(context.Context, *apigatewayv2.GetIntegrationsInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error)
	GetModel(context.Context, *apigatewayv2.GetModelInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelOutput, error)
	GetModelTemplate(context.Context, *apigatewayv2.GetModelTemplateInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModels(context.Context, *apigatewayv2.GetModelsInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error)
	GetRoute(context.Context, *apigatewayv2.GetRouteInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteOutput, error)
	GetRouteResponse(context.Context, *apigatewayv2.GetRouteResponseInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponses(context.Context, *apigatewayv2.GetRouteResponsesInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRoutes(context.Context, *apigatewayv2.GetRoutesInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error)
	GetStage(context.Context, *apigatewayv2.GetStageInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStageOutput, error)
	GetStages(context.Context, *apigatewayv2.GetStagesInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error)
	GetTags(context.Context, *apigatewayv2.GetTagsInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error)
	GetVpcLink(context.Context, *apigatewayv2.GetVpcLinkInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinkOutput, error)
	GetVpcLinks(context.Context, *apigatewayv2.GetVpcLinksInput, ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error)
}

type ApplicationautoscalingClient interface {
	DescribeScalableTargets(context.Context, *applicationautoscaling.DescribeScalableTargetsInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalableTargetsOutput, error)
	DescribeScalingActivities(context.Context, *applicationautoscaling.DescribeScalingActivitiesInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingPolicies(context.Context, *applicationautoscaling.DescribeScalingPoliciesInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
	DescribeScheduledActions(context.Context, *applicationautoscaling.DescribeScheduledActionsInput, ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScheduledActionsOutput, error)
}

type ApprunnerClient interface {
	DescribeAutoScalingConfiguration(context.Context, *apprunner.DescribeAutoScalingConfigurationInput, ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error)
	DescribeCustomDomains(context.Context, *apprunner.DescribeCustomDomainsInput, ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error)
	DescribeObservabilityConfiguration(context.Context, *apprunner.DescribeObservabilityConfigurationInput, ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error)
	DescribeService(context.Context, *apprunner.DescribeServiceInput, ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error)
	DescribeVpcConnector(context.Context, *apprunner.DescribeVpcConnectorInput, ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error)
	DescribeVpcIngressConnection(context.Context, *apprunner.DescribeVpcIngressConnectionInput, ...func(*apprunner.Options)) (*apprunner.DescribeVpcIngressConnectionOutput, error)
	ListAutoScalingConfigurations(context.Context, *apprunner.ListAutoScalingConfigurationsInput, ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error)
	ListConnections(context.Context, *apprunner.ListConnectionsInput, ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error)
	ListObservabilityConfigurations(context.Context, *apprunner.ListObservabilityConfigurationsInput, ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error)
	ListOperations(context.Context, *apprunner.ListOperationsInput, ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error)
	ListServices(context.Context, *apprunner.ListServicesInput, ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error)
	ListTagsForResource(context.Context, *apprunner.ListTagsForResourceInput, ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error)
	ListVpcConnectors(context.Context, *apprunner.ListVpcConnectorsInput, ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error)
	ListVpcIngressConnections(context.Context, *apprunner.ListVpcIngressConnectionsInput, ...func(*apprunner.Options)) (*apprunner.ListVpcIngressConnectionsOutput, error)
}

type AppsyncClient interface {
	GetApiAssociation(context.Context, *appsync.GetApiAssociationInput, ...func(*appsync.Options)) (*appsync.GetApiAssociationOutput, error)
	GetApiCache(context.Context, *appsync.GetApiCacheInput, ...func(*appsync.Options)) (*appsync.GetApiCacheOutput, error)
	GetDataSource(context.Context, *appsync.GetDataSourceInput, ...func(*appsync.Options)) (*appsync.GetDataSourceOutput, error)
	GetDomainName(context.Context, *appsync.GetDomainNameInput, ...func(*appsync.Options)) (*appsync.GetDomainNameOutput, error)
	GetFunction(context.Context, *appsync.GetFunctionInput, ...func(*appsync.Options)) (*appsync.GetFunctionOutput, error)
	GetGraphqlApi(context.Context, *appsync.GetGraphqlApiInput, ...func(*appsync.Options)) (*appsync.GetGraphqlApiOutput, error)
	GetIntrospectionSchema(context.Context, *appsync.GetIntrospectionSchemaInput, ...func(*appsync.Options)) (*appsync.GetIntrospectionSchemaOutput, error)
	GetResolver(context.Context, *appsync.GetResolverInput, ...func(*appsync.Options)) (*appsync.GetResolverOutput, error)
	GetSchemaCreationStatus(context.Context, *appsync.GetSchemaCreationStatusInput, ...func(*appsync.Options)) (*appsync.GetSchemaCreationStatusOutput, error)
	GetType(context.Context, *appsync.GetTypeInput, ...func(*appsync.Options)) (*appsync.GetTypeOutput, error)
	ListApiKeys(context.Context, *appsync.ListApiKeysInput, ...func(*appsync.Options)) (*appsync.ListApiKeysOutput, error)
	ListDataSources(context.Context, *appsync.ListDataSourcesInput, ...func(*appsync.Options)) (*appsync.ListDataSourcesOutput, error)
	ListDomainNames(context.Context, *appsync.ListDomainNamesInput, ...func(*appsync.Options)) (*appsync.ListDomainNamesOutput, error)
	ListFunctions(context.Context, *appsync.ListFunctionsInput, ...func(*appsync.Options)) (*appsync.ListFunctionsOutput, error)
	ListGraphqlApis(context.Context, *appsync.ListGraphqlApisInput, ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error)
	ListResolvers(context.Context, *appsync.ListResolversInput, ...func(*appsync.Options)) (*appsync.ListResolversOutput, error)
	ListResolversByFunction(context.Context, *appsync.ListResolversByFunctionInput, ...func(*appsync.Options)) (*appsync.ListResolversByFunctionOutput, error)
	ListTagsForResource(context.Context, *appsync.ListTagsForResourceInput, ...func(*appsync.Options)) (*appsync.ListTagsForResourceOutput, error)
	ListTypes(context.Context, *appsync.ListTypesInput, ...func(*appsync.Options)) (*appsync.ListTypesOutput, error)
}

type AthenaClient interface {
	BatchGetNamedQuery(context.Context, *athena.BatchGetNamedQueryInput, ...func(*athena.Options)) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetPreparedStatement(context.Context, *athena.BatchGetPreparedStatementInput, ...func(*athena.Options)) (*athena.BatchGetPreparedStatementOutput, error)
	BatchGetQueryExecution(context.Context, *athena.BatchGetQueryExecutionInput, ...func(*athena.Options)) (*athena.BatchGetQueryExecutionOutput, error)
	GetDataCatalog(context.Context, *athena.GetDataCatalogInput, ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error)
	GetDatabase(context.Context, *athena.GetDatabaseInput, ...func(*athena.Options)) (*athena.GetDatabaseOutput, error)
	GetNamedQuery(context.Context, *athena.GetNamedQueryInput, ...func(*athena.Options)) (*athena.GetNamedQueryOutput, error)
	GetPreparedStatement(context.Context, *athena.GetPreparedStatementInput, ...func(*athena.Options)) (*athena.GetPreparedStatementOutput, error)
	GetQueryExecution(context.Context, *athena.GetQueryExecutionInput, ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error)
	GetQueryResults(context.Context, *athena.GetQueryResultsInput, ...func(*athena.Options)) (*athena.GetQueryResultsOutput, error)
	GetQueryRuntimeStatistics(context.Context, *athena.GetQueryRuntimeStatisticsInput, ...func(*athena.Options)) (*athena.GetQueryRuntimeStatisticsOutput, error)
	GetTableMetadata(context.Context, *athena.GetTableMetadataInput, ...func(*athena.Options)) (*athena.GetTableMetadataOutput, error)
	GetWorkGroup(context.Context, *athena.GetWorkGroupInput, ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error)
	ListDataCatalogs(context.Context, *athena.ListDataCatalogsInput, ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error)
	ListDatabases(context.Context, *athena.ListDatabasesInput, ...func(*athena.Options)) (*athena.ListDatabasesOutput, error)
	ListEngineVersions(context.Context, *athena.ListEngineVersionsInput, ...func(*athena.Options)) (*athena.ListEngineVersionsOutput, error)
	ListNamedQueries(context.Context, *athena.ListNamedQueriesInput, ...func(*athena.Options)) (*athena.ListNamedQueriesOutput, error)
	ListPreparedStatements(context.Context, *athena.ListPreparedStatementsInput, ...func(*athena.Options)) (*athena.ListPreparedStatementsOutput, error)
	ListQueryExecutions(context.Context, *athena.ListQueryExecutionsInput, ...func(*athena.Options)) (*athena.ListQueryExecutionsOutput, error)
	ListTableMetadata(context.Context, *athena.ListTableMetadataInput, ...func(*athena.Options)) (*athena.ListTableMetadataOutput, error)
	ListTagsForResource(context.Context, *athena.ListTagsForResourceInput, ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error)
	ListWorkGroups(context.Context, *athena.ListWorkGroupsInput, ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error)
}

type AutoscalingClient interface {
	DescribeAccountLimits(context.Context, *autoscaling.DescribeAccountLimitsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeAccountLimitsOutput, error)
	DescribeAdjustmentTypes(context.Context, *autoscaling.DescribeAdjustmentTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeAdjustmentTypesOutput, error)
	DescribeAutoScalingGroups(context.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribeAutoScalingInstances(context.Context, *autoscaling.DescribeAutoScalingInstancesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
	DescribeAutoScalingNotificationTypes(context.Context, *autoscaling.DescribeAutoScalingNotificationTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
	DescribeInstanceRefreshes(context.Context, *autoscaling.DescribeInstanceRefreshesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeInstanceRefreshesOutput, error)
	DescribeLaunchConfigurations(context.Context, *autoscaling.DescribeLaunchConfigurationsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeLifecycleHookTypes(context.Context, *autoscaling.DescribeLifecycleHookTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
	DescribeLifecycleHooks(context.Context, *autoscaling.DescribeLifecycleHooksInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeLoadBalancerTargetGroups(context.Context, *autoscaling.DescribeLoadBalancerTargetGroupsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLoadBalancers(context.Context, *autoscaling.DescribeLoadBalancersInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeMetricCollectionTypes(context.Context, *autoscaling.DescribeMetricCollectionTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
	DescribeNotificationConfigurations(context.Context, *autoscaling.DescribeNotificationConfigurationsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribePolicies(context.Context, *autoscaling.DescribePoliciesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error)
	DescribeScalingActivities(context.Context, *autoscaling.DescribeScalingActivitiesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingProcessTypes(context.Context, *autoscaling.DescribeScalingProcessTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeScalingProcessTypesOutput, error)
	DescribeScheduledActions(context.Context, *autoscaling.DescribeScheduledActionsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error)
	DescribeTags(context.Context, *autoscaling.DescribeTagsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error)
	DescribeTerminationPolicyTypes(context.Context, *autoscaling.DescribeTerminationPolicyTypesInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
	DescribeWarmPool(context.Context, *autoscaling.DescribeWarmPoolInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeWarmPoolOutput, error)
	GetPredictiveScalingForecast(context.Context, *autoscaling.GetPredictiveScalingForecastInput, ...func(*autoscaling.Options)) (*autoscaling.GetPredictiveScalingForecastOutput, error)
}

type BackupClient interface {
	DescribeBackupJob(context.Context, *backup.DescribeBackupJobInput, ...func(*backup.Options)) (*backup.DescribeBackupJobOutput, error)
	DescribeBackupVault(context.Context, *backup.DescribeBackupVaultInput, ...func(*backup.Options)) (*backup.DescribeBackupVaultOutput, error)
	DescribeCopyJob(context.Context, *backup.DescribeCopyJobInput, ...func(*backup.Options)) (*backup.DescribeCopyJobOutput, error)
	DescribeFramework(context.Context, *backup.DescribeFrameworkInput, ...func(*backup.Options)) (*backup.DescribeFrameworkOutput, error)
	DescribeGlobalSettings(context.Context, *backup.DescribeGlobalSettingsInput, ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error)
	DescribeProtectedResource(context.Context, *backup.DescribeProtectedResourceInput, ...func(*backup.Options)) (*backup.DescribeProtectedResourceOutput, error)
	DescribeRecoveryPoint(context.Context, *backup.DescribeRecoveryPointInput, ...func(*backup.Options)) (*backup.DescribeRecoveryPointOutput, error)
	DescribeRegionSettings(context.Context, *backup.DescribeRegionSettingsInput, ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error)
	DescribeReportJob(context.Context, *backup.DescribeReportJobInput, ...func(*backup.Options)) (*backup.DescribeReportJobOutput, error)
	DescribeReportPlan(context.Context, *backup.DescribeReportPlanInput, ...func(*backup.Options)) (*backup.DescribeReportPlanOutput, error)
	DescribeRestoreJob(context.Context, *backup.DescribeRestoreJobInput, ...func(*backup.Options)) (*backup.DescribeRestoreJobOutput, error)
	GetBackupPlan(context.Context, *backup.GetBackupPlanInput, ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error)
	GetBackupPlanFromJSON(context.Context, *backup.GetBackupPlanFromJSONInput, ...func(*backup.Options)) (*backup.GetBackupPlanFromJSONOutput, error)
	GetBackupPlanFromTemplate(context.Context, *backup.GetBackupPlanFromTemplateInput, ...func(*backup.Options)) (*backup.GetBackupPlanFromTemplateOutput, error)
	GetBackupSelection(context.Context, *backup.GetBackupSelectionInput, ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error)
	GetBackupVaultAccessPolicy(context.Context, *backup.GetBackupVaultAccessPolicyInput, ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error)
	GetBackupVaultNotifications(context.Context, *backup.GetBackupVaultNotificationsInput, ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error)
	GetRecoveryPointRestoreMetadata(context.Context, *backup.GetRecoveryPointRestoreMetadataInput, ...func(*backup.Options)) (*backup.GetRecoveryPointRestoreMetadataOutput, error)
	GetSupportedResourceTypes(context.Context, *backup.GetSupportedResourceTypesInput, ...func(*backup.Options)) (*backup.GetSupportedResourceTypesOutput, error)
	ListBackupJobs(context.Context, *backup.ListBackupJobsInput, ...func(*backup.Options)) (*backup.ListBackupJobsOutput, error)
	ListBackupPlanTemplates(context.Context, *backup.ListBackupPlanTemplatesInput, ...func(*backup.Options)) (*backup.ListBackupPlanTemplatesOutput, error)
	ListBackupPlanVersions(context.Context, *backup.ListBackupPlanVersionsInput, ...func(*backup.Options)) (*backup.ListBackupPlanVersionsOutput, error)
	ListBackupPlans(context.Context, *backup.ListBackupPlansInput, ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error)
	ListBackupSelections(context.Context, *backup.ListBackupSelectionsInput, ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error)
	ListBackupVaults(context.Context, *backup.ListBackupVaultsInput, ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error)
	ListCopyJobs(context.Context, *backup.ListCopyJobsInput, ...func(*backup.Options)) (*backup.ListCopyJobsOutput, error)
	ListFrameworks(context.Context, *backup.ListFrameworksInput, ...func(*backup.Options)) (*backup.ListFrameworksOutput, error)
	ListProtectedResources(context.Context, *backup.ListProtectedResourcesInput, ...func(*backup.Options)) (*backup.ListProtectedResourcesOutput, error)
	ListRecoveryPointsByBackupVault(context.Context, *backup.ListRecoveryPointsByBackupVaultInput, ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error)
	ListRecoveryPointsByResource(context.Context, *backup.ListRecoveryPointsByResourceInput, ...func(*backup.Options)) (*backup.ListRecoveryPointsByResourceOutput, error)
	ListReportJobs(context.Context, *backup.ListReportJobsInput, ...func(*backup.Options)) (*backup.ListReportJobsOutput, error)
	ListReportPlans(context.Context, *backup.ListReportPlansInput, ...func(*backup.Options)) (*backup.ListReportPlansOutput, error)
	ListRestoreJobs(context.Context, *backup.ListRestoreJobsInput, ...func(*backup.Options)) (*backup.ListRestoreJobsOutput, error)
	ListTags(context.Context, *backup.ListTagsInput, ...func(*backup.Options)) (*backup.ListTagsOutput, error)
}

type CloudformationClient interface {
	DescribeAccountLimits(context.Context, *cloudformation.DescribeAccountLimitsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeChangeSet(context.Context, *cloudformation.DescribeChangeSetInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetHooks(context.Context, *cloudformation.DescribeChangeSetHooksInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeChangeSetHooksOutput, error)
	DescribePublisher(context.Context, *cloudformation.DescribePublisherInput, ...func(*cloudformation.Options)) (*cloudformation.DescribePublisherOutput, error)
	DescribeStackDriftDetectionStatus(context.Context, *cloudformation.DescribeStackDriftDetectionStatusInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
	DescribeStackEvents(context.Context, *cloudformation.DescribeStackEventsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackInstance(context.Context, *cloudformation.DescribeStackInstanceInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackResource(context.Context, *cloudformation.DescribeStackResourceInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceDrifts(context.Context, *cloudformation.DescribeStackResourceDriftsInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourceDriftsOutput, error)
	DescribeStackResources(context.Context, *cloudformation.DescribeStackResourcesInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackSet(context.Context, *cloudformation.DescribeStackSetInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetOperation(context.Context, *cloudformation.DescribeStackSetOperationInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStacks(context.Context, *cloudformation.DescribeStacksInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)
	DescribeType(context.Context, *cloudformation.DescribeTypeInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeOutput, error)
	DescribeTypeRegistration(context.Context, *cloudformation.DescribeTypeRegistrationInput, ...func(*cloudformation.Options)) (*cloudformation.DescribeTypeRegistrationOutput, error)
	GetStackPolicy(context.Context, *cloudformation.GetStackPolicyInput, ...func(*cloudformation.Options)) (*cloudformation.GetStackPolicyOutput, error)
	GetTemplate(context.Context, *cloudformation.GetTemplateInput, ...func(*cloudformation.Options)) (*cloudformation.GetTemplateOutput, error)
	GetTemplateSummary(context.Context, *cloudformation.GetTemplateSummaryInput, ...func(*cloudformation.Options)) (*cloudformation.GetTemplateSummaryOutput, error)
	ListChangeSets(context.Context, *cloudformation.ListChangeSetsInput, ...func(*cloudformation.Options)) (*cloudformation.ListChangeSetsOutput, error)
	ListExports(context.Context, *cloudformation.ListExportsInput, ...func(*cloudformation.Options)) (*cloudformation.ListExportsOutput, error)
	ListImports(context.Context, *cloudformation.ListImportsInput, ...func(*cloudformation.Options)) (*cloudformation.ListImportsOutput, error)
	ListStackInstances(context.Context, *cloudformation.ListStackInstancesInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackInstancesOutput, error)
	ListStackResources(context.Context, *cloudformation.ListStackResourcesInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error)
	ListStackSetOperationResults(context.Context, *cloudformation.ListStackSetOperationResultsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperations(context.Context, *cloudformation.ListStackSetOperationsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSets(context.Context, *cloudformation.ListStackSetsInput, ...func(*cloudformation.Options)) (*cloudformation.ListStackSetsOutput, error)
	ListStacks(context.Context, *cloudformation.ListStacksInput, ...func(*cloudformation.Options)) (*cloudformation.ListStacksOutput, error)
	ListTypeRegistrations(context.Context, *cloudformation.ListTypeRegistrationsInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypeRegistrationsOutput, error)
	ListTypeVersions(context.Context, *cloudformation.ListTypeVersionsInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypeVersionsOutput, error)
	ListTypes(context.Context, *cloudformation.ListTypesInput, ...func(*cloudformation.Options)) (*cloudformation.ListTypesOutput, error)
}

type CloudfrontClient interface {
	DescribeFunction(context.Context, *cloudfront.DescribeFunctionInput, ...func(*cloudfront.Options)) (*cloudfront.DescribeFunctionOutput, error)
	GetCachePolicy(context.Context, *cloudfront.GetCachePolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyConfig(context.Context, *cloudfront.GetCachePolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCloudFrontOriginAccessIdentity(context.Context, *cloudfront.GetCloudFrontOriginAccessIdentityInput, ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityConfig(context.Context, *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetDistribution(context.Context, *cloudfront.GetDistributionInput, ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error)
	GetDistributionConfig(context.Context, *cloudfront.GetDistributionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetDistributionConfigOutput, error)
	GetFieldLevelEncryption(context.Context, *cloudfront.GetFieldLevelEncryptionInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionConfig(context.Context, *cloudfront.GetFieldLevelEncryptionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionProfile(context.Context, *cloudfront.GetFieldLevelEncryptionProfileInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileConfig(context.Context, *cloudfront.GetFieldLevelEncryptionProfileConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFunction(context.Context, *cloudfront.GetFunctionInput, ...func(*cloudfront.Options)) (*cloudfront.GetFunctionOutput, error)
	GetInvalidation(context.Context, *cloudfront.GetInvalidationInput, ...func(*cloudfront.Options)) (*cloudfront.GetInvalidationOutput, error)
	GetKeyGroup(context.Context, *cloudfront.GetKeyGroupInput, ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupOutput, error)
	GetKeyGroupConfig(context.Context, *cloudfront.GetKeyGroupConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetKeyGroupConfigOutput, error)
	GetMonitoringSubscription(context.Context, *cloudfront.GetMonitoringSubscriptionInput, ...func(*cloudfront.Options)) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetOriginAccessControl(context.Context, *cloudfront.GetOriginAccessControlInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlOutput, error)
	GetOriginAccessControlConfig(context.Context, *cloudfront.GetOriginAccessControlConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginAccessControlConfigOutput, error)
	GetOriginRequestPolicy(context.Context, *cloudfront.GetOriginRequestPolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyConfig(context.Context, *cloudfront.GetOriginRequestPolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetPublicKey(context.Context, *cloudfront.GetPublicKeyInput, ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyConfig(context.Context, *cloudfront.GetPublicKeyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetRealtimeLogConfig(context.Context, *cloudfront.GetRealtimeLogConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetResponseHeadersPolicy(context.Context, *cloudfront.GetResponseHeadersPolicyInput, ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyOutput, error)
	GetResponseHeadersPolicyConfig(context.Context, *cloudfront.GetResponseHeadersPolicyConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error)
	GetStreamingDistribution(context.Context, *cloudfront.GetStreamingDistributionInput, ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionConfig(context.Context, *cloudfront.GetStreamingDistributionConfigInput, ...func(*cloudfront.Options)) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	ListCachePolicies(context.Context, *cloudfront.ListCachePoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error)
	ListCloudFrontOriginAccessIdentities(context.Context, *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, ...func(*cloudfront.Options)) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListConflictingAliases(context.Context, *cloudfront.ListConflictingAliasesInput, ...func(*cloudfront.Options)) (*cloudfront.ListConflictingAliasesOutput, error)
	ListDistributions(context.Context, *cloudfront.ListDistributionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsByCachePolicyId(context.Context, *cloudfront.ListDistributionsByCachePolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByKeyGroup(context.Context, *cloudfront.ListDistributionsByKeyGroupInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByKeyGroupOutput, error)
	ListDistributionsByOriginRequestPolicyId(context.Context, *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByRealtimeLogConfig(context.Context, *cloudfront.ListDistributionsByRealtimeLogConfigInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByResponseHeadersPolicyId(context.Context, *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error)
	ListDistributionsByWebACLId(context.Context, *cloudfront.ListDistributionsByWebACLIdInput, ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListFieldLevelEncryptionConfigs(context.Context, *cloudfront.ListFieldLevelEncryptionConfigsInput, ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionProfiles(context.Context, *cloudfront.ListFieldLevelEncryptionProfilesInput, ...func(*cloudfront.Options)) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFunctions(context.Context, *cloudfront.ListFunctionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListFunctionsOutput, error)
	ListInvalidations(context.Context, *cloudfront.ListInvalidationsInput, ...func(*cloudfront.Options)) (*cloudfront.ListInvalidationsOutput, error)
	ListKeyGroups(context.Context, *cloudfront.ListKeyGroupsInput, ...func(*cloudfront.Options)) (*cloudfront.ListKeyGroupsOutput, error)
	ListOriginAccessControls(context.Context, *cloudfront.ListOriginAccessControlsInput, ...func(*cloudfront.Options)) (*cloudfront.ListOriginAccessControlsOutput, error)
	ListOriginRequestPolicies(context.Context, *cloudfront.ListOriginRequestPoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListPublicKeys(context.Context, *cloudfront.ListPublicKeysInput, ...func(*cloudfront.Options)) (*cloudfront.ListPublicKeysOutput, error)
	ListRealtimeLogConfigs(context.Context, *cloudfront.ListRealtimeLogConfigsInput, ...func(*cloudfront.Options)) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListResponseHeadersPolicies(context.Context, *cloudfront.ListResponseHeadersPoliciesInput, ...func(*cloudfront.Options)) (*cloudfront.ListResponseHeadersPoliciesOutput, error)
	ListStreamingDistributions(context.Context, *cloudfront.ListStreamingDistributionsInput, ...func(*cloudfront.Options)) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListTagsForResource(context.Context, *cloudfront.ListTagsForResourceInput, ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error)
}

type Cloudhsmv2Client interface {
	DescribeBackups(context.Context, *cloudhsmv2.DescribeBackupsInput, ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeClusters(context.Context, *cloudhsmv2.DescribeClustersInput, ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error)
	ListTags(context.Context, *cloudhsmv2.ListTagsInput, ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ListTagsOutput, error)
}

type CloudtrailClient interface {
	DescribeQuery(context.Context, *cloudtrail.DescribeQueryInput, ...func(*cloudtrail.Options)) (*cloudtrail.DescribeQueryOutput, error)
	DescribeTrails(context.Context, *cloudtrail.DescribeTrailsInput, ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error)
	GetChannel(context.Context, *cloudtrail.GetChannelInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetChannelOutput, error)
	GetEventDataStore(context.Context, *cloudtrail.GetEventDataStoreInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetEventDataStoreOutput, error)
	GetEventSelectors(context.Context, *cloudtrail.GetEventSelectorsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error)
	GetImport(context.Context, *cloudtrail.GetImportInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetImportOutput, error)
	GetInsightSelectors(context.Context, *cloudtrail.GetInsightSelectorsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetInsightSelectorsOutput, error)
	GetQueryResults(context.Context, *cloudtrail.GetQueryResultsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetQueryResultsOutput, error)
	GetTrail(context.Context, *cloudtrail.GetTrailInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailOutput, error)
	GetTrailStatus(context.Context, *cloudtrail.GetTrailStatusInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error)
	ListChannels(context.Context, *cloudtrail.ListChannelsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListChannelsOutput, error)
	ListEventDataStores(context.Context, *cloudtrail.ListEventDataStoresInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListEventDataStoresOutput, error)
	ListImportFailures(context.Context, *cloudtrail.ListImportFailuresInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListImportFailuresOutput, error)
	ListImports(context.Context, *cloudtrail.ListImportsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListImportsOutput, error)
	ListPublicKeys(context.Context, *cloudtrail.ListPublicKeysInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListPublicKeysOutput, error)
	ListQueries(context.Context, *cloudtrail.ListQueriesInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListQueriesOutput, error)
	ListTags(context.Context, *cloudtrail.ListTagsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error)
	ListTrails(context.Context, *cloudtrail.ListTrailsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListTrailsOutput, error)
}

type CloudwatchClient interface {
	DescribeAlarmHistory(context.Context, *cloudwatch.DescribeAlarmHistoryInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarms(context.Context, *cloudwatch.DescribeAlarmsInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsForMetric(context.Context, *cloudwatch.DescribeAlarmsForMetricInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAnomalyDetectors(context.Context, *cloudwatch.DescribeAnomalyDetectorsInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeInsightRules(context.Context, *cloudwatch.DescribeInsightRulesInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeInsightRulesOutput, error)
	GetDashboard(context.Context, *cloudwatch.GetDashboardInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetDashboardOutput, error)
	GetInsightRuleReport(context.Context, *cloudwatch.GetInsightRuleReportInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetMetricData(context.Context, *cloudwatch.GetMetricDataInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricStatistics(context.Context, *cloudwatch.GetMetricStatisticsInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStream(context.Context, *cloudwatch.GetMetricStreamInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStreamOutput, error)
	GetMetricWidgetImage(context.Context, *cloudwatch.GetMetricWidgetImageInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricWidgetImageOutput, error)
	ListDashboards(context.Context, *cloudwatch.ListDashboardsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListDashboardsOutput, error)
	ListManagedInsightRules(context.Context, *cloudwatch.ListManagedInsightRulesInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListManagedInsightRulesOutput, error)
	ListMetricStreams(context.Context, *cloudwatch.ListMetricStreamsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricStreamsOutput, error)
	ListMetrics(context.Context, *cloudwatch.ListMetricsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricsOutput, error)
	ListTagsForResource(context.Context, *cloudwatch.ListTagsForResourceInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListTagsForResourceOutput, error)
}

type CloudwatchlogsClient interface {
	DescribeDestinations(context.Context, *cloudwatchlogs.DescribeDestinationsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDestinationsOutput, error)
	DescribeExportTasks(context.Context, *cloudwatchlogs.DescribeExportTasksInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeExportTasksOutput, error)
	DescribeLogGroups(context.Context, *cloudwatchlogs.DescribeLogGroupsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	DescribeLogStreams(context.Context, *cloudwatchlogs.DescribeLogStreamsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogStreamsOutput, error)
	DescribeMetricFilters(context.Context, *cloudwatchlogs.DescribeMetricFiltersInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
	DescribeQueries(context.Context, *cloudwatchlogs.DescribeQueriesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueriesOutput, error)
	DescribeQueryDefinitions(context.Context, *cloudwatchlogs.DescribeQueryDefinitionsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error)
	DescribeResourcePolicies(context.Context, *cloudwatchlogs.DescribeResourcePoliciesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error)
	DescribeSubscriptionFilters(context.Context, *cloudwatchlogs.DescribeSubscriptionFiltersInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error)
	GetLogEvents(context.Context, *cloudwatchlogs.GetLogEventsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error)
	GetLogGroupFields(context.Context, *cloudwatchlogs.GetLogGroupFieldsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogGroupFieldsOutput, error)
	GetLogRecord(context.Context, *cloudwatchlogs.GetLogRecordInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogRecordOutput, error)
	GetQueryResults(context.Context, *cloudwatchlogs.GetQueryResultsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetQueryResultsOutput, error)
	ListTagsForResource(context.Context, *cloudwatchlogs.ListTagsForResourceInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsForResourceOutput, error)
	ListTagsLogGroup(context.Context, *cloudwatchlogs.ListTagsLogGroupInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error)
}

type CodebuildClient interface {
	BatchGetBuildBatches(context.Context, *codebuild.BatchGetBuildBatchesInput, ...func(*codebuild.Options)) (*codebuild.BatchGetBuildBatchesOutput, error)
	BatchGetBuilds(context.Context, *codebuild.BatchGetBuildsInput, ...func(*codebuild.Options)) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetProjects(context.Context, *codebuild.BatchGetProjectsInput, ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetReportGroups(context.Context, *codebuild.BatchGetReportGroupsInput, ...func(*codebuild.Options)) (*codebuild.BatchGetReportGroupsOutput, error)
	BatchGetReports(context.Context, *codebuild.BatchGetReportsInput, ...func(*codebuild.Options)) (*codebuild.BatchGetReportsOutput, error)
	DescribeCodeCoverages(context.Context, *codebuild.DescribeCodeCoveragesInput, ...func(*codebuild.Options)) (*codebuild.DescribeCodeCoveragesOutput, error)
	DescribeTestCases(context.Context, *codebuild.DescribeTestCasesInput, ...func(*codebuild.Options)) (*codebuild.DescribeTestCasesOutput, error)
	GetReportGroupTrend(context.Context, *codebuild.GetReportGroupTrendInput, ...func(*codebuild.Options)) (*codebuild.GetReportGroupTrendOutput, error)
	GetResourcePolicy(context.Context, *codebuild.GetResourcePolicyInput, ...func(*codebuild.Options)) (*codebuild.GetResourcePolicyOutput, error)
	ListBuildBatches(context.Context, *codebuild.ListBuildBatchesInput, ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesOutput, error)
	ListBuildBatchesForProject(context.Context, *codebuild.ListBuildBatchesForProjectInput, ...func(*codebuild.Options)) (*codebuild.ListBuildBatchesForProjectOutput, error)
	ListBuilds(context.Context, *codebuild.ListBuildsInput, ...func(*codebuild.Options)) (*codebuild.ListBuildsOutput, error)
	ListBuildsForProject(context.Context, *codebuild.ListBuildsForProjectInput, ...func(*codebuild.Options)) (*codebuild.ListBuildsForProjectOutput, error)
	ListCuratedEnvironmentImages(context.Context, *codebuild.ListCuratedEnvironmentImagesInput, ...func(*codebuild.Options)) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListProjects(context.Context, *codebuild.ListProjectsInput, ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error)
	ListReportGroups(context.Context, *codebuild.ListReportGroupsInput, ...func(*codebuild.Options)) (*codebuild.ListReportGroupsOutput, error)
	ListReports(context.Context, *codebuild.ListReportsInput, ...func(*codebuild.Options)) (*codebuild.ListReportsOutput, error)
	ListReportsForReportGroup(context.Context, *codebuild.ListReportsForReportGroupInput, ...func(*codebuild.Options)) (*codebuild.ListReportsForReportGroupOutput, error)
	ListSharedProjects(context.Context, *codebuild.ListSharedProjectsInput, ...func(*codebuild.Options)) (*codebuild.ListSharedProjectsOutput, error)
	ListSharedReportGroups(context.Context, *codebuild.ListSharedReportGroupsInput, ...func(*codebuild.Options)) (*codebuild.ListSharedReportGroupsOutput, error)
	ListSourceCredentials(context.Context, *codebuild.ListSourceCredentialsInput, ...func(*codebuild.Options)) (*codebuild.ListSourceCredentialsOutput, error)
}

type CodepipelineClient interface {
	GetActionType(context.Context, *codepipeline.GetActionTypeInput, ...func(*codepipeline.Options)) (*codepipeline.GetActionTypeOutput, error)
	GetJobDetails(context.Context, *codepipeline.GetJobDetailsInput, ...func(*codepipeline.Options)) (*codepipeline.GetJobDetailsOutput, error)
	GetPipeline(context.Context, *codepipeline.GetPipelineInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error)
	GetPipelineExecution(context.Context, *codepipeline.GetPipelineExecutionInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineState(context.Context, *codepipeline.GetPipelineStateInput, ...func(*codepipeline.Options)) (*codepipeline.GetPipelineStateOutput, error)
	GetThirdPartyJobDetails(context.Context, *codepipeline.GetThirdPartyJobDetailsInput, ...func(*codepipeline.Options)) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	ListActionExecutions(context.Context, *codepipeline.ListActionExecutionsInput, ...func(*codepipeline.Options)) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionTypes(context.Context, *codepipeline.ListActionTypesInput, ...func(*codepipeline.Options)) (*codepipeline.ListActionTypesOutput, error)
	ListPipelineExecutions(context.Context, *codepipeline.ListPipelineExecutionsInput, ...func(*codepipeline.Options)) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelines(context.Context, *codepipeline.ListPipelinesInput, ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error)
	ListTagsForResource(context.Context, *codepipeline.ListTagsForResourceInput, ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error)
	ListWebhooks(context.Context, *codepipeline.ListWebhooksInput, ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error)
}

type CognitoidentityClient interface {
	DescribeIdentity(context.Context, *cognitoidentity.DescribeIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityOutput, error)
	DescribeIdentityPool(context.Context, *cognitoidentity.DescribeIdentityPoolInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error)
	GetCredentialsForIdentity(context.Context, *cognitoidentity.GetCredentialsForIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetId(context.Context, *cognitoidentity.GetIdInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdOutput, error)
	GetIdentityPoolRoles(context.Context, *cognitoidentity.GetIdentityPoolRolesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetOpenIdToken(context.Context, *cognitoidentity.GetOpenIdTokenInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenForDeveloperIdentity(context.Context, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	GetPrincipalTagAttributeMap(context.Context, *cognitoidentity.GetPrincipalTagAttributeMapInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.GetPrincipalTagAttributeMapOutput, error)
	ListIdentities(context.Context, *cognitoidentity.ListIdentitiesInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentityPools(context.Context, *cognitoidentity.ListIdentityPoolsInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListTagsForResource(context.Context, *cognitoidentity.ListTagsForResourceInput, ...func(*cognitoidentity.Options)) (*cognitoidentity.ListTagsForResourceOutput, error)
}

type CognitoidentityproviderClient interface {
	DescribeIdentityProvider(context.Context, *cognitoidentityprovider.DescribeIdentityProviderInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeResourceServer(context.Context, *cognitoidentityprovider.DescribeResourceServerInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeRiskConfiguration(context.Context, *cognitoidentityprovider.DescribeRiskConfigurationInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
	DescribeUserImportJob(context.Context, *cognitoidentityprovider.DescribeUserImportJobInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserPool(context.Context, *cognitoidentityprovider.DescribeUserPoolInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolClient(context.Context, *cognitoidentityprovider.DescribeUserPoolClientInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolDomain(context.Context, *cognitoidentityprovider.DescribeUserPoolDomainInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	GetCSVHeader(context.Context, *cognitoidentityprovider.GetCSVHeaderInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetDevice(context.Context, *cognitoidentityprovider.GetDeviceInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetGroup(context.Context, *cognitoidentityprovider.GetGroupInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetGroupOutput, error)
	GetIdentityProviderByIdentifier(context.Context, *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetSigningCertificate(context.Context, *cognitoidentityprovider.GetSigningCertificateInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
	GetUICustomization(context.Context, *cognitoidentityprovider.GetUICustomizationInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUser(context.Context, *cognitoidentityprovider.GetUserInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserAttributeVerificationCode(context.Context, *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserPoolMfaConfig(context.Context, *cognitoidentityprovider.GetUserPoolMfaConfigInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
	ListDevices(context.Context, *cognitoidentityprovider.ListDevicesInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListGroups(context.Context, *cognitoidentityprovider.ListGroupsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListIdentityProviders(context.Context, *cognitoidentityprovider.ListIdentityProvidersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListResourceServers(context.Context, *cognitoidentityprovider.ListResourceServersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListTagsForResource(context.Context, *cognitoidentityprovider.ListTagsForResourceInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
	ListUserImportJobs(context.Context, *cognitoidentityprovider.ListUserImportJobsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserPoolClients(context.Context, *cognitoidentityprovider.ListUserPoolClientsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPools(context.Context, *cognitoidentityprovider.ListUserPoolsInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUsers(context.Context, *cognitoidentityprovider.ListUsersInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersInGroup(context.Context, *cognitoidentityprovider.ListUsersInGroupInput, ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
}

type ConfigserviceClient interface {
	BatchGetAggregateResourceConfig(context.Context, *configservice.BatchGetAggregateResourceConfigInput, ...func(*configservice.Options)) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetResourceConfig(context.Context, *configservice.BatchGetResourceConfigInput, ...func(*configservice.Options)) (*configservice.BatchGetResourceConfigOutput, error)
	DescribeAggregateComplianceByConfigRules(context.Context, *configservice.DescribeAggregateComplianceByConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConformancePacks(context.Context, *configservice.DescribeAggregateComplianceByConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeAggregateComplianceByConformancePacksOutput, error)
	DescribeAggregationAuthorizations(context.Context, *configservice.DescribeAggregationAuthorizationsInput, ...func(*configservice.Options)) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeComplianceByConfigRule(context.Context, *configservice.DescribeComplianceByConfigRuleInput, ...func(*configservice.Options)) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByResource(context.Context, *configservice.DescribeComplianceByResourceInput, ...func(*configservice.Options)) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeConfigRuleEvaluationStatus(context.Context, *configservice.DescribeConfigRuleEvaluationStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRules(context.Context, *configservice.DescribeConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigurationAggregatorSourcesStatus(context.Context, *configservice.DescribeConfigurationAggregatorSourcesStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregators(context.Context, *configservice.DescribeConfigurationAggregatorsInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationRecorderStatus(context.Context, *configservice.DescribeConfigurationRecorderStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorders(context.Context, *configservice.DescribeConfigurationRecordersInput, ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConformancePackCompliance(context.Context, *configservice.DescribeConformancePackComplianceInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePackComplianceOutput, error)
	DescribeConformancePackStatus(context.Context, *configservice.DescribeConformancePackStatusInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePackStatusOutput, error)
	DescribeConformancePacks(context.Context, *configservice.DescribeConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeConformancePacksOutput, error)
	DescribeDeliveryChannelStatus(context.Context, *configservice.DescribeDeliveryChannelStatusInput, ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannels(context.Context, *configservice.DescribeDeliveryChannelsInput, ...func(*configservice.Options)) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeOrganizationConfigRuleStatuses(context.Context, *configservice.DescribeOrganizationConfigRuleStatusesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error)
	DescribeOrganizationConfigRules(context.Context, *configservice.DescribeOrganizationConfigRulesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConfigRulesOutput, error)
	DescribeOrganizationConformancePackStatuses(context.Context, *configservice.DescribeOrganizationConformancePackStatusesInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error)
	DescribeOrganizationConformancePacks(context.Context, *configservice.DescribeOrganizationConformancePacksInput, ...func(*configservice.Options)) (*configservice.DescribeOrganizationConformancePacksOutput, error)
	DescribePendingAggregationRequests(context.Context, *configservice.DescribePendingAggregationRequestsInput, ...func(*configservice.Options)) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribeRemediationConfigurations(context.Context, *configservice.DescribeRemediationConfigurationsInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationConfigurationsOutput, error)
	DescribeRemediationExceptions(context.Context, *configservice.DescribeRemediationExceptionsInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationExceptionsOutput, error)
	DescribeRemediationExecutionStatus(context.Context, *configservice.DescribeRemediationExecutionStatusInput, ...func(*configservice.Options)) (*configservice.DescribeRemediationExecutionStatusOutput, error)
	DescribeRetentionConfigurations(context.Context, *configservice.DescribeRetentionConfigurationsInput, ...func(*configservice.Options)) (*configservice.DescribeRetentionConfigurationsOutput, error)
	GetAggregateComplianceDetailsByConfigRule(context.Context, *configservice.GetAggregateComplianceDetailsByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateConfigRuleComplianceSummary(context.Context, *configservice.GetAggregateConfigRuleComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConformancePackComplianceSummary(context.Context, *configservice.GetAggregateConformancePackComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetAggregateConformancePackComplianceSummaryOutput, error)
	GetAggregateDiscoveredResourceCounts(context.Context, *configservice.GetAggregateDiscoveredResourceCountsInput, ...func(*configservice.Options)) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateResourceConfig(context.Context, *configservice.GetAggregateResourceConfigInput, ...func(*configservice.Options)) (*configservice.GetAggregateResourceConfigOutput, error)
	GetComplianceDetailsByConfigRule(context.Context, *configservice.GetComplianceDetailsByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByResource(context.Context, *configservice.GetComplianceDetailsByResourceInput, ...func(*configservice.Options)) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceSummaryByConfigRule(context.Context, *configservice.GetComplianceSummaryByConfigRuleInput, ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByResourceType(context.Context, *configservice.GetComplianceSummaryByResourceTypeInput, ...func(*configservice.Options)) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetConformancePackComplianceDetails(context.Context, *configservice.GetConformancePackComplianceDetailsInput, ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceDetailsOutput, error)
	GetConformancePackComplianceSummary(context.Context, *configservice.GetConformancePackComplianceSummaryInput, ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceSummaryOutput, error)
	GetCustomRulePolicy(context.Context, *configservice.GetCustomRulePolicyInput, ...func(*configservice.Options)) (*configservice.GetCustomRulePolicyOutput, error)
	GetDiscoveredResourceCounts(context.Context, *configservice.GetDiscoveredResourceCountsInput, ...func(*configservice.Options)) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetOrganizationConfigRuleDetailedStatus(context.Context, *configservice.GetOrganizationConfigRuleDetailedStatusInput, ...func(*configservice.Options)) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error)
	GetOrganizationConformancePackDetailedStatus(context.Context, *configservice.GetOrganizationConformancePackDetailedStatusInput, ...func(*configservice.Options)) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error)
	GetOrganizationCustomRulePolicy(context.Context, *configservice.GetOrganizationCustomRulePolicyInput, ...func(*configservice.Options)) (*configservice.GetOrganizationCustomRulePolicyOutput, error)
	GetResourceConfigHistory(context.Context, *configservice.GetResourceConfigHistoryInput, ...func(*configservice.Options)) (*configservice.GetResourceConfigHistoryOutput, error)
	GetStoredQuery(context.Context, *configservice.GetStoredQueryInput, ...func(*configservice.Options)) (*configservice.GetStoredQueryOutput, error)
	ListAggregateDiscoveredResources(context.Context, *configservice.ListAggregateDiscoveredResourcesInput, ...func(*configservice.Options)) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListConformancePackComplianceScores(context.Context, *configservice.ListConformancePackComplianceScoresInput, ...func(*configservice.Options)) (*configservice.ListConformancePackComplianceScoresOutput, error)
	ListDiscoveredResources(context.Context, *configservice.ListDiscoveredResourcesInput, ...func(*configservice.Options)) (*configservice.ListDiscoveredResourcesOutput, error)
	ListStoredQueries(context.Context, *configservice.ListStoredQueriesInput, ...func(*configservice.Options)) (*configservice.ListStoredQueriesOutput, error)
	ListTagsForResource(context.Context, *configservice.ListTagsForResourceInput, ...func(*configservice.Options)) (*configservice.ListTagsForResourceOutput, error)
}

type DatabasemigrationserviceClient interface {
	DescribeAccountAttributes(context.Context, *databasemigrationservice.DescribeAccountAttributesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeApplicableIndividualAssessments(context.Context, *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error)
	DescribeCertificates(context.Context, *databasemigrationservice.DescribeCertificatesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeConnections(context.Context, *databasemigrationservice.DescribeConnectionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeEndpointSettings(context.Context, *databasemigrationservice.DescribeEndpointSettingsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointSettingsOutput, error)
	DescribeEndpointTypes(context.Context, *databasemigrationservice.DescribeEndpointTypesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpoints(context.Context, *databasemigrationservice.DescribeEndpointsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEventCategories(context.Context, *databasemigrationservice.DescribeEventCategoriesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *databasemigrationservice.DescribeEventSubscriptionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *databasemigrationservice.DescribeEventsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeFleetAdvisorCollectors(context.Context, *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, error)
	DescribeFleetAdvisorDatabases(context.Context, *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, error)
	DescribeFleetAdvisorLsaAnalysis(context.Context, *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, error)
	DescribeFleetAdvisorSchemaObjectSummary(context.Context, *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, error)
	DescribeFleetAdvisorSchemas(context.Context, *databasemigrationservice.DescribeFleetAdvisorSchemasInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, error)
	DescribeOrderableReplicationInstances(context.Context, *databasemigrationservice.DescribeOrderableReplicationInstancesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribePendingMaintenanceActions(context.Context, *databasemigrationservice.DescribePendingMaintenanceActionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error)
	DescribeRefreshSchemasStatus(context.Context, *databasemigrationservice.DescribeRefreshSchemasStatusInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeReplicationInstanceTaskLogs(context.Context, *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error)
	DescribeReplicationInstances(context.Context, *databasemigrationservice.DescribeReplicationInstancesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationSubnetGroups(context.Context, *databasemigrationservice.DescribeReplicationSubnetGroupsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationTaskAssessmentResults(context.Context, *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error)
	DescribeReplicationTaskAssessmentRuns(context.Context, *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error)
	DescribeReplicationTaskIndividualAssessments(context.Context, *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error)
	DescribeReplicationTasks(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeSchemas(context.Context, *databasemigrationservice.DescribeSchemasInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeTableStatistics(context.Context, *databasemigrationservice.DescribeTableStatisticsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	ListTagsForResource(context.Context, *databasemigrationservice.ListTagsForResourceInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error)
}

type DaxClient interface {
	DescribeClusters(context.Context, *dax.DescribeClustersInput, ...func(*dax.Options)) (*dax.DescribeClustersOutput, error)
	DescribeDefaultParameters(context.Context, *dax.DescribeDefaultParametersInput, ...func(*dax.Options)) (*dax.DescribeDefaultParametersOutput, error)
	DescribeEvents(context.Context, *dax.DescribeEventsInput, ...func(*dax.Options)) (*dax.DescribeEventsOutput, error)
	DescribeParameterGroups(context.Context, *dax.DescribeParameterGroupsInput, ...func(*dax.Options)) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameters(context.Context, *dax.DescribeParametersInput, ...func(*dax.Options)) (*dax.DescribeParametersOutput, error)
	DescribeSubnetGroups(context.Context, *dax.DescribeSubnetGroupsInput, ...func(*dax.Options)) (*dax.DescribeSubnetGroupsOutput, error)
	ListTags(context.Context, *dax.ListTagsInput, ...func(*dax.Options)) (*dax.ListTagsOutput, error)
}

type DirectconnectClient interface {
	DescribeConnectionLoa(context.Context, *directconnect.DescribeConnectionLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionLoaOutput, error)
	DescribeConnections(context.Context, *directconnect.DescribeConnectionsInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error)
	DescribeConnectionsOnInterconnect(context.Context, *directconnect.DescribeConnectionsOnInterconnectInput, ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOnInterconnectOutput, error)
	DescribeCustomerMetadata(context.Context, *directconnect.DescribeCustomerMetadataInput, ...func(*directconnect.Options)) (*directconnect.DescribeCustomerMetadataOutput, error)
	DescribeDirectConnectGatewayAssociationProposals(context.Context, *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error)
	DescribeDirectConnectGatewayAssociations(context.Context, *directconnect.DescribeDirectConnectGatewayAssociationsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error)
	DescribeDirectConnectGatewayAttachments(context.Context, *directconnect.DescribeDirectConnectGatewayAttachmentsInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error)
	DescribeDirectConnectGateways(context.Context, *directconnect.DescribeDirectConnectGatewaysInput, ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
	DescribeHostedConnections(context.Context, *directconnect.DescribeHostedConnectionsInput, ...func(*directconnect.Options)) (*directconnect.DescribeHostedConnectionsOutput, error)
	DescribeInterconnectLoa(context.Context, *directconnect.DescribeInterconnectLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectLoaOutput, error)
	DescribeInterconnects(context.Context, *directconnect.DescribeInterconnectsInput, ...func(*directconnect.Options)) (*directconnect.DescribeInterconnectsOutput, error)
	DescribeLags(context.Context, *directconnect.DescribeLagsInput, ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error)
	DescribeLoa(context.Context, *directconnect.DescribeLoaInput, ...func(*directconnect.Options)) (*directconnect.DescribeLoaOutput, error)
	DescribeLocations(context.Context, *directconnect.DescribeLocationsInput, ...func(*directconnect.Options)) (*directconnect.DescribeLocationsOutput, error)
	DescribeRouterConfiguration(context.Context, *directconnect.DescribeRouterConfigurationInput, ...func(*directconnect.Options)) (*directconnect.DescribeRouterConfigurationOutput, error)
	DescribeTags(context.Context, *directconnect.DescribeTagsInput, ...func(*directconnect.Options)) (*directconnect.DescribeTagsOutput, error)
	DescribeVirtualGateways(context.Context, *directconnect.DescribeVirtualGatewaysInput, ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualInterfaces(context.Context, *directconnect.DescribeVirtualInterfacesInput, ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error)
	ListVirtualInterfaceTestHistory(context.Context, *directconnect.ListVirtualInterfaceTestHistoryInput, ...func(*directconnect.Options)) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error)
}

type DocdbClient interface {
	DescribeCertificates(context.Context, *docdb.DescribeCertificatesInput, ...func(*docdb.Options)) (*docdb.DescribeCertificatesOutput, error)
	DescribeDBClusterParameterGroups(context.Context, *docdb.DescribeDBClusterParameterGroupsInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(context.Context, *docdb.DescribeDBClusterParametersInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(context.Context, *docdb.DescribeDBClusterSnapshotAttributesInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(context.Context, *docdb.DescribeDBClusterSnapshotsInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(context.Context, *docdb.DescribeDBClustersInput, ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error)
	DescribeDBEngineVersions(context.Context, *docdb.DescribeDBEngineVersionsInput, ...func(*docdb.Options)) (*docdb.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstances(context.Context, *docdb.DescribeDBInstancesInput, ...func(*docdb.Options)) (*docdb.DescribeDBInstancesOutput, error)
	DescribeDBSubnetGroups(context.Context, *docdb.DescribeDBSubnetGroupsInput, ...func(*docdb.Options)) (*docdb.DescribeDBSubnetGroupsOutput, error)
	DescribeEngineDefaultClusterParameters(context.Context, *docdb.DescribeEngineDefaultClusterParametersInput, ...func(*docdb.Options)) (*docdb.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEventCategories(context.Context, *docdb.DescribeEventCategoriesInput, ...func(*docdb.Options)) (*docdb.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *docdb.DescribeEventSubscriptionsInput, ...func(*docdb.Options)) (*docdb.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *docdb.DescribeEventsInput, ...func(*docdb.Options)) (*docdb.DescribeEventsOutput, error)
	DescribeGlobalClusters(context.Context, *docdb.DescribeGlobalClustersInput, ...func(*docdb.Options)) (*docdb.DescribeGlobalClustersOutput, error)
	DescribeOrderableDBInstanceOptions(context.Context, *docdb.DescribeOrderableDBInstanceOptionsInput, ...func(*docdb.Options)) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribePendingMaintenanceActions(context.Context, *docdb.DescribePendingMaintenanceActionsInput, ...func(*docdb.Options)) (*docdb.DescribePendingMaintenanceActionsOutput, error)
	ListTagsForResource(context.Context, *docdb.ListTagsForResourceInput, ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error)
}

type DynamodbClient interface {
	BatchGetItem(context.Context, *dynamodb.BatchGetItemInput, ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	DescribeBackup(context.Context, *dynamodb.DescribeBackupInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error)
	DescribeContinuousBackups(context.Context, *dynamodb.DescribeContinuousBackupsInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContributorInsights(context.Context, *dynamodb.DescribeContributorInsightsInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error)
	DescribeEndpoints(context.Context, *dynamodb.DescribeEndpointsInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeExport(context.Context, *dynamodb.DescribeExportInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error)
	DescribeGlobalTable(context.Context, *dynamodb.DescribeGlobalTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettings(context.Context, *dynamodb.DescribeGlobalTableSettingsInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeImport(context.Context, *dynamodb.DescribeImportInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeImportOutput, error)
	DescribeKinesisStreamingDestination(context.Context, *dynamodb.DescribeKinesisStreamingDestinationInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	DescribeLimits(context.Context, *dynamodb.DescribeLimitsInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error)
	DescribeTable(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(context.Context, *dynamodb.DescribeTableReplicaAutoScalingInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTimeToLive(context.Context, *dynamodb.DescribeTimeToLiveInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error)
	GetItem(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	ListBackups(context.Context, *dynamodb.ListBackupsInput, ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error)
	ListContributorInsights(context.Context, *dynamodb.ListContributorInsightsInput, ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error)
	ListExports(context.Context, *dynamodb.ListExportsInput, ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error)
	ListGlobalTables(context.Context, *dynamodb.ListGlobalTablesInput, ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error)
	ListImports(context.Context, *dynamodb.ListImportsInput, ...func(*dynamodb.Options)) (*dynamodb.ListImportsOutput, error)
	ListTables(context.Context, *dynamodb.ListTablesInput, ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	ListTagsOfResource(context.Context, *dynamodb.ListTagsOfResourceInput, ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
}

type Ec2Client interface {
	DescribeAccountAttributes(context.Context, *ec2.DescribeAccountAttributesInput, ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error)
	DescribeAddressTransfers(context.Context, *ec2.DescribeAddressTransfersInput, ...func(*ec2.Options)) (*ec2.DescribeAddressTransfersOutput, error)
	DescribeAddresses(context.Context, *ec2.DescribeAddressesInput, ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	DescribeAddressesAttribute(context.Context, *ec2.DescribeAddressesAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error)
	DescribeAggregateIdFormat(context.Context, *ec2.DescribeAggregateIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error)
	DescribeAvailabilityZones(context.Context, *ec2.DescribeAvailabilityZonesInput, ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeBundleTasks(context.Context, *ec2.DescribeBundleTasksInput, ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error)
	DescribeByoipCidrs(context.Context, *ec2.DescribeByoipCidrsInput, ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCapacityReservationFleets(context.Context, *ec2.DescribeCapacityReservationFleetsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error)
	DescribeCapacityReservations(context.Context, *ec2.DescribeCapacityReservationsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error)
	DescribeCarrierGateways(context.Context, *ec2.DescribeCarrierGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error)
	DescribeClassicLinkInstances(context.Context, *ec2.DescribeClassicLinkInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error)
	DescribeClientVpnAuthorizationRules(context.Context, *ec2.DescribeClientVpnAuthorizationRulesInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error)
	DescribeClientVpnConnections(context.Context, *ec2.DescribeClientVpnConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error)
	DescribeClientVpnEndpoints(context.Context, *ec2.DescribeClientVpnEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error)
	DescribeClientVpnRoutes(context.Context, *ec2.DescribeClientVpnRoutesInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error)
	DescribeClientVpnTargetNetworks(context.Context, *ec2.DescribeClientVpnTargetNetworksInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error)
	DescribeCoipPools(context.Context, *ec2.DescribeCoipPoolsInput, ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error)
	DescribeConversionTasks(context.Context, *ec2.DescribeConversionTasksInput, ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error)
	DescribeCustomerGateways(context.Context, *ec2.DescribeCustomerGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeDhcpOptions(context.Context, *ec2.DescribeDhcpOptionsInput, ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error)
	DescribeEgressOnlyInternetGateways(context.Context, *ec2.DescribeEgressOnlyInternetGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeElasticGpus(context.Context, *ec2.DescribeElasticGpusInput, ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error)
	DescribeExportImageTasks(context.Context, *ec2.DescribeExportImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error)
	DescribeExportTasks(context.Context, *ec2.DescribeExportTasksInput, ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error)
	DescribeFastLaunchImages(context.Context, *ec2.DescribeFastLaunchImagesInput, ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error)
	DescribeFastSnapshotRestores(context.Context, *ec2.DescribeFastSnapshotRestoresInput, ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error)
	DescribeFleetHistory(context.Context, *ec2.DescribeFleetHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error)
	DescribeFleetInstances(context.Context, *ec2.DescribeFleetInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error)
	DescribeFleets(context.Context, *ec2.DescribeFleetsInput, ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error)
	DescribeFlowLogs(context.Context, *ec2.DescribeFlowLogsInput, ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeFpgaImageAttribute(context.Context, *ec2.DescribeFpgaImageAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error)
	DescribeFpgaImages(context.Context, *ec2.DescribeFpgaImagesInput, ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error)
	DescribeHostReservationOfferings(context.Context, *ec2.DescribeHostReservationOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error)
	DescribeHostReservations(context.Context, *ec2.DescribeHostReservationsInput, ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error)
	DescribeHosts(context.Context, *ec2.DescribeHostsInput, ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	DescribeIamInstanceProfileAssociations(context.Context, *ec2.DescribeIamInstanceProfileAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	DescribeIdFormat(context.Context, *ec2.DescribeIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error)
	DescribeIdentityIdFormat(context.Context, *ec2.DescribeIdentityIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error)
	DescribeImageAttribute(context.Context, *ec2.DescribeImageAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImages(context.Context, *ec2.DescribeImagesInput, ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeImportImageTasks(context.Context, *ec2.DescribeImportImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error)
	DescribeImportSnapshotTasks(context.Context, *ec2.DescribeImportSnapshotTasksInput, ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error)
	DescribeInstanceAttribute(context.Context, *ec2.DescribeInstanceAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error)
	DescribeInstanceCreditSpecifications(context.Context, *ec2.DescribeInstanceCreditSpecificationsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error)
	DescribeInstanceEventNotificationAttributes(context.Context, *ec2.DescribeInstanceEventNotificationAttributesInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error)
	DescribeInstanceEventWindows(context.Context, *ec2.DescribeInstanceEventWindowsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error)
	DescribeInstanceStatus(context.Context, *ec2.DescribeInstanceStatusInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstanceTypeOfferings(context.Context, *ec2.DescribeInstanceTypeOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	DescribeInstanceTypes(context.Context, *ec2.DescribeInstanceTypesInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	DescribeInstances(context.Context, *ec2.DescribeInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInternetGateways(context.Context, *ec2.DescribeInternetGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeIpamPools(context.Context, *ec2.DescribeIpamPoolsInput, ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error)
	DescribeIpamScopes(context.Context, *ec2.DescribeIpamScopesInput, ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error)
	DescribeIpams(context.Context, *ec2.DescribeIpamsInput, ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error)
	DescribeIpv6Pools(context.Context, *ec2.DescribeIpv6PoolsInput, ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error)
	DescribeKeyPairs(context.Context, *ec2.DescribeKeyPairsInput, ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	DescribeLaunchTemplateVersions(context.Context, *ec2.DescribeLaunchTemplateVersionsInput, ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error)
	DescribeLaunchTemplates(context.Context, *ec2.DescribeLaunchTemplatesInput, ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error)
	DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(context.Context, *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
	DescribeLocalGatewayRouteTableVpcAssociations(context.Context, *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
	DescribeLocalGatewayRouteTables(context.Context, *ec2.DescribeLocalGatewayRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error)
	DescribeLocalGatewayVirtualInterfaceGroups(context.Context, *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
	DescribeLocalGatewayVirtualInterfaces(context.Context, *ec2.DescribeLocalGatewayVirtualInterfacesInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error)
	DescribeLocalGateways(context.Context, *ec2.DescribeLocalGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error)
	DescribeManagedPrefixLists(context.Context, *ec2.DescribeManagedPrefixListsInput, ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error)
	DescribeMovingAddresses(context.Context, *ec2.DescribeMovingAddressesInput, ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error)
	DescribeNatGateways(context.Context, *ec2.DescribeNatGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkAcls(context.Context, *ec2.DescribeNetworkAclsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkInsightsAccessScopeAnalyses(context.Context, *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error)
	DescribeNetworkInsightsAccessScopes(context.Context, *ec2.DescribeNetworkInsightsAccessScopesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error)
	DescribeNetworkInsightsAnalyses(context.Context, *ec2.DescribeNetworkInsightsAnalysesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error)
	DescribeNetworkInsightsPaths(context.Context, *ec2.DescribeNetworkInsightsPathsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error)
	DescribeNetworkInterfaceAttribute(context.Context, *ec2.DescribeNetworkInterfaceAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	DescribeNetworkInterfacePermissions(context.Context, *ec2.DescribeNetworkInterfacePermissionsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	DescribeNetworkInterfaces(context.Context, *ec2.DescribeNetworkInterfacesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribePlacementGroups(context.Context, *ec2.DescribePlacementGroupsInput, ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error)
	DescribePrefixLists(context.Context, *ec2.DescribePrefixListsInput, ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error)
	DescribePrincipalIdFormat(context.Context, *ec2.DescribePrincipalIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error)
	DescribePublicIpv4Pools(context.Context, *ec2.DescribePublicIpv4PoolsInput, ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error)
	DescribeRegions(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	DescribeReplaceRootVolumeTasks(context.Context, *ec2.DescribeReplaceRootVolumeTasksInput, ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error)
	DescribeReservedInstances(context.Context, *ec2.DescribeReservedInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesListings(context.Context, *ec2.DescribeReservedInstancesListingsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error)
	DescribeReservedInstancesModifications(context.Context, *ec2.DescribeReservedInstancesModificationsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	DescribeReservedInstancesOfferings(context.Context, *ec2.DescribeReservedInstancesOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeRouteTables(context.Context, *ec2.DescribeRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	DescribeScheduledInstanceAvailability(context.Context, *ec2.DescribeScheduledInstanceAvailabilityInput, ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	DescribeScheduledInstances(context.Context, *ec2.DescribeScheduledInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error)
	DescribeSecurityGroupReferences(context.Context, *ec2.DescribeSecurityGroupReferencesInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	DescribeSecurityGroupRules(context.Context, *ec2.DescribeSecurityGroupRulesInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error)
	DescribeSecurityGroups(context.Context, *ec2.DescribeSecurityGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSnapshotAttribute(context.Context, *ec2.DescribeSnapshotAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshotTierStatus(context.Context, *ec2.DescribeSnapshotTierStatusInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error)
	DescribeSnapshots(context.Context, *ec2.DescribeSnapshotsInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSpotDatafeedSubscription(context.Context, *ec2.DescribeSpotDatafeedSubscriptionInput, ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	DescribeSpotFleetInstances(context.Context, *ec2.DescribeSpotFleetInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error)
	DescribeSpotFleetRequestHistory(context.Context, *ec2.DescribeSpotFleetRequestHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	DescribeSpotFleetRequests(context.Context, *ec2.DescribeSpotFleetRequestsInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error)
	DescribeSpotInstanceRequests(context.Context, *ec2.DescribeSpotInstanceRequestsInput, ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	DescribeSpotPriceHistory(context.Context, *ec2.DescribeSpotPriceHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error)
	DescribeStaleSecurityGroups(context.Context, *ec2.DescribeStaleSecurityGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	DescribeStoreImageTasks(context.Context, *ec2.DescribeStoreImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error)
	DescribeSubnets(context.Context, *ec2.DescribeSubnetsInput, ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeTags(context.Context, *ec2.DescribeTagsInput, ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error)
	DescribeTrafficMirrorFilters(context.Context, *ec2.DescribeTrafficMirrorFiltersInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error)
	DescribeTrafficMirrorSessions(context.Context, *ec2.DescribeTrafficMirrorSessionsInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error)
	DescribeTrafficMirrorTargets(context.Context, *ec2.DescribeTrafficMirrorTargetsInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error)
	DescribeTransitGatewayAttachments(context.Context, *ec2.DescribeTransitGatewayAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	DescribeTransitGatewayConnectPeers(context.Context, *ec2.DescribeTransitGatewayConnectPeersInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error)
	DescribeTransitGatewayConnects(context.Context, *ec2.DescribeTransitGatewayConnectsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error)
	DescribeTransitGatewayMulticastDomains(context.Context, *ec2.DescribeTransitGatewayMulticastDomainsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	DescribeTransitGatewayPeeringAttachments(context.Context, *ec2.DescribeTransitGatewayPeeringAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	DescribeTransitGatewayPolicyTables(context.Context, *ec2.DescribeTransitGatewayPolicyTablesInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error)
	DescribeTransitGatewayRouteTableAnnouncements(context.Context, *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error)
	DescribeTransitGatewayRouteTables(context.Context, *ec2.DescribeTransitGatewayRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	DescribeTransitGatewayVpcAttachments(context.Context, *ec2.DescribeTransitGatewayVpcAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	DescribeTransitGateways(context.Context, *ec2.DescribeTransitGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	DescribeTrunkInterfaceAssociations(context.Context, *ec2.DescribeTrunkInterfaceAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error)
	DescribeVolumeAttribute(context.Context, *ec2.DescribeVolumeAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error)
	DescribeVolumeStatus(context.Context, *ec2.DescribeVolumeStatusInput, ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error)
	DescribeVolumes(context.Context, *ec2.DescribeVolumesInput, ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	DescribeVolumesModifications(context.Context, *ec2.DescribeVolumesModificationsInput, ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error)
	DescribeVpcAttribute(context.Context, *ec2.DescribeVpcAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeVpcClassicLink(context.Context, *ec2.DescribeVpcClassicLinkInput, ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error)
	DescribeVpcClassicLinkDnsSupport(context.Context, *ec2.DescribeVpcClassicLinkDnsSupportInput, ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	DescribeVpcEndpointConnectionNotifications(context.Context, *ec2.DescribeVpcEndpointConnectionNotificationsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error)
	DescribeVpcEndpointConnections(context.Context, *ec2.DescribeVpcEndpointConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
	DescribeVpcEndpointServiceConfigurations(context.Context, *ec2.DescribeVpcEndpointServiceConfigurationsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	DescribeVpcEndpointServicePermissions(context.Context, *ec2.DescribeVpcEndpointServicePermissionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
	DescribeVpcEndpointServices(context.Context, *ec2.DescribeVpcEndpointServicesInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcEndpoints(context.Context, *ec2.DescribeVpcEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcPeeringConnections(context.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcs(context.Context, *ec2.DescribeVpcsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpnConnections(context.Context, *ec2.DescribeVpnConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error)
	DescribeVpnGateways(context.Context, *ec2.DescribeVpnGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	GetAssociatedEnclaveCertificateIamRoles(context.Context, *ec2.GetAssociatedEnclaveCertificateIamRolesInput, ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error)
	GetAssociatedIpv6PoolCidrs(context.Context, *ec2.GetAssociatedIpv6PoolCidrsInput, ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error)
	GetCapacityReservationUsage(context.Context, *ec2.GetCapacityReservationUsageInput, ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error)
	GetCoipPoolUsage(context.Context, *ec2.GetCoipPoolUsageInput, ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error)
	GetConsoleOutput(context.Context, *ec2.GetConsoleOutputInput, ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error)
	GetConsoleScreenshot(context.Context, *ec2.GetConsoleScreenshotInput, ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error)
	GetDefaultCreditSpecification(context.Context, *ec2.GetDefaultCreditSpecificationInput, ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error)
	GetEbsDefaultKmsKeyId(context.Context, *ec2.GetEbsDefaultKmsKeyIdInput, ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
	GetEbsEncryptionByDefault(context.Context, *ec2.GetEbsEncryptionByDefaultInput, ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
	GetFlowLogsIntegrationTemplate(context.Context, *ec2.GetFlowLogsIntegrationTemplateInput, ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error)
	GetGroupsForCapacityReservation(context.Context, *ec2.GetGroupsForCapacityReservationInput, ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error)
	GetHostReservationPurchasePreview(context.Context, *ec2.GetHostReservationPurchasePreviewInput, ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	GetInstanceTypesFromInstanceRequirements(context.Context, *ec2.GetInstanceTypesFromInstanceRequirementsInput, ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error)
	GetInstanceUefiData(context.Context, *ec2.GetInstanceUefiDataInput, ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error)
	GetIpamAddressHistory(context.Context, *ec2.GetIpamAddressHistoryInput, ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error)
	GetIpamPoolAllocations(context.Context, *ec2.GetIpamPoolAllocationsInput, ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error)
	GetIpamPoolCidrs(context.Context, *ec2.GetIpamPoolCidrsInput, ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error)
	GetIpamResourceCidrs(context.Context, *ec2.GetIpamResourceCidrsInput, ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error)
	GetLaunchTemplateData(context.Context, *ec2.GetLaunchTemplateDataInput, ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error)
	GetManagedPrefixListAssociations(context.Context, *ec2.GetManagedPrefixListAssociationsInput, ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error)
	GetManagedPrefixListEntries(context.Context, *ec2.GetManagedPrefixListEntriesInput, ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error)
	GetNetworkInsightsAccessScopeAnalysisFindings(context.Context, *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error)
	GetNetworkInsightsAccessScopeContent(context.Context, *ec2.GetNetworkInsightsAccessScopeContentInput, ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error)
	GetPasswordData(context.Context, *ec2.GetPasswordDataInput, ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error)
	GetReservedInstancesExchangeQuote(context.Context, *ec2.GetReservedInstancesExchangeQuoteInput, ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	GetSerialConsoleAccessStatus(context.Context, *ec2.GetSerialConsoleAccessStatusInput, ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error)
	GetSpotPlacementScores(context.Context, *ec2.GetSpotPlacementScoresInput, ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error)
	GetSubnetCidrReservations(context.Context, *ec2.GetSubnetCidrReservationsInput, ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error)
	GetTransitGatewayAttachmentPropagations(context.Context, *ec2.GetTransitGatewayAttachmentPropagationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error)
	GetTransitGatewayMulticastDomainAssociations(context.Context, *ec2.GetTransitGatewayMulticastDomainAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error)
	GetTransitGatewayPolicyTableAssociations(context.Context, *ec2.GetTransitGatewayPolicyTableAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error)
	GetTransitGatewayPolicyTableEntries(context.Context, *ec2.GetTransitGatewayPolicyTableEntriesInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error)
	GetTransitGatewayPrefixListReferences(context.Context, *ec2.GetTransitGatewayPrefixListReferencesInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error)
	GetTransitGatewayRouteTableAssociations(context.Context, *ec2.GetTransitGatewayRouteTableAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error)
	GetTransitGatewayRouteTablePropagations(context.Context, *ec2.GetTransitGatewayRouteTablePropagationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error)
	GetVpnConnectionDeviceSampleConfiguration(context.Context, *ec2.GetVpnConnectionDeviceSampleConfigurationInput, ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error)
	GetVpnConnectionDeviceTypes(context.Context, *ec2.GetVpnConnectionDeviceTypesInput, ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error)
	ListImagesInRecycleBin(context.Context, *ec2.ListImagesInRecycleBinInput, ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error)
	ListSnapshotsInRecycleBin(context.Context, *ec2.ListSnapshotsInRecycleBinInput, ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error)
	SearchLocalGatewayRoutes(context.Context, *ec2.SearchLocalGatewayRoutesInput, ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error)
	SearchTransitGatewayMulticastGroups(context.Context, *ec2.SearchTransitGatewayMulticastGroupsInput, ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error)
	SearchTransitGatewayRoutes(context.Context, *ec2.SearchTransitGatewayRoutesInput, ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error)
}

type EcrClient interface {
	BatchGetImage(context.Context, *ecr.BatchGetImageInput, ...func(*ecr.Options)) (*ecr.BatchGetImageOutput, error)
	BatchGetRepositoryScanningConfiguration(context.Context, *ecr.BatchGetRepositoryScanningConfigurationInput, ...func(*ecr.Options)) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error)
	DescribeImageReplicationStatus(context.Context, *ecr.DescribeImageReplicationStatusInput, ...func(*ecr.Options)) (*ecr.DescribeImageReplicationStatusOutput, error)
	DescribeImageScanFindings(context.Context, *ecr.DescribeImageScanFindingsInput, ...func(*ecr.Options)) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImages(context.Context, *ecr.DescribeImagesInput, ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error)
	DescribePullThroughCacheRules(context.Context, *ecr.DescribePullThroughCacheRulesInput, ...func(*ecr.Options)) (*ecr.DescribePullThroughCacheRulesOutput, error)
	DescribeRegistry(context.Context, *ecr.DescribeRegistryInput, ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error)
	DescribeRepositories(context.Context, *ecr.DescribeRepositoriesInput, ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error)
	GetAuthorizationToken(context.Context, *ecr.GetAuthorizationTokenInput, ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error)
	GetDownloadUrlForLayer(context.Context, *ecr.GetDownloadUrlForLayerInput, ...func(*ecr.Options)) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetLifecyclePolicy(context.Context, *ecr.GetLifecyclePolicyInput, ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyPreview(context.Context, *ecr.GetLifecyclePolicyPreviewInput, ...func(*ecr.Options)) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetRegistryPolicy(context.Context, *ecr.GetRegistryPolicyInput, ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error)
	GetRegistryScanningConfiguration(context.Context, *ecr.GetRegistryScanningConfigurationInput, ...func(*ecr.Options)) (*ecr.GetRegistryScanningConfigurationOutput, error)
	GetRepositoryPolicy(context.Context, *ecr.GetRepositoryPolicyInput, ...func(*ecr.Options)) (*ecr.GetRepositoryPolicyOutput, error)
	ListImages(context.Context, *ecr.ListImagesInput, ...func(*ecr.Options)) (*ecr.ListImagesOutput, error)
	ListTagsForResource(context.Context, *ecr.ListTagsForResourceInput, ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error)
}

type EcrpublicClient interface {
	DescribeImageTags(context.Context, *ecrpublic.DescribeImageTagsInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error)
	DescribeImages(context.Context, *ecrpublic.DescribeImagesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error)
	DescribeRegistries(context.Context, *ecrpublic.DescribeRegistriesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error)
	DescribeRepositories(context.Context, *ecrpublic.DescribeRepositoriesInput, ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error)
	GetAuthorizationToken(context.Context, *ecrpublic.GetAuthorizationTokenInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error)
	GetRegistryCatalogData(context.Context, *ecrpublic.GetRegistryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRegistryCatalogDataOutput, error)
	GetRepositoryCatalogData(context.Context, *ecrpublic.GetRepositoryCatalogDataInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryCatalogDataOutput, error)
	GetRepositoryPolicy(context.Context, *ecrpublic.GetRepositoryPolicyInput, ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error)
	ListTagsForResource(context.Context, *ecrpublic.ListTagsForResourceInput, ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error)
}

type EcsClient interface {
	DescribeCapacityProviders(context.Context, *ecs.DescribeCapacityProvidersInput, ...func(*ecs.Options)) (*ecs.DescribeCapacityProvidersOutput, error)
	DescribeClusters(context.Context, *ecs.DescribeClustersInput, ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error)
	DescribeContainerInstances(context.Context, *ecs.DescribeContainerInstancesInput, ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)
	DescribeServices(context.Context, *ecs.DescribeServicesInput, ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)
	DescribeTaskDefinition(context.Context, *ecs.DescribeTaskDefinitionInput, ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeTaskSets(context.Context, *ecs.DescribeTaskSetsInput, ...func(*ecs.Options)) (*ecs.DescribeTaskSetsOutput, error)
	DescribeTasks(context.Context, *ecs.DescribeTasksInput, ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)
	GetTaskProtection(context.Context, *ecs.GetTaskProtectionInput, ...func(*ecs.Options)) (*ecs.GetTaskProtectionOutput, error)
	ListAccountSettings(context.Context, *ecs.ListAccountSettingsInput, ...func(*ecs.Options)) (*ecs.ListAccountSettingsOutput, error)
	ListAttributes(context.Context, *ecs.ListAttributesInput, ...func(*ecs.Options)) (*ecs.ListAttributesOutput, error)
	ListClusters(context.Context, *ecs.ListClustersInput, ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
	ListContainerInstances(context.Context, *ecs.ListContainerInstancesInput, ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error)
	ListServices(context.Context, *ecs.ListServicesInput, ...func(*ecs.Options)) (*ecs.ListServicesOutput, error)
	ListServicesByNamespace(context.Context, *ecs.ListServicesByNamespaceInput, ...func(*ecs.Options)) (*ecs.ListServicesByNamespaceOutput, error)
	ListTagsForResource(context.Context, *ecs.ListTagsForResourceInput, ...func(*ecs.Options)) (*ecs.ListTagsForResourceOutput, error)
	ListTaskDefinitionFamilies(context.Context, *ecs.ListTaskDefinitionFamiliesInput, ...func(*ecs.Options)) (*ecs.ListTaskDefinitionFamiliesOutput, error)
	ListTaskDefinitions(context.Context, *ecs.ListTaskDefinitionsInput, ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error)
	ListTasks(context.Context, *ecs.ListTasksInput, ...func(*ecs.Options)) (*ecs.ListTasksOutput, error)
}

type EfsClient interface {
	DescribeAccessPoints(context.Context, *efs.DescribeAccessPointsInput, ...func(*efs.Options)) (*efs.DescribeAccessPointsOutput, error)
	DescribeAccountPreferences(context.Context, *efs.DescribeAccountPreferencesInput, ...func(*efs.Options)) (*efs.DescribeAccountPreferencesOutput, error)
	DescribeBackupPolicy(context.Context, *efs.DescribeBackupPolicyInput, ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error)
	DescribeFileSystemPolicy(context.Context, *efs.DescribeFileSystemPolicyInput, ...func(*efs.Options)) (*efs.DescribeFileSystemPolicyOutput, error)
	DescribeFileSystems(context.Context, *efs.DescribeFileSystemsInput, ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
	DescribeLifecycleConfiguration(context.Context, *efs.DescribeLifecycleConfigurationInput, ...func(*efs.Options)) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeMountTargetSecurityGroups(context.Context, *efs.DescribeMountTargetSecurityGroupsInput, ...func(*efs.Options)) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargets(context.Context, *efs.DescribeMountTargetsInput, ...func(*efs.Options)) (*efs.DescribeMountTargetsOutput, error)
	DescribeReplicationConfigurations(context.Context, *efs.DescribeReplicationConfigurationsInput, ...func(*efs.Options)) (*efs.DescribeReplicationConfigurationsOutput, error)
	DescribeTags(context.Context, *efs.DescribeTagsInput, ...func(*efs.Options)) (*efs.DescribeTagsOutput, error)
	ListTagsForResource(context.Context, *efs.ListTagsForResourceInput, ...func(*efs.Options)) (*efs.ListTagsForResourceOutput, error)
}

type EksClient interface {
	DescribeAddon(context.Context, *eks.DescribeAddonInput, ...func(*eks.Options)) (*eks.DescribeAddonOutput, error)
	DescribeAddonVersions(context.Context, *eks.DescribeAddonVersionsInput, ...func(*eks.Options)) (*eks.DescribeAddonVersionsOutput, error)
	DescribeCluster(context.Context, *eks.DescribeClusterInput, ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
	DescribeFargateProfile(context.Context, *eks.DescribeFargateProfileInput, ...func(*eks.Options)) (*eks.DescribeFargateProfileOutput, error)
	DescribeIdentityProviderConfig(context.Context, *eks.DescribeIdentityProviderConfigInput, ...func(*eks.Options)) (*eks.DescribeIdentityProviderConfigOutput, error)
	DescribeNodegroup(context.Context, *eks.DescribeNodegroupInput, ...func(*eks.Options)) (*eks.DescribeNodegroupOutput, error)
	DescribeUpdate(context.Context, *eks.DescribeUpdateInput, ...func(*eks.Options)) (*eks.DescribeUpdateOutput, error)
	ListAddons(context.Context, *eks.ListAddonsInput, ...func(*eks.Options)) (*eks.ListAddonsOutput, error)
	ListClusters(context.Context, *eks.ListClustersInput, ...func(*eks.Options)) (*eks.ListClustersOutput, error)
	ListFargateProfiles(context.Context, *eks.ListFargateProfilesInput, ...func(*eks.Options)) (*eks.ListFargateProfilesOutput, error)
	ListIdentityProviderConfigs(context.Context, *eks.ListIdentityProviderConfigsInput, ...func(*eks.Options)) (*eks.ListIdentityProviderConfigsOutput, error)
	ListNodegroups(context.Context, *eks.ListNodegroupsInput, ...func(*eks.Options)) (*eks.ListNodegroupsOutput, error)
	ListTagsForResource(context.Context, *eks.ListTagsForResourceInput, ...func(*eks.Options)) (*eks.ListTagsForResourceOutput, error)
	ListUpdates(context.Context, *eks.ListUpdatesInput, ...func(*eks.Options)) (*eks.ListUpdatesOutput, error)
}

type ElasticacheClient interface {
	DescribeCacheClusters(context.Context, *elasticache.DescribeCacheClustersInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheEngineVersions(context.Context, *elasticache.DescribeCacheEngineVersionsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheParameterGroups(context.Context, *elasticache.DescribeCacheParameterGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameters(context.Context, *elasticache.DescribeCacheParametersInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheSecurityGroups(context.Context, *elasticache.DescribeCacheSecurityGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
	DescribeCacheSubnetGroups(context.Context, *elasticache.DescribeCacheSubnetGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeEngineDefaultParameters(context.Context, *elasticache.DescribeEngineDefaultParametersInput, ...func(*elasticache.Options)) (*elasticache.DescribeEngineDefaultParametersOutput, error)
	DescribeEvents(context.Context, *elasticache.DescribeEventsInput, ...func(*elasticache.Options)) (*elasticache.DescribeEventsOutput, error)
	DescribeGlobalReplicationGroups(context.Context, *elasticache.DescribeGlobalReplicationGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error)
	DescribeReplicationGroups(context.Context, *elasticache.DescribeReplicationGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReservedCacheNodes(context.Context, *elasticache.DescribeReservedCacheNodesInput, ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesOfferings(context.Context, *elasticache.DescribeReservedCacheNodesOfferingsInput, ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeServiceUpdates(context.Context, *elasticache.DescribeServiceUpdatesInput, ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error)
	DescribeSnapshots(context.Context, *elasticache.DescribeSnapshotsInput, ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeUpdateActions(context.Context, *elasticache.DescribeUpdateActionsInput, ...func(*elasticache.Options)) (*elasticache.DescribeUpdateActionsOutput, error)
	DescribeUserGroups(context.Context, *elasticache.DescribeUserGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error)
	DescribeUsers(context.Context, *elasticache.DescribeUsersInput, ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error)
	ListAllowedNodeTypeModifications(context.Context, *elasticache.ListAllowedNodeTypeModificationsInput, ...func(*elasticache.Options)) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
	ListTagsForResource(context.Context, *elasticache.ListTagsForResourceInput, ...func(*elasticache.Options)) (*elasticache.ListTagsForResourceOutput, error)
}

type ElasticbeanstalkClient interface {
	DescribeAccountAttributes(context.Context, *elasticbeanstalk.DescribeAccountAttributesInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeAccountAttributesOutput, error)
	DescribeApplicationVersions(context.Context, *elasticbeanstalk.DescribeApplicationVersionsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
	DescribeApplications(context.Context, *elasticbeanstalk.DescribeApplicationsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationsOutput, error)
	DescribeConfigurationOptions(context.Context, *elasticbeanstalk.DescribeConfigurationOptionsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
	DescribeConfigurationSettings(context.Context, *elasticbeanstalk.DescribeConfigurationSettingsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
	DescribeEnvironmentHealth(context.Context, *elasticbeanstalk.DescribeEnvironmentHealthInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)
	DescribeEnvironmentManagedActionHistory(context.Context, *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)
	DescribeEnvironmentManagedActions(context.Context, *elasticbeanstalk.DescribeEnvironmentManagedActionsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)
	DescribeEnvironmentResources(context.Context, *elasticbeanstalk.DescribeEnvironmentResourcesInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)
	DescribeEnvironments(context.Context, *elasticbeanstalk.DescribeEnvironmentsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error)
	DescribeEvents(context.Context, *elasticbeanstalk.DescribeEventsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEventsOutput, error)
	DescribeInstancesHealth(context.Context, *elasticbeanstalk.DescribeInstancesHealthInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)
	DescribePlatformVersion(context.Context, *elasticbeanstalk.DescribePlatformVersionInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribePlatformVersionOutput, error)
	ListAvailableSolutionStacks(context.Context, *elasticbeanstalk.ListAvailableSolutionStacksInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)
	ListPlatformBranches(context.Context, *elasticbeanstalk.ListPlatformBranchesInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformBranchesOutput, error)
	ListPlatformVersions(context.Context, *elasticbeanstalk.ListPlatformVersionsInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListPlatformVersionsOutput, error)
	ListTagsForResource(context.Context, *elasticbeanstalk.ListTagsForResourceInput, ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListTagsForResourceOutput, error)
}

type ElasticloadbalancingClient interface {
	DescribeAccountLimits(context.Context, *elasticloadbalancing.DescribeAccountLimitsInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeAccountLimitsOutput, error)
	DescribeInstanceHealth(context.Context, *elasticloadbalancing.DescribeInstanceHealthInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeInstanceHealthOutput, error)
	DescribeLoadBalancerAttributes(context.Context, *elasticloadbalancing.DescribeLoadBalancerAttributesInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerPolicies(context.Context, *elasticloadbalancing.DescribeLoadBalancerPoliciesInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPolicyTypes(context.Context, *elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancers(context.Context, *elasticloadbalancing.DescribeLoadBalancersInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancersOutput, error)
	DescribeTags(context.Context, *elasticloadbalancing.DescribeTagsInput, ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeTagsOutput, error)
}

type Elasticloadbalancingv2Client interface {
	DescribeAccountLimits(context.Context, *elasticloadbalancingv2.DescribeAccountLimitsInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeAccountLimitsOutput, error)
	DescribeListenerCertificates(context.Context, *elasticloadbalancingv2.DescribeListenerCertificatesInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenerCertificatesOutput, error)
	DescribeListeners(context.Context, *elasticloadbalancingv2.DescribeListenersInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenersOutput, error)
	DescribeLoadBalancerAttributes(context.Context, *elasticloadbalancingv2.DescribeLoadBalancerAttributesInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancers(context.Context, *elasticloadbalancingv2.DescribeLoadBalancersInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error)
	DescribeRules(context.Context, *elasticloadbalancingv2.DescribeRulesInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeRulesOutput, error)
	DescribeSSLPolicies(context.Context, *elasticloadbalancingv2.DescribeSSLPoliciesInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeSSLPoliciesOutput, error)
	DescribeTags(context.Context, *elasticloadbalancingv2.DescribeTagsInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTagsOutput, error)
	DescribeTargetGroupAttributes(context.Context, *elasticloadbalancingv2.DescribeTargetGroupAttributesInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroups(context.Context, *elasticloadbalancingv2.DescribeTargetGroupsInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetGroupsOutput, error)
	DescribeTargetHealth(context.Context, *elasticloadbalancingv2.DescribeTargetHealthInput, ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetHealthOutput, error)
}

type ElasticsearchserviceClient interface {
	DescribeDomainAutoTunes(context.Context, *elasticsearchservice.DescribeDomainAutoTunesInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeDomainAutoTunesOutput, error)
	DescribeDomainChangeProgress(context.Context, *elasticsearchservice.DescribeDomainChangeProgressInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeDomainChangeProgressOutput, error)
	DescribeElasticsearchDomain(context.Context, *elasticsearchservice.DescribeElasticsearchDomainInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
	DescribeElasticsearchDomainConfig(context.Context, *elasticsearchservice.DescribeElasticsearchDomainConfigInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error)
	DescribeElasticsearchDomains(context.Context, *elasticsearchservice.DescribeElasticsearchDomainsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error)
	DescribeElasticsearchInstanceTypeLimits(context.Context, *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error)
	DescribeInboundCrossClusterSearchConnections(context.Context, *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error)
	DescribeOutboundCrossClusterSearchConnections(context.Context, *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error)
	DescribePackages(context.Context, *elasticsearchservice.DescribePackagesInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribePackagesOutput, error)
	DescribeReservedElasticsearchInstanceOfferings(context.Context, *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error)
	DescribeReservedElasticsearchInstances(context.Context, *elasticsearchservice.DescribeReservedElasticsearchInstancesInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error)
	DescribeVpcEndpoints(context.Context, *elasticsearchservice.DescribeVpcEndpointsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeVpcEndpointsOutput, error)
	GetCompatibleElasticsearchVersions(context.Context, *elasticsearchservice.GetCompatibleElasticsearchVersionsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error)
	GetPackageVersionHistory(context.Context, *elasticsearchservice.GetPackageVersionHistoryInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetPackageVersionHistoryOutput, error)
	GetUpgradeHistory(context.Context, *elasticsearchservice.GetUpgradeHistoryInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetUpgradeHistoryOutput, error)
	GetUpgradeStatus(context.Context, *elasticsearchservice.GetUpgradeStatusInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.GetUpgradeStatusOutput, error)
	ListDomainNames(context.Context, *elasticsearchservice.ListDomainNamesInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error)
	ListDomainsForPackage(context.Context, *elasticsearchservice.ListDomainsForPackageInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainsForPackageOutput, error)
	ListElasticsearchInstanceTypes(context.Context, *elasticsearchservice.ListElasticsearchInstanceTypesInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error)
	ListElasticsearchVersions(context.Context, *elasticsearchservice.ListElasticsearchVersionsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListElasticsearchVersionsOutput, error)
	ListPackagesForDomain(context.Context, *elasticsearchservice.ListPackagesForDomainInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListPackagesForDomainOutput, error)
	ListTags(context.Context, *elasticsearchservice.ListTagsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error)
	ListVpcEndpointAccess(context.Context, *elasticsearchservice.ListVpcEndpointAccessInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointAccessOutput, error)
	ListVpcEndpoints(context.Context, *elasticsearchservice.ListVpcEndpointsInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointsOutput, error)
	ListVpcEndpointsForDomain(context.Context, *elasticsearchservice.ListVpcEndpointsForDomainInput, ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListVpcEndpointsForDomainOutput, error)
}

type EmrClient interface {
	DescribeCluster(context.Context, *emr.DescribeClusterInput, ...func(*emr.Options)) (*emr.DescribeClusterOutput, error)
	DescribeJobFlows(context.Context, *emr.DescribeJobFlowsInput, ...func(*emr.Options)) (*emr.DescribeJobFlowsOutput, error)
	DescribeNotebookExecution(context.Context, *emr.DescribeNotebookExecutionInput, ...func(*emr.Options)) (*emr.DescribeNotebookExecutionOutput, error)
	DescribeReleaseLabel(context.Context, *emr.DescribeReleaseLabelInput, ...func(*emr.Options)) (*emr.DescribeReleaseLabelOutput, error)
	DescribeSecurityConfiguration(context.Context, *emr.DescribeSecurityConfigurationInput, ...func(*emr.Options)) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeStep(context.Context, *emr.DescribeStepInput, ...func(*emr.Options)) (*emr.DescribeStepOutput, error)
	DescribeStudio(context.Context, *emr.DescribeStudioInput, ...func(*emr.Options)) (*emr.DescribeStudioOutput, error)
	GetAutoTerminationPolicy(context.Context, *emr.GetAutoTerminationPolicyInput, ...func(*emr.Options)) (*emr.GetAutoTerminationPolicyOutput, error)
	GetBlockPublicAccessConfiguration(context.Context, *emr.GetBlockPublicAccessConfigurationInput, ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	GetManagedScalingPolicy(context.Context, *emr.GetManagedScalingPolicyInput, ...func(*emr.Options)) (*emr.GetManagedScalingPolicyOutput, error)
	GetStudioSessionMapping(context.Context, *emr.GetStudioSessionMappingInput, ...func(*emr.Options)) (*emr.GetStudioSessionMappingOutput, error)
	ListBootstrapActions(context.Context, *emr.ListBootstrapActionsInput, ...func(*emr.Options)) (*emr.ListBootstrapActionsOutput, error)
	ListClusters(context.Context, *emr.ListClustersInput, ...func(*emr.Options)) (*emr.ListClustersOutput, error)
	ListInstanceFleets(context.Context, *emr.ListInstanceFleetsInput, ...func(*emr.Options)) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceGroups(context.Context, *emr.ListInstanceGroupsInput, ...func(*emr.Options)) (*emr.ListInstanceGroupsOutput, error)
	ListInstances(context.Context, *emr.ListInstancesInput, ...func(*emr.Options)) (*emr.ListInstancesOutput, error)
	ListNotebookExecutions(context.Context, *emr.ListNotebookExecutionsInput, ...func(*emr.Options)) (*emr.ListNotebookExecutionsOutput, error)
	ListReleaseLabels(context.Context, *emr.ListReleaseLabelsInput, ...func(*emr.Options)) (*emr.ListReleaseLabelsOutput, error)
	ListSecurityConfigurations(context.Context, *emr.ListSecurityConfigurationsInput, ...func(*emr.Options)) (*emr.ListSecurityConfigurationsOutput, error)
	ListSteps(context.Context, *emr.ListStepsInput, ...func(*emr.Options)) (*emr.ListStepsOutput, error)
	ListStudioSessionMappings(context.Context, *emr.ListStudioSessionMappingsInput, ...func(*emr.Options)) (*emr.ListStudioSessionMappingsOutput, error)
	ListStudios(context.Context, *emr.ListStudiosInput, ...func(*emr.Options)) (*emr.ListStudiosOutput, error)
}

type EventbridgeClient interface {
	DescribeApiDestination(context.Context, *eventbridge.DescribeApiDestinationInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeApiDestinationOutput, error)
	DescribeArchive(context.Context, *eventbridge.DescribeArchiveInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeArchiveOutput, error)
	DescribeConnection(context.Context, *eventbridge.DescribeConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeConnectionOutput, error)
	DescribeEndpoint(context.Context, *eventbridge.DescribeEndpointInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEndpointOutput, error)
	DescribeEventBus(context.Context, *eventbridge.DescribeEventBusInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEventBusOutput, error)
	DescribeEventSource(context.Context, *eventbridge.DescribeEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEventSourceOutput, error)
	DescribePartnerEventSource(context.Context, *eventbridge.DescribePartnerEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DescribePartnerEventSourceOutput, error)
	DescribeReplay(context.Context, *eventbridge.DescribeReplayInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeReplayOutput, error)
	DescribeRule(context.Context, *eventbridge.DescribeRuleInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeRuleOutput, error)
	ListApiDestinations(context.Context, *eventbridge.ListApiDestinationsInput, ...func(*eventbridge.Options)) (*eventbridge.ListApiDestinationsOutput, error)
	ListArchives(context.Context, *eventbridge.ListArchivesInput, ...func(*eventbridge.Options)) (*eventbridge.ListArchivesOutput, error)
	ListConnections(context.Context, *eventbridge.ListConnectionsInput, ...func(*eventbridge.Options)) (*eventbridge.ListConnectionsOutput, error)
	ListEndpoints(context.Context, *eventbridge.ListEndpointsInput, ...func(*eventbridge.Options)) (*eventbridge.ListEndpointsOutput, error)
	ListEventBuses(context.Context, *eventbridge.ListEventBusesInput, ...func(*eventbridge.Options)) (*eventbridge.ListEventBusesOutput, error)
	ListEventSources(context.Context, *eventbridge.ListEventSourcesInput, ...func(*eventbridge.Options)) (*eventbridge.ListEventSourcesOutput, error)
	ListPartnerEventSourceAccounts(context.Context, *eventbridge.ListPartnerEventSourceAccountsInput, ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSources(context.Context, *eventbridge.ListPartnerEventSourcesInput, ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourcesOutput, error)
	ListReplays(context.Context, *eventbridge.ListReplaysInput, ...func(*eventbridge.Options)) (*eventbridge.ListReplaysOutput, error)
	ListRuleNamesByTarget(context.Context, *eventbridge.ListRuleNamesByTargetInput, ...func(*eventbridge.Options)) (*eventbridge.ListRuleNamesByTargetOutput, error)
	ListRules(context.Context, *eventbridge.ListRulesInput, ...func(*eventbridge.Options)) (*eventbridge.ListRulesOutput, error)
	ListTagsForResource(context.Context, *eventbridge.ListTagsForResourceInput, ...func(*eventbridge.Options)) (*eventbridge.ListTagsForResourceOutput, error)
	ListTargetsByRule(context.Context, *eventbridge.ListTargetsByRuleInput, ...func(*eventbridge.Options)) (*eventbridge.ListTargetsByRuleOutput, error)
}

type FirehoseClient interface {
	DescribeDeliveryStream(context.Context, *firehose.DescribeDeliveryStreamInput, ...func(*firehose.Options)) (*firehose.DescribeDeliveryStreamOutput, error)
	ListDeliveryStreams(context.Context, *firehose.ListDeliveryStreamsInput, ...func(*firehose.Options)) (*firehose.ListDeliveryStreamsOutput, error)
	ListTagsForDeliveryStream(context.Context, *firehose.ListTagsForDeliveryStreamInput, ...func(*firehose.Options)) (*firehose.ListTagsForDeliveryStreamOutput, error)
}

type FrauddetectorClient interface {
	BatchGetVariable(context.Context, *frauddetector.BatchGetVariableInput, ...func(*frauddetector.Options)) (*frauddetector.BatchGetVariableOutput, error)
	DescribeDetector(context.Context, *frauddetector.DescribeDetectorInput, ...func(*frauddetector.Options)) (*frauddetector.DescribeDetectorOutput, error)
	DescribeModelVersions(context.Context, *frauddetector.DescribeModelVersionsInput, ...func(*frauddetector.Options)) (*frauddetector.DescribeModelVersionsOutput, error)
	GetBatchImportJobs(context.Context, *frauddetector.GetBatchImportJobsInput, ...func(*frauddetector.Options)) (*frauddetector.GetBatchImportJobsOutput, error)
	GetBatchPredictionJobs(context.Context, *frauddetector.GetBatchPredictionJobsInput, ...func(*frauddetector.Options)) (*frauddetector.GetBatchPredictionJobsOutput, error)
	GetDeleteEventsByEventTypeStatus(context.Context, *frauddetector.GetDeleteEventsByEventTypeStatusInput, ...func(*frauddetector.Options)) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error)
	GetDetectorVersion(context.Context, *frauddetector.GetDetectorVersionInput, ...func(*frauddetector.Options)) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectors(context.Context, *frauddetector.GetDetectorsInput, ...func(*frauddetector.Options)) (*frauddetector.GetDetectorsOutput, error)
	GetEntityTypes(context.Context, *frauddetector.GetEntityTypesInput, ...func(*frauddetector.Options)) (*frauddetector.GetEntityTypesOutput, error)
	GetEvent(context.Context, *frauddetector.GetEventInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventOutput, error)
	GetEventPrediction(context.Context, *frauddetector.GetEventPredictionInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionMetadata(context.Context, *frauddetector.GetEventPredictionMetadataInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionMetadataOutput, error)
	GetEventTypes(context.Context, *frauddetector.GetEventTypesInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventTypesOutput, error)
	GetExternalModels(context.Context, *frauddetector.GetExternalModelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetExternalModelsOutput, error)
	GetKMSEncryptionKey(context.Context, *frauddetector.GetKMSEncryptionKeyInput, ...func(*frauddetector.Options)) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetLabels(context.Context, *frauddetector.GetLabelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetLabelsOutput, error)
	GetModelVersion(context.Context, *frauddetector.GetModelVersionInput, ...func(*frauddetector.Options)) (*frauddetector.GetModelVersionOutput, error)
	GetModels(context.Context, *frauddetector.GetModelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetModelsOutput, error)
	GetOutcomes(context.Context, *frauddetector.GetOutcomesInput, ...func(*frauddetector.Options)) (*frauddetector.GetOutcomesOutput, error)
	GetRules(context.Context, *frauddetector.GetRulesInput, ...func(*frauddetector.Options)) (*frauddetector.GetRulesOutput, error)
	GetVariables(context.Context, *frauddetector.GetVariablesInput, ...func(*frauddetector.Options)) (*frauddetector.GetVariablesOutput, error)
	ListEventPredictions(context.Context, *frauddetector.ListEventPredictionsInput, ...func(*frauddetector.Options)) (*frauddetector.ListEventPredictionsOutput, error)
	ListTagsForResource(context.Context, *frauddetector.ListTagsForResourceInput, ...func(*frauddetector.Options)) (*frauddetector.ListTagsForResourceOutput, error)
}

type FsxClient interface {
	DescribeBackups(context.Context, *fsx.DescribeBackupsInput, ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error)
	DescribeDataRepositoryAssociations(context.Context, *fsx.DescribeDataRepositoryAssociationsInput, ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error)
	DescribeDataRepositoryTasks(context.Context, *fsx.DescribeDataRepositoryTasksInput, ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeFileCaches(context.Context, *fsx.DescribeFileCachesInput, ...func(*fsx.Options)) (*fsx.DescribeFileCachesOutput, error)
	DescribeFileSystemAliases(context.Context, *fsx.DescribeFileSystemAliasesInput, ...func(*fsx.Options)) (*fsx.DescribeFileSystemAliasesOutput, error)
	DescribeFileSystems(context.Context, *fsx.DescribeFileSystemsInput, ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error)
	DescribeSnapshots(context.Context, *fsx.DescribeSnapshotsInput, ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error)
	DescribeStorageVirtualMachines(context.Context, *fsx.DescribeStorageVirtualMachinesInput, ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error)
	DescribeVolumes(context.Context, *fsx.DescribeVolumesInput, ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error)
	ListTagsForResource(context.Context, *fsx.ListTagsForResourceInput, ...func(*fsx.Options)) (*fsx.ListTagsForResourceOutput, error)
}

type GlacierClient interface {
	DescribeJob(context.Context, *glacier.DescribeJobInput, ...func(*glacier.Options)) (*glacier.DescribeJobOutput, error)
	DescribeVault(context.Context, *glacier.DescribeVaultInput, ...func(*glacier.Options)) (*glacier.DescribeVaultOutput, error)
	GetDataRetrievalPolicy(context.Context, *glacier.GetDataRetrievalPolicyInput, ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error)
	GetJobOutput(context.Context, *glacier.GetJobOutputInput, ...func(*glacier.Options)) (*glacier.GetJobOutputOutput, error)
	GetVaultAccessPolicy(context.Context, *glacier.GetVaultAccessPolicyInput, ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultLock(context.Context, *glacier.GetVaultLockInput, ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error)
	GetVaultNotifications(context.Context, *glacier.GetVaultNotificationsInput, ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error)
	ListJobs(context.Context, *glacier.ListJobsInput, ...func(*glacier.Options)) (*glacier.ListJobsOutput, error)
	ListMultipartUploads(context.Context, *glacier.ListMultipartUploadsInput, ...func(*glacier.Options)) (*glacier.ListMultipartUploadsOutput, error)
	ListParts(context.Context, *glacier.ListPartsInput, ...func(*glacier.Options)) (*glacier.ListPartsOutput, error)
	ListProvisionedCapacity(context.Context, *glacier.ListProvisionedCapacityInput, ...func(*glacier.Options)) (*glacier.ListProvisionedCapacityOutput, error)
	ListTagsForVault(context.Context, *glacier.ListTagsForVaultInput, ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error)
	ListVaults(context.Context, *glacier.ListVaultsInput, ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error)
}

type GlueClient interface {
	BatchGetBlueprints(context.Context, *glue.BatchGetBlueprintsInput, ...func(*glue.Options)) (*glue.BatchGetBlueprintsOutput, error)
	BatchGetCrawlers(context.Context, *glue.BatchGetCrawlersInput, ...func(*glue.Options)) (*glue.BatchGetCrawlersOutput, error)
	BatchGetCustomEntityTypes(context.Context, *glue.BatchGetCustomEntityTypesInput, ...func(*glue.Options)) (*glue.BatchGetCustomEntityTypesOutput, error)
	BatchGetDevEndpoints(context.Context, *glue.BatchGetDevEndpointsInput, ...func(*glue.Options)) (*glue.BatchGetDevEndpointsOutput, error)
	BatchGetJobs(context.Context, *glue.BatchGetJobsInput, ...func(*glue.Options)) (*glue.BatchGetJobsOutput, error)
	BatchGetPartition(context.Context, *glue.BatchGetPartitionInput, ...func(*glue.Options)) (*glue.BatchGetPartitionOutput, error)
	BatchGetTriggers(context.Context, *glue.BatchGetTriggersInput, ...func(*glue.Options)) (*glue.BatchGetTriggersOutput, error)
	BatchGetWorkflows(context.Context, *glue.BatchGetWorkflowsInput, ...func(*glue.Options)) (*glue.BatchGetWorkflowsOutput, error)
	GetBlueprint(context.Context, *glue.GetBlueprintInput, ...func(*glue.Options)) (*glue.GetBlueprintOutput, error)
	GetBlueprintRun(context.Context, *glue.GetBlueprintRunInput, ...func(*glue.Options)) (*glue.GetBlueprintRunOutput, error)
	GetBlueprintRuns(context.Context, *glue.GetBlueprintRunsInput, ...func(*glue.Options)) (*glue.GetBlueprintRunsOutput, error)
	GetCatalogImportStatus(context.Context, *glue.GetCatalogImportStatusInput, ...func(*glue.Options)) (*glue.GetCatalogImportStatusOutput, error)
	GetClassifier(context.Context, *glue.GetClassifierInput, ...func(*glue.Options)) (*glue.GetClassifierOutput, error)
	GetClassifiers(context.Context, *glue.GetClassifiersInput, ...func(*glue.Options)) (*glue.GetClassifiersOutput, error)
	GetColumnStatisticsForPartition(context.Context, *glue.GetColumnStatisticsForPartitionInput, ...func(*glue.Options)) (*glue.GetColumnStatisticsForPartitionOutput, error)
	GetColumnStatisticsForTable(context.Context, *glue.GetColumnStatisticsForTableInput, ...func(*glue.Options)) (*glue.GetColumnStatisticsForTableOutput, error)
	GetConnection(context.Context, *glue.GetConnectionInput, ...func(*glue.Options)) (*glue.GetConnectionOutput, error)
	GetConnections(context.Context, *glue.GetConnectionsInput, ...func(*glue.Options)) (*glue.GetConnectionsOutput, error)
	GetCrawler(context.Context, *glue.GetCrawlerInput, ...func(*glue.Options)) (*glue.GetCrawlerOutput, error)
	GetCrawlerMetrics(context.Context, *glue.GetCrawlerMetricsInput, ...func(*glue.Options)) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlers(context.Context, *glue.GetCrawlersInput, ...func(*glue.Options)) (*glue.GetCrawlersOutput, error)
	GetCustomEntityType(context.Context, *glue.GetCustomEntityTypeInput, ...func(*glue.Options)) (*glue.GetCustomEntityTypeOutput, error)
	GetDataCatalogEncryptionSettings(context.Context, *glue.GetDataCatalogEncryptionSettingsInput, ...func(*glue.Options)) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
	GetDatabase(context.Context, *glue.GetDatabaseInput, ...func(*glue.Options)) (*glue.GetDatabaseOutput, error)
	GetDatabases(context.Context, *glue.GetDatabasesInput, ...func(*glue.Options)) (*glue.GetDatabasesOutput, error)
	GetDataflowGraph(context.Context, *glue.GetDataflowGraphInput, ...func(*glue.Options)) (*glue.GetDataflowGraphOutput, error)
	GetDevEndpoint(context.Context, *glue.GetDevEndpointInput, ...func(*glue.Options)) (*glue.GetDevEndpointOutput, error)
	GetDevEndpoints(context.Context, *glue.GetDevEndpointsInput, ...func(*glue.Options)) (*glue.GetDevEndpointsOutput, error)
	GetJob(context.Context, *glue.GetJobInput, ...func(*glue.Options)) (*glue.GetJobOutput, error)
	GetJobBookmark(context.Context, *glue.GetJobBookmarkInput, ...func(*glue.Options)) (*glue.GetJobBookmarkOutput, error)
	GetJobRun(context.Context, *glue.GetJobRunInput, ...func(*glue.Options)) (*glue.GetJobRunOutput, error)
	GetJobRuns(context.Context, *glue.GetJobRunsInput, ...func(*glue.Options)) (*glue.GetJobRunsOutput, error)
	GetJobs(context.Context, *glue.GetJobsInput, ...func(*glue.Options)) (*glue.GetJobsOutput, error)
	GetMLTaskRun(context.Context, *glue.GetMLTaskRunInput, ...func(*glue.Options)) (*glue.GetMLTaskRunOutput, error)
	GetMLTaskRuns(context.Context, *glue.GetMLTaskRunsInput, ...func(*glue.Options)) (*glue.GetMLTaskRunsOutput, error)
	GetMLTransform(context.Context, *glue.GetMLTransformInput, ...func(*glue.Options)) (*glue.GetMLTransformOutput, error)
	GetMLTransforms(context.Context, *glue.GetMLTransformsInput, ...func(*glue.Options)) (*glue.GetMLTransformsOutput, error)
	GetMapping(context.Context, *glue.GetMappingInput, ...func(*glue.Options)) (*glue.GetMappingOutput, error)
	GetPartition(context.Context, *glue.GetPartitionInput, ...func(*glue.Options)) (*glue.GetPartitionOutput, error)
	GetPartitionIndexes(context.Context, *glue.GetPartitionIndexesInput, ...func(*glue.Options)) (*glue.GetPartitionIndexesOutput, error)
	GetPartitions(context.Context, *glue.GetPartitionsInput, ...func(*glue.Options)) (*glue.GetPartitionsOutput, error)
	GetPlan(context.Context, *glue.GetPlanInput, ...func(*glue.Options)) (*glue.GetPlanOutput, error)
	GetRegistry(context.Context, *glue.GetRegistryInput, ...func(*glue.Options)) (*glue.GetRegistryOutput, error)
	GetResourcePolicies(context.Context, *glue.GetResourcePoliciesInput, ...func(*glue.Options)) (*glue.GetResourcePoliciesOutput, error)
	GetResourcePolicy(context.Context, *glue.GetResourcePolicyInput, ...func(*glue.Options)) (*glue.GetResourcePolicyOutput, error)
	GetSchema(context.Context, *glue.GetSchemaInput, ...func(*glue.Options)) (*glue.GetSchemaOutput, error)
	GetSchemaByDefinition(context.Context, *glue.GetSchemaByDefinitionInput, ...func(*glue.Options)) (*glue.GetSchemaByDefinitionOutput, error)
	GetSchemaVersion(context.Context, *glue.GetSchemaVersionInput, ...func(*glue.Options)) (*glue.GetSchemaVersionOutput, error)
	GetSchemaVersionsDiff(context.Context, *glue.GetSchemaVersionsDiffInput, ...func(*glue.Options)) (*glue.GetSchemaVersionsDiffOutput, error)
	GetSecurityConfiguration(context.Context, *glue.GetSecurityConfigurationInput, ...func(*glue.Options)) (*glue.GetSecurityConfigurationOutput, error)
	GetSecurityConfigurations(context.Context, *glue.GetSecurityConfigurationsInput, ...func(*glue.Options)) (*glue.GetSecurityConfigurationsOutput, error)
	GetSession(context.Context, *glue.GetSessionInput, ...func(*glue.Options)) (*glue.GetSessionOutput, error)
	GetStatement(context.Context, *glue.GetStatementInput, ...func(*glue.Options)) (*glue.GetStatementOutput, error)
	GetTable(context.Context, *glue.GetTableInput, ...func(*glue.Options)) (*glue.GetTableOutput, error)
	GetTableVersion(context.Context, *glue.GetTableVersionInput, ...func(*glue.Options)) (*glue.GetTableVersionOutput, error)
	GetTableVersions(context.Context, *glue.GetTableVersionsInput, ...func(*glue.Options)) (*glue.GetTableVersionsOutput, error)
	GetTables(context.Context, *glue.GetTablesInput, ...func(*glue.Options)) (*glue.GetTablesOutput, error)
	GetTags(context.Context, *glue.GetTagsInput, ...func(*glue.Options)) (*glue.GetTagsOutput, error)
	GetTrigger(context.Context, *glue.GetTriggerInput, ...func(*glue.Options)) (*glue.GetTriggerOutput, error)
	GetTriggers(context.Context, *glue.GetTriggersInput, ...func(*glue.Options)) (*glue.GetTriggersOutput, error)
	GetUnfilteredPartitionMetadata(context.Context, *glue.GetUnfilteredPartitionMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredPartitionMetadataOutput, error)
	GetUnfilteredPartitionsMetadata(context.Context, *glue.GetUnfilteredPartitionsMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredPartitionsMetadataOutput, error)
	GetUnfilteredTableMetadata(context.Context, *glue.GetUnfilteredTableMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredTableMetadataOutput, error)
	GetUserDefinedFunction(context.Context, *glue.GetUserDefinedFunctionInput, ...func(*glue.Options)) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctions(context.Context, *glue.GetUserDefinedFunctionsInput, ...func(*glue.Options)) (*glue.GetUserDefinedFunctionsOutput, error)
	GetWorkflow(context.Context, *glue.GetWorkflowInput, ...func(*glue.Options)) (*glue.GetWorkflowOutput, error)
	GetWorkflowRun(context.Context, *glue.GetWorkflowRunInput, ...func(*glue.Options)) (*glue.GetWorkflowRunOutput, error)
	GetWorkflowRunProperties(context.Context, *glue.GetWorkflowRunPropertiesInput, ...func(*glue.Options)) (*glue.GetWorkflowRunPropertiesOutput, error)
	GetWorkflowRuns(context.Context, *glue.GetWorkflowRunsInput, ...func(*glue.Options)) (*glue.GetWorkflowRunsOutput, error)
	ListBlueprints(context.Context, *glue.ListBlueprintsInput, ...func(*glue.Options)) (*glue.ListBlueprintsOutput, error)
	ListCrawlers(context.Context, *glue.ListCrawlersInput, ...func(*glue.Options)) (*glue.ListCrawlersOutput, error)
	ListCrawls(context.Context, *glue.ListCrawlsInput, ...func(*glue.Options)) (*glue.ListCrawlsOutput, error)
	ListCustomEntityTypes(context.Context, *glue.ListCustomEntityTypesInput, ...func(*glue.Options)) (*glue.ListCustomEntityTypesOutput, error)
	ListDevEndpoints(context.Context, *glue.ListDevEndpointsInput, ...func(*glue.Options)) (*glue.ListDevEndpointsOutput, error)
	ListJobs(context.Context, *glue.ListJobsInput, ...func(*glue.Options)) (*glue.ListJobsOutput, error)
	ListMLTransforms(context.Context, *glue.ListMLTransformsInput, ...func(*glue.Options)) (*glue.ListMLTransformsOutput, error)
	ListRegistries(context.Context, *glue.ListRegistriesInput, ...func(*glue.Options)) (*glue.ListRegistriesOutput, error)
	ListSchemaVersions(context.Context, *glue.ListSchemaVersionsInput, ...func(*glue.Options)) (*glue.ListSchemaVersionsOutput, error)
	ListSchemas(context.Context, *glue.ListSchemasInput, ...func(*glue.Options)) (*glue.ListSchemasOutput, error)
	ListSessions(context.Context, *glue.ListSessionsInput, ...func(*glue.Options)) (*glue.ListSessionsOutput, error)
	ListStatements(context.Context, *glue.ListStatementsInput, ...func(*glue.Options)) (*glue.ListStatementsOutput, error)
	ListTriggers(context.Context, *glue.ListTriggersInput, ...func(*glue.Options)) (*glue.ListTriggersOutput, error)
	ListWorkflows(context.Context, *glue.ListWorkflowsInput, ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error)
	QuerySchemaVersionMetadata(context.Context, *glue.QuerySchemaVersionMetadataInput, ...func(*glue.Options)) (*glue.QuerySchemaVersionMetadataOutput, error)
	SearchTables(context.Context, *glue.SearchTablesInput, ...func(*glue.Options)) (*glue.SearchTablesOutput, error)
}

type SavingsplansClient interface {
	DescribeSavingsPlanRates(context.Context, *savingsplans.DescribeSavingsPlanRatesInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
	DescribeSavingsPlans(context.Context, *savingsplans.DescribeSavingsPlansInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOutput, error)
	DescribeSavingsPlansOfferingRates(context.Context, *savingsplans.DescribeSavingsPlansOfferingRatesInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
	DescribeSavingsPlansOfferings(context.Context, *savingsplans.DescribeSavingsPlansOfferingsInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
	ListTagsForResource(context.Context, *savingsplans.ListTagsForResourceInput, ...func(*savingsplans.Options)) (*savingsplans.ListTagsForResourceOutput, error)
}

type GuarddutyClient interface {
	DescribeMalwareScans(context.Context, *guardduty.DescribeMalwareScansInput, ...func(*guardduty.Options)) (*guardduty.DescribeMalwareScansOutput, error)
	DescribeOrganizationConfiguration(context.Context, *guardduty.DescribeOrganizationConfigurationInput, ...func(*guardduty.Options)) (*guardduty.DescribeOrganizationConfigurationOutput, error)
	DescribePublishingDestination(context.Context, *guardduty.DescribePublishingDestinationInput, ...func(*guardduty.Options)) (*guardduty.DescribePublishingDestinationOutput, error)
	GetAdministratorAccount(context.Context, *guardduty.GetAdministratorAccountInput, ...func(*guardduty.Options)) (*guardduty.GetAdministratorAccountOutput, error)
	GetDetector(context.Context, *guardduty.GetDetectorInput, ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error)
	GetFilter(context.Context, *guardduty.GetFilterInput, ...func(*guardduty.Options)) (*guardduty.GetFilterOutput, error)
	GetFindings(context.Context, *guardduty.GetFindingsInput, ...func(*guardduty.Options)) (*guardduty.GetFindingsOutput, error)
	GetFindingsStatistics(context.Context, *guardduty.GetFindingsStatisticsInput, ...func(*guardduty.Options)) (*guardduty.GetFindingsStatisticsOutput, error)
	GetIPSet(context.Context, *guardduty.GetIPSetInput, ...func(*guardduty.Options)) (*guardduty.GetIPSetOutput, error)
	GetInvitationsCount(context.Context, *guardduty.GetInvitationsCountInput, ...func(*guardduty.Options)) (*guardduty.GetInvitationsCountOutput, error)
	GetMalwareScanSettings(context.Context, *guardduty.GetMalwareScanSettingsInput, ...func(*guardduty.Options)) (*guardduty.GetMalwareScanSettingsOutput, error)
	GetMasterAccount(context.Context, *guardduty.GetMasterAccountInput, ...func(*guardduty.Options)) (*guardduty.GetMasterAccountOutput, error)
	GetMemberDetectors(context.Context, *guardduty.GetMemberDetectorsInput, ...func(*guardduty.Options)) (*guardduty.GetMemberDetectorsOutput, error)
	GetMembers(context.Context, *guardduty.GetMembersInput, ...func(*guardduty.Options)) (*guardduty.GetMembersOutput, error)
	GetRemainingFreeTrialDays(context.Context, *guardduty.GetRemainingFreeTrialDaysInput, ...func(*guardduty.Options)) (*guardduty.GetRemainingFreeTrialDaysOutput, error)
	GetThreatIntelSet(context.Context, *guardduty.GetThreatIntelSetInput, ...func(*guardduty.Options)) (*guardduty.GetThreatIntelSetOutput, error)
	GetUsageStatistics(context.Context, *guardduty.GetUsageStatisticsInput, ...func(*guardduty.Options)) (*guardduty.GetUsageStatisticsOutput, error)
	ListDetectors(context.Context, *guardduty.ListDetectorsInput, ...func(*guardduty.Options)) (*guardduty.ListDetectorsOutput, error)
	ListFilters(context.Context, *guardduty.ListFiltersInput, ...func(*guardduty.Options)) (*guardduty.ListFiltersOutput, error)
	ListFindings(context.Context, *guardduty.ListFindingsInput, ...func(*guardduty.Options)) (*guardduty.ListFindingsOutput, error)
	ListIPSets(context.Context, *guardduty.ListIPSetsInput, ...func(*guardduty.Options)) (*guardduty.ListIPSetsOutput, error)
	ListInvitations(context.Context, *guardduty.ListInvitationsInput, ...func(*guardduty.Options)) (*guardduty.ListInvitationsOutput, error)
	ListMembers(context.Context, *guardduty.ListMembersInput, ...func(*guardduty.Options)) (*guardduty.ListMembersOutput, error)
	ListOrganizationAdminAccounts(context.Context, *guardduty.ListOrganizationAdminAccountsInput, ...func(*guardduty.Options)) (*guardduty.ListOrganizationAdminAccountsOutput, error)
	ListPublishingDestinations(context.Context, *guardduty.ListPublishingDestinationsInput, ...func(*guardduty.Options)) (*guardduty.ListPublishingDestinationsOutput, error)
	ListTagsForResource(context.Context, *guardduty.ListTagsForResourceInput, ...func(*guardduty.Options)) (*guardduty.ListTagsForResourceOutput, error)
	ListThreatIntelSets(context.Context, *guardduty.ListThreatIntelSetsInput, ...func(*guardduty.Options)) (*guardduty.ListThreatIntelSetsOutput, error)
}

type IdentitystoreClient interface {
	DescribeGroup(context.Context, *identitystore.DescribeGroupInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupOutput, error)
	DescribeGroupMembership(context.Context, *identitystore.DescribeGroupMembershipInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupMembershipOutput, error)
	DescribeUser(context.Context, *identitystore.DescribeUserInput, ...func(*identitystore.Options)) (*identitystore.DescribeUserOutput, error)
	GetGroupId(context.Context, *identitystore.GetGroupIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupIdOutput, error)
	GetGroupMembershipId(context.Context, *identitystore.GetGroupMembershipIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupMembershipIdOutput, error)
	GetUserId(context.Context, *identitystore.GetUserIdInput, ...func(*identitystore.Options)) (*identitystore.GetUserIdOutput, error)
	ListGroupMemberships(context.Context, *identitystore.ListGroupMembershipsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsOutput, error)
	ListGroupMembershipsForMember(context.Context, *identitystore.ListGroupMembershipsForMemberInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsForMemberOutput, error)
	ListGroups(context.Context, *identitystore.ListGroupsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupsOutput, error)
	ListUsers(context.Context, *identitystore.ListUsersInput, ...func(*identitystore.Options)) (*identitystore.ListUsersOutput, error)
}

type InspectorClient interface {
	DescribeAssessmentRuns(context.Context, *inspector.DescribeAssessmentRunsInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentRunsOutput, error)
	DescribeAssessmentTargets(context.Context, *inspector.DescribeAssessmentTargetsInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentTargetsOutput, error)
	DescribeAssessmentTemplates(context.Context, *inspector.DescribeAssessmentTemplatesInput, ...func(*inspector.Options)) (*inspector.DescribeAssessmentTemplatesOutput, error)
	DescribeCrossAccountAccessRole(context.Context, *inspector.DescribeCrossAccountAccessRoleInput, ...func(*inspector.Options)) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
	DescribeExclusions(context.Context, *inspector.DescribeExclusionsInput, ...func(*inspector.Options)) (*inspector.DescribeExclusionsOutput, error)
	DescribeFindings(context.Context, *inspector.DescribeFindingsInput, ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error)
	DescribeResourceGroups(context.Context, *inspector.DescribeResourceGroupsInput, ...func(*inspector.Options)) (*inspector.DescribeResourceGroupsOutput, error)
	DescribeRulesPackages(context.Context, *inspector.DescribeRulesPackagesInput, ...func(*inspector.Options)) (*inspector.DescribeRulesPackagesOutput, error)
	GetAssessmentReport(context.Context, *inspector.GetAssessmentReportInput, ...func(*inspector.Options)) (*inspector.GetAssessmentReportOutput, error)
	GetExclusionsPreview(context.Context, *inspector.GetExclusionsPreviewInput, ...func(*inspector.Options)) (*inspector.GetExclusionsPreviewOutput, error)
	GetTelemetryMetadata(context.Context, *inspector.GetTelemetryMetadataInput, ...func(*inspector.Options)) (*inspector.GetTelemetryMetadataOutput, error)
	ListAssessmentRunAgents(context.Context, *inspector.ListAssessmentRunAgentsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentRunAgentsOutput, error)
	ListAssessmentRuns(context.Context, *inspector.ListAssessmentRunsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentRunsOutput, error)
	ListAssessmentTargets(context.Context, *inspector.ListAssessmentTargetsInput, ...func(*inspector.Options)) (*inspector.ListAssessmentTargetsOutput, error)
	ListAssessmentTemplates(context.Context, *inspector.ListAssessmentTemplatesInput, ...func(*inspector.Options)) (*inspector.ListAssessmentTemplatesOutput, error)
	ListEventSubscriptions(context.Context, *inspector.ListEventSubscriptionsInput, ...func(*inspector.Options)) (*inspector.ListEventSubscriptionsOutput, error)
	ListExclusions(context.Context, *inspector.ListExclusionsInput, ...func(*inspector.Options)) (*inspector.ListExclusionsOutput, error)
	ListFindings(context.Context, *inspector.ListFindingsInput, ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error)
	ListRulesPackages(context.Context, *inspector.ListRulesPackagesInput, ...func(*inspector.Options)) (*inspector.ListRulesPackagesOutput, error)
	ListTagsForResource(context.Context, *inspector.ListTagsForResourceInput, ...func(*inspector.Options)) (*inspector.ListTagsForResourceOutput, error)
}

type Inspector2Client interface {
	BatchGetAccountStatus(context.Context, *inspector2.BatchGetAccountStatusInput, ...func(*inspector2.Options)) (*inspector2.BatchGetAccountStatusOutput, error)
	BatchGetFreeTrialInfo(context.Context, *inspector2.BatchGetFreeTrialInfoInput, ...func(*inspector2.Options)) (*inspector2.BatchGetFreeTrialInfoOutput, error)
	DescribeOrganizationConfiguration(context.Context, *inspector2.DescribeOrganizationConfigurationInput, ...func(*inspector2.Options)) (*inspector2.DescribeOrganizationConfigurationOutput, error)
	GetConfiguration(context.Context, *inspector2.GetConfigurationInput, ...func(*inspector2.Options)) (*inspector2.GetConfigurationOutput, error)
	GetDelegatedAdminAccount(context.Context, *inspector2.GetDelegatedAdminAccountInput, ...func(*inspector2.Options)) (*inspector2.GetDelegatedAdminAccountOutput, error)
	GetFindingsReportStatus(context.Context, *inspector2.GetFindingsReportStatusInput, ...func(*inspector2.Options)) (*inspector2.GetFindingsReportStatusOutput, error)
	GetMember(context.Context, *inspector2.GetMemberInput, ...func(*inspector2.Options)) (*inspector2.GetMemberOutput, error)
	ListAccountPermissions(context.Context, *inspector2.ListAccountPermissionsInput, ...func(*inspector2.Options)) (*inspector2.ListAccountPermissionsOutput, error)
	ListCoverage(context.Context, *inspector2.ListCoverageInput, ...func(*inspector2.Options)) (*inspector2.ListCoverageOutput, error)
	ListCoverageStatistics(context.Context, *inspector2.ListCoverageStatisticsInput, ...func(*inspector2.Options)) (*inspector2.ListCoverageStatisticsOutput, error)
	ListDelegatedAdminAccounts(context.Context, *inspector2.ListDelegatedAdminAccountsInput, ...func(*inspector2.Options)) (*inspector2.ListDelegatedAdminAccountsOutput, error)
	ListFilters(context.Context, *inspector2.ListFiltersInput, ...func(*inspector2.Options)) (*inspector2.ListFiltersOutput, error)
	ListFindingAggregations(context.Context, *inspector2.ListFindingAggregationsInput, ...func(*inspector2.Options)) (*inspector2.ListFindingAggregationsOutput, error)
	ListFindings(context.Context, *inspector2.ListFindingsInput, ...func(*inspector2.Options)) (*inspector2.ListFindingsOutput, error)
	ListMembers(context.Context, *inspector2.ListMembersInput, ...func(*inspector2.Options)) (*inspector2.ListMembersOutput, error)
	ListTagsForResource(context.Context, *inspector2.ListTagsForResourceInput, ...func(*inspector2.Options)) (*inspector2.ListTagsForResourceOutput, error)
	ListUsageTotals(context.Context, *inspector2.ListUsageTotalsInput, ...func(*inspector2.Options)) (*inspector2.ListUsageTotalsOutput, error)
}

type IotClient interface {
	DescribeAccountAuditConfiguration(context.Context, *iot.DescribeAccountAuditConfigurationInput, ...func(*iot.Options)) (*iot.DescribeAccountAuditConfigurationOutput, error)
	DescribeAuditFinding(context.Context, *iot.DescribeAuditFindingInput, ...func(*iot.Options)) (*iot.DescribeAuditFindingOutput, error)
	DescribeAuditMitigationActionsTask(context.Context, *iot.DescribeAuditMitigationActionsTaskInput, ...func(*iot.Options)) (*iot.DescribeAuditMitigationActionsTaskOutput, error)
	DescribeAuditSuppression(context.Context, *iot.DescribeAuditSuppressionInput, ...func(*iot.Options)) (*iot.DescribeAuditSuppressionOutput, error)
	DescribeAuditTask(context.Context, *iot.DescribeAuditTaskInput, ...func(*iot.Options)) (*iot.DescribeAuditTaskOutput, error)
	DescribeAuthorizer(context.Context, *iot.DescribeAuthorizerInput, ...func(*iot.Options)) (*iot.DescribeAuthorizerOutput, error)
	DescribeBillingGroup(context.Context, *iot.DescribeBillingGroupInput, ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error)
	DescribeCACertificate(context.Context, *iot.DescribeCACertificateInput, ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error)
	DescribeCertificate(context.Context, *iot.DescribeCertificateInput, ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error)
	DescribeCustomMetric(context.Context, *iot.DescribeCustomMetricInput, ...func(*iot.Options)) (*iot.DescribeCustomMetricOutput, error)
	DescribeDefaultAuthorizer(context.Context, *iot.DescribeDefaultAuthorizerInput, ...func(*iot.Options)) (*iot.DescribeDefaultAuthorizerOutput, error)
	DescribeDetectMitigationActionsTask(context.Context, *iot.DescribeDetectMitigationActionsTaskInput, ...func(*iot.Options)) (*iot.DescribeDetectMitigationActionsTaskOutput, error)
	DescribeDimension(context.Context, *iot.DescribeDimensionInput, ...func(*iot.Options)) (*iot.DescribeDimensionOutput, error)
	DescribeDomainConfiguration(context.Context, *iot.DescribeDomainConfigurationInput, ...func(*iot.Options)) (*iot.DescribeDomainConfigurationOutput, error)
	DescribeEndpoint(context.Context, *iot.DescribeEndpointInput, ...func(*iot.Options)) (*iot.DescribeEndpointOutput, error)
	DescribeEventConfigurations(context.Context, *iot.DescribeEventConfigurationsInput, ...func(*iot.Options)) (*iot.DescribeEventConfigurationsOutput, error)
	DescribeFleetMetric(context.Context, *iot.DescribeFleetMetricInput, ...func(*iot.Options)) (*iot.DescribeFleetMetricOutput, error)
	DescribeIndex(context.Context, *iot.DescribeIndexInput, ...func(*iot.Options)) (*iot.DescribeIndexOutput, error)
	DescribeJob(context.Context, *iot.DescribeJobInput, ...func(*iot.Options)) (*iot.DescribeJobOutput, error)
	DescribeJobExecution(context.Context, *iot.DescribeJobExecutionInput, ...func(*iot.Options)) (*iot.DescribeJobExecutionOutput, error)
	DescribeJobTemplate(context.Context, *iot.DescribeJobTemplateInput, ...func(*iot.Options)) (*iot.DescribeJobTemplateOutput, error)
	DescribeManagedJobTemplate(context.Context, *iot.DescribeManagedJobTemplateInput, ...func(*iot.Options)) (*iot.DescribeManagedJobTemplateOutput, error)
	DescribeMitigationAction(context.Context, *iot.DescribeMitigationActionInput, ...func(*iot.Options)) (*iot.DescribeMitigationActionOutput, error)
	DescribeProvisioningTemplate(context.Context, *iot.DescribeProvisioningTemplateInput, ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateOutput, error)
	DescribeProvisioningTemplateVersion(context.Context, *iot.DescribeProvisioningTemplateVersionInput, ...func(*iot.Options)) (*iot.DescribeProvisioningTemplateVersionOutput, error)
	DescribeRoleAlias(context.Context, *iot.DescribeRoleAliasInput, ...func(*iot.Options)) (*iot.DescribeRoleAliasOutput, error)
	DescribeScheduledAudit(context.Context, *iot.DescribeScheduledAuditInput, ...func(*iot.Options)) (*iot.DescribeScheduledAuditOutput, error)
	DescribeSecurityProfile(context.Context, *iot.DescribeSecurityProfileInput, ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error)
	DescribeStream(context.Context, *iot.DescribeStreamInput, ...func(*iot.Options)) (*iot.DescribeStreamOutput, error)
	DescribeThing(context.Context, *iot.DescribeThingInput, ...func(*iot.Options)) (*iot.DescribeThingOutput, error)
	DescribeThingGroup(context.Context, *iot.DescribeThingGroupInput, ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error)
	DescribeThingRegistrationTask(context.Context, *iot.DescribeThingRegistrationTaskInput, ...func(*iot.Options)) (*iot.DescribeThingRegistrationTaskOutput, error)
	DescribeThingType(context.Context, *iot.DescribeThingTypeInput, ...func(*iot.Options)) (*iot.DescribeThingTypeOutput, error)
	GetBehaviorModelTrainingSummaries(context.Context, *iot.GetBehaviorModelTrainingSummariesInput, ...func(*iot.Options)) (*iot.GetBehaviorModelTrainingSummariesOutput, error)
	GetBucketsAggregation(context.Context, *iot.GetBucketsAggregationInput, ...func(*iot.Options)) (*iot.GetBucketsAggregationOutput, error)
	GetCardinality(context.Context, *iot.GetCardinalityInput, ...func(*iot.Options)) (*iot.GetCardinalityOutput, error)
	GetEffectivePolicies(context.Context, *iot.GetEffectivePoliciesInput, ...func(*iot.Options)) (*iot.GetEffectivePoliciesOutput, error)
	GetIndexingConfiguration(context.Context, *iot.GetIndexingConfigurationInput, ...func(*iot.Options)) (*iot.GetIndexingConfigurationOutput, error)
	GetJobDocument(context.Context, *iot.GetJobDocumentInput, ...func(*iot.Options)) (*iot.GetJobDocumentOutput, error)
	GetLoggingOptions(context.Context, *iot.GetLoggingOptionsInput, ...func(*iot.Options)) (*iot.GetLoggingOptionsOutput, error)
	GetOTAUpdate(context.Context, *iot.GetOTAUpdateInput, ...func(*iot.Options)) (*iot.GetOTAUpdateOutput, error)
	GetPercentiles(context.Context, *iot.GetPercentilesInput, ...func(*iot.Options)) (*iot.GetPercentilesOutput, error)
	GetPolicy(context.Context, *iot.GetPolicyInput, ...func(*iot.Options)) (*iot.GetPolicyOutput, error)
	GetPolicyVersion(context.Context, *iot.GetPolicyVersionInput, ...func(*iot.Options)) (*iot.GetPolicyVersionOutput, error)
	GetRegistrationCode(context.Context, *iot.GetRegistrationCodeInput, ...func(*iot.Options)) (*iot.GetRegistrationCodeOutput, error)
	GetStatistics(context.Context, *iot.GetStatisticsInput, ...func(*iot.Options)) (*iot.GetStatisticsOutput, error)
	GetTopicRule(context.Context, *iot.GetTopicRuleInput, ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleDestination(context.Context, *iot.GetTopicRuleDestinationInput, ...func(*iot.Options)) (*iot.GetTopicRuleDestinationOutput, error)
	GetV2LoggingOptions(context.Context, *iot.GetV2LoggingOptionsInput, ...func(*iot.Options)) (*iot.GetV2LoggingOptionsOutput, error)
	ListActiveViolations(context.Context, *iot.ListActiveViolationsInput, ...func(*iot.Options)) (*iot.ListActiveViolationsOutput, error)
	ListAttachedPolicies(context.Context, *iot.ListAttachedPoliciesInput, ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error)
	ListAuditFindings(context.Context, *iot.ListAuditFindingsInput, ...func(*iot.Options)) (*iot.ListAuditFindingsOutput, error)
	ListAuditMitigationActionsExecutions(context.Context, *iot.ListAuditMitigationActionsExecutionsInput, ...func(*iot.Options)) (*iot.ListAuditMitigationActionsExecutionsOutput, error)
	ListAuditMitigationActionsTasks(context.Context, *iot.ListAuditMitigationActionsTasksInput, ...func(*iot.Options)) (*iot.ListAuditMitigationActionsTasksOutput, error)
	ListAuditSuppressions(context.Context, *iot.ListAuditSuppressionsInput, ...func(*iot.Options)) (*iot.ListAuditSuppressionsOutput, error)
	ListAuditTasks(context.Context, *iot.ListAuditTasksInput, ...func(*iot.Options)) (*iot.ListAuditTasksOutput, error)
	ListAuthorizers(context.Context, *iot.ListAuthorizersInput, ...func(*iot.Options)) (*iot.ListAuthorizersOutput, error)
	ListBillingGroups(context.Context, *iot.ListBillingGroupsInput, ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error)
	ListCACertificates(context.Context, *iot.ListCACertificatesInput, ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error)
	ListCertificates(context.Context, *iot.ListCertificatesInput, ...func(*iot.Options)) (*iot.ListCertificatesOutput, error)
	ListCertificatesByCA(context.Context, *iot.ListCertificatesByCAInput, ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error)
	ListCustomMetrics(context.Context, *iot.ListCustomMetricsInput, ...func(*iot.Options)) (*iot.ListCustomMetricsOutput, error)
	ListDetectMitigationActionsExecutions(context.Context, *iot.ListDetectMitigationActionsExecutionsInput, ...func(*iot.Options)) (*iot.ListDetectMitigationActionsExecutionsOutput, error)
	ListDetectMitigationActionsTasks(context.Context, *iot.ListDetectMitigationActionsTasksInput, ...func(*iot.Options)) (*iot.ListDetectMitigationActionsTasksOutput, error)
	ListDimensions(context.Context, *iot.ListDimensionsInput, ...func(*iot.Options)) (*iot.ListDimensionsOutput, error)
	ListDomainConfigurations(context.Context, *iot.ListDomainConfigurationsInput, ...func(*iot.Options)) (*iot.ListDomainConfigurationsOutput, error)
	ListFleetMetrics(context.Context, *iot.ListFleetMetricsInput, ...func(*iot.Options)) (*iot.ListFleetMetricsOutput, error)
	ListIndices(context.Context, *iot.ListIndicesInput, ...func(*iot.Options)) (*iot.ListIndicesOutput, error)
	ListJobExecutionsForJob(context.Context, *iot.ListJobExecutionsForJobInput, ...func(*iot.Options)) (*iot.ListJobExecutionsForJobOutput, error)
	ListJobExecutionsForThing(context.Context, *iot.ListJobExecutionsForThingInput, ...func(*iot.Options)) (*iot.ListJobExecutionsForThingOutput, error)
	ListJobTemplates(context.Context, *iot.ListJobTemplatesInput, ...func(*iot.Options)) (*iot.ListJobTemplatesOutput, error)
	ListJobs(context.Context, *iot.ListJobsInput, ...func(*iot.Options)) (*iot.ListJobsOutput, error)
	ListManagedJobTemplates(context.Context, *iot.ListManagedJobTemplatesInput, ...func(*iot.Options)) (*iot.ListManagedJobTemplatesOutput, error)
	ListMetricValues(context.Context, *iot.ListMetricValuesInput, ...func(*iot.Options)) (*iot.ListMetricValuesOutput, error)
	ListMitigationActions(context.Context, *iot.ListMitigationActionsInput, ...func(*iot.Options)) (*iot.ListMitigationActionsOutput, error)
	ListOTAUpdates(context.Context, *iot.ListOTAUpdatesInput, ...func(*iot.Options)) (*iot.ListOTAUpdatesOutput, error)
	ListOutgoingCertificates(context.Context, *iot.ListOutgoingCertificatesInput, ...func(*iot.Options)) (*iot.ListOutgoingCertificatesOutput, error)
	ListPolicies(context.Context, *iot.ListPoliciesInput, ...func(*iot.Options)) (*iot.ListPoliciesOutput, error)
	ListPolicyPrincipals(context.Context, *iot.ListPolicyPrincipalsInput, ...func(*iot.Options)) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyVersions(context.Context, *iot.ListPolicyVersionsInput, ...func(*iot.Options)) (*iot.ListPolicyVersionsOutput, error)
	ListPrincipalPolicies(context.Context, *iot.ListPrincipalPoliciesInput, ...func(*iot.Options)) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalThings(context.Context, *iot.ListPrincipalThingsInput, ...func(*iot.Options)) (*iot.ListPrincipalThingsOutput, error)
	ListProvisioningTemplateVersions(context.Context, *iot.ListProvisioningTemplateVersionsInput, ...func(*iot.Options)) (*iot.ListProvisioningTemplateVersionsOutput, error)
	ListProvisioningTemplates(context.Context, *iot.ListProvisioningTemplatesInput, ...func(*iot.Options)) (*iot.ListProvisioningTemplatesOutput, error)
	ListRoleAliases(context.Context, *iot.ListRoleAliasesInput, ...func(*iot.Options)) (*iot.ListRoleAliasesOutput, error)
	ListScheduledAudits(context.Context, *iot.ListScheduledAuditsInput, ...func(*iot.Options)) (*iot.ListScheduledAuditsOutput, error)
	ListSecurityProfiles(context.Context, *iot.ListSecurityProfilesInput, ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error)
	ListSecurityProfilesForTarget(context.Context, *iot.ListSecurityProfilesForTargetInput, ...func(*iot.Options)) (*iot.ListSecurityProfilesForTargetOutput, error)
	ListStreams(context.Context, *iot.ListStreamsInput, ...func(*iot.Options)) (*iot.ListStreamsOutput, error)
	ListTagsForResource(context.Context, *iot.ListTagsForResourceInput, ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error)
	ListTargetsForPolicy(context.Context, *iot.ListTargetsForPolicyInput, ...func(*iot.Options)) (*iot.ListTargetsForPolicyOutput, error)
	ListTargetsForSecurityProfile(context.Context, *iot.ListTargetsForSecurityProfileInput, ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error)
	ListThingGroups(context.Context, *iot.ListThingGroupsInput, ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error)
	ListThingGroupsForThing(context.Context, *iot.ListThingGroupsForThingInput, ...func(*iot.Options)) (*iot.ListThingGroupsForThingOutput, error)
	ListThingPrincipals(context.Context, *iot.ListThingPrincipalsInput, ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error)
	ListThingRegistrationTaskReports(context.Context, *iot.ListThingRegistrationTaskReportsInput, ...func(*iot.Options)) (*iot.ListThingRegistrationTaskReportsOutput, error)
	ListThingRegistrationTasks(context.Context, *iot.ListThingRegistrationTasksInput, ...func(*iot.Options)) (*iot.ListThingRegistrationTasksOutput, error)
	ListThingTypes(context.Context, *iot.ListThingTypesInput, ...func(*iot.Options)) (*iot.ListThingTypesOutput, error)
	ListThings(context.Context, *iot.ListThingsInput, ...func(*iot.Options)) (*iot.ListThingsOutput, error)
	ListThingsInBillingGroup(context.Context, *iot.ListThingsInBillingGroupInput, ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error)
	ListThingsInThingGroup(context.Context, *iot.ListThingsInThingGroupInput, ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error)
	ListTopicRuleDestinations(context.Context, *iot.ListTopicRuleDestinationsInput, ...func(*iot.Options)) (*iot.ListTopicRuleDestinationsOutput, error)
	ListTopicRules(context.Context, *iot.ListTopicRulesInput, ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error)
	ListV2LoggingLevels(context.Context, *iot.ListV2LoggingLevelsInput, ...func(*iot.Options)) (*iot.ListV2LoggingLevelsOutput, error)
	ListViolationEvents(context.Context, *iot.ListViolationEventsInput, ...func(*iot.Options)) (*iot.ListViolationEventsOutput, error)
	SearchIndex(context.Context, *iot.SearchIndexInput, ...func(*iot.Options)) (*iot.SearchIndexOutput, error)
}

type KinesisClient interface {
	DescribeLimits(context.Context, *kinesis.DescribeLimitsInput, ...func(*kinesis.Options)) (*kinesis.DescribeLimitsOutput, error)
	DescribeStream(context.Context, *kinesis.DescribeStreamInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamOutput, error)
	DescribeStreamConsumer(context.Context, *kinesis.DescribeStreamConsumerInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamConsumerOutput, error)
	DescribeStreamSummary(context.Context, *kinesis.DescribeStreamSummaryInput, ...func(*kinesis.Options)) (*kinesis.DescribeStreamSummaryOutput, error)
	GetRecords(context.Context, *kinesis.GetRecordsInput, ...func(*kinesis.Options)) (*kinesis.GetRecordsOutput, error)
	GetShardIterator(context.Context, *kinesis.GetShardIteratorInput, ...func(*kinesis.Options)) (*kinesis.GetShardIteratorOutput, error)
	ListShards(context.Context, *kinesis.ListShardsInput, ...func(*kinesis.Options)) (*kinesis.ListShardsOutput, error)
	ListStreamConsumers(context.Context, *kinesis.ListStreamConsumersInput, ...func(*kinesis.Options)) (*kinesis.ListStreamConsumersOutput, error)
	ListStreams(context.Context, *kinesis.ListStreamsInput, ...func(*kinesis.Options)) (*kinesis.ListStreamsOutput, error)
	ListTagsForStream(context.Context, *kinesis.ListTagsForStreamInput, ...func(*kinesis.Options)) (*kinesis.ListTagsForStreamOutput, error)
}

type KmsClient interface {
	DescribeCustomKeyStores(context.Context, *kms.DescribeCustomKeyStoresInput, ...func(*kms.Options)) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeKey(context.Context, *kms.DescribeKeyInput, ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	GetKeyPolicy(context.Context, *kms.GetKeyPolicyInput, ...func(*kms.Options)) (*kms.GetKeyPolicyOutput, error)
	GetKeyRotationStatus(context.Context, *kms.GetKeyRotationStatusInput, ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error)
	GetParametersForImport(context.Context, *kms.GetParametersForImportInput, ...func(*kms.Options)) (*kms.GetParametersForImportOutput, error)
	GetPublicKey(context.Context, *kms.GetPublicKeyInput, ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error)
	ListAliases(context.Context, *kms.ListAliasesInput, ...func(*kms.Options)) (*kms.ListAliasesOutput, error)
	ListGrants(context.Context, *kms.ListGrantsInput, ...func(*kms.Options)) (*kms.ListGrantsOutput, error)
	ListKeyPolicies(context.Context, *kms.ListKeyPoliciesInput, ...func(*kms.Options)) (*kms.ListKeyPoliciesOutput, error)
	ListKeys(context.Context, *kms.ListKeysInput, ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	ListResourceTags(context.Context, *kms.ListResourceTagsInput, ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error)
	ListRetirableGrants(context.Context, *kms.ListRetirableGrantsInput, ...func(*kms.Options)) (*kms.ListRetirableGrantsOutput, error)
}

type LambdaClient interface {
	GetAccountSettings(context.Context, *lambda.GetAccountSettingsInput, ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error)
	GetAlias(context.Context, *lambda.GetAliasInput, ...func(*lambda.Options)) (*lambda.GetAliasOutput, error)
	GetCodeSigningConfig(context.Context, *lambda.GetCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	GetEventSourceMapping(context.Context, *lambda.GetEventSourceMappingInput, ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error)
	GetFunction(context.Context, *lambda.GetFunctionInput, ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	GetFunctionCodeSigningConfig(context.Context, *lambda.GetFunctionCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionConcurrency(context.Context, *lambda.GetFunctionConcurrencyInput, ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConfiguration(context.Context, *lambda.GetFunctionConfigurationInput, ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error)
	GetFunctionEventInvokeConfig(context.Context, *lambda.GetFunctionEventInvokeConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionUrlConfig(context.Context, *lambda.GetFunctionUrlConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error)
	GetLayerVersion(context.Context, *lambda.GetLayerVersionInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionByArn(context.Context, *lambda.GetLayerVersionByArnInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionPolicy(context.Context, *lambda.GetLayerVersionPolicyInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	GetPolicy(context.Context, *lambda.GetPolicyInput, ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	GetProvisionedConcurrencyConfig(context.Context, *lambda.GetProvisionedConcurrencyConfigInput, ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetRuntimeManagementConfig(context.Context, *lambda.GetRuntimeManagementConfigInput, ...func(*lambda.Options)) (*lambda.GetRuntimeManagementConfigOutput, error)
	ListAliases(context.Context, *lambda.ListAliasesInput, ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	ListCodeSigningConfigs(context.Context, *lambda.ListCodeSigningConfigsInput, ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error)
	ListEventSourceMappings(context.Context, *lambda.ListEventSourceMappingsInput, ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	ListFunctionEventInvokeConfigs(context.Context, *lambda.ListFunctionEventInvokeConfigsInput, ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionUrlConfigs(context.Context, *lambda.ListFunctionUrlConfigsInput, ...func(*lambda.Options)) (*lambda.ListFunctionUrlConfigsOutput, error)
	ListFunctions(context.Context, *lambda.ListFunctionsInput, ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	ListFunctionsByCodeSigningConfig(context.Context, *lambda.ListFunctionsByCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListLayerVersions(context.Context, *lambda.ListLayerVersionsInput, ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	ListLayers(context.Context, *lambda.ListLayersInput, ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	ListProvisionedConcurrencyConfigs(context.Context, *lambda.ListProvisionedConcurrencyConfigsInput, ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListTags(context.Context, *lambda.ListTagsInput, ...func(*lambda.Options)) (*lambda.ListTagsOutput, error)
	ListVersionsByFunction(context.Context, *lambda.ListVersionsByFunctionInput, ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
}

type LightsailClient interface {
	GetActiveNames(context.Context, *lightsail.GetActiveNamesInput, ...func(*lightsail.Options)) (*lightsail.GetActiveNamesOutput, error)
	GetAlarms(context.Context, *lightsail.GetAlarmsInput, ...func(*lightsail.Options)) (*lightsail.GetAlarmsOutput, error)
	GetAutoSnapshots(context.Context, *lightsail.GetAutoSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetAutoSnapshotsOutput, error)
	GetBlueprints(context.Context, *lightsail.GetBlueprintsInput, ...func(*lightsail.Options)) (*lightsail.GetBlueprintsOutput, error)
	GetBucketAccessKeys(context.Context, *lightsail.GetBucketAccessKeysInput, ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error)
	GetBucketBundles(context.Context, *lightsail.GetBucketBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetBucketBundlesOutput, error)
	GetBucketMetricData(context.Context, *lightsail.GetBucketMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetBucketMetricDataOutput, error)
	GetBuckets(context.Context, *lightsail.GetBucketsInput, ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error)
	GetBundles(context.Context, *lightsail.GetBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetBundlesOutput, error)
	GetCertificates(context.Context, *lightsail.GetCertificatesInput, ...func(*lightsail.Options)) (*lightsail.GetCertificatesOutput, error)
	GetCloudFormationStackRecords(context.Context, *lightsail.GetCloudFormationStackRecordsInput, ...func(*lightsail.Options)) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetContactMethods(context.Context, *lightsail.GetContactMethodsInput, ...func(*lightsail.Options)) (*lightsail.GetContactMethodsOutput, error)
	GetContainerAPIMetadata(context.Context, *lightsail.GetContainerAPIMetadataInput, ...func(*lightsail.Options)) (*lightsail.GetContainerAPIMetadataOutput, error)
	GetContainerImages(context.Context, *lightsail.GetContainerImagesInput, ...func(*lightsail.Options)) (*lightsail.GetContainerImagesOutput, error)
	GetContainerLog(context.Context, *lightsail.GetContainerLogInput, ...func(*lightsail.Options)) (*lightsail.GetContainerLogOutput, error)
	GetContainerServiceDeployments(context.Context, *lightsail.GetContainerServiceDeploymentsInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServiceDeploymentsOutput, error)
	GetContainerServiceMetricData(context.Context, *lightsail.GetContainerServiceMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServiceMetricDataOutput, error)
	GetContainerServicePowers(context.Context, *lightsail.GetContainerServicePowersInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServicePowersOutput, error)
	GetContainerServices(context.Context, *lightsail.GetContainerServicesInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServicesOutput, error)
	GetDisk(context.Context, *lightsail.GetDiskInput, ...func(*lightsail.Options)) (*lightsail.GetDiskOutput, error)
	GetDiskSnapshot(context.Context, *lightsail.GetDiskSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshots(context.Context, *lightsail.GetDiskSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDisks(context.Context, *lightsail.GetDisksInput, ...func(*lightsail.Options)) (*lightsail.GetDisksOutput, error)
	GetDistributionBundles(context.Context, *lightsail.GetDistributionBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionBundlesOutput, error)
	GetDistributionLatestCacheReset(context.Context, *lightsail.GetDistributionLatestCacheResetInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionLatestCacheResetOutput, error)
	GetDistributionMetricData(context.Context, *lightsail.GetDistributionMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionMetricDataOutput, error)
	GetDistributions(context.Context, *lightsail.GetDistributionsInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionsOutput, error)
	GetDomain(context.Context, *lightsail.GetDomainInput, ...func(*lightsail.Options)) (*lightsail.GetDomainOutput, error)
	GetDomains(context.Context, *lightsail.GetDomainsInput, ...func(*lightsail.Options)) (*lightsail.GetDomainsOutput, error)
	GetExportSnapshotRecords(context.Context, *lightsail.GetExportSnapshotRecordsInput, ...func(*lightsail.Options)) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetInstance(context.Context, *lightsail.GetInstanceInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceOutput, error)
	GetInstanceAccessDetails(context.Context, *lightsail.GetInstanceAccessDetailsInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceMetricData(context.Context, *lightsail.GetInstanceMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstancePortStates(context.Context, *lightsail.GetInstancePortStatesInput, ...func(*lightsail.Options)) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstanceSnapshot(context.Context, *lightsail.GetInstanceSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshots(context.Context, *lightsail.GetInstanceSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceState(context.Context, *lightsail.GetInstanceStateInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceStateOutput, error)
	GetInstances(context.Context, *lightsail.GetInstancesInput, ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error)
	GetKeyPair(context.Context, *lightsail.GetKeyPairInput, ...func(*lightsail.Options)) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairs(context.Context, *lightsail.GetKeyPairsInput, ...func(*lightsail.Options)) (*lightsail.GetKeyPairsOutput, error)
	GetLoadBalancer(context.Context, *lightsail.GetLoadBalancerInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerMetricData(context.Context, *lightsail.GetLoadBalancerMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerTlsCertificates(context.Context, *lightsail.GetLoadBalancerTlsCertificatesInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsPolicies(context.Context, *lightsail.GetLoadBalancerTlsPoliciesInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsPoliciesOutput, error)
	GetLoadBalancers(context.Context, *lightsail.GetLoadBalancersInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancersOutput, error)
	GetOperation(context.Context, *lightsail.GetOperationInput, ...func(*lightsail.Options)) (*lightsail.GetOperationOutput, error)
	GetOperations(context.Context, *lightsail.GetOperationsInput, ...func(*lightsail.Options)) (*lightsail.GetOperationsOutput, error)
	GetOperationsForResource(context.Context, *lightsail.GetOperationsForResourceInput, ...func(*lightsail.Options)) (*lightsail.GetOperationsForResourceOutput, error)
	GetRegions(context.Context, *lightsail.GetRegionsInput, ...func(*lightsail.Options)) (*lightsail.GetRegionsOutput, error)
	GetRelationalDatabase(context.Context, *lightsail.GetRelationalDatabaseInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseBlueprints(context.Context, *lightsail.GetRelationalDatabaseBlueprintsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBundles(context.Context, *lightsail.GetRelationalDatabaseBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseEvents(context.Context, *lightsail.GetRelationalDatabaseEventsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseLogEvents(context.Context, *lightsail.GetRelationalDatabaseLogEventsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogStreams(context.Context, *lightsail.GetRelationalDatabaseLogStreamsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseMasterUserPassword(context.Context, *lightsail.GetRelationalDatabaseMasterUserPasswordInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMetricData(context.Context, *lightsail.GetRelationalDatabaseMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseParameters(context.Context, *lightsail.GetRelationalDatabaseParametersInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseSnapshot(context.Context, *lightsail.GetRelationalDatabaseSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshots(context.Context, *lightsail.GetRelationalDatabaseSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabases(context.Context, *lightsail.GetRelationalDatabasesInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error)
	GetStaticIp(context.Context, *lightsail.GetStaticIpInput, ...func(*lightsail.Options)) (*lightsail.GetStaticIpOutput, error)
	GetStaticIps(context.Context, *lightsail.GetStaticIpsInput, ...func(*lightsail.Options)) (*lightsail.GetStaticIpsOutput, error)
}

type MqClient interface {
	DescribeBroker(context.Context, *mq.DescribeBrokerInput, ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error)
	DescribeBrokerEngineTypes(context.Context, *mq.DescribeBrokerEngineTypesInput, ...func(*mq.Options)) (*mq.DescribeBrokerEngineTypesOutput, error)
	DescribeBrokerInstanceOptions(context.Context, *mq.DescribeBrokerInstanceOptionsInput, ...func(*mq.Options)) (*mq.DescribeBrokerInstanceOptionsOutput, error)
	DescribeConfiguration(context.Context, *mq.DescribeConfigurationInput, ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationRevision(context.Context, *mq.DescribeConfigurationRevisionInput, ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error)
	DescribeUser(context.Context, *mq.DescribeUserInput, ...func(*mq.Options)) (*mq.DescribeUserOutput, error)
	ListBrokers(context.Context, *mq.ListBrokersInput, ...func(*mq.Options)) (*mq.ListBrokersOutput, error)
	ListConfigurationRevisions(context.Context, *mq.ListConfigurationRevisionsInput, ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error)
	ListConfigurations(context.Context, *mq.ListConfigurationsInput, ...func(*mq.Options)) (*mq.ListConfigurationsOutput, error)
	ListTags(context.Context, *mq.ListTagsInput, ...func(*mq.Options)) (*mq.ListTagsOutput, error)
	ListUsers(context.Context, *mq.ListUsersInput, ...func(*mq.Options)) (*mq.ListUsersOutput, error)
}

type MwaaClient interface {
	GetEnvironment(context.Context, *mwaa.GetEnvironmentInput, ...func(*mwaa.Options)) (*mwaa.GetEnvironmentOutput, error)
	ListEnvironments(context.Context, *mwaa.ListEnvironmentsInput, ...func(*mwaa.Options)) (*mwaa.ListEnvironmentsOutput, error)
	ListTagsForResource(context.Context, *mwaa.ListTagsForResourceInput, ...func(*mwaa.Options)) (*mwaa.ListTagsForResourceOutput, error)
}

type NeptuneClient interface {
	DescribeDBClusterEndpoints(context.Context, *neptune.DescribeDBClusterEndpointsInput, ...func(*neptune.Options)) (*neptune.DescribeDBClusterEndpointsOutput, error)
	DescribeDBClusterParameterGroups(context.Context, *neptune.DescribeDBClusterParameterGroupsInput, ...func(*neptune.Options)) (*neptune.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(context.Context, *neptune.DescribeDBClusterParametersInput, ...func(*neptune.Options)) (*neptune.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(context.Context, *neptune.DescribeDBClusterSnapshotAttributesInput, ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(context.Context, *neptune.DescribeDBClusterSnapshotsInput, ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(context.Context, *neptune.DescribeDBClustersInput, ...func(*neptune.Options)) (*neptune.DescribeDBClustersOutput, error)
	DescribeDBEngineVersions(context.Context, *neptune.DescribeDBEngineVersionsInput, ...func(*neptune.Options)) (*neptune.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstances(context.Context, *neptune.DescribeDBInstancesInput, ...func(*neptune.Options)) (*neptune.DescribeDBInstancesOutput, error)
	DescribeDBParameterGroups(context.Context, *neptune.DescribeDBParameterGroupsInput, ...func(*neptune.Options)) (*neptune.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameters(context.Context, *neptune.DescribeDBParametersInput, ...func(*neptune.Options)) (*neptune.DescribeDBParametersOutput, error)
	DescribeDBSubnetGroups(context.Context, *neptune.DescribeDBSubnetGroupsInput, ...func(*neptune.Options)) (*neptune.DescribeDBSubnetGroupsOutput, error)
	DescribeEngineDefaultClusterParameters(context.Context, *neptune.DescribeEngineDefaultClusterParametersInput, ...func(*neptune.Options)) (*neptune.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEngineDefaultParameters(context.Context, *neptune.DescribeEngineDefaultParametersInput, ...func(*neptune.Options)) (*neptune.DescribeEngineDefaultParametersOutput, error)
	DescribeEventCategories(context.Context, *neptune.DescribeEventCategoriesInput, ...func(*neptune.Options)) (*neptune.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *neptune.DescribeEventSubscriptionsInput, ...func(*neptune.Options)) (*neptune.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *neptune.DescribeEventsInput, ...func(*neptune.Options)) (*neptune.DescribeEventsOutput, error)
	DescribeGlobalClusters(context.Context, *neptune.DescribeGlobalClustersInput, ...func(*neptune.Options)) (*neptune.DescribeGlobalClustersOutput, error)
	DescribeOrderableDBInstanceOptions(context.Context, *neptune.DescribeOrderableDBInstanceOptionsInput, ...func(*neptune.Options)) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribePendingMaintenanceActions(context.Context, *neptune.DescribePendingMaintenanceActionsInput, ...func(*neptune.Options)) (*neptune.DescribePendingMaintenanceActionsOutput, error)
	DescribeValidDBInstanceModifications(context.Context, *neptune.DescribeValidDBInstanceModificationsInput, ...func(*neptune.Options)) (*neptune.DescribeValidDBInstanceModificationsOutput, error)
	ListTagsForResource(context.Context, *neptune.ListTagsForResourceInput, ...func(*neptune.Options)) (*neptune.ListTagsForResourceOutput, error)
}

type OrganizationsClient interface {
	DescribeAccount(context.Context, *organizations.DescribeAccountInput, ...func(*organizations.Options)) (*organizations.DescribeAccountOutput, error)
	DescribeCreateAccountStatus(context.Context, *organizations.DescribeCreateAccountStatusInput, ...func(*organizations.Options)) (*organizations.DescribeCreateAccountStatusOutput, error)
	DescribeEffectivePolicy(context.Context, *organizations.DescribeEffectivePolicyInput, ...func(*organizations.Options)) (*organizations.DescribeEffectivePolicyOutput, error)
	DescribeHandshake(context.Context, *organizations.DescribeHandshakeInput, ...func(*organizations.Options)) (*organizations.DescribeHandshakeOutput, error)
	DescribeOrganization(context.Context, *organizations.DescribeOrganizationInput, ...func(*organizations.Options)) (*organizations.DescribeOrganizationOutput, error)
	DescribeOrganizationalUnit(context.Context, *organizations.DescribeOrganizationalUnitInput, ...func(*organizations.Options)) (*organizations.DescribeOrganizationalUnitOutput, error)
	DescribePolicy(context.Context, *organizations.DescribePolicyInput, ...func(*organizations.Options)) (*organizations.DescribePolicyOutput, error)
	DescribeResourcePolicy(context.Context, *organizations.DescribeResourcePolicyInput, ...func(*organizations.Options)) (*organizations.DescribeResourcePolicyOutput, error)
	ListAWSServiceAccessForOrganization(context.Context, *organizations.ListAWSServiceAccessForOrganizationInput, ...func(*organizations.Options)) (*organizations.ListAWSServiceAccessForOrganizationOutput, error)
	ListAccounts(context.Context, *organizations.ListAccountsInput, ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
	ListAccountsForParent(context.Context, *organizations.ListAccountsForParentInput, ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
	ListChildren(context.Context, *organizations.ListChildrenInput, ...func(*organizations.Options)) (*organizations.ListChildrenOutput, error)
	ListCreateAccountStatus(context.Context, *organizations.ListCreateAccountStatusInput, ...func(*organizations.Options)) (*organizations.ListCreateAccountStatusOutput, error)
	ListDelegatedAdministrators(context.Context, *organizations.ListDelegatedAdministratorsInput, ...func(*organizations.Options)) (*organizations.ListDelegatedAdministratorsOutput, error)
	ListDelegatedServicesForAccount(context.Context, *organizations.ListDelegatedServicesForAccountInput, ...func(*organizations.Options)) (*organizations.ListDelegatedServicesForAccountOutput, error)
	ListHandshakesForAccount(context.Context, *organizations.ListHandshakesForAccountInput, ...func(*organizations.Options)) (*organizations.ListHandshakesForAccountOutput, error)
	ListHandshakesForOrganization(context.Context, *organizations.ListHandshakesForOrganizationInput, ...func(*organizations.Options)) (*organizations.ListHandshakesForOrganizationOutput, error)
	ListOrganizationalUnitsForParent(context.Context, *organizations.ListOrganizationalUnitsForParentInput, ...func(*organizations.Options)) (*organizations.ListOrganizationalUnitsForParentOutput, error)
	ListParents(context.Context, *organizations.ListParentsInput, ...func(*organizations.Options)) (*organizations.ListParentsOutput, error)
	ListPolicies(context.Context, *organizations.ListPoliciesInput, ...func(*organizations.Options)) (*organizations.ListPoliciesOutput, error)
	ListPoliciesForTarget(context.Context, *organizations.ListPoliciesForTargetInput, ...func(*organizations.Options)) (*organizations.ListPoliciesForTargetOutput, error)
	ListRoots(context.Context, *organizations.ListRootsInput, ...func(*organizations.Options)) (*organizations.ListRootsOutput, error)
	ListTagsForResource(context.Context, *organizations.ListTagsForResourceInput, ...func(*organizations.Options)) (*organizations.ListTagsForResourceOutput, error)
	ListTargetsForPolicy(context.Context, *organizations.ListTargetsForPolicyInput, ...func(*organizations.Options)) (*organizations.ListTargetsForPolicyOutput, error)
}

type QldbClient interface {
	DescribeJournalKinesisStream(context.Context, *qldb.DescribeJournalKinesisStreamInput, ...func(*qldb.Options)) (*qldb.DescribeJournalKinesisStreamOutput, error)
	DescribeJournalS3Export(context.Context, *qldb.DescribeJournalS3ExportInput, ...func(*qldb.Options)) (*qldb.DescribeJournalS3ExportOutput, error)
	DescribeLedger(context.Context, *qldb.DescribeLedgerInput, ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error)
	GetBlock(context.Context, *qldb.GetBlockInput, ...func(*qldb.Options)) (*qldb.GetBlockOutput, error)
	GetDigest(context.Context, *qldb.GetDigestInput, ...func(*qldb.Options)) (*qldb.GetDigestOutput, error)
	GetRevision(context.Context, *qldb.GetRevisionInput, ...func(*qldb.Options)) (*qldb.GetRevisionOutput, error)
	ListJournalKinesisStreamsForLedger(context.Context, *qldb.ListJournalKinesisStreamsForLedgerInput, ...func(*qldb.Options)) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error)
	ListJournalS3Exports(context.Context, *qldb.ListJournalS3ExportsInput, ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsOutput, error)
	ListJournalS3ExportsForLedger(context.Context, *qldb.ListJournalS3ExportsForLedgerInput, ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsForLedgerOutput, error)
	ListLedgers(context.Context, *qldb.ListLedgersInput, ...func(*qldb.Options)) (*qldb.ListLedgersOutput, error)
	ListTagsForResource(context.Context, *qldb.ListTagsForResourceInput, ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error)
}

type RdsClient interface {
	DescribeAccountAttributes(context.Context, *rds.DescribeAccountAttributesInput, ...func(*rds.Options)) (*rds.DescribeAccountAttributesOutput, error)
	DescribeCertificates(context.Context, *rds.DescribeCertificatesInput, ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error)
	DescribeDBClusterBacktracks(context.Context, *rds.DescribeDBClusterBacktracksInput, ...func(*rds.Options)) (*rds.DescribeDBClusterBacktracksOutput, error)
	DescribeDBClusterEndpoints(context.Context, *rds.DescribeDBClusterEndpointsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterEndpointsOutput, error)
	DescribeDBClusterParameterGroups(context.Context, *rds.DescribeDBClusterParameterGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(context.Context, *rds.DescribeDBClusterParametersInput, ...func(*rds.Options)) (*rds.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(context.Context, *rds.DescribeDBClusterSnapshotAttributesInput, ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(context.Context, *rds.DescribeDBClusterSnapshotsInput, ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(context.Context, *rds.DescribeDBClustersInput, ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error)
	DescribeDBEngineVersions(context.Context, *rds.DescribeDBEngineVersionsInput, ...func(*rds.Options)) (*rds.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstanceAutomatedBackups(context.Context, *rds.DescribeDBInstanceAutomatedBackupsInput, ...func(*rds.Options)) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error)
	DescribeDBInstances(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)
	DescribeDBLogFiles(context.Context, *rds.DescribeDBLogFilesInput, ...func(*rds.Options)) (*rds.DescribeDBLogFilesOutput, error)
	DescribeDBParameterGroups(context.Context, *rds.DescribeDBParameterGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameters(context.Context, *rds.DescribeDBParametersInput, ...func(*rds.Options)) (*rds.DescribeDBParametersOutput, error)
	DescribeDBProxies(context.Context, *rds.DescribeDBProxiesInput, ...func(*rds.Options)) (*rds.DescribeDBProxiesOutput, error)
	DescribeDBProxyEndpoints(context.Context, *rds.DescribeDBProxyEndpointsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyEndpointsOutput, error)
	DescribeDBProxyTargetGroups(context.Context, *rds.DescribeDBProxyTargetGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyTargetGroupsOutput, error)
	DescribeDBProxyTargets(context.Context, *rds.DescribeDBProxyTargetsInput, ...func(*rds.Options)) (*rds.DescribeDBProxyTargetsOutput, error)
	DescribeDBSecurityGroups(context.Context, *rds.DescribeDBSecurityGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSnapshotAttributes(context.Context, *rds.DescribeDBSnapshotAttributesInput, ...func(*rds.Options)) (*rds.DescribeDBSnapshotAttributesOutput, error)
	DescribeDBSnapshots(context.Context, *rds.DescribeDBSnapshotsInput, ...func(*rds.Options)) (*rds.DescribeDBSnapshotsOutput, error)
	DescribeDBSubnetGroups(context.Context, *rds.DescribeDBSubnetGroupsInput, ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error)
	DescribeEngineDefaultClusterParameters(context.Context, *rds.DescribeEngineDefaultClusterParametersInput, ...func(*rds.Options)) (*rds.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEngineDefaultParameters(context.Context, *rds.DescribeEngineDefaultParametersInput, ...func(*rds.Options)) (*rds.DescribeEngineDefaultParametersOutput, error)
	DescribeEventCategories(context.Context, *rds.DescribeEventCategoriesInput, ...func(*rds.Options)) (*rds.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *rds.DescribeEventSubscriptionsInput, ...func(*rds.Options)) (*rds.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *rds.DescribeEventsInput, ...func(*rds.Options)) (*rds.DescribeEventsOutput, error)
	DescribeExportTasks(context.Context, *rds.DescribeExportTasksInput, ...func(*rds.Options)) (*rds.DescribeExportTasksOutput, error)
	DescribeGlobalClusters(context.Context, *rds.DescribeGlobalClustersInput, ...func(*rds.Options)) (*rds.DescribeGlobalClustersOutput, error)
	DescribeOptionGroupOptions(context.Context, *rds.DescribeOptionGroupOptionsInput, ...func(*rds.Options)) (*rds.DescribeOptionGroupOptionsOutput, error)
	DescribeOptionGroups(context.Context, *rds.DescribeOptionGroupsInput, ...func(*rds.Options)) (*rds.DescribeOptionGroupsOutput, error)
	DescribeOrderableDBInstanceOptions(context.Context, *rds.DescribeOrderableDBInstanceOptionsInput, ...func(*rds.Options)) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribePendingMaintenanceActions(context.Context, *rds.DescribePendingMaintenanceActionsInput, ...func(*rds.Options)) (*rds.DescribePendingMaintenanceActionsOutput, error)
	DescribeReservedDBInstances(context.Context, *rds.DescribeReservedDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOutput, error)
	DescribeReservedDBInstancesOfferings(context.Context, *rds.DescribeReservedDBInstancesOfferingsInput, ...func(*rds.Options)) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)
	DescribeSourceRegions(context.Context, *rds.DescribeSourceRegionsInput, ...func(*rds.Options)) (*rds.DescribeSourceRegionsOutput, error)
	DescribeValidDBInstanceModifications(context.Context, *rds.DescribeValidDBInstanceModificationsInput, ...func(*rds.Options)) (*rds.DescribeValidDBInstanceModificationsOutput, error)
	ListTagsForResource(context.Context, *rds.ListTagsForResourceInput, ...func(*rds.Options)) (*rds.ListTagsForResourceOutput, error)
}

type RedshiftClient interface {
	DescribeAccountAttributes(context.Context, *redshift.DescribeAccountAttributesInput, ...func(*redshift.Options)) (*redshift.DescribeAccountAttributesOutput, error)
	DescribeAuthenticationProfiles(context.Context, *redshift.DescribeAuthenticationProfilesInput, ...func(*redshift.Options)) (*redshift.DescribeAuthenticationProfilesOutput, error)
	DescribeClusterDbRevisions(context.Context, *redshift.DescribeClusterDbRevisionsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterDbRevisionsOutput, error)
	DescribeClusterParameterGroups(context.Context, *redshift.DescribeClusterParameterGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterParameterGroupsOutput, error)
	DescribeClusterParameters(context.Context, *redshift.DescribeClusterParametersInput, ...func(*redshift.Options)) (*redshift.DescribeClusterParametersOutput, error)
	DescribeClusterSecurityGroups(context.Context, *redshift.DescribeClusterSecurityGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSecurityGroupsOutput, error)
	DescribeClusterSnapshots(context.Context, *redshift.DescribeClusterSnapshotsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSnapshotsOutput, error)
	DescribeClusterSubnetGroups(context.Context, *redshift.DescribeClusterSubnetGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error)
	DescribeClusterTracks(context.Context, *redshift.DescribeClusterTracksInput, ...func(*redshift.Options)) (*redshift.DescribeClusterTracksOutput, error)
	DescribeClusterVersions(context.Context, *redshift.DescribeClusterVersionsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterVersionsOutput, error)
	DescribeClusters(context.Context, *redshift.DescribeClustersInput, ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error)
	DescribeDataShares(context.Context, *redshift.DescribeDataSharesInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesOutput, error)
	DescribeDataSharesForConsumer(context.Context, *redshift.DescribeDataSharesForConsumerInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesForConsumerOutput, error)
	DescribeDataSharesForProducer(context.Context, *redshift.DescribeDataSharesForProducerInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesForProducerOutput, error)
	DescribeDefaultClusterParameters(context.Context, *redshift.DescribeDefaultClusterParametersInput, ...func(*redshift.Options)) (*redshift.DescribeDefaultClusterParametersOutput, error)
	DescribeEndpointAccess(context.Context, *redshift.DescribeEndpointAccessInput, ...func(*redshift.Options)) (*redshift.DescribeEndpointAccessOutput, error)
	DescribeEndpointAuthorization(context.Context, *redshift.DescribeEndpointAuthorizationInput, ...func(*redshift.Options)) (*redshift.DescribeEndpointAuthorizationOutput, error)
	DescribeEventCategories(context.Context, *redshift.DescribeEventCategoriesInput, ...func(*redshift.Options)) (*redshift.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *redshift.DescribeEventSubscriptionsInput, ...func(*redshift.Options)) (*redshift.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *redshift.DescribeEventsInput, ...func(*redshift.Options)) (*redshift.DescribeEventsOutput, error)
	DescribeHsmClientCertificates(context.Context, *redshift.DescribeHsmClientCertificatesInput, ...func(*redshift.Options)) (*redshift.DescribeHsmClientCertificatesOutput, error)
	DescribeHsmConfigurations(context.Context, *redshift.DescribeHsmConfigurationsInput, ...func(*redshift.Options)) (*redshift.DescribeHsmConfigurationsOutput, error)
	DescribeLoggingStatus(context.Context, *redshift.DescribeLoggingStatusInput, ...func(*redshift.Options)) (*redshift.DescribeLoggingStatusOutput, error)
	DescribeNodeConfigurationOptions(context.Context, *redshift.DescribeNodeConfigurationOptionsInput, ...func(*redshift.Options)) (*redshift.DescribeNodeConfigurationOptionsOutput, error)
	DescribeOrderableClusterOptions(context.Context, *redshift.DescribeOrderableClusterOptionsInput, ...func(*redshift.Options)) (*redshift.DescribeOrderableClusterOptionsOutput, error)
	DescribePartners(context.Context, *redshift.DescribePartnersInput, ...func(*redshift.Options)) (*redshift.DescribePartnersOutput, error)
	DescribeReservedNodeExchangeStatus(context.Context, *redshift.DescribeReservedNodeExchangeStatusInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodeExchangeStatusOutput, error)
	DescribeReservedNodeOfferings(context.Context, *redshift.DescribeReservedNodeOfferingsInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodeOfferingsOutput, error)
	DescribeReservedNodes(context.Context, *redshift.DescribeReservedNodesInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodesOutput, error)
	DescribeResize(context.Context, *redshift.DescribeResizeInput, ...func(*redshift.Options)) (*redshift.DescribeResizeOutput, error)
	DescribeScheduledActions(context.Context, *redshift.DescribeScheduledActionsInput, ...func(*redshift.Options)) (*redshift.DescribeScheduledActionsOutput, error)
	DescribeSnapshotCopyGrants(context.Context, *redshift.DescribeSnapshotCopyGrantsInput, ...func(*redshift.Options)) (*redshift.DescribeSnapshotCopyGrantsOutput, error)
	DescribeSnapshotSchedules(context.Context, *redshift.DescribeSnapshotSchedulesInput, ...func(*redshift.Options)) (*redshift.DescribeSnapshotSchedulesOutput, error)
	DescribeStorage(context.Context, *redshift.DescribeStorageInput, ...func(*redshift.Options)) (*redshift.DescribeStorageOutput, error)
	DescribeTableRestoreStatus(context.Context, *redshift.DescribeTableRestoreStatusInput, ...func(*redshift.Options)) (*redshift.DescribeTableRestoreStatusOutput, error)
	DescribeTags(context.Context, *redshift.DescribeTagsInput, ...func(*redshift.Options)) (*redshift.DescribeTagsOutput, error)
	DescribeUsageLimits(context.Context, *redshift.DescribeUsageLimitsInput, ...func(*redshift.Options)) (*redshift.DescribeUsageLimitsOutput, error)
	GetClusterCredentials(context.Context, *redshift.GetClusterCredentialsInput, ...func(*redshift.Options)) (*redshift.GetClusterCredentialsOutput, error)
	GetClusterCredentialsWithIAM(context.Context, *redshift.GetClusterCredentialsWithIAMInput, ...func(*redshift.Options)) (*redshift.GetClusterCredentialsWithIAMOutput, error)
	GetReservedNodeExchangeConfigurationOptions(context.Context, *redshift.GetReservedNodeExchangeConfigurationOptionsInput, ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, error)
	GetReservedNodeExchangeOfferings(context.Context, *redshift.GetReservedNodeExchangeOfferingsInput, ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeOfferingsOutput, error)
}

type ResourcegroupsClient interface {
	GetGroup(context.Context, *resourcegroups.GetGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error)
	GetGroupConfiguration(context.Context, *resourcegroups.GetGroupConfigurationInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error)
	GetGroupQuery(context.Context, *resourcegroups.GetGroupQueryInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error)
	GetTags(context.Context, *resourcegroups.GetTagsInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error)
	ListGroupResources(context.Context, *resourcegroups.ListGroupResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error)
	ListGroups(context.Context, *resourcegroups.ListGroupsInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error)
	SearchResources(context.Context, *resourcegroups.SearchResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error)
}

type Route53Client interface {
	GetAccountLimit(context.Context, *route53.GetAccountLimitInput, ...func(*route53.Options)) (*route53.GetAccountLimitOutput, error)
	GetChange(context.Context, *route53.GetChangeInput, ...func(*route53.Options)) (*route53.GetChangeOutput, error)
	GetCheckerIpRanges(context.Context, *route53.GetCheckerIpRangesInput, ...func(*route53.Options)) (*route53.GetCheckerIpRangesOutput, error)
	GetDNSSEC(context.Context, *route53.GetDNSSECInput, ...func(*route53.Options)) (*route53.GetDNSSECOutput, error)
	GetGeoLocation(context.Context, *route53.GetGeoLocationInput, ...func(*route53.Options)) (*route53.GetGeoLocationOutput, error)
	GetHealthCheck(context.Context, *route53.GetHealthCheckInput, ...func(*route53.Options)) (*route53.GetHealthCheckOutput, error)
	GetHealthCheckCount(context.Context, *route53.GetHealthCheckCountInput, ...func(*route53.Options)) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckLastFailureReason(context.Context, *route53.GetHealthCheckLastFailureReasonInput, ...func(*route53.Options)) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckStatus(context.Context, *route53.GetHealthCheckStatusInput, ...func(*route53.Options)) (*route53.GetHealthCheckStatusOutput, error)
	GetHostedZone(context.Context, *route53.GetHostedZoneInput, ...func(*route53.Options)) (*route53.GetHostedZoneOutput, error)
	GetHostedZoneCount(context.Context, *route53.GetHostedZoneCountInput, ...func(*route53.Options)) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneLimit(context.Context, *route53.GetHostedZoneLimitInput, ...func(*route53.Options)) (*route53.GetHostedZoneLimitOutput, error)
	GetQueryLoggingConfig(context.Context, *route53.GetQueryLoggingConfigInput, ...func(*route53.Options)) (*route53.GetQueryLoggingConfigOutput, error)
	GetReusableDelegationSet(context.Context, *route53.GetReusableDelegationSetInput, ...func(*route53.Options)) (*route53.GetReusableDelegationSetOutput, error)
	GetReusableDelegationSetLimit(context.Context, *route53.GetReusableDelegationSetLimitInput, ...func(*route53.Options)) (*route53.GetReusableDelegationSetLimitOutput, error)
	GetTrafficPolicy(context.Context, *route53.GetTrafficPolicyInput, ...func(*route53.Options)) (*route53.GetTrafficPolicyOutput, error)
	GetTrafficPolicyInstance(context.Context, *route53.GetTrafficPolicyInstanceInput, ...func(*route53.Options)) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyInstanceCount(context.Context, *route53.GetTrafficPolicyInstanceCountInput, ...func(*route53.Options)) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	ListCidrBlocks(context.Context, *route53.ListCidrBlocksInput, ...func(*route53.Options)) (*route53.ListCidrBlocksOutput, error)
	ListCidrCollections(context.Context, *route53.ListCidrCollectionsInput, ...func(*route53.Options)) (*route53.ListCidrCollectionsOutput, error)
	ListCidrLocations(context.Context, *route53.ListCidrLocationsInput, ...func(*route53.Options)) (*route53.ListCidrLocationsOutput, error)
	ListGeoLocations(context.Context, *route53.ListGeoLocationsInput, ...func(*route53.Options)) (*route53.ListGeoLocationsOutput, error)
	ListHealthChecks(context.Context, *route53.ListHealthChecksInput, ...func(*route53.Options)) (*route53.ListHealthChecksOutput, error)
	ListHostedZones(context.Context, *route53.ListHostedZonesInput, ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error)
	ListHostedZonesByName(context.Context, *route53.ListHostedZonesByNameInput, ...func(*route53.Options)) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesByVPC(context.Context, *route53.ListHostedZonesByVPCInput, ...func(*route53.Options)) (*route53.ListHostedZonesByVPCOutput, error)
	ListQueryLoggingConfigs(context.Context, *route53.ListQueryLoggingConfigsInput, ...func(*route53.Options)) (*route53.ListQueryLoggingConfigsOutput, error)
	ListResourceRecordSets(context.Context, *route53.ListResourceRecordSetsInput, ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
	ListReusableDelegationSets(context.Context, *route53.ListReusableDelegationSetsInput, ...func(*route53.Options)) (*route53.ListReusableDelegationSetsOutput, error)
	ListTagsForResource(context.Context, *route53.ListTagsForResourceInput, ...func(*route53.Options)) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResources(context.Context, *route53.ListTagsForResourcesInput, ...func(*route53.Options)) (*route53.ListTagsForResourcesOutput, error)
	ListTrafficPolicies(context.Context, *route53.ListTrafficPoliciesInput, ...func(*route53.Options)) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPolicyInstances(context.Context, *route53.ListTrafficPolicyInstancesInput, ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyInstancesByHostedZone(context.Context, *route53.ListTrafficPolicyInstancesByHostedZoneInput, ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByPolicy(context.Context, *route53.ListTrafficPolicyInstancesByPolicyInput, ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyVersions(context.Context, *route53.ListTrafficPolicyVersionsInput, ...func(*route53.Options)) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListVPCAssociationAuthorizations(context.Context, *route53.ListVPCAssociationAuthorizationsInput, ...func(*route53.Options)) (*route53.ListVPCAssociationAuthorizationsOutput, error)
}

type Route53domainsClient interface {
	GetContactReachabilityStatus(context.Context, *route53domains.GetContactReachabilityStatusInput, ...func(*route53domains.Options)) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetDomainDetail(context.Context, *route53domains.GetDomainDetailInput, ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error)
	GetDomainSuggestions(context.Context, *route53domains.GetDomainSuggestionsInput, ...func(*route53domains.Options)) (*route53domains.GetDomainSuggestionsOutput, error)
	GetOperationDetail(context.Context, *route53domains.GetOperationDetailInput, ...func(*route53domains.Options)) (*route53domains.GetOperationDetailOutput, error)
	ListDomains(context.Context, *route53domains.ListDomainsInput, ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error)
	ListOperations(context.Context, *route53domains.ListOperationsInput, ...func(*route53domains.Options)) (*route53domains.ListOperationsOutput, error)
	ListPrices(context.Context, *route53domains.ListPricesInput, ...func(*route53domains.Options)) (*route53domains.ListPricesOutput, error)
	ListTagsForDomain(context.Context, *route53domains.ListTagsForDomainInput, ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error)
}

type S3Client interface {
	GetBucketAccelerateConfiguration(context.Context, *s3.GetBucketAccelerateConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAcl(context.Context, *s3.GetBucketAclInput, ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketAnalyticsConfiguration(context.Context, *s3.GetBucketAnalyticsConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketCors(context.Context, *s3.GetBucketCorsInput, ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketEncryption(context.Context, *s3.GetBucketEncryptionInput, ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetBucketIntelligentTieringConfiguration(context.Context, *s3.GetBucketIntelligentTieringConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error)
	GetBucketInventoryConfiguration(context.Context, *s3.GetBucketInventoryConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketLifecycleConfiguration(context.Context, *s3.GetBucketLifecycleConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLocation(context.Context, *s3.GetBucketLocationInput, ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketLogging(context.Context, *s3.GetBucketLoggingInput, ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
	GetBucketMetricsConfiguration(context.Context, *s3.GetBucketMetricsConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketNotificationConfiguration(context.Context, *s3.GetBucketNotificationConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error)
	GetBucketOwnershipControls(context.Context, *s3.GetBucketOwnershipControlsInput, ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error)
	GetBucketPolicy(context.Context, *s3.GetBucketPolicyInput, ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyStatus(context.Context, *s3.GetBucketPolicyStatusInput, ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error)
	GetBucketReplication(context.Context, *s3.GetBucketReplicationInput, ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error)
	GetBucketRequestPayment(context.Context, *s3.GetBucketRequestPaymentInput, ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketTagging(context.Context, *s3.GetBucketTaggingInput, ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error)
	GetBucketVersioning(context.Context, *s3.GetBucketVersioningInput, ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetBucketWebsite(context.Context, *s3.GetBucketWebsiteInput, ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error)
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectAcl(context.Context, *s3.GetObjectAclInput, ...func(*s3.Options)) (*s3.GetObjectAclOutput, error)
	GetObjectAttributes(context.Context, *s3.GetObjectAttributesInput, ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
	GetObjectLegalHold(context.Context, *s3.GetObjectLegalHoldInput, ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error)
	GetObjectLockConfiguration(context.Context, *s3.GetObjectLockConfigurationInput, ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error)
	GetObjectRetention(context.Context, *s3.GetObjectRetentionInput, ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error)
	GetObjectTagging(context.Context, *s3.GetObjectTaggingInput, ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error)
	GetObjectTorrent(context.Context, *s3.GetObjectTorrentInput, ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error)
	GetPublicAccessBlock(context.Context, *s3.GetPublicAccessBlockInput, ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error)
	ListBucketAnalyticsConfigurations(context.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketIntelligentTieringConfigurations(context.Context, *s3.ListBucketIntelligentTieringConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error)
	ListBucketInventoryConfigurations(context.Context, *s3.ListBucketInventoryConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketMetricsConfigurations(context.Context, *s3.ListBucketMetricsConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBuckets(context.Context, *s3.ListBucketsInput, ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	ListMultipartUploads(context.Context, *s3.ListMultipartUploadsInput, ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error)
	ListObjectVersions(context.Context, *s3.ListObjectVersionsInput, ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error)
	ListObjects(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) (*s3.ListObjectsOutput, error)
	ListObjectsV2(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	ListParts(context.Context, *s3.ListPartsInput, ...func(*s3.Options)) (*s3.ListPartsOutput, error)
}

type S3controlClient interface {
	DescribeJob(context.Context, *s3control.DescribeJobInput, ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error)
	DescribeMultiRegionAccessPointOperation(context.Context, *s3control.DescribeMultiRegionAccessPointOperationInput, ...func(*s3control.Options)) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error)
	GetAccessPoint(context.Context, *s3control.GetAccessPointInput, ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error)
	GetAccessPointConfigurationForObjectLambda(context.Context, *s3control.GetAccessPointConfigurationForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	GetAccessPointForObjectLambda(context.Context, *s3control.GetAccessPointForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	GetAccessPointPolicy(context.Context, *s3control.GetAccessPointPolicyInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyForObjectLambda(context.Context, *s3control.GetAccessPointPolicyForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	GetAccessPointPolicyStatus(context.Context, *s3control.GetAccessPointPolicyStatusInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusForObjectLambda(context.Context, *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	GetBucket(context.Context, *s3control.GetBucketInput, ...func(*s3control.Options)) (*s3control.GetBucketOutput, error)
	GetBucketLifecycleConfiguration(context.Context, *s3control.GetBucketLifecycleConfigurationInput, ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketPolicy(context.Context, *s3control.GetBucketPolicyInput, ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error)
	GetBucketTagging(context.Context, *s3control.GetBucketTaggingInput, ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error)
	GetBucketVersioning(context.Context, *s3control.GetBucketVersioningInput, ...func(*s3control.Options)) (*s3control.GetBucketVersioningOutput, error)
	GetJobTagging(context.Context, *s3control.GetJobTaggingInput, ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error)
	GetMultiRegionAccessPoint(context.Context, *s3control.GetMultiRegionAccessPointInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointOutput, error)
	GetMultiRegionAccessPointPolicy(context.Context, *s3control.GetMultiRegionAccessPointPolicyInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyOutput, error)
	GetMultiRegionAccessPointPolicyStatus(context.Context, *s3control.GetMultiRegionAccessPointPolicyStatusInput, ...func(*s3control.Options)) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error)
	GetPublicAccessBlock(context.Context, *s3control.GetPublicAccessBlockInput, ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
	GetStorageLensConfiguration(context.Context, *s3control.GetStorageLensConfigurationInput, ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationTagging(context.Context, *s3control.GetStorageLensConfigurationTaggingInput, ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	ListAccessPoints(context.Context, *s3control.ListAccessPointsInput, ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsForObjectLambda(context.Context, *s3control.ListAccessPointsForObjectLambdaInput, ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	ListJobs(context.Context, *s3control.ListJobsInput, ...func(*s3control.Options)) (*s3control.ListJobsOutput, error)
	ListMultiRegionAccessPoints(context.Context, *s3control.ListMultiRegionAccessPointsInput, ...func(*s3control.Options)) (*s3control.ListMultiRegionAccessPointsOutput, error)
	ListRegionalBuckets(context.Context, *s3control.ListRegionalBucketsInput, ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error)
	ListStorageLensConfigurations(context.Context, *s3control.ListStorageLensConfigurationsInput, ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error)
}

type S3managerClient interface {
	GetBucketRegion(context.Context, string, ...func(*s3.Options)) (string, error)
}

type SagemakerClient interface {
	DescribeAction(context.Context, *sagemaker.DescribeActionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeActionOutput, error)
	DescribeAlgorithm(context.Context, *sagemaker.DescribeAlgorithmInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAlgorithmOutput, error)
	DescribeApp(context.Context, *sagemaker.DescribeAppInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAppOutput, error)
	DescribeAppImageConfig(context.Context, *sagemaker.DescribeAppImageConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAppImageConfigOutput, error)
	DescribeArtifact(context.Context, *sagemaker.DescribeArtifactInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeArtifactOutput, error)
	DescribeAutoMLJob(context.Context, *sagemaker.DescribeAutoMLJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAutoMLJobOutput, error)
	DescribeCodeRepository(context.Context, *sagemaker.DescribeCodeRepositoryInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeCodeRepositoryOutput, error)
	DescribeCompilationJob(context.Context, *sagemaker.DescribeCompilationJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeCompilationJobOutput, error)
	DescribeContext(context.Context, *sagemaker.DescribeContextInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeContextOutput, error)
	DescribeDataQualityJobDefinition(context.Context, *sagemaker.DescribeDataQualityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error)
	DescribeDevice(context.Context, *sagemaker.DescribeDeviceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceOutput, error)
	DescribeDeviceFleet(context.Context, *sagemaker.DescribeDeviceFleetInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceFleetOutput, error)
	DescribeDomain(context.Context, *sagemaker.DescribeDomainInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDomainOutput, error)
	DescribeEdgeDeploymentPlan(context.Context, *sagemaker.DescribeEdgeDeploymentPlanInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgeDeploymentPlanOutput, error)
	DescribeEdgePackagingJob(context.Context, *sagemaker.DescribeEdgePackagingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgePackagingJobOutput, error)
	DescribeEndpoint(context.Context, *sagemaker.DescribeEndpointInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointOutput, error)
	DescribeEndpointConfig(context.Context, *sagemaker.DescribeEndpointConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error)
	DescribeExperiment(context.Context, *sagemaker.DescribeExperimentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeExperimentOutput, error)
	DescribeFeatureGroup(context.Context, *sagemaker.DescribeFeatureGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureGroupOutput, error)
	DescribeFeatureMetadata(context.Context, *sagemaker.DescribeFeatureMetadataInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureMetadataOutput, error)
	DescribeFlowDefinition(context.Context, *sagemaker.DescribeFlowDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFlowDefinitionOutput, error)
	DescribeHumanTaskUi(context.Context, *sagemaker.DescribeHumanTaskUiInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHumanTaskUiOutput, error)
	DescribeHyperParameterTuningJob(context.Context, *sagemaker.DescribeHyperParameterTuningJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHyperParameterTuningJobOutput, error)
	DescribeImage(context.Context, *sagemaker.DescribeImageInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeImageOutput, error)
	DescribeImageVersion(context.Context, *sagemaker.DescribeImageVersionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeImageVersionOutput, error)
	DescribeInferenceRecommendationsJob(context.Context, *sagemaker.DescribeInferenceRecommendationsJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeInferenceRecommendationsJobOutput, error)
	DescribeLabelingJob(context.Context, *sagemaker.DescribeLabelingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeLabelingJobOutput, error)
	DescribeLineageGroup(context.Context, *sagemaker.DescribeLineageGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeLineageGroupOutput, error)
	DescribeModel(context.Context, *sagemaker.DescribeModelInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error)
	DescribeModelBiasJobDefinition(context.Context, *sagemaker.DescribeModelBiasJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelBiasJobDefinitionOutput, error)
	DescribeModelExplainabilityJobDefinition(context.Context, *sagemaker.DescribeModelExplainabilityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelExplainabilityJobDefinitionOutput, error)
	DescribeModelPackage(context.Context, *sagemaker.DescribeModelPackageInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageOutput, error)
	DescribeModelPackageGroup(context.Context, *sagemaker.DescribeModelPackageGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageGroupOutput, error)
	DescribeModelQualityJobDefinition(context.Context, *sagemaker.DescribeModelQualityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelQualityJobDefinitionOutput, error)
	DescribeMonitoringSchedule(context.Context, *sagemaker.DescribeMonitoringScheduleInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeMonitoringScheduleOutput, error)
	DescribeNotebookInstance(context.Context, *sagemaker.DescribeNotebookInstanceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error)
	DescribeNotebookInstanceLifecycleConfig(context.Context, *sagemaker.DescribeNotebookInstanceLifecycleConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error)
	DescribePipeline(context.Context, *sagemaker.DescribePipelineInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineOutput, error)
	DescribePipelineDefinitionForExecution(context.Context, *sagemaker.DescribePipelineDefinitionForExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineDefinitionForExecutionOutput, error)
	DescribePipelineExecution(context.Context, *sagemaker.DescribePipelineExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineExecutionOutput, error)
	DescribeProcessingJob(context.Context, *sagemaker.DescribeProcessingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeProcessingJobOutput, error)
	DescribeProject(context.Context, *sagemaker.DescribeProjectInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeProjectOutput, error)
	DescribeStudioLifecycleConfig(context.Context, *sagemaker.DescribeStudioLifecycleConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeStudioLifecycleConfigOutput, error)
	DescribeSubscribedWorkteam(context.Context, *sagemaker.DescribeSubscribedWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeSubscribedWorkteamOutput, error)
	DescribeTrainingJob(context.Context, *sagemaker.DescribeTrainingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error)
	DescribeTransformJob(context.Context, *sagemaker.DescribeTransformJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTransformJobOutput, error)
	DescribeTrial(context.Context, *sagemaker.DescribeTrialInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialOutput, error)
	DescribeTrialComponent(context.Context, *sagemaker.DescribeTrialComponentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialComponentOutput, error)
	DescribeUserProfile(context.Context, *sagemaker.DescribeUserProfileInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeUserProfileOutput, error)
	DescribeWorkforce(context.Context, *sagemaker.DescribeWorkforceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkforceOutput, error)
	DescribeWorkteam(context.Context, *sagemaker.DescribeWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkteamOutput, error)
	GetDeviceFleetReport(context.Context, *sagemaker.GetDeviceFleetReportInput, ...func(*sagemaker.Options)) (*sagemaker.GetDeviceFleetReportOutput, error)
	GetLineageGroupPolicy(context.Context, *sagemaker.GetLineageGroupPolicyInput, ...func(*sagemaker.Options)) (*sagemaker.GetLineageGroupPolicyOutput, error)
	GetModelPackageGroupPolicy(context.Context, *sagemaker.GetModelPackageGroupPolicyInput, ...func(*sagemaker.Options)) (*sagemaker.GetModelPackageGroupPolicyOutput, error)
	GetSagemakerServicecatalogPortfolioStatus(context.Context, *sagemaker.GetSagemakerServicecatalogPortfolioStatusInput, ...func(*sagemaker.Options)) (*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput, error)
	GetSearchSuggestions(context.Context, *sagemaker.GetSearchSuggestionsInput, ...func(*sagemaker.Options)) (*sagemaker.GetSearchSuggestionsOutput, error)
	ListActions(context.Context, *sagemaker.ListActionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListActionsOutput, error)
	ListAlgorithms(context.Context, *sagemaker.ListAlgorithmsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAlgorithmsOutput, error)
	ListAppImageConfigs(context.Context, *sagemaker.ListAppImageConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAppImageConfigsOutput, error)
	ListApps(context.Context, *sagemaker.ListAppsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAppsOutput, error)
	ListArtifacts(context.Context, *sagemaker.ListArtifactsInput, ...func(*sagemaker.Options)) (*sagemaker.ListArtifactsOutput, error)
	ListAssociations(context.Context, *sagemaker.ListAssociationsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAssociationsOutput, error)
	ListAutoMLJobs(context.Context, *sagemaker.ListAutoMLJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAutoMLJobsOutput, error)
	ListCandidatesForAutoMLJob(context.Context, *sagemaker.ListCandidatesForAutoMLJobInput, ...func(*sagemaker.Options)) (*sagemaker.ListCandidatesForAutoMLJobOutput, error)
	ListCodeRepositories(context.Context, *sagemaker.ListCodeRepositoriesInput, ...func(*sagemaker.Options)) (*sagemaker.ListCodeRepositoriesOutput, error)
	ListCompilationJobs(context.Context, *sagemaker.ListCompilationJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListCompilationJobsOutput, error)
	ListContexts(context.Context, *sagemaker.ListContextsInput, ...func(*sagemaker.Options)) (*sagemaker.ListContextsOutput, error)
	ListDataQualityJobDefinitions(context.Context, *sagemaker.ListDataQualityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDataQualityJobDefinitionsOutput, error)
	ListDeviceFleets(context.Context, *sagemaker.ListDeviceFleetsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDeviceFleetsOutput, error)
	ListDevices(context.Context, *sagemaker.ListDevicesInput, ...func(*sagemaker.Options)) (*sagemaker.ListDevicesOutput, error)
	ListDomains(context.Context, *sagemaker.ListDomainsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDomainsOutput, error)
	ListEdgeDeploymentPlans(context.Context, *sagemaker.ListEdgeDeploymentPlansInput, ...func(*sagemaker.Options)) (*sagemaker.ListEdgeDeploymentPlansOutput, error)
	ListEdgePackagingJobs(context.Context, *sagemaker.ListEdgePackagingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEdgePackagingJobsOutput, error)
	ListEndpointConfigs(context.Context, *sagemaker.ListEndpointConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error)
	ListEndpoints(context.Context, *sagemaker.ListEndpointsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEndpointsOutput, error)
	ListExperiments(context.Context, *sagemaker.ListExperimentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListExperimentsOutput, error)
	ListFeatureGroups(context.Context, *sagemaker.ListFeatureGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListFeatureGroupsOutput, error)
	ListFlowDefinitions(context.Context, *sagemaker.ListFlowDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListFlowDefinitionsOutput, error)
	ListHumanTaskUis(context.Context, *sagemaker.ListHumanTaskUisInput, ...func(*sagemaker.Options)) (*sagemaker.ListHumanTaskUisOutput, error)
	ListHyperParameterTuningJobs(context.Context, *sagemaker.ListHyperParameterTuningJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListHyperParameterTuningJobsOutput, error)
	ListImageVersions(context.Context, *sagemaker.ListImageVersionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListImageVersionsOutput, error)
	ListImages(context.Context, *sagemaker.ListImagesInput, ...func(*sagemaker.Options)) (*sagemaker.ListImagesOutput, error)
	ListInferenceRecommendationsJobSteps(context.Context, *sagemaker.ListInferenceRecommendationsJobStepsInput, ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobStepsOutput, error)
	ListInferenceRecommendationsJobs(context.Context, *sagemaker.ListInferenceRecommendationsJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobsOutput, error)
	ListLabelingJobs(context.Context, *sagemaker.ListLabelingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsOutput, error)
	ListLabelingJobsForWorkteam(context.Context, *sagemaker.ListLabelingJobsForWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsForWorkteamOutput, error)
	ListLineageGroups(context.Context, *sagemaker.ListLineageGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListLineageGroupsOutput, error)
	ListModelBiasJobDefinitions(context.Context, *sagemaker.ListModelBiasJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelBiasJobDefinitionsOutput, error)
	ListModelExplainabilityJobDefinitions(context.Context, *sagemaker.ListModelExplainabilityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelExplainabilityJobDefinitionsOutput, error)
	ListModelMetadata(context.Context, *sagemaker.ListModelMetadataInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelMetadataOutput, error)
	ListModelPackageGroups(context.Context, *sagemaker.ListModelPackageGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelPackageGroupsOutput, error)
	ListModelPackages(context.Context, *sagemaker.ListModelPackagesInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelPackagesOutput, error)
	ListModelQualityJobDefinitions(context.Context, *sagemaker.ListModelQualityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelQualityJobDefinitionsOutput, error)
	ListModels(context.Context, *sagemaker.ListModelsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error)
	ListMonitoringExecutions(context.Context, *sagemaker.ListMonitoringExecutionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringExecutionsOutput, error)
	ListMonitoringSchedules(context.Context, *sagemaker.ListMonitoringSchedulesInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringSchedulesOutput, error)
	ListNotebookInstanceLifecycleConfigs(context.Context, *sagemaker.ListNotebookInstanceLifecycleConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error)
	ListNotebookInstances(context.Context, *sagemaker.ListNotebookInstancesInput, ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error)
	ListPipelineExecutionSteps(context.Context, *sagemaker.ListPipelineExecutionStepsInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionStepsOutput, error)
	ListPipelineExecutions(context.Context, *sagemaker.ListPipelineExecutionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionsOutput, error)
	ListPipelineParametersForExecution(context.Context, *sagemaker.ListPipelineParametersForExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineParametersForExecutionOutput, error)
	ListPipelines(context.Context, *sagemaker.ListPipelinesInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelinesOutput, error)
	ListProcessingJobs(context.Context, *sagemaker.ListProcessingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListProcessingJobsOutput, error)
	ListProjects(context.Context, *sagemaker.ListProjectsInput, ...func(*sagemaker.Options)) (*sagemaker.ListProjectsOutput, error)
	ListStageDevices(context.Context, *sagemaker.ListStageDevicesInput, ...func(*sagemaker.Options)) (*sagemaker.ListStageDevicesOutput, error)
	ListStudioLifecycleConfigs(context.Context, *sagemaker.ListStudioLifecycleConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListStudioLifecycleConfigsOutput, error)
	ListSubscribedWorkteams(context.Context, *sagemaker.ListSubscribedWorkteamsInput, ...func(*sagemaker.Options)) (*sagemaker.ListSubscribedWorkteamsOutput, error)
	ListTags(context.Context, *sagemaker.ListTagsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error)
	ListTrainingJobs(context.Context, *sagemaker.ListTrainingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error)
	ListTrainingJobsForHyperParameterTuningJob(context.Context, *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error)
	ListTransformJobs(context.Context, *sagemaker.ListTransformJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTransformJobsOutput, error)
	ListTrialComponents(context.Context, *sagemaker.ListTrialComponentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrialComponentsOutput, error)
	ListTrials(context.Context, *sagemaker.ListTrialsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrialsOutput, error)
	ListUserProfiles(context.Context, *sagemaker.ListUserProfilesInput, ...func(*sagemaker.Options)) (*sagemaker.ListUserProfilesOutput, error)
	ListWorkforces(context.Context, *sagemaker.ListWorkforcesInput, ...func(*sagemaker.Options)) (*sagemaker.ListWorkforcesOutput, error)
	ListWorkteams(context.Context, *sagemaker.ListWorkteamsInput, ...func(*sagemaker.Options)) (*sagemaker.ListWorkteamsOutput, error)
	Search(context.Context, *sagemaker.SearchInput, ...func(*sagemaker.Options)) (*sagemaker.SearchOutput, error)
}

type SecretsmanagerClient interface {
	DescribeSecret(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error)
	GetRandomPassword(context.Context, *secretsmanager.GetRandomPasswordInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error)
	GetResourcePolicy(context.Context, *secretsmanager.GetResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error)
	GetSecretValue(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
	ListSecretVersionIds(context.Context, *secretsmanager.ListSecretVersionIdsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecrets(context.Context, *secretsmanager.ListSecretsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)
}

type ServicecatalogClient interface {
	DescribeConstraint(context.Context, *servicecatalog.DescribeConstraintInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeConstraintOutput, error)
	DescribeCopyProductStatus(context.Context, *servicecatalog.DescribeCopyProductStatusInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeCopyProductStatusOutput, error)
	DescribePortfolio(context.Context, *servicecatalog.DescribePortfolioInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioOutput, error)
	DescribePortfolioShareStatus(context.Context, *servicecatalog.DescribePortfolioShareStatusInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioShareStatusOutput, error)
	DescribePortfolioShares(context.Context, *servicecatalog.DescribePortfolioSharesInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioSharesOutput, error)
	DescribeProduct(context.Context, *servicecatalog.DescribeProductInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductOutput, error)
	DescribeProductAsAdmin(context.Context, *servicecatalog.DescribeProductAsAdminInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductAsAdminOutput, error)
	DescribeProductView(context.Context, *servicecatalog.DescribeProductViewInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductViewOutput, error)
	DescribeProvisionedProduct(context.Context, *servicecatalog.DescribeProvisionedProductInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisionedProductOutput, error)
	DescribeProvisionedProductPlan(context.Context, *servicecatalog.DescribeProvisionedProductPlanInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisionedProductPlanOutput, error)
	DescribeProvisioningArtifact(context.Context, *servicecatalog.DescribeProvisioningArtifactInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
	DescribeProvisioningParameters(context.Context, *servicecatalog.DescribeProvisioningParametersInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisioningParametersOutput, error)
	DescribeRecord(context.Context, *servicecatalog.DescribeRecordInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeRecordOutput, error)
	DescribeServiceAction(context.Context, *servicecatalog.DescribeServiceActionInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeServiceActionOutput, error)
	DescribeServiceActionExecutionParameters(context.Context, *servicecatalog.DescribeServiceActionExecutionParametersInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error)
	DescribeTagOption(context.Context, *servicecatalog.DescribeTagOptionInput, ...func(*servicecatalog.Options)) (*servicecatalog.DescribeTagOptionOutput, error)
	GetAWSOrganizationsAccessStatus(context.Context, *servicecatalog.GetAWSOrganizationsAccessStatusInput, ...func(*servicecatalog.Options)) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error)
	GetProvisionedProductOutputs(context.Context, *servicecatalog.GetProvisionedProductOutputsInput, ...func(*servicecatalog.Options)) (*servicecatalog.GetProvisionedProductOutputsOutput, error)
	ListAcceptedPortfolioShares(context.Context, *servicecatalog.ListAcceptedPortfolioSharesInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
	ListBudgetsForResource(context.Context, *servicecatalog.ListBudgetsForResourceInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListBudgetsForResourceOutput, error)
	ListConstraintsForPortfolio(context.Context, *servicecatalog.ListConstraintsForPortfolioInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
	ListLaunchPaths(context.Context, *servicecatalog.ListLaunchPathsInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListLaunchPathsOutput, error)
	ListOrganizationPortfolioAccess(context.Context, *servicecatalog.ListOrganizationPortfolioAccessInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error)
	ListPortfolioAccess(context.Context, *servicecatalog.ListPortfolioAccessInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfolioAccessOutput, error)
	ListPortfolios(context.Context, *servicecatalog.ListPortfoliosInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfoliosOutput, error)
	ListPortfoliosForProduct(context.Context, *servicecatalog.ListPortfoliosForProductInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfoliosForProductOutput, error)
	ListPrincipalsForPortfolio(context.Context, *servicecatalog.ListPrincipalsForPortfolioInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
	ListProvisionedProductPlans(context.Context, *servicecatalog.ListProvisionedProductPlansInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisionedProductPlansOutput, error)
	ListProvisioningArtifacts(context.Context, *servicecatalog.ListProvisioningArtifactsInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisioningArtifactsOutput, error)
	ListProvisioningArtifactsForServiceAction(context.Context, *servicecatalog.ListProvisioningArtifactsForServiceActionInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error)
	ListRecordHistory(context.Context, *servicecatalog.ListRecordHistoryInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListRecordHistoryOutput, error)
	ListResourcesForTagOption(context.Context, *servicecatalog.ListResourcesForTagOptionInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListResourcesForTagOptionOutput, error)
	ListServiceActions(context.Context, *servicecatalog.ListServiceActionsInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListServiceActionsOutput, error)
	ListServiceActionsForProvisioningArtifact(context.Context, *servicecatalog.ListServiceActionsForProvisioningArtifactInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error)
	ListStackInstancesForProvisionedProduct(context.Context, *servicecatalog.ListStackInstancesForProvisionedProductInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error)
	ListTagOptions(context.Context, *servicecatalog.ListTagOptionsInput, ...func(*servicecatalog.Options)) (*servicecatalog.ListTagOptionsOutput, error)
	SearchProducts(context.Context, *servicecatalog.SearchProductsInput, ...func(*servicecatalog.Options)) (*servicecatalog.SearchProductsOutput, error)
	SearchProductsAsAdmin(context.Context, *servicecatalog.SearchProductsAsAdminInput, ...func(*servicecatalog.Options)) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProvisionedProducts(context.Context, *servicecatalog.SearchProvisionedProductsInput, ...func(*servicecatalog.Options)) (*servicecatalog.SearchProvisionedProductsOutput, error)
}

type ServicecatalogappregistryClient interface {
	GetApplication(context.Context, *servicecatalogappregistry.GetApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error)
	GetAssociatedResource(context.Context, *servicecatalogappregistry.GetAssociatedResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error)
	GetAttributeGroup(context.Context, *servicecatalogappregistry.GetAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error)
	ListApplications(context.Context, *servicecatalogappregistry.ListApplicationsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error)
	ListAssociatedAttributeGroups(context.Context, *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error)
	ListAssociatedResources(context.Context, *servicecatalogappregistry.ListAssociatedResourcesInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error)
	ListAttributeGroups(context.Context, *servicecatalogappregistry.ListAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error)
	ListAttributeGroupsForApplication(context.Context, *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error)
	ListTagsForResource(context.Context, *servicecatalogappregistry.ListTagsForResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error)
}

type ServicequotasClient interface {
	GetAWSDefaultServiceQuota(context.Context, *servicequotas.GetAWSDefaultServiceQuotaInput, ...func(*servicequotas.Options)) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
	GetAssociationForServiceQuotaTemplate(context.Context, *servicequotas.GetAssociationForServiceQuotaTemplateInput, ...func(*servicequotas.Options)) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
	GetRequestedServiceQuotaChange(context.Context, *servicequotas.GetRequestedServiceQuotaChangeInput, ...func(*servicequotas.Options)) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
	GetServiceQuota(context.Context, *servicequotas.GetServiceQuotaInput, ...func(*servicequotas.Options)) (*servicequotas.GetServiceQuotaOutput, error)
	GetServiceQuotaIncreaseRequestFromTemplate(context.Context, *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, ...func(*servicequotas.Options)) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
	ListAWSDefaultServiceQuotas(context.Context, *servicequotas.ListAWSDefaultServiceQuotasInput, ...func(*servicequotas.Options)) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
	ListRequestedServiceQuotaChangeHistory(context.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, ...func(*servicequotas.Options)) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
	ListRequestedServiceQuotaChangeHistoryByQuota(context.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, ...func(*servicequotas.Options)) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
	ListServiceQuotaIncreaseRequestsInTemplate(context.Context, *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
	ListServiceQuotas(context.Context, *servicequotas.ListServiceQuotasInput, ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error)
	ListServices(context.Context, *servicequotas.ListServicesInput, ...func(*servicequotas.Options)) (*servicequotas.ListServicesOutput, error)
	ListTagsForResource(context.Context, *servicequotas.ListTagsForResourceInput, ...func(*servicequotas.Options)) (*servicequotas.ListTagsForResourceOutput, error)
}

type Sesv2Client interface {
	BatchGetMetricData(context.Context, *sesv2.BatchGetMetricDataInput, ...func(*sesv2.Options)) (*sesv2.BatchGetMetricDataOutput, error)
	GetAccount(context.Context, *sesv2.GetAccountInput, ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error)
	GetBlacklistReports(context.Context, *sesv2.GetBlacklistReportsInput, ...func(*sesv2.Options)) (*sesv2.GetBlacklistReportsOutput, error)
	GetConfigurationSet(context.Context, *sesv2.GetConfigurationSetInput, ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetEventDestinations(context.Context, *sesv2.GetConfigurationSetEventDestinationsInput, ...func(*sesv2.Options)) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetContact(context.Context, *sesv2.GetContactInput, ...func(*sesv2.Options)) (*sesv2.GetContactOutput, error)
	GetContactList(context.Context, *sesv2.GetContactListInput, ...func(*sesv2.Options)) (*sesv2.GetContactListOutput, error)
	GetCustomVerificationEmailTemplate(context.Context, *sesv2.GetCustomVerificationEmailTemplateInput, ...func(*sesv2.Options)) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetDedicatedIp(context.Context, *sesv2.GetDedicatedIpInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpPool(context.Context, *sesv2.GetDedicatedIpPoolInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpPoolOutput, error)
	GetDedicatedIps(context.Context, *sesv2.GetDedicatedIpsInput, ...func(*sesv2.Options)) (*sesv2.GetDedicatedIpsOutput, error)
	GetDeliverabilityDashboardOptions(context.Context, *sesv2.GetDeliverabilityDashboardOptionsInput, ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityTestReport(context.Context, *sesv2.GetDeliverabilityTestReportInput, ...func(*sesv2.Options)) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDomainDeliverabilityCampaign(context.Context, *sesv2.GetDomainDeliverabilityCampaignInput, ...func(*sesv2.Options)) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainStatisticsReport(context.Context, *sesv2.GetDomainStatisticsReportInput, ...func(*sesv2.Options)) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetEmailIdentity(context.Context, *sesv2.GetEmailIdentityInput, ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityPolicies(context.Context, *sesv2.GetEmailIdentityPoliciesInput, ...func(*sesv2.Options)) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailTemplate(context.Context, *sesv2.GetEmailTemplateInput, ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error)
	GetImportJob(context.Context, *sesv2.GetImportJobInput, ...func(*sesv2.Options)) (*sesv2.GetImportJobOutput, error)
	GetSuppressedDestination(context.Context, *sesv2.GetSuppressedDestinationInput, ...func(*sesv2.Options)) (*sesv2.GetSuppressedDestinationOutput, error)
	ListConfigurationSets(context.Context, *sesv2.ListConfigurationSetsInput, ...func(*sesv2.Options)) (*sesv2.ListConfigurationSetsOutput, error)
	ListContactLists(context.Context, *sesv2.ListContactListsInput, ...func(*sesv2.Options)) (*sesv2.ListContactListsOutput, error)
	ListContacts(context.Context, *sesv2.ListContactsInput, ...func(*sesv2.Options)) (*sesv2.ListContactsOutput, error)
	ListCustomVerificationEmailTemplates(context.Context, *sesv2.ListCustomVerificationEmailTemplatesInput, ...func(*sesv2.Options)) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListDedicatedIpPools(context.Context, *sesv2.ListDedicatedIpPoolsInput, ...func(*sesv2.Options)) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDeliverabilityTestReports(context.Context, *sesv2.ListDeliverabilityTestReportsInput, ...func(*sesv2.Options)) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDomainDeliverabilityCampaigns(context.Context, *sesv2.ListDomainDeliverabilityCampaignsInput, ...func(*sesv2.Options)) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListEmailIdentities(context.Context, *sesv2.ListEmailIdentitiesInput, ...func(*sesv2.Options)) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailTemplates(context.Context, *sesv2.ListEmailTemplatesInput, ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error)
	ListImportJobs(context.Context, *sesv2.ListImportJobsInput, ...func(*sesv2.Options)) (*sesv2.ListImportJobsOutput, error)
	ListRecommendations(context.Context, *sesv2.ListRecommendationsInput, ...func(*sesv2.Options)) (*sesv2.ListRecommendationsOutput, error)
	ListSuppressedDestinations(context.Context, *sesv2.ListSuppressedDestinationsInput, ...func(*sesv2.Options)) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListTagsForResource(context.Context, *sesv2.ListTagsForResourceInput, ...func(*sesv2.Options)) (*sesv2.ListTagsForResourceOutput, error)
}

type ShieldClient interface {
	DescribeAttack(context.Context, *shield.DescribeAttackInput, ...func(*shield.Options)) (*shield.DescribeAttackOutput, error)
	DescribeAttackStatistics(context.Context, *shield.DescribeAttackStatisticsInput, ...func(*shield.Options)) (*shield.DescribeAttackStatisticsOutput, error)
	DescribeDRTAccess(context.Context, *shield.DescribeDRTAccessInput, ...func(*shield.Options)) (*shield.DescribeDRTAccessOutput, error)
	DescribeEmergencyContactSettings(context.Context, *shield.DescribeEmergencyContactSettingsInput, ...func(*shield.Options)) (*shield.DescribeEmergencyContactSettingsOutput, error)
	DescribeProtection(context.Context, *shield.DescribeProtectionInput, ...func(*shield.Options)) (*shield.DescribeProtectionOutput, error)
	DescribeProtectionGroup(context.Context, *shield.DescribeProtectionGroupInput, ...func(*shield.Options)) (*shield.DescribeProtectionGroupOutput, error)
	DescribeSubscription(context.Context, *shield.DescribeSubscriptionInput, ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error)
	GetSubscriptionState(context.Context, *shield.GetSubscriptionStateInput, ...func(*shield.Options)) (*shield.GetSubscriptionStateOutput, error)
	ListAttacks(context.Context, *shield.ListAttacksInput, ...func(*shield.Options)) (*shield.ListAttacksOutput, error)
	ListProtectionGroups(context.Context, *shield.ListProtectionGroupsInput, ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error)
	ListProtections(context.Context, *shield.ListProtectionsInput, ...func(*shield.Options)) (*shield.ListProtectionsOutput, error)
	ListResourcesInProtectionGroup(context.Context, *shield.ListResourcesInProtectionGroupInput, ...func(*shield.Options)) (*shield.ListResourcesInProtectionGroupOutput, error)
	ListTagsForResource(context.Context, *shield.ListTagsForResourceInput, ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error)
}

type SnsClient interface {
	GetDataProtectionPolicy(context.Context, *sns.GetDataProtectionPolicyInput, ...func(*sns.Options)) (*sns.GetDataProtectionPolicyOutput, error)
	GetEndpointAttributes(context.Context, *sns.GetEndpointAttributesInput, ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error)
	GetPlatformApplicationAttributes(context.Context, *sns.GetPlatformApplicationAttributesInput, ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetSMSAttributes(context.Context, *sns.GetSMSAttributesInput, ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error)
	GetSMSSandboxAccountStatus(context.Context, *sns.GetSMSSandboxAccountStatusInput, ...func(*sns.Options)) (*sns.GetSMSSandboxAccountStatusOutput, error)
	GetSubscriptionAttributes(context.Context, *sns.GetSubscriptionAttributesInput, ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributes(context.Context, *sns.GetTopicAttributesInput, ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListEndpointsByPlatformApplication(context.Context, *sns.ListEndpointsByPlatformApplicationInput, ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListOriginationNumbers(context.Context, *sns.ListOriginationNumbersInput, ...func(*sns.Options)) (*sns.ListOriginationNumbersOutput, error)
	ListPhoneNumbersOptedOut(context.Context, *sns.ListPhoneNumbersOptedOutInput, ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPlatformApplications(context.Context, *sns.ListPlatformApplicationsInput, ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error)
	ListSMSSandboxPhoneNumbers(context.Context, *sns.ListSMSSandboxPhoneNumbersInput, ...func(*sns.Options)) (*sns.ListSMSSandboxPhoneNumbersOutput, error)
	ListSubscriptions(context.Context, *sns.ListSubscriptionsInput, ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsByTopic(context.Context, *sns.ListSubscriptionsByTopicInput, ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error)
	ListTagsForResource(context.Context, *sns.ListTagsForResourceInput, ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	ListTopics(context.Context, *sns.ListTopicsInput, ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
}

type SqsClient interface {
	GetQueueAttributes(context.Context, *sqs.GetQueueAttributesInput, ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error)
	GetQueueUrl(context.Context, *sqs.GetQueueUrlInput, ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	ListDeadLetterSourceQueues(context.Context, *sqs.ListDeadLetterSourceQueuesInput, ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListQueueTags(context.Context, *sqs.ListQueueTagsInput, ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error)
	ListQueues(context.Context, *sqs.ListQueuesInput, ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
}

type SsmClient interface {
	DescribeActivations(context.Context, *ssm.DescribeActivationsInput, ...func(*ssm.Options)) (*ssm.DescribeActivationsOutput, error)
	DescribeAssociation(context.Context, *ssm.DescribeAssociationInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationOutput, error)
	DescribeAssociationExecutionTargets(context.Context, *ssm.DescribeAssociationExecutionTargetsInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionTargetsOutput, error)
	DescribeAssociationExecutions(context.Context, *ssm.DescribeAssociationExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionsOutput, error)
	DescribeAutomationExecutions(context.Context, *ssm.DescribeAutomationExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAutomationExecutionsOutput, error)
	DescribeAutomationStepExecutions(context.Context, *ssm.DescribeAutomationStepExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAutomationStepExecutionsOutput, error)
	DescribeAvailablePatches(context.Context, *ssm.DescribeAvailablePatchesInput, ...func(*ssm.Options)) (*ssm.DescribeAvailablePatchesOutput, error)
	DescribeDocument(context.Context, *ssm.DescribeDocumentInput, ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error)
	DescribeDocumentPermission(context.Context, *ssm.DescribeDocumentPermissionInput, ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error)
	DescribeEffectiveInstanceAssociations(context.Context, *ssm.DescribeEffectiveInstanceAssociationsInput, ...func(*ssm.Options)) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error)
	DescribeEffectivePatchesForPatchBaseline(context.Context, *ssm.DescribeEffectivePatchesForPatchBaselineInput, ...func(*ssm.Options)) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error)
	DescribeInstanceAssociationsStatus(context.Context, *ssm.DescribeInstanceAssociationsStatusInput, ...func(*ssm.Options)) (*ssm.DescribeInstanceAssociationsStatusOutput, error)
	DescribeInstanceInformation(context.Context, *ssm.DescribeInstanceInformationInput, ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error)
	DescribeInstancePatchStates(context.Context, *ssm.DescribeInstancePatchStatesInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesOutput, error)
	DescribeInstancePatchStatesForPatchGroup(context.Context, *ssm.DescribeInstancePatchStatesForPatchGroupInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error)
	DescribeInstancePatches(context.Context, *ssm.DescribeInstancePatchesInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchesOutput, error)
	DescribeInventoryDeletions(context.Context, *ssm.DescribeInventoryDeletionsInput, ...func(*ssm.Options)) (*ssm.DescribeInventoryDeletionsOutput, error)
	DescribeMaintenanceWindowExecutionTaskInvocations(context.Context, *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)
	DescribeMaintenanceWindowExecutionTasks(context.Context, *ssm.DescribeMaintenanceWindowExecutionTasksInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error)
	DescribeMaintenanceWindowExecutions(context.Context, *ssm.DescribeMaintenanceWindowExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error)
	DescribeMaintenanceWindowSchedule(context.Context, *ssm.DescribeMaintenanceWindowScheduleInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowScheduleOutput, error)
	DescribeMaintenanceWindowTargets(context.Context, *ssm.DescribeMaintenanceWindowTargetsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTargetsOutput, error)
	DescribeMaintenanceWindowTasks(context.Context, *ssm.DescribeMaintenanceWindowTasksInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTasksOutput, error)
	DescribeMaintenanceWindows(context.Context, *ssm.DescribeMaintenanceWindowsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsOutput, error)
	DescribeMaintenanceWindowsForTarget(context.Context, *ssm.DescribeMaintenanceWindowsForTargetInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error)
	DescribeOpsItems(context.Context, *ssm.DescribeOpsItemsInput, ...func(*ssm.Options)) (*ssm.DescribeOpsItemsOutput, error)
	DescribeParameters(context.Context, *ssm.DescribeParametersInput, ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error)
	DescribePatchBaselines(context.Context, *ssm.DescribePatchBaselinesInput, ...func(*ssm.Options)) (*ssm.DescribePatchBaselinesOutput, error)
	DescribePatchGroupState(context.Context, *ssm.DescribePatchGroupStateInput, ...func(*ssm.Options)) (*ssm.DescribePatchGroupStateOutput, error)
	DescribePatchGroups(context.Context, *ssm.DescribePatchGroupsInput, ...func(*ssm.Options)) (*ssm.DescribePatchGroupsOutput, error)
	DescribePatchProperties(context.Context, *ssm.DescribePatchPropertiesInput, ...func(*ssm.Options)) (*ssm.DescribePatchPropertiesOutput, error)
	DescribeSessions(context.Context, *ssm.DescribeSessionsInput, ...func(*ssm.Options)) (*ssm.DescribeSessionsOutput, error)
	GetAutomationExecution(context.Context, *ssm.GetAutomationExecutionInput, ...func(*ssm.Options)) (*ssm.GetAutomationExecutionOutput, error)
	GetCalendarState(context.Context, *ssm.GetCalendarStateInput, ...func(*ssm.Options)) (*ssm.GetCalendarStateOutput, error)
	GetCommandInvocation(context.Context, *ssm.GetCommandInvocationInput, ...func(*ssm.Options)) (*ssm.GetCommandInvocationOutput, error)
	GetConnectionStatus(context.Context, *ssm.GetConnectionStatusInput, ...func(*ssm.Options)) (*ssm.GetConnectionStatusOutput, error)
	GetDefaultPatchBaseline(context.Context, *ssm.GetDefaultPatchBaselineInput, ...func(*ssm.Options)) (*ssm.GetDefaultPatchBaselineOutput, error)
	GetDeployablePatchSnapshotForInstance(context.Context, *ssm.GetDeployablePatchSnapshotForInstanceInput, ...func(*ssm.Options)) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error)
	GetDocument(context.Context, *ssm.GetDocumentInput, ...func(*ssm.Options)) (*ssm.GetDocumentOutput, error)
	GetInventory(context.Context, *ssm.GetInventoryInput, ...func(*ssm.Options)) (*ssm.GetInventoryOutput, error)
	GetInventorySchema(context.Context, *ssm.GetInventorySchemaInput, ...func(*ssm.Options)) (*ssm.GetInventorySchemaOutput, error)
	GetMaintenanceWindow(context.Context, *ssm.GetMaintenanceWindowInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowOutput, error)
	GetMaintenanceWindowExecution(context.Context, *ssm.GetMaintenanceWindowExecutionInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionOutput, error)
	GetMaintenanceWindowExecutionTask(context.Context, *ssm.GetMaintenanceWindowExecutionTaskInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error)
	GetMaintenanceWindowExecutionTaskInvocation(context.Context, *ssm.GetMaintenanceWindowExecutionTaskInvocationInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error)
	GetMaintenanceWindowTask(context.Context, *ssm.GetMaintenanceWindowTaskInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowTaskOutput, error)
	GetOpsItem(context.Context, *ssm.GetOpsItemInput, ...func(*ssm.Options)) (*ssm.GetOpsItemOutput, error)
	GetOpsMetadata(context.Context, *ssm.GetOpsMetadataInput, ...func(*ssm.Options)) (*ssm.GetOpsMetadataOutput, error)
	GetOpsSummary(context.Context, *ssm.GetOpsSummaryInput, ...func(*ssm.Options)) (*ssm.GetOpsSummaryOutput, error)
	GetParameter(context.Context, *ssm.GetParameterInput, ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
	GetParameterHistory(context.Context, *ssm.GetParameterHistoryInput, ...func(*ssm.Options)) (*ssm.GetParameterHistoryOutput, error)
	GetParameters(context.Context, *ssm.GetParametersInput, ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
	GetParametersByPath(context.Context, *ssm.GetParametersByPathInput, ...func(*ssm.Options)) (*ssm.GetParametersByPathOutput, error)
	GetPatchBaseline(context.Context, *ssm.GetPatchBaselineInput, ...func(*ssm.Options)) (*ssm.GetPatchBaselineOutput, error)
	GetPatchBaselineForPatchGroup(context.Context, *ssm.GetPatchBaselineForPatchGroupInput, ...func(*ssm.Options)) (*ssm.GetPatchBaselineForPatchGroupOutput, error)
	GetServiceSetting(context.Context, *ssm.GetServiceSettingInput, ...func(*ssm.Options)) (*ssm.GetServiceSettingOutput, error)
	ListAssociationVersions(context.Context, *ssm.ListAssociationVersionsInput, ...func(*ssm.Options)) (*ssm.ListAssociationVersionsOutput, error)
	ListAssociations(context.Context, *ssm.ListAssociationsInput, ...func(*ssm.Options)) (*ssm.ListAssociationsOutput, error)
	ListCommandInvocations(context.Context, *ssm.ListCommandInvocationsInput, ...func(*ssm.Options)) (*ssm.ListCommandInvocationsOutput, error)
	ListCommands(context.Context, *ssm.ListCommandsInput, ...func(*ssm.Options)) (*ssm.ListCommandsOutput, error)
	ListComplianceItems(context.Context, *ssm.ListComplianceItemsInput, ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error)
	ListComplianceSummaries(context.Context, *ssm.ListComplianceSummariesInput, ...func(*ssm.Options)) (*ssm.ListComplianceSummariesOutput, error)
	ListDocumentMetadataHistory(context.Context, *ssm.ListDocumentMetadataHistoryInput, ...func(*ssm.Options)) (*ssm.ListDocumentMetadataHistoryOutput, error)
	ListDocumentVersions(context.Context, *ssm.ListDocumentVersionsInput, ...func(*ssm.Options)) (*ssm.ListDocumentVersionsOutput, error)
	ListDocuments(context.Context, *ssm.ListDocumentsInput, ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error)
	ListInventoryEntries(context.Context, *ssm.ListInventoryEntriesInput, ...func(*ssm.Options)) (*ssm.ListInventoryEntriesOutput, error)
	ListOpsItemEvents(context.Context, *ssm.ListOpsItemEventsInput, ...func(*ssm.Options)) (*ssm.ListOpsItemEventsOutput, error)
	ListOpsItemRelatedItems(context.Context, *ssm.ListOpsItemRelatedItemsInput, ...func(*ssm.Options)) (*ssm.ListOpsItemRelatedItemsOutput, error)
	ListOpsMetadata(context.Context, *ssm.ListOpsMetadataInput, ...func(*ssm.Options)) (*ssm.ListOpsMetadataOutput, error)
	ListResourceComplianceSummaries(context.Context, *ssm.ListResourceComplianceSummariesInput, ...func(*ssm.Options)) (*ssm.ListResourceComplianceSummariesOutput, error)
	ListResourceDataSync(context.Context, *ssm.ListResourceDataSyncInput, ...func(*ssm.Options)) (*ssm.ListResourceDataSyncOutput, error)
	ListTagsForResource(context.Context, *ssm.ListTagsForResourceInput, ...func(*ssm.Options)) (*ssm.ListTagsForResourceOutput, error)
}

type SsoadminClient interface {
	DescribeAccountAssignmentCreationStatus(context.Context, *ssoadmin.DescribeAccountAssignmentCreationStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error)
	DescribeAccountAssignmentDeletionStatus(context.Context, *ssoadmin.DescribeAccountAssignmentDeletionStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error)
	DescribeInstanceAccessControlAttributeConfiguration(context.Context, *ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput, error)
	DescribePermissionSet(context.Context, *ssoadmin.DescribePermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetOutput, error)
	DescribePermissionSetProvisioningStatus(context.Context, *ssoadmin.DescribePermissionSetProvisioningStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error)
	GetInlinePolicyForPermissionSet(context.Context, *ssoadmin.GetInlinePolicyForPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error)
	GetPermissionsBoundaryForPermissionSet(context.Context, *ssoadmin.GetPermissionsBoundaryForPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput, error)
	ListAccountAssignmentCreationStatus(context.Context, *ssoadmin.ListAccountAssignmentCreationStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error)
	ListAccountAssignmentDeletionStatus(context.Context, *ssoadmin.ListAccountAssignmentDeletionStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error)
	ListAccountAssignments(context.Context, *ssoadmin.ListAccountAssignmentsInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentsOutput, error)
	ListAccountsForProvisionedPermissionSet(context.Context, *ssoadmin.ListAccountsForProvisionedPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error)
	ListCustomerManagedPolicyReferencesInPermissionSet(context.Context, *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, error)
	ListInstances(context.Context, *ssoadmin.ListInstancesInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListInstancesOutput, error)
	ListManagedPoliciesInPermissionSet(context.Context, *ssoadmin.ListManagedPoliciesInPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error)
	ListPermissionSetProvisioningStatus(context.Context, *ssoadmin.ListPermissionSetProvisioningStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error)
	ListPermissionSets(context.Context, *ssoadmin.ListPermissionSetsInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsOutput, error)
	ListPermissionSetsProvisionedToAccount(context.Context, *ssoadmin.ListPermissionSetsProvisionedToAccountInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error)
	ListTagsForResource(context.Context, *ssoadmin.ListTagsForResourceInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListTagsForResourceOutput, error)
}

type TransferClient interface {
	DescribeAccess(context.Context, *transfer.DescribeAccessInput, ...func(*transfer.Options)) (*transfer.DescribeAccessOutput, error)
	DescribeAgreement(context.Context, *transfer.DescribeAgreementInput, ...func(*transfer.Options)) (*transfer.DescribeAgreementOutput, error)
	DescribeCertificate(context.Context, *transfer.DescribeCertificateInput, ...func(*transfer.Options)) (*transfer.DescribeCertificateOutput, error)
	DescribeConnector(context.Context, *transfer.DescribeConnectorInput, ...func(*transfer.Options)) (*transfer.DescribeConnectorOutput, error)
	DescribeExecution(context.Context, *transfer.DescribeExecutionInput, ...func(*transfer.Options)) (*transfer.DescribeExecutionOutput, error)
	DescribeHostKey(context.Context, *transfer.DescribeHostKeyInput, ...func(*transfer.Options)) (*transfer.DescribeHostKeyOutput, error)
	DescribeProfile(context.Context, *transfer.DescribeProfileInput, ...func(*transfer.Options)) (*transfer.DescribeProfileOutput, error)
	DescribeSecurityPolicy(context.Context, *transfer.DescribeSecurityPolicyInput, ...func(*transfer.Options)) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeServer(context.Context, *transfer.DescribeServerInput, ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error)
	DescribeUser(context.Context, *transfer.DescribeUserInput, ...func(*transfer.Options)) (*transfer.DescribeUserOutput, error)
	DescribeWorkflow(context.Context, *transfer.DescribeWorkflowInput, ...func(*transfer.Options)) (*transfer.DescribeWorkflowOutput, error)
	ListAccesses(context.Context, *transfer.ListAccessesInput, ...func(*transfer.Options)) (*transfer.ListAccessesOutput, error)
	ListAgreements(context.Context, *transfer.ListAgreementsInput, ...func(*transfer.Options)) (*transfer.ListAgreementsOutput, error)
	ListCertificates(context.Context, *transfer.ListCertificatesInput, ...func(*transfer.Options)) (*transfer.ListCertificatesOutput, error)
	ListConnectors(context.Context, *transfer.ListConnectorsInput, ...func(*transfer.Options)) (*transfer.ListConnectorsOutput, error)
	ListExecutions(context.Context, *transfer.ListExecutionsInput, ...func(*transfer.Options)) (*transfer.ListExecutionsOutput, error)
	ListHostKeys(context.Context, *transfer.ListHostKeysInput, ...func(*transfer.Options)) (*transfer.ListHostKeysOutput, error)
	ListProfiles(context.Context, *transfer.ListProfilesInput, ...func(*transfer.Options)) (*transfer.ListProfilesOutput, error)
	ListSecurityPolicies(context.Context, *transfer.ListSecurityPoliciesInput, ...func(*transfer.Options)) (*transfer.ListSecurityPoliciesOutput, error)
	ListServers(context.Context, *transfer.ListServersInput, ...func(*transfer.Options)) (*transfer.ListServersOutput, error)
	ListTagsForResource(context.Context, *transfer.ListTagsForResourceInput, ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error)
	ListUsers(context.Context, *transfer.ListUsersInput, ...func(*transfer.Options)) (*transfer.ListUsersOutput, error)
	ListWorkflows(context.Context, *transfer.ListWorkflowsInput, ...func(*transfer.Options)) (*transfer.ListWorkflowsOutput, error)
}

type WafClient interface {
	GetByteMatchSet(context.Context, *waf.GetByteMatchSetInput, ...func(*waf.Options)) (*waf.GetByteMatchSetOutput, error)
	GetChangeToken(context.Context, *waf.GetChangeTokenInput, ...func(*waf.Options)) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenStatus(context.Context, *waf.GetChangeTokenStatusInput, ...func(*waf.Options)) (*waf.GetChangeTokenStatusOutput, error)
	GetGeoMatchSet(context.Context, *waf.GetGeoMatchSetInput, ...func(*waf.Options)) (*waf.GetGeoMatchSetOutput, error)
	GetIPSet(context.Context, *waf.GetIPSetInput, ...func(*waf.Options)) (*waf.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *waf.GetLoggingConfigurationInput, ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error)
	GetPermissionPolicy(context.Context, *waf.GetPermissionPolicyInput, ...func(*waf.Options)) (*waf.GetPermissionPolicyOutput, error)
	GetRateBasedRule(context.Context, *waf.GetRateBasedRuleInput, ...func(*waf.Options)) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleManagedKeys(context.Context, *waf.GetRateBasedRuleManagedKeysInput, ...func(*waf.Options)) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRegexMatchSet(context.Context, *waf.GetRegexMatchSetInput, ...func(*waf.Options)) (*waf.GetRegexMatchSetOutput, error)
	GetRegexPatternSet(context.Context, *waf.GetRegexPatternSetInput, ...func(*waf.Options)) (*waf.GetRegexPatternSetOutput, error)
	GetRule(context.Context, *waf.GetRuleInput, ...func(*waf.Options)) (*waf.GetRuleOutput, error)
	GetRuleGroup(context.Context, *waf.GetRuleGroupInput, ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *waf.GetSampledRequestsInput, ...func(*waf.Options)) (*waf.GetSampledRequestsOutput, error)
	GetSizeConstraintSet(context.Context, *waf.GetSizeConstraintSetInput, ...func(*waf.Options)) (*waf.GetSizeConstraintSetOutput, error)
	GetSqlInjectionMatchSet(context.Context, *waf.GetSqlInjectionMatchSetInput, ...func(*waf.Options)) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetWebACL(context.Context, *waf.GetWebACLInput, ...func(*waf.Options)) (*waf.GetWebACLOutput, error)
	GetXssMatchSet(context.Context, *waf.GetXssMatchSetInput, ...func(*waf.Options)) (*waf.GetXssMatchSetOutput, error)
	ListActivatedRulesInRuleGroup(context.Context, *waf.ListActivatedRulesInRuleGroupInput, ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListByteMatchSets(context.Context, *waf.ListByteMatchSetsInput, ...func(*waf.Options)) (*waf.ListByteMatchSetsOutput, error)
	ListGeoMatchSets(context.Context, *waf.ListGeoMatchSetsInput, ...func(*waf.Options)) (*waf.ListGeoMatchSetsOutput, error)
	ListIPSets(context.Context, *waf.ListIPSetsInput, ...func(*waf.Options)) (*waf.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *waf.ListLoggingConfigurationsInput, ...func(*waf.Options)) (*waf.ListLoggingConfigurationsOutput, error)
	ListRateBasedRules(context.Context, *waf.ListRateBasedRulesInput, ...func(*waf.Options)) (*waf.ListRateBasedRulesOutput, error)
	ListRegexMatchSets(context.Context, *waf.ListRegexMatchSetsInput, ...func(*waf.Options)) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexPatternSets(context.Context, *waf.ListRegexPatternSetsInput, ...func(*waf.Options)) (*waf.ListRegexPatternSetsOutput, error)
	ListRuleGroups(context.Context, *waf.ListRuleGroupsInput, ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error)
	ListRules(context.Context, *waf.ListRulesInput, ...func(*waf.Options)) (*waf.ListRulesOutput, error)
	ListSizeConstraintSets(context.Context, *waf.ListSizeConstraintSetsInput, ...func(*waf.Options)) (*waf.ListSizeConstraintSetsOutput, error)
	ListSqlInjectionMatchSets(context.Context, *waf.ListSqlInjectionMatchSetsInput, ...func(*waf.Options)) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSubscribedRuleGroups(context.Context, *waf.ListSubscribedRuleGroupsInput, ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *waf.ListTagsForResourceInput, ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *waf.ListWebACLsInput, ...func(*waf.Options)) (*waf.ListWebACLsOutput, error)
	ListXssMatchSets(context.Context, *waf.ListXssMatchSetsInput, ...func(*waf.Options)) (*waf.ListXssMatchSetsOutput, error)
}

type WafregionalClient interface {
	GetByteMatchSet(context.Context, *wafregional.GetByteMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetByteMatchSetOutput, error)
	GetChangeToken(context.Context, *wafregional.GetChangeTokenInput, ...func(*wafregional.Options)) (*wafregional.GetChangeTokenOutput, error)
	GetChangeTokenStatus(context.Context, *wafregional.GetChangeTokenStatusInput, ...func(*wafregional.Options)) (*wafregional.GetChangeTokenStatusOutput, error)
	GetGeoMatchSet(context.Context, *wafregional.GetGeoMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetGeoMatchSetOutput, error)
	GetIPSet(context.Context, *wafregional.GetIPSetInput, ...func(*wafregional.Options)) (*wafregional.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *wafregional.GetLoggingConfigurationInput, ...func(*wafregional.Options)) (*wafregional.GetLoggingConfigurationOutput, error)
	GetPermissionPolicy(context.Context, *wafregional.GetPermissionPolicyInput, ...func(*wafregional.Options)) (*wafregional.GetPermissionPolicyOutput, error)
	GetRateBasedRule(context.Context, *wafregional.GetRateBasedRuleInput, ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error)
	GetRateBasedRuleManagedKeys(context.Context, *wafregional.GetRateBasedRuleManagedKeysInput, ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleManagedKeysOutput, error)
	GetRegexMatchSet(context.Context, *wafregional.GetRegexMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetRegexMatchSetOutput, error)
	GetRegexPatternSet(context.Context, *wafregional.GetRegexPatternSetInput, ...func(*wafregional.Options)) (*wafregional.GetRegexPatternSetOutput, error)
	GetRule(context.Context, *wafregional.GetRuleInput, ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error)
	GetRuleGroup(context.Context, *wafregional.GetRuleGroupInput, ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *wafregional.GetSampledRequestsInput, ...func(*wafregional.Options)) (*wafregional.GetSampledRequestsOutput, error)
	GetSizeConstraintSet(context.Context, *wafregional.GetSizeConstraintSetInput, ...func(*wafregional.Options)) (*wafregional.GetSizeConstraintSetOutput, error)
	GetSqlInjectionMatchSet(context.Context, *wafregional.GetSqlInjectionMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetSqlInjectionMatchSetOutput, error)
	GetWebACL(context.Context, *wafregional.GetWebACLInput, ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error)
	GetWebACLForResource(context.Context, *wafregional.GetWebACLForResourceInput, ...func(*wafregional.Options)) (*wafregional.GetWebACLForResourceOutput, error)
	GetXssMatchSet(context.Context, *wafregional.GetXssMatchSetInput, ...func(*wafregional.Options)) (*wafregional.GetXssMatchSetOutput, error)
	ListActivatedRulesInRuleGroup(context.Context, *wafregional.ListActivatedRulesInRuleGroupInput, ...func(*wafregional.Options)) (*wafregional.ListActivatedRulesInRuleGroupOutput, error)
	ListByteMatchSets(context.Context, *wafregional.ListByteMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListByteMatchSetsOutput, error)
	ListGeoMatchSets(context.Context, *wafregional.ListGeoMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListGeoMatchSetsOutput, error)
	ListIPSets(context.Context, *wafregional.ListIPSetsInput, ...func(*wafregional.Options)) (*wafregional.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *wafregional.ListLoggingConfigurationsInput, ...func(*wafregional.Options)) (*wafregional.ListLoggingConfigurationsOutput, error)
	ListRateBasedRules(context.Context, *wafregional.ListRateBasedRulesInput, ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error)
	ListRegexMatchSets(context.Context, *wafregional.ListRegexMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListRegexMatchSetsOutput, error)
	ListRegexPatternSets(context.Context, *wafregional.ListRegexPatternSetsInput, ...func(*wafregional.Options)) (*wafregional.ListRegexPatternSetsOutput, error)
	ListResourcesForWebACL(context.Context, *wafregional.ListResourcesForWebACLInput, ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error)
	ListRuleGroups(context.Context, *wafregional.ListRuleGroupsInput, ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error)
	ListRules(context.Context, *wafregional.ListRulesInput, ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error)
	ListSizeConstraintSets(context.Context, *wafregional.ListSizeConstraintSetsInput, ...func(*wafregional.Options)) (*wafregional.ListSizeConstraintSetsOutput, error)
	ListSqlInjectionMatchSets(context.Context, *wafregional.ListSqlInjectionMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListSqlInjectionMatchSetsOutput, error)
	ListSubscribedRuleGroups(context.Context, *wafregional.ListSubscribedRuleGroupsInput, ...func(*wafregional.Options)) (*wafregional.ListSubscribedRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *wafregional.ListTagsForResourceInput, ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *wafregional.ListWebACLsInput, ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error)
	ListXssMatchSets(context.Context, *wafregional.ListXssMatchSetsInput, ...func(*wafregional.Options)) (*wafregional.ListXssMatchSetsOutput, error)
}

type Wafv2Client interface {
	DescribeManagedRuleGroup(context.Context, *wafv2.DescribeManagedRuleGroupInput, ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error)
	GetIPSet(context.Context, *wafv2.GetIPSetInput, ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error)
	GetLoggingConfiguration(context.Context, *wafv2.GetLoggingConfigurationInput, ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error)
	GetManagedRuleSet(context.Context, *wafv2.GetManagedRuleSetInput, ...func(*wafv2.Options)) (*wafv2.GetManagedRuleSetOutput, error)
	GetMobileSdkRelease(context.Context, *wafv2.GetMobileSdkReleaseInput, ...func(*wafv2.Options)) (*wafv2.GetMobileSdkReleaseOutput, error)
	GetPermissionPolicy(context.Context, *wafv2.GetPermissionPolicyInput, ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error)
	GetRateBasedStatementManagedKeys(context.Context, *wafv2.GetRateBasedStatementManagedKeysInput, ...func(*wafv2.Options)) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
	GetRegexPatternSet(context.Context, *wafv2.GetRegexPatternSetInput, ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error)
	GetRuleGroup(context.Context, *wafv2.GetRuleGroupInput, ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error)
	GetSampledRequests(context.Context, *wafv2.GetSampledRequestsInput, ...func(*wafv2.Options)) (*wafv2.GetSampledRequestsOutput, error)
	GetWebACL(context.Context, *wafv2.GetWebACLInput, ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error)
	GetWebACLForResource(context.Context, *wafv2.GetWebACLForResourceInput, ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error)
	ListAvailableManagedRuleGroupVersions(context.Context, *wafv2.ListAvailableManagedRuleGroupVersionsInput, ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error)
	ListAvailableManagedRuleGroups(context.Context, *wafv2.ListAvailableManagedRuleGroupsInput, ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListIPSets(context.Context, *wafv2.ListIPSetsInput, ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error)
	ListLoggingConfigurations(context.Context, *wafv2.ListLoggingConfigurationsInput, ...func(*wafv2.Options)) (*wafv2.ListLoggingConfigurationsOutput, error)
	ListManagedRuleSets(context.Context, *wafv2.ListManagedRuleSetsInput, ...func(*wafv2.Options)) (*wafv2.ListManagedRuleSetsOutput, error)
	ListMobileSdkReleases(context.Context, *wafv2.ListMobileSdkReleasesInput, ...func(*wafv2.Options)) (*wafv2.ListMobileSdkReleasesOutput, error)
	ListRegexPatternSets(context.Context, *wafv2.ListRegexPatternSetsInput, ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error)
	ListResourcesForWebACL(context.Context, *wafv2.ListResourcesForWebACLInput, ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error)
	ListRuleGroups(context.Context, *wafv2.ListRuleGroupsInput, ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error)
	ListTagsForResource(context.Context, *wafv2.ListTagsForResourceInput, ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error)
	ListWebACLs(context.Context, *wafv2.ListWebACLsInput, ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error)
}

type WorkspacesClient interface {
	DescribeAccount(context.Context, *workspaces.DescribeAccountInput, ...func(*workspaces.Options)) (*workspaces.DescribeAccountOutput, error)
	DescribeAccountModifications(context.Context, *workspaces.DescribeAccountModificationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeAccountModificationsOutput, error)
	DescribeClientBranding(context.Context, *workspaces.DescribeClientBrandingInput, ...func(*workspaces.Options)) (*workspaces.DescribeClientBrandingOutput, error)
	DescribeClientProperties(context.Context, *workspaces.DescribeClientPropertiesInput, ...func(*workspaces.Options)) (*workspaces.DescribeClientPropertiesOutput, error)
	DescribeConnectClientAddIns(context.Context, *workspaces.DescribeConnectClientAddInsInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectClientAddInsOutput, error)
	DescribeConnectionAliasPermissions(context.Context, *workspaces.DescribeConnectionAliasPermissionsInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasPermissionsOutput, error)
	DescribeConnectionAliases(context.Context, *workspaces.DescribeConnectionAliasesInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasesOutput, error)
	DescribeIpGroups(context.Context, *workspaces.DescribeIpGroupsInput, ...func(*workspaces.Options)) (*workspaces.DescribeIpGroupsOutput, error)
	DescribeTags(context.Context, *workspaces.DescribeTagsInput, ...func(*workspaces.Options)) (*workspaces.DescribeTagsOutput, error)
	DescribeWorkspaceBundles(context.Context, *workspaces.DescribeWorkspaceBundlesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceDirectories(context.Context, *workspaces.DescribeWorkspaceDirectoriesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceImagePermissions(context.Context, *workspaces.DescribeWorkspaceImagePermissionsInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error)
	DescribeWorkspaceImages(context.Context, *workspaces.DescribeWorkspaceImagesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagesOutput, error)
	DescribeWorkspaceSnapshots(context.Context, *workspaces.DescribeWorkspaceSnapshotsInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceSnapshotsOutput, error)
	DescribeWorkspaces(context.Context, *workspaces.DescribeWorkspacesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesConnectionStatus(context.Context, *workspaces.DescribeWorkspacesConnectionStatusInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	ListAvailableManagementCidrRanges(context.Context, *workspaces.ListAvailableManagementCidrRangesInput, ...func(*workspaces.Options)) (*workspaces.ListAvailableManagementCidrRangesOutput, error)
}

type XrayClient interface {
	BatchGetTraces(context.Context, *xray.BatchGetTracesInput, ...func(*xray.Options)) (*xray.BatchGetTracesOutput, error)
	GetEncryptionConfig(context.Context, *xray.GetEncryptionConfigInput, ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error)
	GetGroup(context.Context, *xray.GetGroupInput, ...func(*xray.Options)) (*xray.GetGroupOutput, error)
	GetGroups(context.Context, *xray.GetGroupsInput, ...func(*xray.Options)) (*xray.GetGroupsOutput, error)
	GetInsight(context.Context, *xray.GetInsightInput, ...func(*xray.Options)) (*xray.GetInsightOutput, error)
	GetInsightEvents(context.Context, *xray.GetInsightEventsInput, ...func(*xray.Options)) (*xray.GetInsightEventsOutput, error)
	GetInsightImpactGraph(context.Context, *xray.GetInsightImpactGraphInput, ...func(*xray.Options)) (*xray.GetInsightImpactGraphOutput, error)
	GetInsightSummaries(context.Context, *xray.GetInsightSummariesInput, ...func(*xray.Options)) (*xray.GetInsightSummariesOutput, error)
	GetSamplingRules(context.Context, *xray.GetSamplingRulesInput, ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error)
	GetSamplingStatisticSummaries(context.Context, *xray.GetSamplingStatisticSummariesInput, ...func(*xray.Options)) (*xray.GetSamplingStatisticSummariesOutput, error)
	GetSamplingTargets(context.Context, *xray.GetSamplingTargetsInput, ...func(*xray.Options)) (*xray.GetSamplingTargetsOutput, error)
	GetServiceGraph(context.Context, *xray.GetServiceGraphInput, ...func(*xray.Options)) (*xray.GetServiceGraphOutput, error)
	GetTimeSeriesServiceStatistics(context.Context, *xray.GetTimeSeriesServiceStatisticsInput, ...func(*xray.Options)) (*xray.GetTimeSeriesServiceStatisticsOutput, error)
	GetTraceGraph(context.Context, *xray.GetTraceGraphInput, ...func(*xray.Options)) (*xray.GetTraceGraphOutput, error)
	GetTraceSummaries(context.Context, *xray.GetTraceSummariesInput, ...func(*xray.Options)) (*xray.GetTraceSummariesOutput, error)
	ListResourcePolicies(context.Context, *xray.ListResourcePoliciesInput, ...func(*xray.Options)) (*xray.ListResourcePoliciesOutput, error)
	ListTagsForResource(context.Context, *xray.ListTagsForResourceInput, ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error)
}

type AccountClient interface {
	GetAlternateContact(context.Context, *account.GetAlternateContactInput, ...func(*account.Options)) (*account.GetAlternateContactOutput, error)
	GetContactInformation(context.Context, *account.GetContactInformationInput, ...func(*account.Options)) (*account.GetContactInformationOutput, error)
}

type AcmClient interface {
	DescribeCertificate(context.Context, *acm.DescribeCertificateInput, ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error)
	GetAccountConfiguration(context.Context, *acm.GetAccountConfigurationInput, ...func(*acm.Options)) (*acm.GetAccountConfigurationOutput, error)
	GetCertificate(context.Context, *acm.GetCertificateInput, ...func(*acm.Options)) (*acm.GetCertificateOutput, error)
	ListCertificates(context.Context, *acm.ListCertificatesInput, ...func(*acm.Options)) (*acm.ListCertificatesOutput, error)
	ListTagsForCertificate(context.Context, *acm.ListTagsForCertificateInput, ...func(*acm.Options)) (*acm.ListTagsForCertificateOutput, error)
}

type AmpClient interface {
	DescribeAlertManagerDefinition(context.Context, *amp.DescribeAlertManagerDefinitionInput, ...func(*amp.Options)) (*amp.DescribeAlertManagerDefinitionOutput, error)
	DescribeLoggingConfiguration(context.Context, *amp.DescribeLoggingConfigurationInput, ...func(*amp.Options)) (*amp.DescribeLoggingConfigurationOutput, error)
	DescribeRuleGroupsNamespace(context.Context, *amp.DescribeRuleGroupsNamespaceInput, ...func(*amp.Options)) (*amp.DescribeRuleGroupsNamespaceOutput, error)
	DescribeWorkspace(context.Context, *amp.DescribeWorkspaceInput, ...func(*amp.Options)) (*amp.DescribeWorkspaceOutput, error)
	ListRuleGroupsNamespaces(context.Context, *amp.ListRuleGroupsNamespacesInput, ...func(*amp.Options)) (*amp.ListRuleGroupsNamespacesOutput, error)
	ListTagsForResource(context.Context, *amp.ListTagsForResourceInput, ...func(*amp.Options)) (*amp.ListTagsForResourceOutput, error)
	ListWorkspaces(context.Context, *amp.ListWorkspacesInput, ...func(*amp.Options)) (*amp.ListWorkspacesOutput, error)
}

type AppstreamClient interface {
	DescribeAppBlocks(context.Context, *appstream.DescribeAppBlocksInput, ...func(*appstream.Options)) (*appstream.DescribeAppBlocksOutput, error)
	DescribeApplicationFleetAssociations(context.Context, *appstream.DescribeApplicationFleetAssociationsInput, ...func(*appstream.Options)) (*appstream.DescribeApplicationFleetAssociationsOutput, error)
	DescribeApplications(context.Context, *appstream.DescribeApplicationsInput, ...func(*appstream.Options)) (*appstream.DescribeApplicationsOutput, error)
	DescribeDirectoryConfigs(context.Context, *appstream.DescribeDirectoryConfigsInput, ...func(*appstream.Options)) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeEntitlements(context.Context, *appstream.DescribeEntitlementsInput, ...func(*appstream.Options)) (*appstream.DescribeEntitlementsOutput, error)
	DescribeFleets(context.Context, *appstream.DescribeFleetsInput, ...func(*appstream.Options)) (*appstream.DescribeFleetsOutput, error)
	DescribeImageBuilders(context.Context, *appstream.DescribeImageBuildersInput, ...func(*appstream.Options)) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImagePermissions(context.Context, *appstream.DescribeImagePermissionsInput, ...func(*appstream.Options)) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImages(context.Context, *appstream.DescribeImagesInput, ...func(*appstream.Options)) (*appstream.DescribeImagesOutput, error)
	DescribeSessions(context.Context, *appstream.DescribeSessionsInput, ...func(*appstream.Options)) (*appstream.DescribeSessionsOutput, error)
	DescribeStacks(context.Context, *appstream.DescribeStacksInput, ...func(*appstream.Options)) (*appstream.DescribeStacksOutput, error)
	DescribeUsageReportSubscriptions(context.Context, *appstream.DescribeUsageReportSubscriptionsInput, ...func(*appstream.Options)) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUserStackAssociations(context.Context, *appstream.DescribeUserStackAssociationsInput, ...func(*appstream.Options)) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUsers(context.Context, *appstream.DescribeUsersInput, ...func(*appstream.Options)) (*appstream.DescribeUsersOutput, error)
	ListAssociatedFleets(context.Context, *appstream.ListAssociatedFleetsInput, ...func(*appstream.Options)) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedStacks(context.Context, *appstream.ListAssociatedStacksInput, ...func(*appstream.Options)) (*appstream.ListAssociatedStacksOutput, error)
	ListEntitledApplications(context.Context, *appstream.ListEntitledApplicationsInput, ...func(*appstream.Options)) (*appstream.ListEntitledApplicationsOutput, error)
	ListTagsForResource(context.Context, *appstream.ListTagsForResourceInput, ...func(*appstream.Options)) (*appstream.ListTagsForResourceOutput, error)
}

type ElastictranscoderClient interface {
	ListJobsByPipeline(context.Context, *elastictranscoder.ListJobsByPipelineInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByStatus(context.Context, *elastictranscoder.ListJobsByStatusInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListPipelines(context.Context, *elastictranscoder.ListPipelinesInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPipelinesOutput, error)
	ListPresets(context.Context, *elastictranscoder.ListPresetsInput, ...func(*elastictranscoder.Options)) (*elastictranscoder.ListPresetsOutput, error)
}

type IAMClient interface {
	GenerateCredentialReport(context.Context, *iam.GenerateCredentialReportInput, ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error)
	GetAccessKeyLastUsed(context.Context, *iam.GetAccessKeyLastUsedInput, ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccountAuthorizationDetails(context.Context, *iam.GetAccountAuthorizationDetailsInput, ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountPasswordPolicy(context.Context, *iam.GetAccountPasswordPolicyInput, ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountSummary(context.Context, *iam.GetAccountSummaryInput, ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error)
	GetContextKeysForCustomPolicy(context.Context, *iam.GetContextKeysForCustomPolicyInput, ...func(*iam.Options)) (*iam.GetContextKeysForCustomPolicyOutput, error)
	GetContextKeysForPrincipalPolicy(context.Context, *iam.GetContextKeysForPrincipalPolicyInput, ...func(*iam.Options)) (*iam.GetContextKeysForPrincipalPolicyOutput, error)
	GetCredentialReport(context.Context, *iam.GetCredentialReportInput, ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error)
	GetGroup(context.Context, *iam.GetGroupInput, ...func(*iam.Options)) (*iam.GetGroupOutput, error)
	GetGroupPolicy(context.Context, *iam.GetGroupPolicyInput, ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error)
	GetInstanceProfile(context.Context, *iam.GetInstanceProfileInput, ...func(*iam.Options)) (*iam.GetInstanceProfileOutput, error)
	GetLoginProfile(context.Context, *iam.GetLoginProfileInput, ...func(*iam.Options)) (*iam.GetLoginProfileOutput, error)
	GetOpenIDConnectProvider(context.Context, *iam.GetOpenIDConnectProviderInput, ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOrganizationsAccessReport(context.Context, *iam.GetOrganizationsAccessReportInput, ...func(*iam.Options)) (*iam.GetOrganizationsAccessReportOutput, error)
	GetPolicy(context.Context, *iam.GetPolicyInput, ...func(*iam.Options)) (*iam.GetPolicyOutput, error)
	GetPolicyVersion(context.Context, *iam.GetPolicyVersionInput, ...func(*iam.Options)) (*iam.GetPolicyVersionOutput, error)
	GetRole(context.Context, *iam.GetRoleInput, ...func(*iam.Options)) (*iam.GetRoleOutput, error)
	GetRolePolicy(context.Context, *iam.GetRolePolicyInput, ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error)
	GetSAMLProvider(context.Context, *iam.GetSAMLProviderInput, ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error)
	GetSSHPublicKey(context.Context, *iam.GetSSHPublicKeyInput, ...func(*iam.Options)) (*iam.GetSSHPublicKeyOutput, error)
	GetServerCertificate(context.Context, *iam.GetServerCertificateInput, ...func(*iam.Options)) (*iam.GetServerCertificateOutput, error)
	GetServiceLastAccessedDetails(context.Context, *iam.GetServiceLastAccessedDetailsInput, ...func(*iam.Options)) (*iam.GetServiceLastAccessedDetailsOutput, error)
	GetServiceLastAccessedDetailsWithEntities(context.Context, *iam.GetServiceLastAccessedDetailsWithEntitiesInput, ...func(*iam.Options)) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
	GetServiceLinkedRoleDeletionStatus(context.Context, *iam.GetServiceLinkedRoleDeletionStatusInput, ...func(*iam.Options)) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetUser(context.Context, *iam.GetUserInput, ...func(*iam.Options)) (*iam.GetUserOutput, error)
	GetUserPolicy(context.Context, *iam.GetUserPolicyInput, ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error)
	ListAccessKeys(context.Context, *iam.ListAccessKeysInput, ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)
	ListAccountAliases(context.Context, *iam.ListAccountAliasesInput, ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)
	ListAttachedGroupPolicies(context.Context, *iam.ListAttachedGroupPoliciesInput, ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedRolePolicies(context.Context, *iam.ListAttachedRolePoliciesInput, ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedUserPolicies(context.Context, *iam.ListAttachedUserPoliciesInput, ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error)
	ListEntitiesForPolicy(context.Context, *iam.ListEntitiesForPolicyInput, ...func(*iam.Options)) (*iam.ListEntitiesForPolicyOutput, error)
	ListGroupPolicies(context.Context, *iam.ListGroupPoliciesInput, ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error)
	ListGroups(context.Context, *iam.ListGroupsInput, ...func(*iam.Options)) (*iam.ListGroupsOutput, error)
	ListGroupsForUser(context.Context, *iam.ListGroupsForUserInput, ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error)
	ListInstanceProfileTags(context.Context, *iam.ListInstanceProfileTagsInput, ...func(*iam.Options)) (*iam.ListInstanceProfileTagsOutput, error)
	ListInstanceProfiles(context.Context, *iam.ListInstanceProfilesInput, ...func(*iam.Options)) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesForRole(context.Context, *iam.ListInstanceProfilesForRoleInput, ...func(*iam.Options)) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListMFADeviceTags(context.Context, *iam.ListMFADeviceTagsInput, ...func(*iam.Options)) (*iam.ListMFADeviceTagsOutput, error)
	ListMFADevices(context.Context, *iam.ListMFADevicesInput, ...func(*iam.Options)) (*iam.ListMFADevicesOutput, error)
	ListOpenIDConnectProviderTags(context.Context, *iam.ListOpenIDConnectProviderTagsInput, ...func(*iam.Options)) (*iam.ListOpenIDConnectProviderTagsOutput, error)
	ListOpenIDConnectProviders(context.Context, *iam.ListOpenIDConnectProvidersInput, ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListPolicies(context.Context, *iam.ListPoliciesInput, ...func(*iam.Options)) (*iam.ListPoliciesOutput, error)
	ListPoliciesGrantingServiceAccess(context.Context, *iam.ListPoliciesGrantingServiceAccessInput, ...func(*iam.Options)) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
	ListPolicyTags(context.Context, *iam.ListPolicyTagsInput, ...func(*iam.Options)) (*iam.ListPolicyTagsOutput, error)
	ListPolicyVersions(context.Context, *iam.ListPolicyVersionsInput, ...func(*iam.Options)) (*iam.ListPolicyVersionsOutput, error)
	ListRolePolicies(context.Context, *iam.ListRolePoliciesInput, ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error)
	ListRoleTags(context.Context, *iam.ListRoleTagsInput, ...func(*iam.Options)) (*iam.ListRoleTagsOutput, error)
	ListRoles(context.Context, *iam.ListRolesInput, ...func(*iam.Options)) (*iam.ListRolesOutput, error)
	ListSAMLProviderTags(context.Context, *iam.ListSAMLProviderTagsInput, ...func(*iam.Options)) (*iam.ListSAMLProviderTagsOutput, error)
	ListSAMLProviders(context.Context, *iam.ListSAMLProvidersInput, ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error)
	ListSSHPublicKeys(context.Context, *iam.ListSSHPublicKeysInput, ...func(*iam.Options)) (*iam.ListSSHPublicKeysOutput, error)
	ListServerCertificateTags(context.Context, *iam.ListServerCertificateTagsInput, ...func(*iam.Options)) (*iam.ListServerCertificateTagsOutput, error)
	ListServerCertificates(context.Context, *iam.ListServerCertificatesInput, ...func(*iam.Options)) (*iam.ListServerCertificatesOutput, error)
	ListServiceSpecificCredentials(context.Context, *iam.ListServiceSpecificCredentialsInput, ...func(*iam.Options)) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListSigningCertificates(context.Context, *iam.ListSigningCertificatesInput, ...func(*iam.Options)) (*iam.ListSigningCertificatesOutput, error)
	ListUserPolicies(context.Context, *iam.ListUserPoliciesInput, ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error)
	ListUserTags(context.Context, *iam.ListUserTagsInput, ...func(*iam.Options)) (*iam.ListUserTagsOutput, error)
	ListUsers(context.Context, *iam.ListUsersInput, ...func(*iam.Options)) (*iam.ListUsersOutput, error)
	ListVirtualMFADevices(context.Context, *iam.ListVirtualMFADevicesInput, ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error)
}

type QuicksightClient interface {
	DescribeAccountCustomization(context.Context, *quicksight.DescribeAccountCustomizationInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountCustomizationOutput, error)
	DescribeAccountSettings(context.Context, *quicksight.DescribeAccountSettingsInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountSettingsOutput, error)
	DescribeAccountSubscription(context.Context, *quicksight.DescribeAccountSubscriptionInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountSubscriptionOutput, error)
	DescribeAnalysis(context.Context, *quicksight.DescribeAnalysisInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisOutput, error)
	DescribeAnalysisDefinition(context.Context, *quicksight.DescribeAnalysisDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisDefinitionOutput, error)
	DescribeAnalysisPermissions(context.Context, *quicksight.DescribeAnalysisPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisPermissionsOutput, error)
	DescribeDashboard(context.Context, *quicksight.DescribeDashboardInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardOutput, error)
	DescribeDashboardDefinition(context.Context, *quicksight.DescribeDashboardDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardDefinitionOutput, error)
	DescribeDashboardPermissions(context.Context, *quicksight.DescribeDashboardPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardPermissionsOutput, error)
	DescribeDataSet(context.Context, *quicksight.DescribeDataSetInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSetOutput, error)
	DescribeDataSetPermissions(context.Context, *quicksight.DescribeDataSetPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSetPermissionsOutput, error)
	DescribeDataSource(context.Context, *quicksight.DescribeDataSourceInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSourceOutput, error)
	DescribeDataSourcePermissions(context.Context, *quicksight.DescribeDataSourcePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSourcePermissionsOutput, error)
	DescribeFolder(context.Context, *quicksight.DescribeFolderInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderOutput, error)
	DescribeFolderPermissions(context.Context, *quicksight.DescribeFolderPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderPermissionsOutput, error)
	DescribeFolderResolvedPermissions(context.Context, *quicksight.DescribeFolderResolvedPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderResolvedPermissionsOutput, error)
	DescribeGroup(context.Context, *quicksight.DescribeGroupInput, ...func(*quicksight.Options)) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupMembership(context.Context, *quicksight.DescribeGroupMembershipInput, ...func(*quicksight.Options)) (*quicksight.DescribeGroupMembershipOutput, error)
	DescribeIAMPolicyAssignment(context.Context, *quicksight.DescribeIAMPolicyAssignmentInput, ...func(*quicksight.Options)) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
	DescribeIngestion(context.Context, *quicksight.DescribeIngestionInput, ...func(*quicksight.Options)) (*quicksight.DescribeIngestionOutput, error)
	DescribeIpRestriction(context.Context, *quicksight.DescribeIpRestrictionInput, ...func(*quicksight.Options)) (*quicksight.DescribeIpRestrictionOutput, error)
	DescribeNamespace(context.Context, *quicksight.DescribeNamespaceInput, ...func(*quicksight.Options)) (*quicksight.DescribeNamespaceOutput, error)
	DescribeTemplate(context.Context, *quicksight.DescribeTemplateInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateOutput, error)
	DescribeTemplateAlias(context.Context, *quicksight.DescribeTemplateAliasInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateAliasOutput, error)
	DescribeTemplateDefinition(context.Context, *quicksight.DescribeTemplateDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateDefinitionOutput, error)
	DescribeTemplatePermissions(context.Context, *quicksight.DescribeTemplatePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplatePermissionsOutput, error)
	DescribeTheme(context.Context, *quicksight.DescribeThemeInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemeOutput, error)
	DescribeThemeAlias(context.Context, *quicksight.DescribeThemeAliasInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemeAliasOutput, error)
	DescribeThemePermissions(context.Context, *quicksight.DescribeThemePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemePermissionsOutput, error)
	DescribeUser(context.Context, *quicksight.DescribeUserInput, ...func(*quicksight.Options)) (*quicksight.DescribeUserOutput, error)
	GetDashboardEmbedUrl(context.Context, *quicksight.GetDashboardEmbedUrlInput, ...func(*quicksight.Options)) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetSessionEmbedUrl(context.Context, *quicksight.GetSessionEmbedUrlInput, ...func(*quicksight.Options)) (*quicksight.GetSessionEmbedUrlOutput, error)
	ListAnalyses(context.Context, *quicksight.ListAnalysesInput, ...func(*quicksight.Options)) (*quicksight.ListAnalysesOutput, error)
	ListDashboardVersions(context.Context, *quicksight.ListDashboardVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListDashboardVersionsOutput, error)
	ListDashboards(context.Context, *quicksight.ListDashboardsInput, ...func(*quicksight.Options)) (*quicksight.ListDashboardsOutput, error)
	ListDataSets(context.Context, *quicksight.ListDataSetsInput, ...func(*quicksight.Options)) (*quicksight.ListDataSetsOutput, error)
	ListDataSources(context.Context, *quicksight.ListDataSourcesInput, ...func(*quicksight.Options)) (*quicksight.ListDataSourcesOutput, error)
	ListFolderMembers(context.Context, *quicksight.ListFolderMembersInput, ...func(*quicksight.Options)) (*quicksight.ListFolderMembersOutput, error)
	ListFolders(context.Context, *quicksight.ListFoldersInput, ...func(*quicksight.Options)) (*quicksight.ListFoldersOutput, error)
	ListGroupMemberships(context.Context, *quicksight.ListGroupMembershipsInput, ...func(*quicksight.Options)) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroups(context.Context, *quicksight.ListGroupsInput, ...func(*quicksight.Options)) (*quicksight.ListGroupsOutput, error)
	ListIAMPolicyAssignments(context.Context, *quicksight.ListIAMPolicyAssignmentsInput, ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
	ListIAMPolicyAssignmentsForUser(context.Context, *quicksight.ListIAMPolicyAssignmentsForUserInput, ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
	ListIngestions(context.Context, *quicksight.ListIngestionsInput, ...func(*quicksight.Options)) (*quicksight.ListIngestionsOutput, error)
	ListNamespaces(context.Context, *quicksight.ListNamespacesInput, ...func(*quicksight.Options)) (*quicksight.ListNamespacesOutput, error)
	ListTagsForResource(context.Context, *quicksight.ListTagsForResourceInput, ...func(*quicksight.Options)) (*quicksight.ListTagsForResourceOutput, error)
	ListTemplateAliases(context.Context, *quicksight.ListTemplateAliasesInput, ...func(*quicksight.Options)) (*quicksight.ListTemplateAliasesOutput, error)
	ListTemplateVersions(context.Context, *quicksight.ListTemplateVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListTemplateVersionsOutput, error)
	ListTemplates(context.Context, *quicksight.ListTemplatesInput, ...func(*quicksight.Options)) (*quicksight.ListTemplatesOutput, error)
	ListThemeAliases(context.Context, *quicksight.ListThemeAliasesInput, ...func(*quicksight.Options)) (*quicksight.ListThemeAliasesOutput, error)
	ListThemeVersions(context.Context, *quicksight.ListThemeVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListThemeVersionsOutput, error)
	ListThemes(context.Context, *quicksight.ListThemesInput, ...func(*quicksight.Options)) (*quicksight.ListThemesOutput, error)
	ListUserGroups(context.Context, *quicksight.ListUserGroupsInput, ...func(*quicksight.Options)) (*quicksight.ListUserGroupsOutput, error)
	ListUsers(context.Context, *quicksight.ListUsersInput, ...func(*quicksight.Options)) (*quicksight.ListUsersOutput, error)
	SearchAnalyses(context.Context, *quicksight.SearchAnalysesInput, ...func(*quicksight.Options)) (*quicksight.SearchAnalysesOutput, error)
	SearchDashboards(context.Context, *quicksight.SearchDashboardsInput, ...func(*quicksight.Options)) (*quicksight.SearchDashboardsOutput, error)
	SearchDataSets(context.Context, *quicksight.SearchDataSetsInput, ...func(*quicksight.Options)) (*quicksight.SearchDataSetsOutput, error)
	SearchDataSources(context.Context, *quicksight.SearchDataSourcesInput, ...func(*quicksight.Options)) (*quicksight.SearchDataSourcesOutput, error)
	SearchFolders(context.Context, *quicksight.SearchFoldersInput, ...func(*quicksight.Options)) (*quicksight.SearchFoldersOutput, error)
	SearchGroups(context.Context, *quicksight.SearchGroupsInput, ...func(*quicksight.Options)) (*quicksight.SearchGroupsOutput, error)
}

type RAMClient interface {
	GetPermission(context.Context, *ram.GetPermissionInput, ...func(*ram.Options)) (*ram.GetPermissionOutput, error)
	GetResourcePolicies(context.Context, *ram.GetResourcePoliciesInput, ...func(*ram.Options)) (*ram.GetResourcePoliciesOutput, error)
	GetResourceShareAssociations(context.Context, *ram.GetResourceShareAssociationsInput, ...func(*ram.Options)) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareInvitations(context.Context, *ram.GetResourceShareInvitationsInput, ...func(*ram.Options)) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShares(context.Context, *ram.GetResourceSharesInput, ...func(*ram.Options)) (*ram.GetResourceSharesOutput, error)
	ListPendingInvitationResources(context.Context, *ram.ListPendingInvitationResourcesInput, ...func(*ram.Options)) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPermissionVersions(context.Context, *ram.ListPermissionVersionsInput, ...func(*ram.Options)) (*ram.ListPermissionVersionsOutput, error)
	ListPermissions(context.Context, *ram.ListPermissionsInput, ...func(*ram.Options)) (*ram.ListPermissionsOutput, error)
	ListPrincipals(context.Context, *ram.ListPrincipalsInput, ...func(*ram.Options)) (*ram.ListPrincipalsOutput, error)
	ListResourceSharePermissions(context.Context, *ram.ListResourceSharePermissionsInput, ...func(*ram.Options)) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceTypes(context.Context, *ram.ListResourceTypesInput, ...func(*ram.Options)) (*ram.ListResourceTypesOutput, error)
	ListResources(context.Context, *ram.ListResourcesInput, ...func(*ram.Options)) (*ram.ListResourcesOutput, error)
}

type SecurityhubClient interface {
	DescribeActionTargets(context.Context, *securityhub.DescribeActionTargetsInput, ...func(*securityhub.Options)) (*securityhub.DescribeActionTargetsOutput, error)
	DescribeHub(context.Context, *securityhub.DescribeHubInput, ...func(*securityhub.Options)) (*securityhub.DescribeHubOutput, error)
	DescribeOrganizationConfiguration(context.Context, *securityhub.DescribeOrganizationConfigurationInput, ...func(*securityhub.Options)) (*securityhub.DescribeOrganizationConfigurationOutput, error)
	DescribeProducts(context.Context, *securityhub.DescribeProductsInput, ...func(*securityhub.Options)) (*securityhub.DescribeProductsOutput, error)
	DescribeStandards(context.Context, *securityhub.DescribeStandardsInput, ...func(*securityhub.Options)) (*securityhub.DescribeStandardsOutput, error)
	DescribeStandardsControls(context.Context, *securityhub.DescribeStandardsControlsInput, ...func(*securityhub.Options)) (*securityhub.DescribeStandardsControlsOutput, error)
	GetAdministratorAccount(context.Context, *securityhub.GetAdministratorAccountInput, ...func(*securityhub.Options)) (*securityhub.GetAdministratorAccountOutput, error)
	GetEnabledStandards(context.Context, *securityhub.GetEnabledStandardsInput, ...func(*securityhub.Options)) (*securityhub.GetEnabledStandardsOutput, error)
	GetFindingAggregator(context.Context, *securityhub.GetFindingAggregatorInput, ...func(*securityhub.Options)) (*securityhub.GetFindingAggregatorOutput, error)
	GetFindings(context.Context, *securityhub.GetFindingsInput, ...func(*securityhub.Options)) (*securityhub.GetFindingsOutput, error)
	GetInsightResults(context.Context, *securityhub.GetInsightResultsInput, ...func(*securityhub.Options)) (*securityhub.GetInsightResultsOutput, error)
	GetInsights(context.Context, *securityhub.GetInsightsInput, ...func(*securityhub.Options)) (*securityhub.GetInsightsOutput, error)
	GetInvitationsCount(context.Context, *securityhub.GetInvitationsCountInput, ...func(*securityhub.Options)) (*securityhub.GetInvitationsCountOutput, error)
	GetMasterAccount(context.Context, *securityhub.GetMasterAccountInput, ...func(*securityhub.Options)) (*securityhub.GetMasterAccountOutput, error)
	GetMembers(context.Context, *securityhub.GetMembersInput, ...func(*securityhub.Options)) (*securityhub.GetMembersOutput, error)
	ListEnabledProductsForImport(context.Context, *securityhub.ListEnabledProductsForImportInput, ...func(*securityhub.Options)) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListFindingAggregators(context.Context, *securityhub.ListFindingAggregatorsInput, ...func(*securityhub.Options)) (*securityhub.ListFindingAggregatorsOutput, error)
	ListInvitations(context.Context, *securityhub.ListInvitationsInput, ...func(*securityhub.Options)) (*securityhub.ListInvitationsOutput, error)
	ListMembers(context.Context, *securityhub.ListMembersInput, ...func(*securityhub.Options)) (*securityhub.ListMembersOutput, error)
	ListOrganizationAdminAccounts(context.Context, *securityhub.ListOrganizationAdminAccountsInput, ...func(*securityhub.Options)) (*securityhub.ListOrganizationAdminAccountsOutput, error)
	ListTagsForResource(context.Context, *securityhub.ListTagsForResourceInput, ...func(*securityhub.Options)) (*securityhub.ListTagsForResourceOutput, error)
}

type SchedulerClient interface {
	GetSchedule(context.Context, *scheduler.GetScheduleInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)
	GetScheduleGroup(context.Context, *scheduler.GetScheduleGroupInput, ...func(*scheduler.Options)) (*scheduler.GetScheduleGroupOutput, error)
	ListScheduleGroups(context.Context, *scheduler.ListScheduleGroupsInput, ...func(*scheduler.Options)) (*scheduler.ListScheduleGroupsOutput, error)
	ListSchedules(context.Context, *scheduler.ListSchedulesInput, ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)
	ListTagsForResource(context.Context, *scheduler.ListTagsForResourceInput, ...func(*scheduler.Options)) (*scheduler.ListTagsForResourceOutput, error)
}

type SesClient interface {
	DescribeActiveReceiptRuleSet(context.Context, *ses.DescribeActiveReceiptRuleSetInput, ...func(*ses.Options)) (*ses.DescribeActiveReceiptRuleSetOutput, error)
	DescribeConfigurationSet(context.Context, *ses.DescribeConfigurationSetInput, ...func(*ses.Options)) (*ses.DescribeConfigurationSetOutput, error)
	DescribeReceiptRule(context.Context, *ses.DescribeReceiptRuleInput, ...func(*ses.Options)) (*ses.DescribeReceiptRuleOutput, error)
	DescribeReceiptRuleSet(context.Context, *ses.DescribeReceiptRuleSetInput, ...func(*ses.Options)) (*ses.DescribeReceiptRuleSetOutput, error)
	GetAccountSendingEnabled(context.Context, *ses.GetAccountSendingEnabledInput, ...func(*ses.Options)) (*ses.GetAccountSendingEnabledOutput, error)
	GetCustomVerificationEmailTemplate(context.Context, *ses.GetCustomVerificationEmailTemplateInput, ...func(*ses.Options)) (*ses.GetCustomVerificationEmailTemplateOutput, error)
	GetIdentityDkimAttributes(context.Context, *ses.GetIdentityDkimAttributesInput, ...func(*ses.Options)) (*ses.GetIdentityDkimAttributesOutput, error)
	GetIdentityMailFromDomainAttributes(context.Context, *ses.GetIdentityMailFromDomainAttributesInput, ...func(*ses.Options)) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
	GetIdentityNotificationAttributes(context.Context, *ses.GetIdentityNotificationAttributesInput, ...func(*ses.Options)) (*ses.GetIdentityNotificationAttributesOutput, error)
	GetIdentityPolicies(context.Context, *ses.GetIdentityPoliciesInput, ...func(*ses.Options)) (*ses.GetIdentityPoliciesOutput, error)
	GetIdentityVerificationAttributes(context.Context, *ses.GetIdentityVerificationAttributesInput, ...func(*ses.Options)) (*ses.GetIdentityVerificationAttributesOutput, error)
	GetSendQuota(context.Context, *ses.GetSendQuotaInput, ...func(*ses.Options)) (*ses.GetSendQuotaOutput, error)
	GetSendStatistics(context.Context, *ses.GetSendStatisticsInput, ...func(*ses.Options)) (*ses.GetSendStatisticsOutput, error)
	GetTemplate(context.Context, *ses.GetTemplateInput, ...func(*ses.Options)) (*ses.GetTemplateOutput, error)
	ListConfigurationSets(context.Context, *ses.ListConfigurationSetsInput, ...func(*ses.Options)) (*ses.ListConfigurationSetsOutput, error)
	ListCustomVerificationEmailTemplates(context.Context, *ses.ListCustomVerificationEmailTemplatesInput, ...func(*ses.Options)) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
	ListIdentities(context.Context, *ses.ListIdentitiesInput, ...func(*ses.Options)) (*ses.ListIdentitiesOutput, error)
	ListIdentityPolicies(context.Context, *ses.ListIdentityPoliciesInput, ...func(*ses.Options)) (*ses.ListIdentityPoliciesOutput, error)
	ListReceiptFilters(context.Context, *ses.ListReceiptFiltersInput, ...func(*ses.Options)) (*ses.ListReceiptFiltersOutput, error)
	ListReceiptRuleSets(context.Context, *ses.ListReceiptRuleSetsInput, ...func(*ses.Options)) (*ses.ListReceiptRuleSetsOutput, error)
	ListTemplates(context.Context, *ses.ListTemplatesInput, ...func(*ses.Options)) (*ses.ListTemplatesOutput, error)
	ListVerifiedEmailAddresses(context.Context, *ses.ListVerifiedEmailAddressesInput, ...func(*ses.Options)) (*ses.ListVerifiedEmailAddressesOutput, error)
}

type SfnClient interface {
	DescribeActivity(context.Context, *sfn.DescribeActivityInput, ...func(*sfn.Options)) (*sfn.DescribeActivityOutput, error)
	DescribeExecution(context.Context, *sfn.DescribeExecutionInput, ...func(*sfn.Options)) (*sfn.DescribeExecutionOutput, error)
	DescribeMapRun(context.Context, *sfn.DescribeMapRunInput, ...func(*sfn.Options)) (*sfn.DescribeMapRunOutput, error)
	DescribeStateMachine(context.Context, *sfn.DescribeStateMachineInput, ...func(*sfn.Options)) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineForExecution(context.Context, *sfn.DescribeStateMachineForExecutionInput, ...func(*sfn.Options)) (*sfn.DescribeStateMachineForExecutionOutput, error)
	GetActivityTask(context.Context, *sfn.GetActivityTaskInput, ...func(*sfn.Options)) (*sfn.GetActivityTaskOutput, error)
	GetExecutionHistory(context.Context, *sfn.GetExecutionHistoryInput, ...func(*sfn.Options)) (*sfn.GetExecutionHistoryOutput, error)
	ListActivities(context.Context, *sfn.ListActivitiesInput, ...func(*sfn.Options)) (*sfn.ListActivitiesOutput, error)
	ListExecutions(context.Context, *sfn.ListExecutionsInput, ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error)
	ListMapRuns(context.Context, *sfn.ListMapRunsInput, ...func(*sfn.Options)) (*sfn.ListMapRunsOutput, error)
	ListStateMachines(context.Context, *sfn.ListStateMachinesInput, ...func(*sfn.Options)) (*sfn.ListStateMachinesOutput, error)
	ListTagsForResource(context.Context, *sfn.ListTagsForResourceInput, ...func(*sfn.Options)) (*sfn.ListTagsForResourceOutput, error)
}

type TimestreamwriteClient interface {
	DescribeDatabase(context.Context, *timestreamwrite.DescribeDatabaseInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeDatabaseOutput, error)
	DescribeEndpoints(context.Context, *timestreamwrite.DescribeEndpointsInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeEndpointsOutput, error)
	DescribeTable(context.Context, *timestreamwrite.DescribeTableInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeTableOutput, error)
	ListDatabases(context.Context, *timestreamwrite.ListDatabasesInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListDatabasesOutput, error)
	ListTables(context.Context, *timestreamwrite.ListTablesInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTablesOutput, error)
	ListTagsForResource(context.Context, *timestreamwrite.ListTagsForResourceInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTagsForResourceOutput, error)
}

type KafkaClient interface {
	DescribeCluster(context.Context, *kafka.DescribeClusterInput, ...func(*kafka.Options)) (*kafka.DescribeClusterOutput, error)
	DescribeClusterOperation(context.Context, *kafka.DescribeClusterOperationInput, ...func(*kafka.Options)) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterV2(context.Context, *kafka.DescribeClusterV2Input, ...func(*kafka.Options)) (*kafka.DescribeClusterV2Output, error)
	DescribeConfiguration(context.Context, *kafka.DescribeConfigurationInput, ...func(*kafka.Options)) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationRevision(context.Context, *kafka.DescribeConfigurationRevisionInput, ...func(*kafka.Options)) (*kafka.DescribeConfigurationRevisionOutput, error)
	GetBootstrapBrokers(context.Context, *kafka.GetBootstrapBrokersInput, ...func(*kafka.Options)) (*kafka.GetBootstrapBrokersOutput, error)
	GetCompatibleKafkaVersions(context.Context, *kafka.GetCompatibleKafkaVersionsInput, ...func(*kafka.Options)) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	ListClusterOperations(context.Context, *kafka.ListClusterOperationsInput, ...func(*kafka.Options)) (*kafka.ListClusterOperationsOutput, error)
	ListClusters(context.Context, *kafka.ListClustersInput, ...func(*kafka.Options)) (*kafka.ListClustersOutput, error)
	ListClustersV2(context.Context, *kafka.ListClustersV2Input, ...func(*kafka.Options)) (*kafka.ListClustersV2Output, error)
	ListConfigurationRevisions(context.Context, *kafka.ListConfigurationRevisionsInput, ...func(*kafka.Options)) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurations(context.Context, *kafka.ListConfigurationsInput, ...func(*kafka.Options)) (*kafka.ListConfigurationsOutput, error)
	ListKafkaVersions(context.Context, *kafka.ListKafkaVersionsInput, ...func(*kafka.Options)) (*kafka.ListKafkaVersionsOutput, error)
	ListNodes(context.Context, *kafka.ListNodesInput, ...func(*kafka.Options)) (*kafka.ListNodesOutput, error)
	ListScramSecrets(context.Context, *kafka.ListScramSecretsInput, ...func(*kafka.Options)) (*kafka.ListScramSecretsOutput, error)
	ListTagsForResource(context.Context, *kafka.ListTagsForResourceInput, ...func(*kafka.Options)) (*kafka.ListTagsForResourceOutput, error)
}
