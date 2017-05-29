package errors

/**
*    Copyright (C) 2017 Ethan Frey
**/

import abci "github.com/tendermint/abci/types"

const (
	msgDecoding        = "Error decoding input"
	msgUnauthorized    = "Unauthorized"
	msgInvalidAddress  = "Invalid Address"
	msgInvalidCoins    = "Invalid Coins"
	msgInvalidSequence = "Invalid Sequence"
	msgNoInputs        = "No Input Coins"
	msgNoOutputs       = "No Output Coins"
)

func DecodingError() TMError {
	return New(msgDecoding, abci.CodeType_EncodingError)
}

func Unauthorized() TMError {
	return New(msgUnauthorized, abci.CodeType_Unauthorized)
}

func InvalidAddress() TMError {
	return New(msgInvalidAddress, abci.CodeType_BaseInvalidInput)
}

func InvalidCoins() TMError {
	return New(msgInvalidCoins, abci.CodeType_BaseInvalidInput)
}

func InvalidSequence() TMError {
	return New(msgInvalidSequence, abci.CodeType_BaseInvalidInput)
}

func NoInputs() TMError {
	return New(msgNoInputs, abci.CodeType_BaseInvalidInput)
}

func NoOutputs() TMError {
	return New(msgNoOutputs, abci.CodeType_BaseInvalidOutput)
}