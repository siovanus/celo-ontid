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

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/celo-ontid/config"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
)

func invokeRegIdWithPublicKey(ontSdk *sdk.OntologySdk, user *sdk.Account, ontId string) bool {
	method := "regIDWithPublicKey"
	contractAddress, err := common.AddressFromHexString(config.DefConfig.ContractAddress)
	if err != nil {
		fmt.Println("invokeRegIdWithPublicKey, common.AddressFromHexString error :", err)
		return false
	}
	pk := keypair.SerializePublicKey(user.PublicKey)
	txHash, err := ontSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit, user, user,
		contractAddress, []interface{}{method, []interface{}{ontId, pk}})
	if err != nil {
		fmt.Println("invokeRegIdWithPublicKey, InvokeNeoVMContract error :", err)
		return false
	}
	fmt.Println("invokeRegIdWithPublicKey txHash is :", txHash.ToHexString())
	return true
}

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
		fmt.Println("invokeBindCelo, InvokeNeoVMContract error :", err)
		return false
	}
	fmt.Println("invokeBindCelo txHash is :", txHash.ToHexString())
	return true
}

func invokeSetCeloDefault(ontSdk *sdk.OntologySdk, user *sdk.Account, ontId string, index uint32, celoAddress []byte) bool {
	method := "setCeloDefault"
	contractAddress, err := common.AddressFromHexString(config.DefConfig.ContractAddress)
	if err != nil {
		fmt.Println("invokeSetCeloDefault, common.AddressFromHexString error :", err)
		return false
	}
	txHash, err := ontSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit, user, user,
		contractAddress, []interface{}{method, []interface{}{ontId, index, celoAddress}})
	if err != nil {
		fmt.Println("invokeSetCeloDefault, InvokeNeoVMContract error :", err)
		return false
	}
	fmt.Println("invokeSetCeloDefault txHash is :", txHash.ToHexString())
	return true
}

func invokeGetCeloDefault(ontSdk *sdk.OntologySdk, ontId string) bool {
	method := "getCeloDefault"
	contractAddress, err := common.AddressFromHexString(config.DefConfig.ContractAddress)
	if err != nil {
		fmt.Println("invokeGetCeloDefault, common.AddressFromHexString error :", err)
		return false
	}
	res, err := ontSdk.NeoVM.PreExecInvokeNeoVMContract(contractAddress, []interface{}{method, []interface{}{ontId}})
	if err != nil {
		fmt.Println("invokeGetCeloDefault, PreExecInvokeNeoVMContract error :", err)
		return false
	}
	r, err := res.Result.ToByteArray()
	if err != nil {
		fmt.Println("invokeGetCeloDefault, res.Result.ToByteArray error :", err)
		return false
	}
	fmt.Println("invokeGetCeloDefault txHash is :", ethcommon.BytesToAddress(r).Hex())
	return true
}
