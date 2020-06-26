package operations_test

import (
	"context"
	"testing"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	client := testutil.NewFixedClient(t)
	capability, err := client.Capability().Retrieve(context.Background(), &omise.RetrieveCapabilityParams{})

	r.Nil(t, err)
	r.Equal(t, "TH", capability.Country)
	r.Contains(t, capability.Banks, "kbank")
	r.True(t, capability.ZeroInterestInstallments)
	r.Equal(t, "card", capability.PaymentMethods[0].Name)
}

func TestCapability_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)

	capability, err := client.Capability().Retrieve(context.Background(), &omise.RetrieveCapabilityParams{})

	r.Nil(t, err)
	r.Equal(t, capability.Object, "capability")
	testutil.LogObj(t, capability)
}
