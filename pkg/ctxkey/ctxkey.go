package ctxkey

import (
	"context"
	"fmt"
	"reflect"
)

type stringer[T any] struct{ v T }

func (v stringer[T]) String() string { return fmt.Sprint(v.v) }

type CtxKey[ValueT any] struct {
	name         *stringer[string]
	defaultValue *ValueT
}

func New[ValueT any](name string, defaultValue ValueT) CtxKey[ValueT] {
	key := CtxKey[ValueT]{name: new(stringer[string])}

	if name == "" {
		name = reflect.TypeFor[ValueT]().String()
	}

	key.name.v = name

	v := reflect.ValueOf(defaultValue)
	if v.IsValid() && !v.IsZero() {
		key.defaultValue = &defaultValue
	}

	return key
}

func (key CtxKey[ValueT]) contextKey() any {
	if key.name == nil {
		return reflect.TypeFor[ValueT]()
	}

	return key.name
}

func (k CtxKey[ValueT]) WithValue(
	parent context.Context,
	val ValueT,
) context.Context {
	key := k.contextKey()
	stringer := stringer[ValueT]{val}

	return context.WithValue(parent, key, stringer)
}

func (k CtxKey[ValueT]) ValueOk(ctx context.Context) (ValueT, bool) {
	vv, ok := ctx.Value(k.contextKey()).(stringer[ValueT])
	if !ok && k.defaultValue != nil {
		vv.v = *k.defaultValue
	}

	return vv.v, ok
}

func (k CtxKey[ValueT]) Value(ctx context.Context) ValueT {
	v, _ := k.ValueOk(ctx)

	return v
}

func (k CtxKey[ValueT]) Has(ctx context.Context) bool {
	_, ok := k.ValueOk(ctx)

	return ok
}

func (key CtxKey[ValueT]) String() string {
	if key.name == nil {
		return reflect.TypeFor[ValueT]().String()
	}

	return key.name.String()
}
