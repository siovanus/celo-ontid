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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/celo-ontid/common"
	sdk "github.com/ontio/ontology-go-sdk"
)

type regIdWithPublicKeyParam struct {
	OntId string
	Path  string
}

func regIdWithPublicKey(ontSdk *sdk.OntologySdk) bool {
	data, err := ioutil.ReadFile("./params/regIdWithPublicKey.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed: ", err)
		return false
	}
	regIdWithPublicKeyParam := new(regIdWithPublicKeyParam)
	err = json.Unmarshal(data, regIdWithPublicKeyParam)
	if err != nil {
		fmt.Println("json.Unmarshal failed: ", err)
		return false
	}
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, regIdWithPublicKeyParam.Path)
	if !ok {
		return false
	}
	ok = invokeRegIdWithPublicKey(ontSdk, user, regIdWithPublicKeyParam.OntId)
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

type bindCeloParam struct {
	OntId       string
	Index       uint32
	Path        string
	CeloAddress string
}

func bindCelo(ontSdk *sdk.OntologySdk) bool {
	data, err := ioutil.ReadFile("./params/bindCelo.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed: ", err)
		return false
	}
	bindCeloParam := new(bindCeloParam)
	err = json.Unmarshal(data, bindCeloParam)
	if err != nil {
		fmt.Println("json.Unmarshal failed: ", err)
		return false
	}
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, bindCeloParam.Path)
	if !ok {
		return false
	}
	celoAddress := ethcommon.HexToAddress(bindCeloParam.CeloAddress)
	ok = invokeBindCelo(ontSdk, user, bindCeloParam.OntId, bindCeloParam.Index, celoAddress[:])
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

type setCeloDefaultParam struct {
	OntId       string
	Index       uint32
	Path        string
	CeloAddress string
}

func setCeloDefault(ontSdk *sdk.OntologySdk) bool {
	data, err := ioutil.ReadFile("./params/setCeloDefault.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed: ", err)
		return false
	}
	setCeloDefaultParam := new(setCeloDefaultParam)
	err = json.Unmarshal(data, setCeloDefaultParam)
	if err != nil {
		fmt.Println("json.Unmarshal failed: ", err)
		return false
	}
	time.Sleep(1 * time.Second)
	user, ok := common.GetAccountByPassword(ontSdk, setCeloDefaultParam.Path)
	if !ok {
		return false
	}
	celoAddress := ethcommon.HexToAddress(setCeloDefaultParam.CeloAddress)
	ok = invokeSetCeloDefault(ontSdk, user, setCeloDefaultParam.OntId, setCeloDefaultParam.Index, celoAddress[:])
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

type getCeloDefaultParam struct {
	OntId string
}

func getCeloDefault(ontSdk *sdk.OntologySdk) bool {
	data, err := ioutil.ReadFile("./params/getCeloDefault.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed: ", err)
		return false
	}
	getCeloDefaultParam := new(getCeloDefaultParam)
	err = json.Unmarshal(data, getCeloDefaultParam)
	if err != nil {
		fmt.Println("json.Unmarshal failed: ", err)
		return false
	}
	ok := invokeGetCeloDefault(ontSdk, getCeloDefaultParam.OntId)
	if !ok {
		return false
	}
	return true
}
