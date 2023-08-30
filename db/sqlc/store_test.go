package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(testDB)

	// First will crateCart data in Store in Database
	cart := CreateRandomCart(t)

	n := 1

	// run concurrent go routine
	errs := make(chan error)
	results := make(chan QuantityTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.QuantityTx(context.Background(), QuantityTxParams{
				NameProduct: cart.NameProduct,
				Quantity:    cart.Quantity,
			})
			errs <- err
			results <- result
		}()
	}
	// check result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check Result
		Quantity := result.ID
		require.NotEmpty(t, Quantity)
		require.Equal(t, cart.ID, Quantity.ID)
		require.Equal(t, cart.NameProduct, Quantity.NameProduct)
		require.Equal(t, cart.Quantity, Quantity.Quantity)

		_, err = store.GetCart(context.Background(), Quantity.ID)
		require.NoError(t, err)

	}
	// second will getQuantity Data in table quantity in database

	// third will getQuantity Data in table quantity in database after experiencing a change in quantity

}
