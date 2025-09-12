package v5

import (
	_ "github.com/mobilemindtech/go-payments/api"
	pagarme "github.com/mobilemindtech/go-payments/pagarme/v5"
	gopayments "github.com/mobilemindtech/go-payments/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test -v  github.com/mobilemindtech/go-payments/tests/pagarme/v5 -run TestPagarmev5InvoiceList
func TestPagarmev5InvoiceList(t *testing.T) {

	Pagarme := pagarme.
		NewPagarmeInvoice("pt-BR",
			pagarme.NewAuthentication(gopayments.SecretKey, gopayments.PublicKey))

	Pagarme.DebugOn()

	result := Pagarme.List(pagarme.NewInvoiceQuery())

	assert.False(t, result.IsLeft())
	assert.True(t, result.Right().NonEmpty())
}
