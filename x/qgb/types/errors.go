package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrDuplicate                             = sdkerrors.Register(ModuleName, 2, "duplicate")
	ErrEmpty                                 = sdkerrors.Register(ModuleName, 6, "empty")
	ErrNoValidators                          = sdkerrors.Register(ModuleName, 12, "no bonded validators in active set")
	ErrInvalidValAddress                     = sdkerrors.Register(ModuleName, 13, "invalid validator address in current valset %v")
	ErrInvalidEVMAddress                     = sdkerrors.Register(ModuleName, 14, "discovered invalid EVM address stored for validator %v")
	ErrInvalidValset                         = sdkerrors.Register(ModuleName, 15, "generated invalid valset")
	ErrAttestationNotValsetRequest           = sdkerrors.Register(ModuleName, 16, "attestation is not a valset request")
	ErrAttestationNotFound                   = sdkerrors.Register(ModuleName, 18, "attestation not found")
	ErrNilAttestation                        = sdkerrors.Register(ModuleName, 22, "nil attestation")
	ErrUnmarshalllAttestation                = sdkerrors.Register(ModuleName, 26, "couldn't unmarshall attestation from store")
	ErrNonceHigherThanLatestAttestationNonce = sdkerrors.Register(ModuleName, 27, "the provided nonce is higher than the latest attestation nonce")
	ErrNoValsetBeforeNonceOne                = sdkerrors.Register(ModuleName, 28, "there is no valset before attestation nonce 1")
	ErrDataCommitmentNotGenerated			 = sdkerrors.Register(ModuleName, 29, "there is no data commitment has been generated")
	ErrDataCommitmentNotFound				 = sdkerrors.Register(ModuleName, 30, "data commitment not found")
)
