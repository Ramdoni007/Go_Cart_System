package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Ramdoni007/Go_CartSystem/util"
	"github.com/stretchr/testify/require"
)

func CreateCartForUpdate(t *testing.T) Cart {
	arg := CreateCartParams{
		NameProduct: "LAPTOP ACE GAMEMAX 007",
		Quantity:    util.RandomQuantity(),
	}
	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	return cart
}

func CreateRandomCart(t *testing.T) Cart {
	arg := CreateCartParams{
		NameProduct: "Laptop ACE COLD 220078",
		Quantity:    util.RandomQuantity(),
	}
	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.NameProduct, cart.NameProduct)
	require.Equal(t, arg.Quantity, cart.Quantity)

	return cart

}

func TestCreateCart(t *testing.T) {
	arg := CreateCartParams{
		NameProduct: "Laptop ROG454YXXttt",
		Quantity:    100,
	}

	cart, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.NameProduct, cart.NameProduct)
	require.Equal(t, arg.Quantity, cart.Quantity)

}

func TestGetCart(t *testing.T) {
	arg := CreateCartParams{
		NameProduct: "ASUS 3454",
		Quantity:    22,
	}
	cart1, err := testQueris.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart1)
	cart2, err := testQueris.GetCart(context.Background(), cart1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

}

func TestUpdateCart(t *testing.T) {
	cart := CreateCartParams{
		NameProduct: "Laptop xyzxyz",
		Quantity:    util.RandomQuantity(),
	}

	arg := UpdateCartForQuantityParams{
		Quantity: cart.Quantity,
	}

	cart1, err := testQueris.CreateCart(context.Background(), cart)
	require.NoError(t, err)
	require.NotEmpty(t, cart1)
	cart2, err := testQueris.UpdateCartForQuantity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

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
