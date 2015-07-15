package main

import (
	"gopkg.in/guregu/null.v2"
)

type Post struct {
	ID          null.Int    `db:"id" json:"UserID" xml:"UserID,omitempty"`
	Username    null.String `db:"username" json:"UserName" xml:"UserName,omitempty"`
	Password    null.String `db:"password" json:"-" xml:"-"`
	Name        null.String `db:"name" json:"Name" xml:"Name,omitempty"`
	Address     null.String `db:"address" json:"Address" xml:"Address,omitempty"`
	City        null.String `db:"city" json:"City" xml:"city,omitempty"`
	State       null.String `db:"state" json:"State" xml:"State,omitempty"`
	Zip         null.String `db:"zip" json:"ZipCode" xml:"ZipCode,omitempty"`
	Country     null.String `db:"country" json:"Country" xml:"Country,omitempty"`
	CreatedDate null.String `db:"createdDate" json:"CreatedDate" xml:"CreatedDate,omitempty"`
	UpdatedDate null.String `db:"updatedDate" json:"UpdatedDate" xml:"UpdatedDate,omitempty"`
}
