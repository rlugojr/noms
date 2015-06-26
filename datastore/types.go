// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package datastore

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// RootSet

type RootSet struct {
	s types.Set
}

type RootSetIterCallback (func (p Root) (stop bool))

func NewRootSet() RootSet {
	return RootSet{types.NewSet()}
}

func RootSetFromVal(p types.Value) RootSet {
	return RootSet{p.(types.Set)}
}

func (s RootSet) NomsValue() types.Set {
	return s.s
}

func (s RootSet) Equals(p RootSet) bool {
	return s.s.Equals(p.s)
}

func (s RootSet) Ref() ref.Ref {
	return s.s.Ref()
}

func (s RootSet) Empty() bool {
	return s.s.Empty()
}

func (s RootSet) Len() uint64 {
	return s.s.Len()
}

func (s RootSet) Has(p Root) bool {
	return s.s.Has(p.NomsValue())
}

func (s RootSet) Iter(cb RootSetIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(RootFromVal(v))
	})
}

func (s RootSet) Insert(p ...Root) RootSet {
	return RootSet{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s RootSet) Remove(p ...Root) RootSet {
	return RootSet{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s RootSet) Union(others ...RootSet) RootSet {
	return RootSet{s.s.Union(s.fromStructSlice(others)...)}
}

func (s RootSet) Subtract(others ...RootSet) RootSet {
	return RootSet{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s RootSet) Any() Root {
	return RootFromVal(s.s.Any())
}

func (s RootSet) fromStructSlice(p []RootSet) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s RootSet) fromElemSlice(p []Root) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

// Root

type Root struct {
	m types.Map
}

func NewRoot() Root {
	return Root{types.NewMap()}
}

func RootFromVal(v types.Value) Root {
	return Root{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Root) NomsValue() types.Map {
	return s.m
}

func (s Root) Equals(p Root) bool {
	return s.m.Equals(p.m)
}

func (s Root) Ref() ref.Ref {
	return s.m.Ref()
}
func(s Root) Value() types.Value {
	return s.m.Get(types.NewString("value")).(types.Value)
}

func (s Root) SetValue(p types.Value) Root {
	return RootFromVal(s.m.Set(types.NewString("value"), p))
}
func(s Root) Parents() types.Set {
	return s.m.Get(types.NewString("parents")).(types.Set)
}

func (s Root) SetParents(p types.Set) Root {
	return RootFromVal(s.m.Set(types.NewString("parents"), p))
}
