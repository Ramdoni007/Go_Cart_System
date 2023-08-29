package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore(testDB)

	// First will crateCart data in Store in Database
	arg := CreateCartParams{
		ID:          17,
		NameProduct: "Laptop ThinkPad",
		Quantity:    2,
	}

	cart, err := store.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.ID, cart.ID)
	require.Equal(t, arg.NameProduct, cart.NameProduct)
	require.Equal(t, arg.Quantity, cart.Quantity)

	// second will getQuantity Data in table quantity in database

	cart2, err := store.GetQuantity(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	// third will getQuantity Data in table quantity in database after experiencing a change in quantity
	cart3, err := store.GetQuantityForUpdate(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, cart3)

}
