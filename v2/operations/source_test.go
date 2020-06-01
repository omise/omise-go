package operations_test

import (
	"testing"

	omise "github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	. "github.com/omise/omise-go/v2/operations"
	r "github.com/stretchr/testify/require"
)

func TestSource(t *testing.T) {
	const (
		SourceID = "src_test_5a444nhh27tlyv81u40"
	)
	client := testutil.NewFixedClient(t)

	source := &omise.Source{}
	client.MustDo(source, &RetrieveSource{SourceID: SourceID})
	r.Equal(t, SourceID, source.ID)

	source = &omise.Source{}
	client.MustDo(source, &CreateSource{})
	r.Equal(t, SourceID, source.ID)
	r.Equal(t, omise.SourceBillPaymentTescoLotus, source.Type)
	r.Equal(t, "qr", source.ScannableCode.Type)
	r.Equal(t, "https://omise.co/barcode.jpg", source.ScannableCode.Image.DownloadURI)
}

func TestSource_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	var source omise.Source
	client.MustDo(&source, &CreateSource{
		Type:     omise.SourceAlipay,
		Amount:   10000,
		Currency: "thb",
	})

	r.Equal(t, omise.SourceAlipay, source.Type)
	t.Logf("%#v\n", source)
}
