package db

import (
	"context"
	"database/sql"
	"github.com/Ramdoni007/Go_CartSystem/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateCartForUpdate(t *testing.T) Cart {
	arg := CreateCartParams{
		ID:          18,
		NameProduct: "LAPTOP ROG 454YZ",
		Quantity:    5,
	}
	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	return cart
}

func CreateRandomCart(t *testing.T) Cart {
	arg := CreateCartParams{
		ID:          util.RandomInt(1, 100),
		NameProduct: util.RandomString(12),
		Quantity:    util.RandomInt(1, 100),
	}
	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.ID, cart.ID)
	require.Equal(t, arg.NameProduct, cart.NameProduct)
	require.Equal(t, arg.Quantity, cart.Quantity)

	return cart

}

func TestCreateCart(t *testing.T) {
	arg := CreateCartParams{
		ID:          1,
		NameProduct: "Laptop ROG",
		Quantity:    2,
	}

	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.ID, cart.ID)
	require.Equal(t, arg.NameProduct, cart.NameProduct)
	require.Equal(t, arg.Quantity, cart.Quantity)

}

func TestGetCart(t *testing.T) {
	cart1 := CreateRandomCart(t)
	cart2, err := testQueris.GetCart(context.Background(), cart1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	require.Equal(t, cart1.ID, cart2.ID)
	require.Equal(t, cart1.NameProduct, cart2.NameProduct)
	require.Equal(t, cart1.Quantity, cart2.Quantity)

}

func TestUpdateCart(t *testing.T) {
	cart1 := CreateCartForUpdate(t)

	arg := UpdateCartParams{
		ID:          16,
		NameProduct: cart1.NameProduct,
		Quantity:    cart1.Quantity,
	}
	cart2, err := testQueris.UpdateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	require.Equal(t, cart1.ID, cart2.ID)
	require.Equal(t, cart1.NameProduct, cart2.NameProduct)
	require.Equal(t, cart1.Quantity, cart2.Quantity)
}

func TestDeletedCart(t *testing.T) {
	cart1 := CreateCartForUpdate(t)
	err := testQueris.DeleteCart(context.Background(), cart1.ID)
	require.NoError(t, err)

	cart2, err := testQueris.GetCart(context.Background(), cart1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cart2)
}
