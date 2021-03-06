package core_types

import (
	amino "github.com/tendermint/go-amino-x"
	"github.com/tendermint/classic/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
