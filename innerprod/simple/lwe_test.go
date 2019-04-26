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

package simple_test

import (
	"github.com/fentec-project/gofe"
	"math/big"
	"testing"

	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/simple"
	"github.com/fentec-project/gofe/sample"
	"github.com/stretchr/testify/assert"
)

func TestSimple_LWE(t *testing.T) {

	n, l,_b := gofe.GetParams()
	b := big.NewInt(int64(_b))

	x, y, xy := testVectorData(l, b, b)

	simpleLWE, err := simple.NewLWE(l, b, b, n, 128)
	assert.NoError(t, err)

	SK, err := simpleLWE.GenerateSecretKey()
	assert.NoError(t, err)

	PK, err := simpleLWE.GeneratePublicKey(SK)
	assert.NoError(t, err)

	skY, err := simpleLWE.DeriveKey(y, SK)
	assert.NoError(t, err)

	cipher, err := simpleLWE.Encrypt(x, PK)
	assert.NoError(t, err)

	xyDecrypted, err := simpleLWE.Decrypt(cipher, skY, y)
	assert.NoError(t, err)
	assert.Equal(t, xy.Cmp(xyDecrypted), 0, "obtained incorrect inner product")
}

// testVectorData returns random vectors x, y, each containing
// elements up to the respective bound.
// It also returns the dot product of the vectors.
func testVectorData(len int, boundX, boundY *big.Int) (data.Vector, data.Vector, *big.Int) {
	samplerX := sample.NewUniformRange(new(big.Int).Neg(boundX), boundX)
	samplerY := sample.NewUniformRange(new(big.Int).Neg(boundY), boundY)
	x, _ := data.NewRandomVector(len, samplerX)
	y, _ := data.NewRandomVector(len, samplerY)
	xy, _ := x.Dot(y)

	return x, y, xy
}
