package yo_BJ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type yo_BJ struct {
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

// New returns a new instance of translator for the 'yo_BJ' locale
func New() locales.Translator {
	return &yo_BJ{
		locale:                 "yo_BJ",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x53, 0x68, 0xc9, 0x9b, 0xcc, 0x81, 0x72, 0xc9, 0x9b, 0xcc, 0x81}, {0xc3, 0x88, 0x72, 0xc3, 0xa8, 0x6c, 0xc3, 0xa8}, {0xc6, 0x90, 0x72, 0xc9, 0x9b, 0xcc, 0x80, 0x6e, 0xc3, 0xa0}, {0xc3, 0x8c, 0x67, 0x62, 0xc3, 0xa9}, {0xc6, 0x90, 0xcc, 0x80, 0x62, 0x69, 0x62, 0x69}, {0xc3, 0x92, 0x6b, 0xc3, 0xba, 0x64, 0x75}, {0x41, 0x67, 0xc9, 0x9b, 0x6d, 0xc9, 0x94}, {0xc3, 0x92, 0x67, 0xc3, 0xba, 0x6e}, {0x4f, 0x77, 0x65, 0x77, 0x65}, {0xc6, 0x86, 0xcc, 0x80, 0x77, 0xc3, 0xa0, 0x72, 0xc3, 0xa0}, {0x42, 0xc3, 0xa9, 0x6c, 0xc3, 0xba}, {0xc6, 0x86, 0xcc, 0x80, 0x70, 0xc9, 0x9b, 0xcc, 0x80}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0x53, 0x68, 0xc9, 0x9b, 0xcc, 0x81, 0x72, 0xc9, 0x9b, 0xcc, 0x81}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc3, 0x88, 0x72, 0xc3, 0xa8, 0x6c, 0xc3, 0xa8}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc6, 0x90, 0x72, 0xc9, 0x9b, 0xcc, 0x80, 0x6e, 0xc3, 0xa0}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc3, 0x8c, 0x67, 0x62, 0xc3, 0xa9}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc6, 0x90, 0xcc, 0x80, 0x62, 0x69, 0x62, 0x69}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc3, 0x92, 0x6b, 0xc3, 0xba, 0x64, 0x75}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0x41, 0x67, 0xc9, 0x9b, 0x6d, 0xc9, 0x94}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc3, 0x92, 0x67, 0xc3, 0xba, 0x6e}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0x4f, 0x77, 0x65, 0x77, 0x65}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc6, 0x86, 0xcc, 0x80, 0x77, 0xc3, 0xa0, 0x72, 0xc3, 0xa0}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0x42, 0xc3, 0xa9, 0x6c, 0xc3, 0xba}, {0x4f, 0x73, 0x68, 0xc3, 0xb9, 0x20, 0xc6, 0x86, 0xcc, 0x80, 0x70, 0xc9, 0x9b, 0xcc, 0x80}},
		daysAbbreviated:        [][]uint8{{0xc3, 0x80, 0xc3, 0xac, 0x6b, 0xc3, 0xba}, {0x41, 0x6a, 0xc3, 0xa9}, {0xc3, 0x8c, 0x73, 0xc9, 0x9b, 0xcc, 0x81, 0x67, 0x75, 0x6e}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x72, 0xc3, 0xba}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x62, 0xc9, 0x94}, {0xc6, 0x90, 0x74, 0xc3, 0xac}, {0xc3, 0x80, 0x62, 0xc3, 0xa1, 0x6d, 0xc9, 0x9b, 0xcc, 0x81, 0x74, 0x61}},
		daysWide:               [][]uint8{{0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x20, 0xc3, 0x80, 0xc3, 0xac, 0x6b, 0xc3, 0xba}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x20, 0x41, 0x6a, 0xc3, 0xa9}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x20, 0xc3, 0x8c, 0x73, 0xc9, 0x9b, 0xcc, 0x81, 0x67, 0x75, 0x6e}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x72, 0xc3, 0xba}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x62, 0xc9, 0x94}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x20, 0xc6, 0x90, 0x74, 0xc3, 0xac}, {0xc6, 0x86, 0x6a, 0xc9, 0x94, 0xcc, 0x81, 0x20, 0xc3, 0x80, 0x62, 0xc3, 0xa1, 0x6d, 0xc9, 0x9b, 0xcc, 0x81, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0xc3, 0x80, 0xc3, 0xa1, 0x72, 0xc9, 0x94, 0xcc, 0x80}, {0xc6, 0x86, 0xcc, 0x80, 0x73, 0xc3, 0xa1, 0x6e}},
		periodsWide:            [][]uint8{{0xc3, 0x80, 0xc3, 0xa1, 0x72, 0xc9, 0x94, 0xcc, 0x80}, {0xc6, 0x86, 0xcc, 0x80, 0x73, 0xc3, 0xa1, 0x6e}},
		erasAbbreviated:        [][]uint8{{0x53, 0x4b}, {0x4c, 0x4b}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x53, 0x61, 0x6a, 0x75, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x69}, {0x4c, 0x65, 0x68, 0x69, 0x6e, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x69}},
		timezones:              map[string][]uint8{"SRT": {0x53, 0x52, 0x54}, "CDT": {0x43, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "CAT": {0x43, 0x41, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "HAT": {0x48, 0x41, 0x54}, "WIB": {0x57, 0x49, 0x42}, "AST": {0x41, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "ADT": {0x41, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "CST": {0x43, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "WIT": {0x57, 0x49, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "ECT": {0x45, 0x43, 0x54}, "EAT": {0x45, 0x41, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "SGT": {0x53, 0x47, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "BT": {0x42, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "ARST": {0x41, 0x52, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "WAT": {0x57, 0x41, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "GFT": {0x47, 0x46, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "WAST": {0x57, 0x41, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "MST": {0x4d, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "IST": {0x49, 0x53, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (yo *yo_BJ) Locale() string {
	return yo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsCardinal() []locales.PluralRule {
	return yo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsOrdinal() []locales.PluralRule {
	return yo.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yo_BJ'
func (yo *yo_BJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yo *yo_BJ) MonthAbbreviated(month time.Month) []byte {
	return yo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yo *yo_BJ) MonthsAbbreviated() [][]byte {
	return yo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yo *yo_BJ) MonthNarrow(month time.Month) []byte {
	return yo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yo *yo_BJ) MonthsNarrow() [][]byte {
	return yo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yo *yo_BJ) MonthWide(month time.Month) []byte {
	return yo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yo *yo_BJ) MonthsWide() [][]byte {
	return yo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return yo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yo *yo_BJ) WeekdaysAbbreviated() [][]byte {
	return yo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayNarrow(weekday time.Weekday) []byte {
	return yo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yo *yo_BJ) WeekdaysNarrow() [][]byte {
	return yo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayShort(weekday time.Weekday) []byte {
	return yo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yo *yo_BJ) WeekdaysShort() [][]byte {
	return yo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayWide(weekday time.Weekday) []byte {
	return yo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yo *yo_BJ) WeekdaysWide() [][]byte {
	return yo.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yo *yo_BJ) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(yo.decimal) + len(yo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(yo.decimal) - 1; j >= 0; j-- {
				b = append(b, yo.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(yo.group) - 1; j >= 0; j-- {
					b = append(b, yo.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		for j := len(yo.minus) - 1; j >= 0; j-- {
			b = append(b, yo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(yo.decimal) + len(yo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(yo.decimal) - 1; j >= 0; j-- {
				b = append(b, yo.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(yo.group) - 1; j >= 0; j-- {
					b = append(b, yo.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(yo.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, yo.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, yo.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, yo.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, yo.periodsAbbreviated[0]...)
	} else {
		b = append(b, yo.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, yo.periodsAbbreviated[0]...)
	} else {
		b = append(b, yo.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, yo.periodsAbbreviated[0]...)
	} else {
		b = append(b, yo.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'yo_BJ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yo *yo_BJ) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, yo.periodsAbbreviated[0]...)
	} else {
		b = append(b, yo.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
