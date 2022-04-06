package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCoinSend = "coin_send"

var _ sdk.Msg = &MsgCoinSend{}

func NewMsgCoinSend(creator string, receivers []Receiver) *MsgCoinSend {
	return &MsgCoinSend{
		Creator:   creator,
		Receivers: receivers,
	}
}

func (msg *MsgCoinSend) Route() string {
	return RouterKey
}

func (msg *MsgCoinSend) Type() string {
	return TypeMsgCoinSend
}

func (msg *MsgCoinSend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCoinSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCoinSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
