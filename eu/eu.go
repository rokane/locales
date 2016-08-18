package eu

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type eu struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'eu' locale
func New() locales.Translator {
	return &eu{
		locale:                 "eu",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0xe2, 0x82, 0xa7}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentPrefix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x75, 0x72, 0x74, 0x2e}, {0x6f, 0x74, 0x73, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x61, 0x70, 0x69, 0x2e}, {0x6d, 0x61, 0x69, 0x2e}, {0x65, 0x6b, 0x61, 0x2e}, {0x75, 0x7a, 0x74, 0x2e}, {0x61, 0x62, 0x75, 0x2e}, {0x69, 0x72, 0x61, 0x2e}, {0x75, 0x72, 0x72, 0x2e}, {0x61, 0x7a, 0x61, 0x2e}, {0x61, 0x62, 0x65, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x55}, {0x4f}, {0x4d}, {0x41}, {0x4d}, {0x45}, {0x55}, {0x41}, {0x49}, {0x55}, {0x41}, {0x41}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x75, 0x72, 0x74, 0x61, 0x72, 0x72, 0x69, 0x6c, 0x61, 0x6b}, {0x6f, 0x74, 0x73, 0x61, 0x69, 0x6c, 0x61, 0x6b}, {0x6d, 0x61, 0x72, 0x74, 0x78, 0x6f, 0x61, 0x6b}, {0x61, 0x70, 0x69, 0x72, 0x69, 0x6c, 0x61, 0x6b}, {0x6d, 0x61, 0x69, 0x61, 0x74, 0x7a, 0x61, 0x6b}, {0x65, 0x6b, 0x61, 0x69, 0x6e, 0x61, 0x6b}, {0x75, 0x7a, 0x74, 0x61, 0x69, 0x6c, 0x61, 0x6b}, {0x61, 0x62, 0x75, 0x7a, 0x74, 0x75, 0x61, 0x6b}, {0x69, 0x72, 0x61, 0x69, 0x6c, 0x61, 0x6b}, {0x75, 0x72, 0x72, 0x69, 0x61, 0x6b}, {0x61, 0x7a, 0x61, 0x72, 0x6f, 0x61, 0x6b}, {0x61, 0x62, 0x65, 0x6e, 0x64, 0x75, 0x61, 0x6b}},
		daysAbbreviated:        [][]uint8{{0x69, 0x67, 0x2e}, {0x61, 0x6c, 0x2e}, {0x61, 0x72, 0x2e}, {0x61, 0x7a, 0x2e}, {0x6f, 0x67, 0x2e}, {0x6f, 0x72, 0x2e}, {0x6c, 0x72, 0x2e}},
		daysNarrow:             [][]uint8{{0x49}, {0x41}, {0x41}, {0x41}, {0x4f}, {0x4f}, {0x4c}},
		daysShort:              [][]uint8{{0x69, 0x67, 0x2e}, {0x61, 0x6c, 0x2e}, {0x61, 0x72, 0x2e}, {0x61, 0x7a, 0x2e}, {0x6f, 0x67, 0x2e}, {0x6f, 0x72, 0x2e}, {0x6c, 0x72, 0x2e}},
		daysWide:               [][]uint8{{0x69, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x61}, {0x61, 0x73, 0x74, 0x65, 0x6c, 0x65, 0x68, 0x65, 0x6e, 0x61}, {0x61, 0x73, 0x74, 0x65, 0x61, 0x72, 0x74, 0x65, 0x61}, {0x61, 0x73, 0x74, 0x65, 0x61, 0x7a, 0x6b, 0x65, 0x6e, 0x61}, {0x6f, 0x73, 0x74, 0x65, 0x67, 0x75, 0x6e, 0x61}, {0x6f, 0x73, 0x74, 0x69, 0x72, 0x61, 0x6c, 0x61}, {0x6c, 0x61, 0x72, 0x75, 0x6e, 0x62, 0x61, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x67}, {0x61}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x4b, 0x2e, 0x61, 0x2e}, {0x4b, 0x2e, 0x6f, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x4b, 0x2e, 0x61, 0x2e}, {0x4b, 0x2e, 0x6f, 0x2e}},
		timezones:              map[string][]uint8{"ACWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x2d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "CAT": {0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "SAST": {0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x68, 0x65, 0x67, 0x6f, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x72, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "COT": {0x4b, 0x6f, 0x6c, 0x6f, 0x6e, 0x62, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "WITA": {0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "MESZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "HAT": {0x54, 0x65, 0x72, 0x6e, 0x75, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "EAT": {0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "NZST": {0x5a, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61, 0x20, 0x42, 0x65, 0x72, 0x72, 0x69, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "WARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "GFT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x65, 0x73, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x67, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "WAST": {0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "PST": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x50, 0x61, 0x7a, 0x69, 0x66, 0x69, 0x6b, 0x6f, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "ACWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x2d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x6f, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "CLT": {0x54, 0x78, 0x69, 0x6c, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "CST": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "CDT": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "COST": {0x4b, 0x6f, 0x6c, 0x6f, 0x6e, 0x62, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "AEST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "AWST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "CLST": {0x54, 0x78, 0x69, 0x6c, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "WEZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "HNT": {0x54, 0x65, 0x72, 0x6e, 0x75, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x67, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "ACDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "PDT": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x50, 0x61, 0x7a, 0x69, 0x66, 0x69, 0x6b, 0x6f, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "WIT": {0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "WIB": {0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x67, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "AWDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "WAT": {0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "JST": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "MST": {0x4d, 0x53, 0x54}, "EST": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x6f, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x67, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "NZDT": {0x5a, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61, 0x20, 0x42, 0x65, 0x72, 0x72, 0x69, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x72, 0x20, 0x75, 0x68, 0x61, 0x72, 0x74, 0x65, 0x65, 0x74, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "MEZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "AEDT": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "ECT": {0x45, 0x6b, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x72, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x67, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x72, 0x20, 0x75, 0x68, 0x61, 0x72, 0x74, 0x65, 0x65, 0x74, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "WESZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "OEZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "ACST": {0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x65, 0x72, 0x64, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "MDT": {0x4d, 0x44, 0x54}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "OESZ": {0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "JDT": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x67, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "WART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x62, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x20, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x72, 0x61}, "EDT": {0x49, 0x70, 0x61, 0x72, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6b, 0x6f, 0x20, 0x65, 0x6b, 0x69, 0x61, 0x6c, 0x64, 0x65, 0x6b, 0x6f, 0x20, 0x75, 0x64, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x6b, 0x6f, 0x20, 0x6f, 0x72, 0x64, 0x75, 0x61}},
	}
}

// Locale returns the current translators string locale
func (eu *eu) Locale() string {
	return eu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'eu'
func (eu *eu) PluralsCardinal() []locales.PluralRule {
	return eu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'eu'
func (eu *eu) PluralsOrdinal() []locales.PluralRule {
	return eu.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'eu'
func (eu *eu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'eu'
func (eu *eu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'eu'
func (eu *eu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (eu *eu) MonthAbbreviated(month time.Month) []byte {
	return eu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (eu *eu) MonthsAbbreviated() [][]byte {
	return eu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (eu *eu) MonthNarrow(month time.Month) []byte {
	return eu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (eu *eu) MonthsNarrow() [][]byte {
	return eu.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (eu *eu) MonthWide(month time.Month) []byte {
	return eu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (eu *eu) MonthsWide() [][]byte {
	return eu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (eu *eu) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return eu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (eu *eu) WeekdaysAbbreviated() [][]byte {
	return eu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (eu *eu) WeekdayNarrow(weekday time.Weekday) []byte {
	return eu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (eu *eu) WeekdaysNarrow() [][]byte {
	return eu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (eu *eu) WeekdayShort(weekday time.Weekday) []byte {
	return eu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (eu *eu) WeekdaysShort() [][]byte {
	return eu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (eu *eu) WeekdayWide(weekday time.Weekday) []byte {
	return eu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (eu *eu) WeekdaysWide() [][]byte {
	return eu.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'eu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(eu.decimal) + len(eu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'eu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (eu *eu) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(eu.decimal) + len(eu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eu.minus[0])
	}

	for j := len(eu.percentPrefix) - 1; j >= 0; j-- {
		b = append(b, eu.percentPrefix[j])
	}

	b = append(b, eu.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eu.currencies[currency]
	l := len(s) + len(eu.decimal) + len(eu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, eu.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'eu'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eu.currencies[currency]
	l := len(s) + len(eu.decimal) + len(eu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, eu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(eu.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, eu.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, eu.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, eu.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, eu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x65, 0x27}...)
	b = append(b, []byte{0x27, 0x6b, 0x6f, 0x27, 0x20}...)
	b = append(b, eu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x65, 0x27}...)
	b = append(b, []byte{0x27, 0x6b, 0x6f, 0x27, 0x20}...)
	b = append(b, eu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, eu.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'eu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (eu *eu) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := eu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return b
}
