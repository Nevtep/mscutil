package mscutil

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)
func TestSerializeToKey(t *testing.T) {
	// TODO: Verify values are the same as ruby implementation
	tests := []struct {
		address string
		values  SimpleSend
	}{
		{"01000000000000000200000000000000320000000000000000000000000000", SimpleSend{CurrencyId: 2, Sequence: 76, TransactionType: 0, Amount: 50}},
		{"010000000000000002000000000000305e0000000000000000000000000000", SimpleSend{CurrencyId: 2, Sequence: 48, TransactionType: 0, Amount: 12382}},
		{"0100000000000000010000000000018fee0000000000000000000000000000", SimpleSend{CurrencyId: 1, Sequence: 48, TransactionType: 0, Amount: 102382}},
	}

	for _, pair := range tests {
		v := pair.values.SerializeToKey()
		if v != pair.address {
			t.Error("For", pair.address,
				"Expected", pair.address,
				"Got", v,
			)

		}
	}
}
func TestMakeStringArray(t *testing.T) {
	tests := []struct {
		value    []string
		expected []string
	}{
		{[]string{"100", "4"}, []string{"0", "1", "0", "0"}},
		{[]string{"3292", "8"}, []string{"0", "0", "0", "0", "3", "2", "9", "2"}},
	}
	for _, pair := range tests {
		toint, _ := strconv.Atoi(pair.value[1])
		v := makeStringArray(pair.value[0], toint)
		// TODO: DeepEqual is pretty slow. If we end up needing this more, make something custom.
		if !reflect.DeepEqual(v, pair.expected) {
			t.Error("Expected", pair.expected,
				"For", pair.value,
				"Got", v)
		}
	}
}

func TestMakeBinary(t *testing.T) {
	tests := []struct {
		value    uint32
		expected []byte
	}{
		{3, []byte{0, 0, 0, 3}},
		{256, []byte{0, 0, 1, 0}},
		{123456, []byte{0, 1, 226, 64}},
		{1239923456, []byte{73, 231, 187, 0}},
	}

	for _, pair := range tests {
		v := makeBinary(pair.value)
		if !bytes.Equal(v, pair.expected) {
			t.Error("Expected", pair.expected,
				"For", pair.value,
				"Got", v)
		}
	}
}

func TestMakeBinary64(t *testing.T) {
	tests := []struct {
		value    uint64
		expected []byte
	}{
		{3, []byte{0, 0, 0, 0, 0, 0, 0, 3}},
		{256, []byte{0, 0, 0, 0, 0, 0, 1, 0}},
		{123456, []byte{0, 0, 0, 0, 0, 1, 226, 64}},
		{1239923456, []byte{0, 0, 0, 0, 73, 231, 187, 0}},
		{221239923456, []byte{0, 0, 0, 51, 130, 237, 83, 0}},
	}

	for _, pair := range tests {
		v := makeBinary(pair.value)
		if !bytes.Equal(v, pair.expected) {
			t.Error("Expected", pair.expected,
				"For", pair.value,
				"Got", v)
		}
	}
}

func TestDecodeFromPulicKeys(t *testing.T){
	tests := []struct{
		publicKeys []string
		values SimpleSend
	}{
		{[]string{"0100000000000000020000000005f5e1", "000000000000000000000000000000"}, SimpleSend{CurrencyId: 2, Sequence: 1, TransactionType: 0, Amount: 1e8}},
	}
	for _, pair := range tests {
		v := DecodeFromPublicKeys(pair.publicKeys)
		if v != pair.values {
			t.Error("For", pair.publicKeys,
				"Expected", pair.values,
				"Got", v,
			)

		}
	}
}

func TestDecodeFromAddress(t *testing.T) {
	tests := []struct {
		address string
		values  SimpleSend
	}{
		{"17vrMab8gQx72eCEaUxJzL4fg5VwEUumJQ", SimpleSend{CurrencyId: 2, Sequence: 76, TransactionType: 0, Amount: 50}},
		{"15NoSD4F1ULYHPfSiV1dp1kr9n2bBffGGd", SimpleSend{CurrencyId: 2, Sequence: 48, TransactionType: 0, Amount: 12382}},
		{"15NoSD4F1ULYHGW3TK6khj6NEZsPAmHf41", SimpleSend{CurrencyId: 1, Sequence: 48, TransactionType: 0, Amount: 102382}},
	}

	for _, pair := range tests {
		v := DecodeFromAddress(pair.address)
		if v != pair.values {
			t.Error("For", pair.address,
				"Expected", pair.values,
				"Got", v,
			)

		}
	}
}

func TestEncodeToAddress(t *testing.T) {
	tests := []struct {
		address string
		values  SimpleSend
	}{
		{"17vrMab8gQx72eCEaUxJzL4fg5VwDuND4T", SimpleSend{CurrencyId: 2, Sequence: 76, TransactionType: 0, Amount: 50}},
		{"15NoSD4F1ULYHPfSiV1dp1kr9n2b9Npxf1", SimpleSend{CurrencyId: 2, Sequence: 48, TransactionType: 0, Amount: 12382}},
		{"15NoSD4F1ULYHGW3TK6khj6NEZsP9ariEK", SimpleSend{CurrencyId: 1, Sequence: 48, TransactionType: 0, Amount: 102382}},
	}

	for _, pair := range tests {
		v := pair.values.SerializeToAddress()
		if v != pair.address {
			t.Error("For", pair.address,
				"Expected", pair.address,
				"Got", v,
			)

		}
	}
}
