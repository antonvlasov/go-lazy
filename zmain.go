// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package lazy

import "sync"

type OfString struct {
	New        func() string
	lockString sync.Mutex
	value      string
}

func (this *OfString) Value() string {
	this.lockString.Lock()
	defer this.lockString.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}

type OfInt struct {
	New     func() int
	lockInt sync.Mutex
	value   int
}

func (this *OfInt) Value() int {
	this.lockInt.Lock()
	defer this.lockInt.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}

type OfBool struct {
	New      func() bool
	lockBool sync.Mutex
	value    bool
}

func (this *OfBool) Value() bool {
	this.lockBool.Lock()
	defer this.lockBool.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}

type OfAny struct {
	New     func() Any
	lockAny sync.Mutex
	value   Any
}

func (this *OfAny) Value() Any {
	this.lockAny.Lock()
	defer this.lockAny.Unlock()
	if this.New != nil {
		this.value = this.New()
		this.New = nil
	}
	return this.value
}
