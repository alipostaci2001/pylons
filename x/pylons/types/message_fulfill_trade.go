package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFulfillTrade{}

func NewMsgFulfillTrade(creator string, id string, items []ItemRef) *MsgFulfillTrade {
	return &MsgFulfillTrade{
		Creator: creator,
		ID:      id,
		Items:   items,
	}
}

func (msg *MsgFulfillTrade) Route() string {
	return RouterKey
}

func (msg *MsgFulfillTrade) Type() string {
	return "FulfillTrade"
}

func (msg *MsgFulfillTrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFulfillTrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFulfillTrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
