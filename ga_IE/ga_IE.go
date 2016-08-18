package ga_IE

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ga_IE struct {
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

// New returns a new instance of translator for the 'ga_IE' locale
func New() locales.Translator {
	return &ga_IE{
		locale:                 "ga_IE",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x45, 0x61, 0x6e}, {0x46, 0x65, 0x61, 0x62, 0x68}, {0x4d, 0xc3, 0xa1, 0x72, 0x74, 0x61}, {0x41, 0x69, 0x62}, {0x42, 0x65, 0x61, 0x6c}, {0x4d, 0x65, 0x69, 0x74, 0x68}, {0x49, 0xc3, 0xba, 0x69, 0x6c}, {0x4c, 0xc3, 0xba, 0x6e}, {0x4d, 0x46, 0xc3, 0xb3, 0x6d, 0x68}, {0x44, 0x46, 0xc3, 0xb3, 0x6d, 0x68}, {0x53, 0x61, 0x6d, 0x68}, {0x4e, 0x6f, 0x6c, 0x6c}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x45}, {0x46}, {0x4d}, {0x41}, {0x42}, {0x4d}, {0x49}, {0x4c}, {0x4d}, {0x44}, {0x53}, {0x4e}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x45, 0x61, 0x6e, 0xc3, 0xa1, 0x69, 0x72}, {0x46, 0x65, 0x61, 0x62, 0x68, 0x72, 0x61}, {0x4d, 0xc3, 0xa1, 0x72, 0x74, 0x61}, {0x41, 0x69, 0x62, 0x72, 0x65, 0xc3, 0xa1, 0x6e}, {0x42, 0x65, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x65}, {0x4d, 0x65, 0x69, 0x74, 0x68, 0x65, 0x61, 0x6d, 0x68}, {0x49, 0xc3, 0xba, 0x69, 0x6c}, {0x4c, 0xc3, 0xba, 0x6e, 0x61, 0x73, 0x61}, {0x4d, 0x65, 0xc3, 0xa1, 0x6e, 0x20, 0x46, 0xc3, 0xb3, 0x6d, 0x68, 0x61, 0x69, 0x72}, {0x44, 0x65, 0x69, 0x72, 0x65, 0x61, 0x64, 0x68, 0x20, 0x46, 0xc3, 0xb3, 0x6d, 0x68, 0x61, 0x69, 0x72}, {0x53, 0x61, 0x6d, 0x68, 0x61, 0x69, 0x6e}, {0x4e, 0x6f, 0x6c, 0x6c, 0x61, 0x69, 0x67}},
		daysAbbreviated:        [][]uint8{{0x44, 0x6f, 0x6d, 0x68}, {0x4c, 0x75, 0x61, 0x6e}, {0x4d, 0xc3, 0xa1, 0x69, 0x72, 0x74}, {0x43, 0xc3, 0xa9, 0x61, 0x64}, {0x44, 0xc3, 0xa9, 0x61, 0x72}, {0x41, 0x6f, 0x69, 0x6e, 0x65}, {0x53, 0x61, 0x74, 0x68}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x43}, {0x44}, {0x41}, {0x53}},
		daysShort:              [][]uint8{{0x44, 0x6f}, {0x4c, 0x75}, {0x4d, 0xc3, 0xa1}, {0x43, 0xc3, 0xa9}, {0x44, 0xc3, 0xa9}, {0x41, 0x6f}, {0x53, 0x61}},
		daysWide:               [][]uint8{{0x44, 0xc3, 0xa9, 0x20, 0x44, 0x6f, 0x6d, 0x68, 0x6e, 0x61, 0x69, 0x67, 0x68}, {0x44, 0xc3, 0xa9, 0x20, 0x4c, 0x75, 0x61, 0x69, 0x6e}, {0x44, 0xc3, 0xa9, 0x20, 0x4d, 0xc3, 0xa1, 0x69, 0x72, 0x74}, {0x44, 0xc3, 0xa9, 0x20, 0x43, 0xc3, 0xa9, 0x61, 0x64, 0x61, 0x6f, 0x69, 0x6e}, {0x44, 0xc3, 0xa9, 0x61, 0x72, 0x64, 0x61, 0x6f, 0x69, 0x6e}, {0x44, 0xc3, 0xa9, 0x20, 0x68, 0x41, 0x6f, 0x69, 0x6e, 0x65}, {0x44, 0xc3, 0xa9, 0x20, 0x53, 0x61, 0x74, 0x68, 0x61, 0x69, 0x72, 0x6e}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x52, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{{0x52, 0x43}, {0x41, 0x44}},
		erasWide:               [][]uint8{{0x52, 0x6f, 0x69, 0x6d, 0x68, 0x20, 0x43, 0x68, 0x72, 0xc3, 0xad, 0x6f, 0x73, 0x74}, {0x41, 0x6e, 0x6e, 0x6f, 0x20, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x69}},
		timezones:              map[string][]uint8{"ECT": {0x41, 0x6d, 0x20, 0x45, 0x61, 0x63, 0x75, 0x61, 0x64, 0xc3, 0xb3, 0x72}, "HADT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x48, 0x61, 0x76, 0xc3, 0xa1, 0xc3, 0xad, 0x2d, 0x41, 0x69, 0x6c, 0x69, 0xc3, 0xba, 0x69, 0x74}, "UYT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x55, 0x72, 0x61, 0x67, 0x75, 0x61}, "OEZ": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "CHADT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "GMT": {0x4d, 0x65, 0xc3, 0xa1, 0x6e, 0x2d, 0x41, 0x6d, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "UYST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x55, 0x72, 0x61, 0x67, 0x75, 0x61}, "COT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb3, 0x69, 0x6d, 0x65}, "AST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x67, 0x68}, "NZDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x4e, 0x75, 0x61, 0x2d, 0x53, 0x68, 0xc3, 0xa9, 0x61, 0x6c, 0x61, 0x69, 0x6e, 0x6e, 0x65}, "CHAST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AEST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "TMT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x54, 0x75, 0x69, 0x72, 0x63, 0x6d, 0xc3, 0xa9, 0x61, 0x6e, 0x61, 0x73, 0x74, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "GYT": {0x41, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x47, 0x75, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "WAST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x66, 0x72, 0x61, 0x69, 0x63, 0x65}, "HAT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x54, 0x68, 0x61, 0x6c, 0x61, 0x6d, 0x68, 0x20, 0x61, 0x6e, 0x20, 0xc3, 0x89, 0x69, 0x73, 0x63}, "ACST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "ACDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "COST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb3, 0x69, 0x6d, 0x65}, "MESZ": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "LHST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "LHDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "VET": {0x41, 0x6d, 0x20, 0x56, 0x65, 0x69, 0x6e, 0x69, 0x73, 0xc3, 0xa9, 0x61, 0x6c, 0x61}, "WIT": {0x41, 0x6d, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x49, 0x6e, 0x64, 0x69, 0x6e, 0xc3, 0xa9, 0x69, 0x73, 0x65}, "HKST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x43, 0x6f, 0x6e, 0x67}, "JST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x65, 0x61, 0x70, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "EAT": {0x41, 0x6d, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x66, 0x72, 0x61, 0x69, 0x63, 0x65}, "CDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x6e, 0x61, 0x63, 0x68}, "OESZ": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "CLT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x65}, "WIB": {0x41, 0x6d, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x49, 0x6e, 0x64, 0x69, 0x6e, 0xc3, 0xa9, 0x69, 0x73, 0x65}, "HAST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x48, 0x61, 0x76, 0xc3, 0xa1, 0xc3, 0xad, 0x2d, 0x41, 0x69, 0x6c, 0x69, 0xc3, 0xba, 0x69, 0x74}, "MDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x68, 0x61, 0x63, 0x61, 0x6f}, "GFT": {0x41, 0x6d, 0x20, 0x47, 0x68, 0x75, 0xc3, 0xa1, 0x69, 0x6e, 0x20, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x69, 0x6e, 0x63, 0x65}, "ACWDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4d, 0x68, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "TMST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x54, 0x75, 0x69, 0x72, 0x63, 0x6d, 0xc3, 0xa9, 0x61, 0x6e, 0x61, 0x73, 0x74, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "JDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x65, 0x61, 0x70, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "CAT": {0x41, 0x6d, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x66, 0x72, 0x61, 0x69, 0x63, 0x65}, "CLST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x53, 0x69, 0x6c, 0x65}, "NZST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x4e, 0x75, 0x61, 0x2d, 0x53, 0x68, 0xc3, 0xa9, 0x61, 0x6c, 0x61, 0x69, 0x6e, 0x6e, 0x65}, "WARST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x69, 0x72, 0x67, 0x69, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x65}, "AKST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "AEDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x65, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "AWDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "WAT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x66, 0x72, 0x61, 0x69, 0x63, 0x65}, "BT": {0x41, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x42, 0xc3, 0xba, 0x74, 0xc3, 0xa1, 0x69, 0x6e, 0x65}, "WITA": {0x41, 0x6d, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x49, 0x6e, 0x64, 0x69, 0x6e, 0xc3, 0xa9, 0x69, 0x73, 0x65}, "ARST": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x69, 0x72, 0x67, 0x69, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x65}, "EST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x69, 0x72}, "PDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x69, 0x67, 0xc3, 0xa9, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x69, 0xc3, 0xba, 0x69, 0x6e}, "WART": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x69, 0x72, 0x67, 0x69, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x65}, "BOT": {0x41, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x42, 0x6f, 0x6c, 0x61, 0x69, 0x76, 0x65}, "EDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x4f, 0x69, 0x72, 0x74, 0x68, 0x69, 0x72}, "HKT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x43, 0x6f, 0x6e, 0x67}, "HNT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x54, 0x68, 0x61, 0x6c, 0x61, 0x6d, 0x68, 0x20, 0x61, 0x6e, 0x20, 0xc3, 0x89, 0x69, 0x73, 0x63}, "SGT": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x53, 0x68, 0x69, 0x6e, 0x67, 0x65, 0x61, 0x70, 0xc3, 0xb3, 0x72}, "ART": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x69, 0x72, 0x67, 0x69, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x65}, "WEZ": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "MST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4d, 0x68, 0x61, 0x63, 0x61, 0x6f}, "MYT": {0x41, 0x6d, 0x20, 0x6e, 0x61, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x65, 0x69, 0x73, 0x69, 0x61}, "ADT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x67, 0x68}, "AWST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "IST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x49, 0x6e, 0x64, 0x69, 0x61}, "∅∅∅": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x63, 0x72, 0x65}, "CST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x6e, 0x61, 0x63, 0x68}, "PST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x69, 0x67, 0xc3, 0xa9, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x69, 0xc3, 0xba, 0x69, 0x6e}, "ACWST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4d, 0x68, 0x65, 0xc3, 0xa1, 0x6e, 0x69, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x69, 0x6c, 0x65}, "SAST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x41, 0x66, 0x72, 0x61, 0x69, 0x63, 0x65, 0x20, 0x54, 0x68, 0x65, 0x61, 0x73}, "WESZ": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x49, 0x61, 0x72, 0x74, 0x68, 0x61, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "SRT": {0x41, 0x6d, 0x20, 0x53, 0x68, 0x75, 0x72, 0x61, 0x6e, 0x61, 0x6d}, "ChST": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x53, 0x65, 0x61, 0x6d, 0xc3, 0xb3, 0x72, 0x61, 0x63, 0x68}, "MEZ": {0x41, 0x6d, 0x20, 0x43, 0x61, 0x69, 0x67, 0x68, 0x64, 0x65, 0xc3, 0xa1, 0x6e, 0x61, 0x63, 0x68, 0x20, 0x4c, 0xc3, 0xa1, 0x72, 0x20, 0x6e, 0x61, 0x20, 0x68, 0x45, 0x6f, 0x72, 0x70, 0x61}, "AKDT": {0x41, 0x6d, 0x20, 0x53, 0x61, 0x6d, 0x68, 0x72, 0x61, 0x69, 0x64, 0x68, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}},
	}
}

// Locale returns the current translators string locale
func (ga *ga_IE) Locale() string {
	return ga.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ga_IE'
func (ga *ga_IE) PluralsCardinal() []locales.PluralRule {
	return ga.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ga_IE'
func (ga *ga_IE) PluralsOrdinal() []locales.PluralRule {
	return ga.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ga_IE'
func (ga *ga_IE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ga_IE'
func (ga *ga_IE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ga_IE'
func (ga *ga_IE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ga *ga_IE) MonthAbbreviated(month time.Month) []byte {
	return ga.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ga *ga_IE) MonthsAbbreviated() [][]byte {
	return ga.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ga *ga_IE) MonthNarrow(month time.Month) []byte {
	return ga.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ga *ga_IE) MonthsNarrow() [][]byte {
	return ga.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ga *ga_IE) MonthWide(month time.Month) []byte {
	return ga.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ga *ga_IE) MonthsWide() [][]byte {
	return ga.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ga *ga_IE) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ga.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ga *ga_IE) WeekdaysAbbreviated() [][]byte {
	return ga.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ga *ga_IE) WeekdayNarrow(weekday time.Weekday) []byte {
	return ga.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ga *ga_IE) WeekdaysNarrow() [][]byte {
	return ga.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ga *ga_IE) WeekdayShort(weekday time.Weekday) []byte {
	return ga.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ga *ga_IE) WeekdaysShort() [][]byte {
	return ga.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ga *ga_IE) WeekdayWide(weekday time.Weekday) []byte {
	return ga.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ga *ga_IE) WeekdaysWide() [][]byte {
	return ga.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ga_IE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ga_IE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ga *ga_IE) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ga.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ga.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ga_IE'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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

		for j := len(ga.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ga.currencyNegativePrefix[j])
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
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ga.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ga.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ga_IE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ga *ga_IE) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ga.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
