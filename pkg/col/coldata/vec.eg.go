// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package coldata

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/cockroach/pkg/util/json"
	"github.com/cockroachdb/errors"
)

// Workaround for bazel auto-generated code. goimports does not automatically
// pick up the right packages when run within the bazel sandbox.
var (
	_ = typeconv.DatumVecCanonicalTypeFamily
	_ apd.Context
	_ duration.Duration
	_ json.JSON
	_ = colexecerror.InternalError
	_ = errors.AssertionFailedf
)

// TypedVecs represents a slice of Vecs that have been converted into the typed
// columns. The idea is that every Vec is stored both in Vecs slice as well as
// in the typed slice, in order. Components that know the type of the vector
// they are working with can then access the typed column directly, avoiding
// expensive type casts.
type TypedVecs struct {
	Vecs  []Vec
	Nulls []*Nulls

	// Fields below need to be accessed by an index mapped via ColsMap.
	BoolCols      []Bools
	BytesCols     []*Bytes
	DecimalCols   []Decimals
	Int16Cols     []Int16s
	Int32Cols     []Int32s
	Int64Cols     []Int64s
	Float64Cols   []Float64s
	TimestampCols []Times
	IntervalCols  []Durations
	JSONCols      []*JSONs
	DatumCols     []DatumVec
	// ColsMap contains the positions of the corresponding vectors in the slice
	// for the same types. For example, if we have a batch with
	//   types = [Int64, Int64, Bool, Bytes, Bool, Int64],
	// then ColsMap will be
	//                      [0, 1, 0, 0, 1, 2]
	//                       ^  ^  ^  ^  ^  ^
	//                       |  |  |  |  |  |
	//                       |  |  |  |  |  3rd among all Int64's
	//                       |  |  |  |  2nd among all Bool's
	//                       |  |  |  1st among all Bytes's
	//                       |  |  1st among all Bool's
	//                       |  2nd among all Int64's
	//                       1st among all Int64's
	ColsMap []int
}

// SetBatch updates TypedVecs to represent all vectors from batch.
func (v *TypedVecs) SetBatch(batch Batch) {
	v.Vecs = batch.ColVecs()
	if cap(v.Nulls) < len(v.Vecs) {
		v.Nulls = make([]*Nulls, len(v.Vecs))
		v.ColsMap = make([]int, len(v.Vecs))
	} else {
		v.Nulls = v.Nulls[:len(v.Vecs)]
		v.ColsMap = v.ColsMap[:len(v.Vecs)]
	}
	v.BoolCols = v.BoolCols[:0]
	v.BytesCols = v.BytesCols[:0]
	v.DecimalCols = v.DecimalCols[:0]
	v.Int16Cols = v.Int16Cols[:0]
	v.Int32Cols = v.Int32Cols[:0]
	v.Int64Cols = v.Int64Cols[:0]
	v.Float64Cols = v.Float64Cols[:0]
	v.TimestampCols = v.TimestampCols[:0]
	v.IntervalCols = v.IntervalCols[:0]
	v.JSONCols = v.JSONCols[:0]
	v.DatumCols = v.DatumCols[:0]
	for i, vec := range v.Vecs {
		v.Nulls[i] = vec.Nulls()
		switch vec.CanonicalTypeFamily() {
		case types.BoolFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.BoolCols)
				v.BoolCols = append(v.BoolCols, vec.Bool())
			}
		case types.BytesFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.BytesCols)
				v.BytesCols = append(v.BytesCols, vec.Bytes())
			}
		case types.DecimalFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.DecimalCols)
				v.DecimalCols = append(v.DecimalCols, vec.Decimal())
			}
		case types.IntFamily:
			switch vec.Type().Width() {
			case 16:
				v.ColsMap[i] = len(v.Int16Cols)
				v.Int16Cols = append(v.Int16Cols, vec.Int16())
			case 32:
				v.ColsMap[i] = len(v.Int32Cols)
				v.Int32Cols = append(v.Int32Cols, vec.Int32())
			case -1:
			default:
				v.ColsMap[i] = len(v.Int64Cols)
				v.Int64Cols = append(v.Int64Cols, vec.Int64())
			}
		case types.FloatFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.Float64Cols)
				v.Float64Cols = append(v.Float64Cols, vec.Float64())
			}
		case types.TimestampTZFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.TimestampCols)
				v.TimestampCols = append(v.TimestampCols, vec.Timestamp())
			}
		case types.IntervalFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.IntervalCols)
				v.IntervalCols = append(v.IntervalCols, vec.Interval())
			}
		case types.JsonFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.JSONCols)
				v.JSONCols = append(v.JSONCols, vec.JSON())
			}
		case typeconv.DatumVecCanonicalTypeFamily:
			switch vec.Type().Width() {
			case -1:
			default:
				v.ColsMap[i] = len(v.DatumCols)
				v.DatumCols = append(v.DatumCols, vec.Datum())
			}
		default:
			colexecerror.InternalError(errors.AssertionFailedf("unhandled type %s", vec.Type()))
		}
	}
}

// Reset performs a deep reset of v while keeping the references to the slices.
func (v *TypedVecs) Reset() {
	v.Vecs = nil
	for i := range v.Nulls {
		v.Nulls[i] = nil
	}
	for i := range v.BoolCols {
		v.BoolCols[i] = nil
	}
	for i := range v.BytesCols {
		v.BytesCols[i] = nil
	}
	for i := range v.DecimalCols {
		v.DecimalCols[i] = nil
	}
	for i := range v.Int16Cols {
		v.Int16Cols[i] = nil
	}
	for i := range v.Int32Cols {
		v.Int32Cols[i] = nil
	}
	for i := range v.Int64Cols {
		v.Int64Cols[i] = nil
	}
	for i := range v.Float64Cols {
		v.Float64Cols[i] = nil
	}
	for i := range v.TimestampCols {
		v.TimestampCols[i] = nil
	}
	for i := range v.IntervalCols {
		v.IntervalCols[i] = nil
	}
	for i := range v.JSONCols {
		v.JSONCols[i] = nil
	}
	for i := range v.DatumCols {
		v.DatumCols[i] = nil
	}
}

func (m *memColumn) Append(args SliceArgs) {
	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bool()
			toCol := m.Bool()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bytes()
			toCol := m.Bytes()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol.appendSliceWithSel(fromCol, args.DestIdx, sel)
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Decimal()
			toCol := m.Decimal()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				{
					__desiredCap := args.DestIdx + args.SrcEndIdx - args.SrcStartIdx
					if cap(toCol) >= __desiredCap {
						toCol = toCol[:__desiredCap]
					} else {
						__prevCap := cap(toCol)
						__capToAllocate := __desiredCap
						if __capToAllocate < 2*__prevCap {
							__capToAllocate = 2 * __prevCap
						}
						__new_slice := make([]apd.Decimal, __desiredCap, __capToAllocate)
						copy(__new_slice, toCol[:args.DestIdx])
						toCol = __new_slice
					}
					__src_slice := fromCol[args.SrcStartIdx:args.SrcEndIdx]
					__dst_slice := toCol[args.DestIdx:]
					_ = __dst_slice[len(__src_slice)-1]
					for __i := range __src_slice {
						//gcassert:bce
						__dst_slice[__i].Set(&__src_slice[__i])
					}
				}
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, apd.Decimal{})
					toCol[len(toCol)-1].Set(&val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			fromCol := args.Src.Int16()
			toCol := m.Int16()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		case 32:
			fromCol := args.Src.Int32()
			toCol := m.Int32()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		case -1:
		default:
			fromCol := args.Src.Int64()
			toCol := m.Int64()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Float64()
			toCol := m.Float64()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Timestamp()
			toCol := m.Timestamp()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Interval()
			toCol := m.Interval()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol = append(toCol[:args.DestIdx], fromCol[args.SrcStartIdx:args.SrcEndIdx]...)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol = append(toCol, val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case types.JsonFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.JSON()
			toCol := m.JSON()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol.appendSliceWithSel(fromCol, args.DestIdx, sel)
			}
			m.nulls.set(args)
			m.col = toCol
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Datum()
			toCol := m.Datum()
			// NOTE: it is unfortunate that we always append whole slice without paying
			// attention to whether the values are NULL. However, if we do start paying
			// attention, the performance suffers dramatically, so we choose to copy
			// over "actual" as well as "garbage" values.
			if args.Sel == nil {
				toCol.AppendSlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			} else {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				toCol = toCol.Window(0, args.DestIdx)
				for _, selIdx := range sel {
					val := fromCol.Get(selIdx)
					toCol.AppendVal(val)
				}
			}
			m.nulls.set(args)
			m.col = toCol
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", m.t))
	}
}

func (m *memColumn) Copy(args SliceArgs) {
	if args.SrcStartIdx == args.SrcEndIdx {
		// Nothing to copy, so return early.
		return
	}
	if m.Nulls().MaybeHasNulls() {
		// We're about to overwrite this entire range, so unset all the nulls.
		m.Nulls().UnsetNullRange(args.DestIdx, args.DestIdx+(args.SrcEndIdx-args.SrcStartIdx))
	}

	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bool()
			toCol := m.Bool()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Bytes()
			toCol := m.Bytes()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							toCol.Copy(fromCol, i+args.DestIdx, selIdx)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					toCol.Copy(fromCol, i+args.DestIdx, selIdx)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Decimal()
			toCol := m.Decimal()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			fromCol := args.Src.Int16()
			toCol := m.Int16()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		case 32:
			fromCol := args.Src.Int32()
			toCol := m.Int32()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		case -1:
		default:
			fromCol := args.Src.Int64()
			toCol := m.Int64()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Float64()
			toCol := m.Float64()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Timestamp()
			toCol := m.Timestamp()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Interval()
			toCol := m.Interval()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				toCol = toCol[args.DestIdx:]
				_ = toCol[n-1]
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							//gcassert:bce
							toCol.Set(i, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					//gcassert:bce
					toCol.Set(i, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case types.JsonFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.JSON()
			toCol := m.JSON()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							toCol.Copy(fromCol, i+args.DestIdx, selIdx)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					toCol.Copy(fromCol, i+args.DestIdx, selIdx)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := args.Src.Datum()
			toCol := m.Datum()
			if args.Sel != nil {
				sel := args.Sel[args.SrcStartIdx:args.SrcEndIdx]
				n := len(sel)
				if args.Src.MaybeHasNulls() {
					nulls := args.Src.Nulls()
					for i := 0; i < n; i++ {
						//gcassert:bce
						selIdx := sel[i]
						if nulls.NullAt(selIdx) {
							m.nulls.SetNull(i + args.DestIdx)
						} else {
							v := fromCol.Get(selIdx)
							toCol.Set(i+args.DestIdx, v)
						}
					}
					return
				}
				// No Nulls.
				for i := 0; i < n; i++ {
					//gcassert:bce
					selIdx := sel[i]
					v := fromCol.Get(selIdx)
					toCol.Set(i+args.DestIdx, v)
				}
				return
			}
			// No Sel.
			toCol.CopySlice(fromCol, args.DestIdx, args.SrcStartIdx, args.SrcEndIdx)
			m.nulls.set(args)
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", m.t))
	}
}

func (m *memColumn) CopyWithReorderedSource(src Vec, sel, order []int) {
	if len(sel) == 0 {
		return
	}
	if m.nulls.MaybeHasNulls() {
		m.nulls.UnsetNulls()
	}
	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Bool()
			toCol := m.Bool()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Bytes()
			toCol := m.Bytes()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Decimal()
			toCol := m.Decimal()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			fromCol := src.Int16()
			toCol := m.Int16()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		case 32:
			fromCol := src.Int32()
			toCol := m.Int32()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		case -1:
		default:
			fromCol := src.Int64()
			toCol := m.Int64()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Float64()
			toCol := m.Float64()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Timestamp()
			toCol := m.Timestamp()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Interval()
			toCol := m.Interval()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case types.JsonFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.JSON()
			toCol := m.JSON()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			fromCol := src.Datum()
			toCol := m.Datum()
			n := len(sel)
			_ = sel[n-1]
			if src.MaybeHasNulls() {
				nulls := src.Nulls()
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					if nulls.NullAt(srcIdx) {
						m.nulls.SetNull(destIdx)
					} else {
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			} else {
				for i := 0; i < n; i++ {
					//gcassert:bce
					destIdx := sel[i]
					srcIdx := order[destIdx]
					{
						v := fromCol.Get(srcIdx)
						toCol.Set(destIdx, v)
					}
				}
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", m.t))
	}
}

func (m *memColumn) Window(start int, end int) Vec {
	switch m.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Bool()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.BytesFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Bytes()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.DecimalFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Decimal()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.IntFamily:
		switch m.t.Width() {
		case 16:
			col := m.Int16()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		case 32:
			col := m.Int32()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		case -1:
		default:
			col := m.Int64()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.FloatFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Float64()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.TimestampTZFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Timestamp()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.IntervalFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Interval()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case types.JsonFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.JSON()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch m.t.Width() {
		case -1:
		default:
			col := m.Datum()
			return &memColumn{
				t:                   m.t,
				canonicalTypeFamily: m.canonicalTypeFamily,
				col:                 col.Window(start, end),
				nulls:               m.nulls.Slice(start, end),
			}
		}
	}
	panic(fmt.Sprintf("unhandled type %s", m.t))
}

// SetValueAt is an inefficient helper to set the value in a Vec when the type
// is unknown.
func SetValueAt(v Vec, elem interface{}, rowIdx int) {
	switch t := v.Type(); v.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bool()
			newVal := elem.(bool)
			target.Set(rowIdx, newVal)
		}
	case types.BytesFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bytes()
			newVal := elem.([]byte)
			target.Set(rowIdx, newVal)
		}
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Decimal()
			newVal := elem.(apd.Decimal)
			target.Set(rowIdx, newVal)
		}
	case types.IntFamily:
		switch t.Width() {
		case 16:
			target := v.Int16()
			newVal := elem.(int16)
			target.Set(rowIdx, newVal)
		case 32:
			target := v.Int32()
			newVal := elem.(int32)
			target.Set(rowIdx, newVal)
		case -1:
		default:
			target := v.Int64()
			newVal := elem.(int64)
			target.Set(rowIdx, newVal)
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Float64()
			newVal := elem.(float64)
			target.Set(rowIdx, newVal)
		}
	case types.TimestampTZFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Timestamp()
			newVal := elem.(time.Time)
			target.Set(rowIdx, newVal)
		}
	case types.IntervalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Interval()
			newVal := elem.(duration.Duration)
			target.Set(rowIdx, newVal)
		}
	case types.JsonFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.JSON()
			newVal := elem.(json.JSON)
			target.Set(rowIdx, newVal)
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Datum()
			newVal := elem.(interface{})
			target.Set(rowIdx, newVal)
		}
	default:
		panic(fmt.Sprintf("unhandled type %s", t))
	}
}

// GetValueAt is an inefficient helper to get the value in a Vec when the type
// is unknown.
func GetValueAt(v Vec, rowIdx int) interface{} {
	if v.Nulls().NullAt(rowIdx) {
		return nil
	}
	t := v.Type()
	switch v.CanonicalTypeFamily() {
	case types.BoolFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bool()
			return target.Get(rowIdx)
		}
	case types.BytesFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Bytes()
			return target.Get(rowIdx)
		}
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Decimal()
			return target.Get(rowIdx)
		}
	case types.IntFamily:
		switch t.Width() {
		case 16:
			target := v.Int16()
			return target.Get(rowIdx)
		case 32:
			target := v.Int32()
			return target.Get(rowIdx)
		case -1:
		default:
			target := v.Int64()
			return target.Get(rowIdx)
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Float64()
			return target.Get(rowIdx)
		}
	case types.TimestampTZFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Timestamp()
			return target.Get(rowIdx)
		}
	case types.IntervalFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Interval()
			return target.Get(rowIdx)
		}
	case types.JsonFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.JSON()
			return target.Get(rowIdx)
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch t.Width() {
		case -1:
		default:
			target := v.Datum()
			return target.Get(rowIdx)
		}
	}
	panic(fmt.Sprintf("unhandled type %s", t))
}
