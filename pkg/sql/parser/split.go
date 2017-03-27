// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Matt Jibson

package parser

import "bytes"

// Split represents a SPLIT statement. Only one of Table and Index can be set.
type Split struct {
	Table *NormalizableTableName
	Index *TableNameWithIndex
	// Each row contains values for the columns in the PK or index (or a prefix
	// of the columns).
	Rows *Select
}

// Format implements the NodeFormatter interface.
func (node *Split) Format(buf *bytes.Buffer, f FmtFlags) {
	buf.WriteString("ALTER ")
	if node.Index != nil {
		buf.WriteString("INDEX ")
		FormatNode(buf, f, node.Index)
	} else {
		buf.WriteString("TABLE ")
		FormatNode(buf, f, node.Table)
	}
	buf.WriteString(" SPLIT AT ")
	FormatNode(buf, f, node.Rows)
}
