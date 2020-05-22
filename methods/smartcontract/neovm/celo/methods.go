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
