package hu

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type hu struct {
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

// New returns a new instance of translator for the 'hu' locale
func New() locales.Translator {
	return &hu{
		locale:                 "hu",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x46, 0x74}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x72, 0x2e}, {0x6d, 0xc3, 0xa1, 0x72, 0x63, 0x2e}, {0xc3, 0xa1, 0x70, 0x72, 0x2e}, {0x6d, 0xc3, 0xa1, 0x6a, 0x2e}, {0x6a, 0xc3, 0xba, 0x6e, 0x2e}, {0x6a, 0xc3, 0xba, 0x6c, 0x2e}, {0x61, 0x75, 0x67, 0x2e}, {0x73, 0x7a, 0x65, 0x70, 0x74, 0x2e}, {0x6f, 0x6b, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0xc3, 0x81}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53, 0x7a}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0xc3, 0xa1, 0x72}, {0x66, 0x65, 0x62, 0x72, 0x75, 0xc3, 0xa1, 0x72}, {0x6d, 0xc3, 0xa1, 0x72, 0x63, 0x69, 0x75, 0x73}, {0xc3, 0xa1, 0x70, 0x72, 0x69, 0x6c, 0x69, 0x73}, {0x6d, 0xc3, 0xa1, 0x6a, 0x75, 0x73}, {0x6a, 0xc3, 0xba, 0x6e, 0x69, 0x75, 0x73}, {0x6a, 0xc3, 0xba, 0x6c, 0x69, 0x75, 0x73}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x7a, 0x74, 0x75, 0x73}, {0x73, 0x7a, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x6f, 0x6b, 0x74, 0xc3, 0xb3, 0x62, 0x65, 0x72}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x56}, {0x48}, {0x4b}, {0x53, 0x7a, 0x65}, {0x43, 0x73}, {0x50}, {0x53, 0x7a, 0x6f}},
		daysNarrow:             [][]uint8{{0x56}, {0x48}, {0x4b}, {0x53, 0x7a}, {0x43, 0x73}, {0x50}, {0x53, 0x7a}},
		daysShort:              [][]uint8{{0x56}, {0x48}, {0x4b}, {0x53, 0x7a, 0x65}, {0x43, 0x73}, {0x50}, {0x53, 0x7a, 0x6f}},
		daysWide:               [][]uint8{{0x76, 0x61, 0x73, 0xc3, 0xa1, 0x72, 0x6e, 0x61, 0x70}, {0x68, 0xc3, 0xa9, 0x74, 0x66, 0xc5, 0x91}, {0x6b, 0x65, 0x64, 0x64}, {0x73, 0x7a, 0x65, 0x72, 0x64, 0x61}, {0x63, 0x73, 0xc3, 0xbc, 0x74, 0xc3, 0xb6, 0x72, 0x74, 0xc3, 0xb6, 0x6b}, {0x70, 0xc3, 0xa9, 0x6e, 0x74, 0x65, 0x6b}, {0x73, 0x7a, 0x6f, 0x6d, 0x62, 0x61, 0x74}},
		periodsAbbreviated:     [][]uint8{{0x64, 0x65, 0x2e}, {0x64, 0x75, 0x2e}},
		periodsNarrow:          [][]uint8{{0x64, 0x65, 0x2e}, {0x64, 0x75, 0x2e}},
		periodsWide:            [][]uint8{{0x64, 0x65, 0x2e}, {0x64, 0x75, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x69, 0x2e, 0x20, 0x65, 0x2e}, {0x69, 0x2e, 0x20, 0x73, 0x7a, 0x2e}},
		erasNarrow:             [][]uint8{{0x69, 0x65, 0x2e}, {0x69, 0x73, 0x7a, 0x2e}},
		erasWide:               [][]uint8{{0x69, 0x64, 0xc5, 0x91, 0x73, 0x7a, 0xc3, 0xa1, 0x6d, 0xc3, 0xad, 0x74, 0xc3, 0xa1, 0x73, 0x75, 0x6e, 0x6b, 0x20, 0x65, 0x6c, 0xc5, 0x91, 0x74, 0x74}, {0x69, 0x64, 0xc5, 0x91, 0x73, 0x7a, 0xc3, 0xa1, 0x6d, 0xc3, 0xad, 0x74, 0xc3, 0xa1, 0x73, 0x75, 0x6e, 0x6b, 0x20, 0x73, 0x7a, 0x65, 0x72, 0x69, 0x6e, 0x74}},
		timezones:              map[string][]uint8{"WIB": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "EST": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x20, 0xc3, 0xa1, 0x6c, 0x6c, 0x61, 0x6d, 0x6f, 0x6b, 0x62, 0x65, 0x6c, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "IST": {0x69, 0x6e, 0x64, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AKDT": {0x61, 0x6c, 0x61, 0x73, 0x7a, 0x6b, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WESZ": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WIT": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x69, 0x61, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ECT": {0x65, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "HAT": {0xc3, 0xba, 0x6a, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CST": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x73, 0xc5, 0x91, 0x20, 0xc3, 0xa1, 0x6c, 0x6c, 0x61, 0x6d, 0x6f, 0x6b, 0x62, 0x65, 0x6c, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "MEZ": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AEDT": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "UYST": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "TMST": {0x74, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x7a, 0x74, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "NZDT": {0xc3, 0xba, 0x6a, 0x2d, 0x7a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "SRT": {0x73, 0x7a, 0x75, 0x72, 0x69, 0x6e, 0xc3, 0xa1, 0x6d, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "MDT": {0x4d, 0x61, 0x63, 0x61, 0x75, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "SGT": {0x73, 0x7a, 0x69, 0x6e, 0x67, 0x61, 0x70, 0xc3, 0xba, 0x72, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "PDT": {0x63, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x2d, 0xc3, 0xb3, 0x63, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "MYT": {0x6d, 0x61, 0x6c, 0x61, 0x6a, 0x7a, 0x69, 0x61, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ACWDT": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "MESZ": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ACST": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "∅∅∅": {0x41, 0x63, 0x72, 0x65, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "COT": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WITA": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x69, 0x61, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WEZ": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CDT": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x73, 0xc5, 0x91, 0x20, 0xc3, 0xa1, 0x6c, 0x6c, 0x61, 0x6d, 0x6f, 0x6b, 0x62, 0x65, 0x6c, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "UYT": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AWDT": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WAT": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "JDT": {0x6a, 0x61, 0x70, 0xc3, 0xa1, 0x6e, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CHAST": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "HAST": {0x68, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "EDT": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x20, 0xc3, 0xa1, 0x6c, 0x6c, 0x61, 0x6d, 0x6f, 0x6b, 0x62, 0x65, 0x6c, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "BOT": {0x62, 0x6f, 0x6c, 0xc3, 0xad, 0x76, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ACWST": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CLST": {0x63, 0x68, 0x69, 0x6c, 0x65, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x73, 0x7a, 0x69, 0x67, 0x65, 0x74, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ART": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "OEZ": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AST": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x2d, 0xc3, 0xb3, 0x63, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "COST": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "MST": {0x4d, 0x61, 0x63, 0x61, 0x75, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "WART": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ADT": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x2d, 0xc3, 0xb3, 0x63, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "HKST": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "HADT": {0x68, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "GFT": {0x66, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x2d, 0x67, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CAT": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ChST": {0x63, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0xc3, 0xb3, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "BT": {0x62, 0x75, 0x74, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AKST": {0x61, 0x6c, 0x61, 0x73, 0x7a, 0x6b, 0x61, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "WARST": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "PST": {0x63, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x73, 0x2d, 0xc3, 0xb3, 0x63, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "HKT": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "TMT": {0x74, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x7a, 0x74, 0xc3, 0xa1, 0x6e, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "SAST": {0x64, 0xc3, 0xa9, 0x6c, 0x2d, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "WAST": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "EAT": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "JST": {0x6a, 0x61, 0x70, 0xc3, 0xa1, 0x6e, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "GMT": {0x67, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x69, 0x20, 0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x69, 0x64, 0xc5, 0x91, 0x2c, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "NZST": {0xc3, 0xba, 0x6a, 0x2d, 0x7a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "HNT": {0xc3, 0xba, 0x6a, 0x2d, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x20, 0x7a, 0xc3, 0xb3, 0x6e, 0x61, 0x69, 0x64, 0xc5, 0x91}, "CHADT": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ACDT": {0x6b, 0xc3, 0xb6, 0x7a, 0xc3, 0xa9, 0x70, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AEST": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "AWST": {0x6e, 0x79, 0x75, 0x67, 0x61, 0x74, 0x2d, 0x61, 0x75, 0x73, 0x7a, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x69, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "CLT": {0x63, 0x68, 0x69, 0x6c, 0x65, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x73, 0x7a, 0x69, 0x67, 0x65, 0x74, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "GYT": {0x67, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x69, 0x20, 0x74, 0xc3, 0xa9, 0x6c, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "OESZ": {0x6b, 0x65, 0x6c, 0x65, 0x74, 0x2d, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x61, 0x69, 0x20, 0x6e, 0x79, 0xc3, 0xa1, 0x72, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}, "VET": {0x76, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x69, 0x20, 0x69, 0x64, 0xc5, 0x91}},
	}
}

// Locale returns the current translators string locale
func (hu *hu) Locale() string {
	return hu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hu'
func (hu *hu) PluralsCardinal() []locales.PluralRule {
	return hu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hu'
func (hu *hu) PluralsOrdinal() []locales.PluralRule {
	return hu.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hu'
func (hu *hu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hu'
func (hu *hu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 5 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hu'
func (hu *hu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := hu.CardinalPluralRule(num1, v1)
	end := hu.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hu *hu) MonthAbbreviated(month time.Month) []byte {
	return hu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hu *hu) MonthsAbbreviated() [][]byte {
	return hu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hu *hu) MonthNarrow(month time.Month) []byte {
	return hu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hu *hu) MonthsNarrow() [][]byte {
	return hu.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hu *hu) MonthWide(month time.Month) []byte {
	return hu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hu *hu) MonthsWide() [][]byte {
	return hu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hu *hu) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return hu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hu *hu) WeekdaysAbbreviated() [][]byte {
	return hu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hu *hu) WeekdayNarrow(weekday time.Weekday) []byte {
	return hu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hu *hu) WeekdaysNarrow() [][]byte {
	return hu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hu *hu) WeekdayShort(weekday time.Weekday) []byte {
	return hu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hu *hu) WeekdaysShort() [][]byte {
	return hu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hu *hu) WeekdayWide(weekday time.Weekday) []byte {
	return hu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hu *hu) WeekdaysWide() [][]byte {
	return hu.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hu.decimal) + len(hu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(hu.group) - 1; j >= 0; j-- {
					b = append(b, hu.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hu *hu) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(hu.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hu.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hu.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hu.currencies[currency]
	l := len(s) + len(hu.decimal) + len(hu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(hu.group) - 1; j >= 0; j-- {
					b = append(b, hu.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hu.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hu'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hu.currencies[currency]
	l := len(s) + len(hu.decimal) + len(hu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(hu.group) - 1; j >= 0; j-- {
					b = append(b, hu.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, hu.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hu.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hu.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x2c, 0x20}...)
	b = append(b, hu.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'hu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (hu *hu) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
