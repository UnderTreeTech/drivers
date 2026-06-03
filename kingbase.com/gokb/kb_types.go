package gokb

import (
	"database/sql"

	"github.com/UnderTreeTech/drivers/kingbase.com/gokb/oid"
)

type bindStruct struct {
	out    sql.Out
	isOut  bool
	isBoth bool
	typ    oid.Oid
}

type CursorString struct {
	CursorName string
}
