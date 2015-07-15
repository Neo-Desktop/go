package main

import (
	_ "gopkg.in/guregu/null.v2"
	"gopkg.in/guregu/null.v2/zero"
)

type User struct {
	ID          zero.Int    `db:"id"`
	Username    zero.String `db:"username"`
	Name        zero.String `db:"name"`
	Password    zero.String `db:"password" json:"-" xml:"-"`
	Address     zero.String `db:"address"`
	City        zero.String `db:"city"`
	State       zero.String `db:"state"`
	Zip         zero.String `db:"zip"`
	Country     zero.String `db:"country"`
	CreatedDate zero.String `db:"createdDate"`
	UpdatedDate zero.String `db:"updatedDate"`
}
