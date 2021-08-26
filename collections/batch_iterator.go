/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package collections

// BatchIterator is an iterator for interface{}
type BatchIterator struct {
	Items		[]interface{}
	BatchSize	int
	Cursor		int
}

// HasNext returns if we should process another batch.
func (bi *BatchIterator) HasNext() bool {
	return bi.Cursor < (len(bi.Items) - 1)
}

// Next yields the next batch.
func (bi *BatchIterator) Next() []interface{} {
	if bi.BatchSize == 0 {
		return nil
	}
	if bi.Cursor >= len(bi.Items) {
		return nil
	}

	if (bi.Cursor + bi.BatchSize) < len(bi.Items) {
		output := bi.Items[bi.Cursor : bi.Cursor+bi.BatchSize]
		bi.Cursor += len(output)
		return output
	}

	output := bi.Items[bi.Cursor:]
	bi.Cursor += len(output)
	return output
}
