package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/step3 module sentinel errors
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1000, "name does not exist")
)
