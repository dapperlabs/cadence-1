/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolExpression_MarshalJSON(t *testing.T) {

	expr := &BoolExpression{
		Value: false,
		Range: Range{
			StartPos: Position{Offset: 1, Line: 2, Column: 3},
			EndPos:   Position{Offset: 4, Line: 5, Column: 6},
		},
	}

	actual, err := json.Marshal(expr)
	require.NoError(t, err)

	assert.JSONEq(t,
		`
        {
            "Type": "BoolExpression",
            "Value": false,
            "StartPos": {"Offset": 1, "Line": 2, "Column": 3}, 
            "EndPos": {"Offset": 4, "Line": 5, "Column": 6}
        }
        `,
		string(actual),
	)
}

func TestNilExpression_MarshalJSON(t *testing.T) {

	expr := &NilExpression{
		Pos: Position{Offset: 1, Line: 2, Column: 3},
	}

	actual, err := json.Marshal(expr)
	require.NoError(t, err)

	assert.JSONEq(t,
		`
        {
            "Type": "NilExpression",
            "StartPos": {"Offset": 1, "Line": 2, "Column": 3}, 
            "EndPos": {"Offset": 3, "Line": 2, "Column": 5}
        }
        `,
		string(actual),
	)

}