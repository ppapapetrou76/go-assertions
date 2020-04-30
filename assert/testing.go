package assert

import (
	"testing"
)

// FluentT wraps the testing.T pointer to provide a better experience to the library users
type FluentT struct {
	t *testing.T
}

// NewFluentT creates a new Fluent testing
func NewFluentT(t *testing.T) *FluentT { // nolint
	return &FluentT{
		t: t,
	}
}

// AssertThatString initializes an assertable string to be used for asserting string properties
func (t FluentT) AssertThatString(actual string) AssertableString {
	return ThatString(t.t, actual)
}

// AssertThatBool initializes an assertable bool to be used for asserting bool properties
func (t FluentT) AssertThatBool(actual bool) AssertableBool {
	return ThatBool(t.t, actual)
}

// AssertThatInt initializes an assertable int to be used for asserting int properties
func (t FluentT) AssertThatInt(actual int) AssertableInt {
	return ThatInt(t.t, actual)
}

// AssertThatSlice initializes an assertable slice to be used for asserting slice properties
func (t FluentT) AssertThatSlice(actual interface{}, opts ...SliceOpt) AssertableSlice {
	return ThatSlice(t.t, actual, opts...)
}

// AssertThatStruct initializes an assertable struct to be used for asserting struct properties
func (t FluentT) AssertThatStruct(actual interface{}) AssertableStruct {
	return ThatStruct(t.t, actual)
}

// AssertThat initializes an assertable object to be used for asserting properties of any type
// The things we can assert using this type are limited
func (t FluentT) AssertThat(actual interface{}) Assertable {
	return That(t.t, actual)
}