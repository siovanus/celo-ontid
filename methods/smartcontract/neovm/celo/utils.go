/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package celo

import (
	"fmt"
	"github.com/ontio/celo-ontid/config"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)

func invokeBindCelo(ontSdk *sdk.OntologySdk, user *sdk.Account, ontId string, index uint32, celoAddress []byte) bool {
	method := "bindCelo"
	contractAddress, err := common.AddressFromHexString(config.DefConfig.ContractAddress)
	if err != nil {
		fmt.Println("invokeBindCelo, common.AddressFromHexString error :", err)
		return false
	}
	txHash, err := ontSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit, user, user,
		contractAddress, []interface{}{method, []interface{}{ontId, index, celoAddress}})
	if err != nil {
		fmt.Println("invokeBindCelo, NewNativeInvokeTransaction error :", err)
		return false
	}

	fmt.Println("invokeBindCelo txHash is :", txHash.ToHexString())
	return true
}
