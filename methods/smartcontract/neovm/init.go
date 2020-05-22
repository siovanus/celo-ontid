package neovm

import (
	"github.com/ontio/celo-ontid/methods/smartcontract/neovm/celo"
)

func RegisterNeoVm() {
	celo.RegisterCelo()
}
