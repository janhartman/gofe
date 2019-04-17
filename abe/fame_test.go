/*
 * Copyright (c) 2018 XLAB d.o.o
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package abe_test

import (
	"github.com/fentec-project/gofe"
	"strconv"
	"testing"

	"github.com/fentec-project/gofe/abe"
	"github.com/stretchr/testify/assert"
)

func TestFAME(t *testing.T) {
	l, _, _ := gofe.GetParams()

	// create a new FAME struct with the universe of attributes
	// denoted by integer
	a := abe.NewFAME()

	// generate a public key and a secret key for the scheme
	pubKey, secKey, err := a.GenerateMasterKeys()
	if err != nil {
		t.Fatalf("Failed to generate master keys: %v", err)
	}

	// create a message to be encrypted
	msg := "Attack at dawn!"

	boolExp := ""
	for i := 0; i < l; i++ {
		boolExp += strconv.Itoa(i)
		if i < l -1 {
			boolExp += " AND "
		}
	}

	// create a msp struct out of a boolean expression representing the
	// policy specifying which attributes are needed to decrypt the ciphertext;
	// note that safety of the encryption is only proved if the mapping
	// msp.RowToAttrib from the rows of msp.Mat to attributes is injective, i.e.
	// only boolean expressions in which each attribute appears at most once
	// are allowed - if expressions with multiple appearances of an attribute
	// are needed, then this attribute can be split into more sub-attributes
	msp, err := abe.BooleanToMSP(boolExp, false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	// encrypt the message msg with the decryption policy specified by the
	// msp structure
	cipher, err := a.Encrypt(msg, msp, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gamma := make([]int, l)
	for i := 0; i < l; i++ {
		gamma[i] = i
	}

	// generate keys for decryption for an entity with
	// attributes gamma
	keys, err := a.GenerateAttribKeys(gamma, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes
	msgCheck, err := a.Decrypt(cipher, keys, pubKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, msgCheck)

}
