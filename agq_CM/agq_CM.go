package agq_CM

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type agq_CM struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	decimal            []byte
	group              []byte
	minus              []byte
	percent            []byte
	perMille           []byte
	timeSeparator      []byte
	inifinity          []byte
	currencies         [][]byte // idx = enum of currency code
	monthsAbbreviated  [][]byte
	monthsNarrow       [][]byte
	monthsWide         [][]byte
	daysAbbreviated    [][]byte
	daysNarrow         [][]byte
	daysShort          [][]byte
	daysWide           [][]byte
	periodsAbbreviated [][]byte
	periodsNarrow      [][]byte
	periodsShort       [][]byte
	periodsWide        [][]byte
	erasAbbreviated    [][]byte
	erasNarrow         [][]byte
	erasWide           [][]byte
	timezones          map[string][]byte
}

// New returns a new instance of translator for the 'agq_CM' locale
func New() locales.Translator {
	return &agq_CM{
		locale:             "agq_CM",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		decimal:            []byte{0x2c},
		group:              []byte{0xc2, 0xa0},
		minus:              []byte{},
		percent:            []byte{},
		perMille:           []byte{},
		timeSeparator:      []byte{0x3a},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0x6e, 0xc3, 0xb9, 0x6d}, {0x6b, 0xc9, 0xa8, 0x7a}, {0x74, 0xc9, 0xa8, 0x64}, {0x74, 0x61, 0x61}, {0x73, 0x65, 0x65}, {0x6e, 0x7a, 0x75}, {0x64, 0x75, 0x6d}, {0x66, 0xc9, 0x94, 0x65}, {0x64, 0x7a, 0x75}, {0x6c, 0xc9, 0x94, 0x6d}, {0x6b, 0x61, 0x61}, {0x66, 0x77, 0x6f}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0x6e}, {0x6b}, {0x74}, {0x74}, {0x73}, {0x7a}, {0x6b}, {0x66}, {0x64}, {0x6c}, {0x63}, {0x66}},
		monthsWide:         [][]uint8{[]uint8(nil), {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x6e, 0xc3, 0xb9, 0x6d}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x6b, 0xc6, 0x97, 0xcc, 0x80, 0x7a, 0xc3, 0xb9, 0xca, 0x94}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x74, 0xc6, 0x97, 0xcc, 0x80, 0x64, 0xca, 0x89, 0xcc, 0x80, 0x67, 0x68, 0xc3, 0xa0}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x74, 0xc7, 0x8e, 0x61, 0x66, 0xca, 0x89, 0xcc, 0x84, 0x67, 0x68, 0xc4, 0x81}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc3, 0xa8, 0x73, 0xc3, 0xa8, 0x65}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x6e, 0x7a, 0xc3, 0xb9, 0x67, 0x68, 0xc3, 0xb2}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x64, 0xc3, 0xb9, 0x6d, 0x6c, 0x6f}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x6b, 0x77, 0xc3, 0xae, 0x66, 0xc9, 0x94, 0xcc, 0x80, 0x65}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x74, 0xc6, 0x97, 0xcc, 0x80, 0x66, 0xca, 0x89, 0xcc, 0x80, 0x67, 0x68, 0xc3, 0xa0, 0x64, 0x7a, 0x75, 0x67, 0x68, 0xc3, 0xb9}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x67, 0x68, 0xc7, 0x94, 0x75, 0x77, 0x65, 0x6c, 0xc9, 0x94, 0xcc, 0x80, 0x6d}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc9, 0x94, 0xcc, 0x80, 0x63, 0x68, 0x77, 0x61, 0xca, 0x94, 0xc3, 0xa0, 0x6b, 0x61, 0x61, 0x20, 0x77, 0x6f}, {0x6e, 0x64, 0x7a, 0xc9, 0x94, 0xcc, 0x80, 0xc5, 0x8b, 0xc3, 0xa8, 0x66, 0x77, 0xc3, 0xb2, 0x6f}},
		daysAbbreviated:    [][]uint8{{0x6e, 0x74, 0x73}, {0x6b, 0x70, 0x61}, {0x67, 0x68, 0xc9, 0x94}, {0x74, 0xc9, 0x94, 0x6d}, {0x75, 0x6d, 0x65}, {0x67, 0x68, 0xc9, 0xa8}, {0x64, 0x7a, 0x6b}},
		daysNarrow:         [][]uint8{{0x6e}, {0x6b}, {0x67}, {0x74}, {0x75}, {0x67}, {0x64}},
		daysWide:           [][]uint8{{0x74, 0x73, 0x75, 0xca, 0x94, 0x6e, 0x74, 0x73, 0xc9, 0xa8}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x75, 0x6b, 0x70, 0xc3, 0xa0}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x75, 0x67, 0x68, 0xc9, 0x94, 0x65}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x75, 0x74, 0xc9, 0x94, 0xcc, 0x80, 0x6d, 0x6c, 0xc3, 0xb2}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x75, 0x6d, 0xc3, 0xa8}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x75, 0x67, 0x68, 0xc9, 0xa8, 0xcc, 0x82, 0x6d}, {0x74, 0x73, 0x75, 0xca, 0x94, 0x6e, 0x64, 0x7a, 0xc9, 0xa8, 0x6b, 0xc9, 0x94, 0xca, 0x94, 0xc9, 0x94}},
		periodsAbbreviated: [][]uint8{{0x61, 0x2e, 0x67}, {0x61, 0x2e, 0x6b}},
		periodsWide:        [][]uint8{{0x61, 0x2e, 0x67}, {0x61, 0x2e, 0x6b}},
		erasAbbreviated:    [][]uint8{{0x53, 0x4b}, {0x42, 0x4b}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0x53, 0xc4, 0x9b, 0x65, 0x20, 0x4b, 0xc9, 0xa8, 0xcc, 0x80, 0x6c, 0x65, 0x73, 0x74, 0x6f}, {0x42, 0xc7, 0x8e, 0x61, 0x20, 0x4b, 0xc9, 0xa8, 0xcc, 0x80, 0x6c, 0x65, 0x73, 0x74, 0x6f}},
		timezones:          map[string][]uint8{"PDT": {0x50, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "ADT": {0x41, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "BT": {0x42, 0x54}, "CST": {0x43, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "ChST": {0x43, 0x68, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "SGT": {0x53, 0x47, 0x54}, "ART": {0x41, 0x52, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "HAT": {0x48, 0x41, 0x54}, "CDT": {0x43, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "CAT": {0x43, 0x41, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "ECT": {0x45, 0x43, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "VET": {0x56, 0x45, 0x54}, "GFT": {0x47, 0x46, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "MST": {0x4d, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "IST": {0x49, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "EDT": {0x45, 0x44, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CLT": {0x43, 0x4c, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WART": {0x57, 0x41, 0x52, 0x54}, "PST": {0x50, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "AEST": {0x41, 0x45, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "JST": {0x4a, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (agq *agq_CM) Locale() string {
	return agq.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'agq_CM'
func (agq *agq_CM) PluralsCardinal() []locales.PluralRule {
	return agq.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'agq_CM'
func (agq *agq_CM) PluralsOrdinal() []locales.PluralRule {
	return agq.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'agq_CM'
func (agq *agq_CM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'agq_CM'
func (agq *agq_CM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'agq_CM'
func (agq *agq_CM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (agq *agq_CM) MonthAbbreviated(month time.Month) []byte {
	return agq.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (agq *agq_CM) MonthsAbbreviated() [][]byte {
	return agq.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (agq *agq_CM) MonthNarrow(month time.Month) []byte {
	return agq.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (agq *agq_CM) MonthsNarrow() [][]byte {
	return agq.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (agq *agq_CM) MonthWide(month time.Month) []byte {
	return agq.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (agq *agq_CM) MonthsWide() [][]byte {
	return agq.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (agq *agq_CM) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return agq.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (agq *agq_CM) WeekdaysAbbreviated() [][]byte {
	return agq.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (agq *agq_CM) WeekdayNarrow(weekday time.Weekday) []byte {
	return agq.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (agq *agq_CM) WeekdaysNarrow() [][]byte {
	return agq.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (agq *agq_CM) WeekdayShort(weekday time.Weekday) []byte {
	return agq.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (agq *agq_CM) WeekdaysShort() [][]byte {
	return agq.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (agq *agq_CM) WeekdayWide(weekday time.Weekday) []byte {
	return agq.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (agq *agq_CM) WeekdaysWide() [][]byte {
	return agq.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'agq_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(agq.decimal) + len(agq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, agq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(agq.group) - 1; j >= 0; j-- {
					b = append(b, agq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(agq.minus) - 1; j >= 0; j-- {
			b = append(b, agq.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'agq_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (agq *agq_CM) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(agq.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, agq.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(agq.minus) - 1; j >= 0; j-- {
			b = append(b, agq.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, agq.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := agq.currencies[currency]
	l := len(s) + len(agq.decimal) + len(agq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, agq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(agq.group) - 1; j >= 0; j-- {
					b = append(b, agq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(agq.minus) - 1; j >= 0; j-- {
			b = append(b, agq.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, agq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'agq_CM'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := agq.currencies[currency]
	l := len(s) + len(agq.decimal) + len(agq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, agq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(agq.group) - 1; j >= 0; j-- {
					b = append(b, agq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(agq.minus) - 1; j >= 0; j-- {
			b = append(b, agq.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, agq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, agq.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, agq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, agq.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, agq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'agq_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (agq *agq_CM) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, agq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := agq.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
