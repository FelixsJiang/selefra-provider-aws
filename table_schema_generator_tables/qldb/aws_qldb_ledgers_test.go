package qldb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildLedgersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockQLDBClient(ctrl)

	ledger := types.LedgerSummary{}
	if err := faker.FakeObject(&ledger); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListLedgers(gomock.Any(), &qldb.ListLedgersInput{}, gomock.Any()).AnyTimes().Return(
		&qldb.ListLedgersOutput{Ledgers: []types.LedgerSummary{ledger}}, nil)

	var resource qldb.DescribeLedgerOutput
	if err := faker.FakeObject(&resource); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLedger(
		gomock.Any(),
		&qldb.DescribeLedgerInput{Name: ledger.Name},
		gomock.Any(),
	).AnyTimes().Return(
		&resource,
		nil,
	)

	tags := &qldb.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)

	s3 := types.JournalS3ExportDescription{}
	if err := faker.FakeObject(&s3); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListJournalS3ExportsForLedger(gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().Return(&qldb.ListJournalS3ExportsForLedgerOutput{
		JournalS3Exports: []types.JournalS3ExportDescription{s3},
	}, nil)

	ke := types.JournalKinesisStreamDescription{}
	if err := faker.FakeObject(&ke); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListJournalKinesisStreamsForLedger(gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().Return(&qldb.ListJournalKinesisStreamsForLedgerOutput{
		Streams: []types.JournalKinesisStreamDescription{ke},
	}, nil)

	return aws_client.AwsServices{QLDB: m}
}

func TestQldbLedgers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsQldbLedgersGenerator{}), buildLedgersMock, aws_client.TestOptions{})
}
