package gd_GB

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type gd_GB struct {
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

// New returns a new instance of translator for the 'gd_GB' locale
func New() locales.Translator {
	return &gd_GB{
		locale:                 "gd_GB",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0xd9, 0xab},
		group:                  []byte{0xd9, 0xac},
		minus:                  []byte{0xe2, 0x80, 0x8f, 0x2d},
		percent:                []byte{0xd9, 0xaa},
		perMille:               []byte{0xd8, 0x89},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x46, 0x61, 0x6f, 0x69}, {0x47, 0x65, 0x61, 0x72, 0x72}, {0x4d, 0xc3, 0xa0, 0x72, 0x74}, {0x47, 0x69, 0x62, 0x6c}, {0x43, 0xc3, 0xa8, 0x69, 0x74}, {0xc3, 0x92, 0x67, 0x6d, 0x68}, {0x49, 0x75, 0x63, 0x68}, {0x4c, 0xc3, 0xb9, 0x6e, 0x61}, {0x53, 0x75, 0x6c, 0x74}, {0x44, 0xc3, 0xa0, 0x6d, 0x68}, {0x53, 0x61, 0x6d, 0x68}, {0x44, 0xc3, 0xb9, 0x62, 0x68}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x46}, {0x47}, {0x4d}, {0x47}, {0x43}, {0xc3, 0x92}, {0x49}, {0x4c}, {0x53}, {0x44}, {0x53}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x64, 0x68, 0x65, 0x6e, 0x20, 0x46, 0x68, 0x61, 0x6f, 0x69, 0x6c, 0x6c, 0x65, 0x61, 0x63, 0x68}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x47, 0x68, 0x65, 0x61, 0x72, 0x72, 0x61, 0x6e}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x4d, 0x68, 0xc3, 0xa0, 0x72, 0x74}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x47, 0x68, 0x69, 0x62, 0x6c, 0x65, 0x61, 0x6e}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x43, 0x68, 0xc3, 0xa8, 0x69, 0x74, 0x65, 0x61, 0x6e}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0xc3, 0x92, 0x67, 0x6d, 0x68, 0x69, 0x6f, 0x73}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x49, 0x75, 0x63, 0x68, 0x61, 0x72}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x4c, 0xc3, 0xb9, 0x6e, 0x61, 0x73, 0x74, 0x61, 0x6c}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x2d, 0x53, 0x75, 0x6c, 0x74, 0x61, 0x69, 0x6e}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x44, 0xc3, 0xa0, 0x6d, 0x68, 0x61, 0x69, 0x72}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x2d, 0x53, 0x61, 0x6d, 0x68, 0x61, 0x69, 0x6e}, {0x64, 0x68, 0x65, 0x6e, 0x20, 0x44, 0xc3, 0xb9, 0x62, 0x68, 0x6c, 0x61, 0x63, 0x68, 0x64}},
		daysAbbreviated:        [][]uint8{{0x44, 0x69, 0x44}, {0x44, 0x69, 0x4c}, {0x44, 0x69, 0x4d}, {0x44, 0x69, 0x43}, {0x44, 0x69, 0x61}, {0x44, 0x69, 0x68}, {0x44, 0x69, 0x53}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x43}, {0x41}, {0x48}, {0x53}},
		daysShort:              [][]uint8{{0x44, 0xc3, 0xb2}, {0x4c, 0x75}, {0x4d, 0xc3, 0xa0}, {0x43, 0x69}, {0x44, 0x61}, {0x68, 0x41}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x44, 0x69, 0x44, 0xc3, 0xb2, 0x6d, 0x68, 0x6e, 0x61, 0x69, 0x63, 0x68}, {0x44, 0x69, 0x4c, 0x75, 0x61, 0x69, 0x6e}, {0x44, 0x69, 0x4d, 0xc3, 0xa0, 0x69, 0x72, 0x74}, {0x44, 0x69, 0x43, 0x69, 0x61, 0x64, 0x61, 0x69, 0x6e}, {0x44, 0x69, 0x61, 0x72, 0x44, 0x61, 0x6f, 0x69, 0x6e}, {0x44, 0x69, 0x68, 0x41, 0x6f, 0x69, 0x6e, 0x65}, {0x44, 0x69, 0x53, 0x61, 0x74, 0x68, 0x61, 0x69, 0x72, 0x6e, 0x65}},
		periodsAbbreviated:     [][]uint8{{0x6d}, {0x66}},
		periodsNarrow:          [][]uint8{{0x6d}, {0x66}},
		periodsWide:            [][]uint8{{0x6d}, {0x66}},
		erasAbbreviated:        [][]uint8{{0x52, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{{0x52}, {0x41}},
		erasWide:               [][]uint8{{0x52, 0x6f, 0x20, 0x43, 0x68, 0x72, 0xc3, 0xac, 0x6f, 0x73, 0x74, 0x61}, {0x41, 0x6e, 0x20, 0x64, 0xc3, 0xa8, 0x69, 0x64, 0x68, 0x20, 0x43, 0x68, 0x72, 0xc3, 0xac, 0x6f, 0x73, 0x74, 0x61}},
		timezones:              map[string][]uint8{"WESZ": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "GFT": {0xc3, 0x80, 0x6d, 0x20, 0x47, 0x75, 0x69, 0x64, 0x68, 0x65, 0xc3, 0xa0, 0x6e, 0x61, 0x20, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x69, 0x6e, 0x67, 0x65}, "HKT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "HAST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x45, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x61, 0x6e, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0xe2, 0x80, 0x99, 0x69, 0x20, 0xe2, 0x80, 0x99, 0x73, 0x20, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x61, 0x63, 0x68}, "MEZ": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61}, "ACDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61}, "GYT": {0xc3, 0x80, 0x6d, 0x20, 0x47, 0x75, 0x69, 0x64, 0x68, 0x65, 0xc3, 0xa0, 0x6e, 0x61}, "AKDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "BOT": {0xc3, 0x80, 0x6d, 0x20, 0x42, 0x6f, 0x69, 0x6c, 0x69, 0x62, 0x68, 0x69, 0x61}, "UYST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x64, 0x68}, "UYT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69, 0x64, 0x68}, "WAT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x41, 0x66, 0x72, 0x61, 0x67, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "CLT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x65}, "ACST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61}, "AEST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "HNT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x54, 0x61, 0x6c, 0x61, 0x6d, 0x68, 0x20, 0x61, 0x6e, 0x20, 0xc3, 0x88, 0x69, 0x73, 0x67}, "CLST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x65}, "ARST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x2d, 0x41, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65}, "SRT": {0xc3, 0x80, 0x6d, 0x20, 0x53, 0x75, 0x72, 0x61, 0x6e, 0x61, 0x69, 0x6d}, "CDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x69, 0x6d, 0x65, 0x69, 0x72, 0x65, 0x61, 0x67, 0x61, 0x20, 0x61, 0x20, 0x54, 0x75, 0x61, 0x74, 0x68}, "WIT": {0xc3, 0x80, 0x6d, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x49, 0x6e, 0x6e, 0x64, 0x2d, 0x49, 0x6e, 0x6e, 0x73, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "ChST": {0xc3, 0x80, 0x6d, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "AKST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "VET": {0xc3, 0x80, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x42, 0x68, 0x65, 0x69, 0x6e, 0x69, 0x73, 0x65, 0x61, 0x6c, 0x61}, "OESZ": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "ADT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x61, 0xe2, 0x80, 0x99, 0x20, 0x43, 0x68, 0x75, 0x61, 0x69, 0x6e, 0x20, 0x53, 0x69, 0x61, 0x72}, "HADT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x45, 0x69, 0x6c, 0x65, 0x61, 0x6e, 0x61, 0x6e, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0xe2, 0x80, 0x99, 0x69, 0x20, 0xe2, 0x80, 0x99, 0x73, 0x20, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x61, 0x63, 0x68}, "MST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4d, 0x61, 0x63, 0xc3, 0xa0, 0x74, 0x68, 0x75}, "EDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x69, 0x6d, 0x65, 0x69, 0x72, 0x65, 0x61, 0x67, 0x61, 0x20, 0x61, 0x20, 0x54, 0x75, 0x61, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "WAST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x66, 0x72, 0x61, 0x67, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "WIB": {0xc3, 0x80, 0x6d, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x49, 0x6e, 0x6e, 0x64, 0x2d, 0x49, 0x6e, 0x6e, 0x73, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "LHST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "∅∅∅": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x42, 0x68, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61}, "PDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x61, 0xe2, 0x80, 0x99, 0x20, 0x43, 0x68, 0x75, 0x61, 0x69, 0x6e, 0x20, 0x53, 0xc3, 0xa8, 0x69, 0x6d, 0x68}, "AWST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "BT": {0xc3, 0x80, 0x6d, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa0, 0x69, 0x6e}, "COT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x69, 0x6d, 0x62, 0x69, 0x61}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AWDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "TMT": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x61, 0x73, 0x74, 0xc3, 0xa0, 0x69, 0x6e}, "CAT": {0xc3, 0x80, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x66, 0x72, 0x61, 0x67, 0x61}, "IST": {0xc3, 0x80, 0x6d, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x49, 0x6e, 0x6e, 0x73, 0x65, 0x61, 0x63, 0x68, 0x61, 0x6e}, "ART": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x2d, 0x41, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65}, "PST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x61, 0xe2, 0x80, 0x99, 0x20, 0x43, 0x68, 0x75, 0x61, 0x69, 0x6e, 0x20, 0x53, 0xc3, 0xa8, 0x69, 0x6d, 0x68}, "WITA": {0xc3, 0x80, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x6e, 0x61, 0x6e, 0x20, 0x49, 0x6e, 0x6e, 0x64, 0x2d, 0x49, 0x6e, 0x6e, 0x73, 0x65}, "ACWDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "JDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x65, 0x61, 0x70, 0x61, 0x69, 0x6e, 0x65}, "NZST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x53, 0x68, 0x65, 0x61, 0x6c, 0x61, 0x69, 0x6e, 0x6e, 0x20, 0x4e, 0x75, 0x61, 0x69, 0x64, 0x68}, "SGT": {0xc3, 0x80, 0x6d, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x65, 0x61, 0x70, 0xc3, 0xb2, 0x72}, "WART": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x2d, 0x41, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x20, 0x53, 0x69, 0x61, 0x72, 0x61, 0x69, 0x63, 0x68}, "WARST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x2d, 0x41, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x20, 0x53, 0x69, 0x61, 0x72, 0x61, 0x69, 0x63, 0x68}, "AEDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "MYT": {0xc3, 0x80, 0x6d, 0x20, 0x4d, 0x68, 0x61, 0x6c, 0x61, 0x69, 0x64, 0x68, 0x73, 0x65, 0x61}, "OEZ": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "MESZ": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61}, "HAT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x54, 0x61, 0x6c, 0x61, 0x6d, 0x68, 0x20, 0x61, 0x6e, 0x20, 0xc3, 0x88, 0x69, 0x73, 0x67}, "LHDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WEZ": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x52, 0x6f, 0x69, 0x6e, 0x6e, 0x2d, 0x45, 0xc3, 0xb2, 0x72, 0x70, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "EST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x41, 0x69, 0x6d, 0x65, 0x69, 0x72, 0x65, 0x61, 0x67, 0x61, 0x20, 0x61, 0x20, 0x54, 0x75, 0x61, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "ACWST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x69, 0x6c, 0x69, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x49, 0x61, 0x72}, "HKST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CHADT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "EAT": {0xc3, 0x80, 0x6d, 0x20, 0x41, 0x66, 0x72, 0x61, 0x67, 0x61, 0x20, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x72}, "MDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x61, 0x63, 0xc3, 0xa0, 0x74, 0x68, 0x75}, "COST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x69, 0x6d, 0x62, 0x69, 0x61}, "AST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x61, 0xe2, 0x80, 0x99, 0x20, 0x43, 0x68, 0x75, 0x61, 0x69, 0x6e, 0x20, 0x53, 0x69, 0x61, 0x72}, "NZDT": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x53, 0x68, 0x65, 0x61, 0x6c, 0x61, 0x69, 0x6e, 0x6e, 0x20, 0x4e, 0x75, 0x61, 0x69, 0x64, 0x68}, "TMST": {0x54, 0xc3, 0xac, 0x64, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6d, 0x61, 0x6e, 0x61, 0x73, 0x74, 0xc3, 0xa0, 0x69, 0x6e}, "ECT": {0xc3, 0x80, 0x6d, 0x20, 0x45, 0x61, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x69, 0x72}, "JST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x65, 0x61, 0x70, 0x61, 0x69, 0x6e, 0x65}, "SAST": {0xc3, 0x80, 0x6d, 0x20, 0x41, 0x66, 0x72, 0x61, 0x67, 0x61, 0x20, 0x61, 0x20, 0x44, 0x65, 0x61, 0x73}, "CST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x4d, 0x65, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x20, 0x41, 0x69, 0x6d, 0x65, 0x69, 0x72, 0x65, 0x61, 0x67, 0x61, 0x20, 0x61, 0x20, 0x54, 0x75, 0x61, 0x74, 0x68}, "CHAST": {0x42, 0x75, 0x6e, 0x2d, 0xc3, 0xa0, 0x6d, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}},
	}
}

// Locale returns the current translators string locale
func (gd *gd_GB) Locale() string {
	return gd.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gd_GB'
func (gd *gd_GB) PluralsCardinal() []locales.PluralRule {
	return gd.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gd_GB'
func (gd *gd_GB) PluralsOrdinal() []locales.PluralRule {
	return gd.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gd_GB'
func (gd *gd_GB) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 11 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 12 {
		return locales.PluralRuleTwo
	} else if (n >= 3 && n <= 10) || (n >= 13 && n <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gd_GB'
func (gd *gd_GB) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gd_GB'
func (gd *gd_GB) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (gd *gd_GB) MonthAbbreviated(month time.Month) []byte {
	return gd.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (gd *gd_GB) MonthsAbbreviated() [][]byte {
	return gd.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (gd *gd_GB) MonthNarrow(month time.Month) []byte {
	return gd.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (gd *gd_GB) MonthsNarrow() [][]byte {
	return gd.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (gd *gd_GB) MonthWide(month time.Month) []byte {
	return gd.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (gd *gd_GB) MonthsWide() [][]byte {
	return gd.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (gd *gd_GB) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return gd.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (gd *gd_GB) WeekdaysAbbreviated() [][]byte {
	return gd.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (gd *gd_GB) WeekdayNarrow(weekday time.Weekday) []byte {
	return gd.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (gd *gd_GB) WeekdaysNarrow() [][]byte {
	return gd.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (gd *gd_GB) WeekdayShort(weekday time.Weekday) []byte {
	return gd.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (gd *gd_GB) WeekdaysShort() [][]byte {
	return gd.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (gd *gd_GB) WeekdayWide(weekday time.Weekday) []byte {
	return gd.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (gd *gd_GB) WeekdaysWide() [][]byte {
	return gd.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'gd_GB' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gd.decimal) + len(gd.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(gd.decimal) - 1; j >= 0; j-- {
				b = append(b, gd.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gd.group) - 1; j >= 0; j-- {
					b = append(b, gd.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(gd.minus) - 1; j >= 0; j-- {
			b = append(b, gd.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'gd_GB' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (gd *gd_GB) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gd.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(gd.decimal) - 1; j >= 0; j-- {
				b = append(b, gd.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(gd.minus) - 1; j >= 0; j-- {
			b = append(b, gd.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, gd.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gd.currencies[currency]
	l := len(s) + len(gd.decimal) + len(gd.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(gd.decimal) - 1; j >= 0; j-- {
				b = append(b, gd.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gd.group) - 1; j >= 0; j-- {
					b = append(b, gd.group[j])
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
		for j := len(gd.minus) - 1; j >= 0; j-- {
			b = append(b, gd.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'gd_GB'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gd.currencies[currency]
	l := len(s) + len(gd.decimal) + len(gd.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(gd.decimal) - 1; j >= 0; j-- {
				b = append(b, gd.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(gd.group) - 1; j >= 0; j-- {
					b = append(b, gd.group[j])
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

		for j := len(gd.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, gd.currencyNegativePrefix[j])
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
			b = append(b, gd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, gd.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gd.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x64, 0x27, 0x6d, 0x68, 0x27, 0x20}...)
	b = append(b, gd.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, gd.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20, 0x64, 0x27, 0x6d, 0x68, 0x27, 0x20}...)
	b = append(b, gd.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'gd_GB'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gd *gd_GB) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := gd.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
