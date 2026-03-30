package core

import "io"

type Trasaction struct {
}

func (tx *Trasaction) DecodeBinary(r io.Reader) error { return nil }

func (tx *Trasaction) EncodeBinary(w io.Writer) error { return nil }
