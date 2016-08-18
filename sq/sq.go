package sq

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sq struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
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

// New returns a new instance of translator for the 'sq' locale
func New() locales.Translator {
	return &sq{
		locale:                 "sq",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 5, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x4c, 0x65, 0x6b, 0xc3, 0xab}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x53, 0x68, 0x6b}, {0x4d, 0x61, 0x72}, {0x50, 0x72, 0x69}, {0x4d, 0x61, 0x6a}, {0x51, 0x65, 0x72}, {0x4b, 0x6f, 0x72}, {0x47, 0x73, 0x68}, {0x53, 0x68, 0x74}, {0x54, 0x65, 0x74}, {0x4e, 0xc3, 0xab, 0x6e}, {0x44, 0x68, 0x6a}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x53}, {0x4d}, {0x50}, {0x4d}, {0x51}, {0x4b}, {0x47}, {0x53}, {0x54}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x61, 0x72}, {0x73, 0x68, 0x6b, 0x75, 0x72, 0x74}, {0x6d, 0x61, 0x72, 0x73}, {0x70, 0x72, 0x69, 0x6c, 0x6c}, {0x6d, 0x61, 0x6a}, {0x71, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x72}, {0x6b, 0x6f, 0x72, 0x72, 0x69, 0x6b}, {0x67, 0x75, 0x73, 0x68, 0x74}, {0x73, 0x68, 0x74, 0x61, 0x74, 0x6f, 0x72}, {0x74, 0x65, 0x74, 0x6f, 0x72}, {0x6e, 0xc3, 0xab, 0x6e, 0x74, 0x6f, 0x72}, {0x64, 0x68, 0x6a, 0x65, 0x74, 0x6f, 0x72}},
		daysAbbreviated:        [][]uint8{{0x44, 0x69, 0x65}, {0x48, 0xc3, 0xab, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0xc3, 0xab, 0x72}, {0x45, 0x6e, 0x6a}, {0x50, 0x72, 0x65}, {0x53, 0x68, 0x74}},
		daysNarrow:             [][]uint8{{0x44}, {0x48}, {0x4d}, {0x4d}, {0x45}, {0x50}, {0x53}},
		daysShort:              [][]uint8{{0x44, 0x69, 0x65}, {0x48, 0xc3, 0xab, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0xc3, 0xab, 0x72}, {0x45, 0x6e, 0x6a}, {0x50, 0x72, 0x65}, {0x53, 0x68, 0x74}},
		daysWide:               [][]uint8{{0x65, 0x20, 0x64, 0x69, 0x65, 0x6c}, {0x65, 0x20, 0x68, 0xc3, 0xab, 0x6e, 0xc3, 0xab}, {0x65, 0x20, 0x6d, 0x61, 0x72, 0x74, 0xc3, 0xab}, {0x65, 0x20, 0x6d, 0xc3, 0xab, 0x72, 0x6b, 0x75, 0x72, 0xc3, 0xab}, {0x65, 0x20, 0x65, 0x6e, 0x6a, 0x74, 0x65}, {0x65, 0x20, 0x70, 0x72, 0x65, 0x6d, 0x74, 0x65}, {0x65, 0x20, 0x73, 0x68, 0x74, 0x75, 0x6e, 0xc3, 0xab}},
		periodsAbbreviated:     [][]uint8{{0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x64, 0x69, 0x74, 0x65, 0x73}, {0x65, 0x20, 0x70, 0x61, 0x73, 0x64, 0x69, 0x74, 0x65, 0x73}},
		periodsNarrow:          [][]uint8{{0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x64, 0x69, 0x74, 0x65, 0x73}, {0x65, 0x20, 0x70, 0x61, 0x73, 0x64, 0x69, 0x74, 0x65, 0x73}},
		periodsWide:            [][]uint8{{0x65, 0x20, 0x70, 0x61, 0x72, 0x61, 0x64, 0x69, 0x74, 0x65, 0x73}, {0x65, 0x20, 0x70, 0x61, 0x73, 0x64, 0x69, 0x74, 0x65, 0x73}},
		erasAbbreviated:        [][]uint8{{0x70, 0x2e, 0x65, 0x2e, 0x72, 0x2e}, {0x65, 0x2e, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x70, 0x2e, 0x65, 0x2e, 0x72, 0x2e}, {0x65, 0x2e, 0x72, 0x2e}},
		erasWide:               [][]uint8{{0x70, 0x61, 0x72, 0x61, 0x20, 0x65, 0x72, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x72, 0x65}, {0x65, 0x72, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x72, 0x65}},
		timezones:              map[string][]uint8{"AKST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0xc3, 0xab, 0x73}, "PDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x65, 0x20, 0x65, 0x20, 0x42, 0x72, 0x65, 0x67, 0x75, 0x74, 0x20, 0x74, 0xc3, 0xab, 0x20, 0x50, 0x61, 0x71, 0xc3, 0xab, 0x73, 0x6f, 0x72, 0x69, 0x74}, "AEDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "NZST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x52, 0x65}, "HAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x49, 0x73, 0x68, 0x75, 0x6a, 0x76, 0x65, 0x20, 0x48, 0x61, 0x75, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "BOT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x73, 0xc3, 0xab}, "AEST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "∅∅∅": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x45, 0x6a, 0x6b, 0x72, 0x69, 0x74, 0x20, 0x5b, 0x41, 0x6b, 0x6f, 0x5d}, "EST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x53, 0x48, 0x42, 0x41, 0x2d, 0x73, 0xc3, 0xab, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "MST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4d, 0x61, 0x6b, 0x61, 0x6f, 0x73}, "GMT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x4d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x69, 0x74, 0x20, 0x74, 0xc3, 0xab, 0x20, 0x47, 0x72, 0x69, 0x6e, 0x75, 0x69, 0xc3, 0xa7, 0x69, 0x74}, "ACDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "LHST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x75, 0x69, 0x74}, "ADT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x75, 0x74}, "ARST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x72, 0x67, 0x6a, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0xc3, 0xab, 0x73}, "AKDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x6c, 0x73, 0x61, 0x73, 0x6b, 0xc3, 0xab, 0x73}, "HKST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x2d, 0x4b, 0x6f, 0x6e, 0x67, 0x75, 0x74}, "HAT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4e, 0x6a, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x20, 0x5b, 0x54, 0x6f, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x52, 0x65, 0x5d}, "UYST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x74}, "SRT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x69, 0x74}, "WESZ": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "VET": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0xc3, 0xab, 0x73}, "OESZ": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "ChST": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x4b, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x73}, "CHAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x61, 0x74, 0x61, 0x6d, 0x69, 0x74}, "CLT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x69, 0x6c, 0x69, 0x74}, "CHADT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x61, 0x74, 0x61, 0x6d, 0x69, 0x74}, "LHDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x75, 0x69, 0x74}, "WARST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x72, 0x67, 0x6a, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "UYT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x74}, "WIT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "COST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x73, 0xc3, 0xab}, "JST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x73, 0xc3, 0xab}, "MESZ": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "JDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x73, 0xc3, 0xab}, "NZDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x52, 0x65}, "CDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x53, 0x48, 0x42, 0x41, 0x2d, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "CAT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "ECT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x45, 0x6b, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x74}, "MDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4d, 0x61, 0x6b, 0x61, 0x6f, 0x73}, "AST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x75, 0x74}, "ACWDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x6f, 0x2d, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "TMST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x74}, "WART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x72, 0x67, 0x6a, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "EDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x53, 0x48, 0x42, 0x41, 0x2d, 0x73, 0xc3, 0xab, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "OEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "WITA": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "AWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "TMT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x74}, "CLST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x69, 0x6c, 0x69, 0x74}, "EAT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x4c, 0x69, 0x6e, 0x64, 0x6f, 0x72, 0x65}, "BT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x42, 0x75, 0x74, 0x61, 0x6e, 0x69, 0x74}, "HADT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x49, 0x73, 0x68, 0x75, 0x6a, 0x76, 0x65, 0x20, 0x48, 0x61, 0x75, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "WEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "COT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x73, 0xc3, 0xab}, "HKT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x2d, 0x4b, 0x6f, 0x6e, 0x67, 0x75, 0x74}, "HNT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x4e, 0x6a, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x69, 0x74, 0x20, 0x5b, 0x54, 0x6f, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x73, 0xc3, 0xab, 0x20, 0x52, 0x65, 0x5d}, "ART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x72, 0x67, 0x6a, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0xc3, 0xab, 0x73}, "GFT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x47, 0x75, 0x61, 0x6a, 0x61, 0x6e, 0xc3, 0xab, 0x73, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x7a, 0x65}, "ACWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x6f, 0x2d, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "GYT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x47, 0x75, 0x61, 0x6a, 0x61, 0x6e, 0xc3, 0xab, 0x73}, "CST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x53, 0x48, 0x42, 0x41, 0x2d, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "PST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x65, 0x20, 0x65, 0x20, 0x42, 0x72, 0x65, 0x67, 0x75, 0x74, 0x20, 0x74, 0xc3, 0xab, 0x20, 0x50, 0x61, 0x71, 0xc3, 0xab, 0x73, 0x6f, 0x72, 0x69, 0x74}, "SAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x4a, 0x75, 0x67, 0x6f, 0x72, 0x65}, "WAST": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "IST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x73, 0xc3, 0xab}, "WAT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0xc3, 0xab, 0x73, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "WIB": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "ACST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "MYT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x6a, 0x7a, 0x69, 0x73, 0xc3, 0xab}, "AWDT": {0x4f, 0x72, 0x61, 0x20, 0x76, 0x65, 0x72, 0x6f, 0x72, 0x65, 0x20, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0xc3, 0xab, 0x20, 0x50, 0x65, 0x72, 0xc3, 0xab, 0x6e, 0x64, 0x69, 0x6d, 0x6f, 0x72, 0x65}, "MEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x65, 0x20, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xab, 0x73, 0x20, 0x51, 0x65, 0x6e, 0x64, 0x72, 0x6f, 0x72, 0x65}, "SGT": {0x4f, 0x72, 0x61, 0x20, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x69, 0x74}},
	}
}

// Locale returns the current translators string locale
func (sq *sq) Locale() string {
	return sq.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sq'
func (sq *sq) PluralsCardinal() []locales.PluralRule {
	return sq.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sq'
func (sq *sq) PluralsOrdinal() []locales.PluralRule {
	return sq.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sq'
func (sq *sq) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sq'
func (sq *sq) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if n == 1 {
		return locales.PluralRuleOne
	} else if nMod10 == 4 && nMod100 != 14 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sq'
func (sq *sq) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sq.CardinalPluralRule(num1, v1)
	end := sq.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sq *sq) MonthAbbreviated(month time.Month) []byte {
	return sq.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sq *sq) MonthsAbbreviated() [][]byte {
	return sq.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sq *sq) MonthNarrow(month time.Month) []byte {
	return sq.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sq *sq) MonthsNarrow() [][]byte {
	return sq.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sq *sq) MonthWide(month time.Month) []byte {
	return sq.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sq *sq) MonthsWide() [][]byte {
	return sq.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sq *sq) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return sq.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sq *sq) WeekdaysAbbreviated() [][]byte {
	return sq.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sq *sq) WeekdayNarrow(weekday time.Weekday) []byte {
	return sq.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sq *sq) WeekdaysNarrow() [][]byte {
	return sq.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sq *sq) WeekdayShort(weekday time.Weekday) []byte {
	return sq.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sq *sq) WeekdaysShort() [][]byte {
	return sq.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sq *sq) WeekdayWide(weekday time.Weekday) []byte {
	return sq.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sq *sq) WeekdaysWide() [][]byte {
	return sq.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sq' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sq.decimal) + len(sq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sq' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sq *sq) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sq.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sq.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(sq.decimal) + len(sq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sq.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sq'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(sq.decimal) + len(sq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(sq.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sq.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sq.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sq.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, sq.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'sq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sq *sq) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sq.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
