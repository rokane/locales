package ki

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ki struct {
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

// New returns a new instance of translator for the 'ki' locale
func New() locales.Translator {
	return &ki{
		locale:                 "ki",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x73, 0x68}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x45, 0x4e}, {0x57, 0x4b, 0x52}, {0x57, 0x47, 0x54}, {0x57, 0x4b, 0x4e}, {0x57, 0x54, 0x4e}, {0x57, 0x54, 0x44}, {0x57, 0x4d, 0x4a}, {0x57, 0x4e, 0x4e}, {0x57, 0x4b, 0x44}, {0x57, 0x49, 0x4b}, {0x57, 0x4d, 0x57}, {0x44, 0x49, 0x54}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x4b}, {0x47}, {0x4b}, {0x47}, {0x47}, {0x4d}, {0x4b}, {0x4b}, {0x49}, {0x49}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4e, 0x6a, 0x65, 0x6e, 0x75, 0x61, 0x72, 0xc4, 0xa9}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x6b, 0x65, 0x72, 0xc4, 0xa9}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x67, 0x61, 0x74, 0x61, 0x74, 0xc5, 0xa9}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x6b, 0x61, 0x6e, 0x61}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x67, 0x61, 0x74, 0x61, 0x6e, 0x6f}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x67, 0x61, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x74, 0xc5, 0xa9}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x6d, 0xc5, 0xa9, 0x67, 0x77, 0x61, 0x6e, 0x6a, 0x61}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x61}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x6b, 0x65, 0x6e, 0x64, 0x61}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x69, 0x6b, 0xc5, 0xa9, 0x6d, 0x69}, {0x4d, 0x77, 0x65, 0x72, 0x65, 0x20, 0x77, 0x61, 0x20, 0x69, 0x6b, 0xc5, 0xa9, 0x6d, 0x69, 0x20, 0x6e, 0x61, 0x20, 0xc5, 0xa9, 0x6d, 0x77, 0x65}, {0x4e, 0x64, 0x69, 0x74, 0x68, 0x65, 0x6d, 0x62, 0x61}},
		daysAbbreviated:        [][]uint8{{0x4b, 0x4d, 0x41}, {0x4e, 0x54, 0x54}, {0x4e, 0x4d, 0x4e}, {0x4e, 0x4d, 0x54}, {0x41, 0x52, 0x54}, {0x4e, 0x4d, 0x41}, {0x4e, 0x4d, 0x4d}},
		daysNarrow:             [][]uint8{{0x4b}, {0x4e}, {0x4e}, {0x4e}, {0x41}, {0x4e}, {0x4e}},
		daysWide:               [][]uint8{{0x4b, 0x69, 0x75, 0x6d, 0x69, 0x61}, {0x4e, 0x6a, 0x75, 0x6d, 0x61, 0x74, 0x61, 0x74, 0xc5, 0xa9}, {0x4e, 0x6a, 0x75, 0x6d, 0x61, 0x69, 0x6e, 0x65}, {0x4e, 0x6a, 0x75, 0x6d, 0x61, 0x74, 0x61, 0x6e, 0x61}, {0x41, 0x72, 0x61, 0x6d, 0x69, 0x74, 0x68, 0x69}, {0x4e, 0x6a, 0x75, 0x6d, 0x61, 0x61}, {0x4e, 0x6a, 0x75, 0x6d, 0x61, 0x6d, 0x6f, 0x74, 0x68, 0x69}},
		periodsAbbreviated:     [][]uint8{{0x4b, 0x69, 0x72, 0x6f, 0x6b, 0x6f}, {0x48, 0x77, 0x61, 0xc4, 0xa9, 0x2d, 0x69, 0x6e, 0xc4, 0xa9}},
		periodsWide:            [][]uint8{{0x4b, 0x69, 0x72, 0x6f, 0x6b, 0x6f}, {0x48, 0x77, 0x61, 0xc4, 0xa9, 0x2d, 0x69, 0x6e, 0xc4, 0xa9}},
		erasAbbreviated:        [][]uint8{{0x4d, 0x4b}, {0x54, 0x4b}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x4d, 0x62, 0x65, 0x72, 0x65, 0x20, 0x79, 0x61, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x54, 0x68, 0x75, 0x74, 0x68, 0x61, 0x20, 0x77, 0x61, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:              map[string][]uint8{"SRT": {0x53, 0x52, 0x54}, "CDT": {0x43, 0x44, 0x54}, "AST": {0x41, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "SGT": {0x53, 0x47, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "CST": {0x43, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "ART": {0x41, 0x52, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "COT": {0x43, 0x4f, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "ADT": {0x41, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "HNT": {0x48, 0x4e, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "ARST": {0x41, 0x52, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "BT": {0x42, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "EST": {0x45, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "MYT": {0x4d, 0x59, 0x54}, "VET": {0x56, 0x45, 0x54}, "PDT": {0x50, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "AWST": {0x41, 0x57, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "GFT": {0x47, 0x46, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "MST": {0x4d, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "UYT": {0x55, 0x59, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "IST": {0x49, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (ki *ki) Locale() string {
	return ki.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ki'
func (ki *ki) PluralsCardinal() []locales.PluralRule {
	return ki.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ki'
func (ki *ki) PluralsOrdinal() []locales.PluralRule {
	return ki.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ki'
func (ki *ki) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ki'
func (ki *ki) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ki'
func (ki *ki) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ki *ki) MonthAbbreviated(month time.Month) []byte {
	return ki.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ki *ki) MonthsAbbreviated() [][]byte {
	return ki.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ki *ki) MonthNarrow(month time.Month) []byte {
	return ki.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ki *ki) MonthsNarrow() [][]byte {
	return ki.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ki *ki) MonthWide(month time.Month) []byte {
	return ki.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ki *ki) MonthsWide() [][]byte {
	return ki.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ki *ki) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ki.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ki *ki) WeekdaysAbbreviated() [][]byte {
	return ki.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ki *ki) WeekdayNarrow(weekday time.Weekday) []byte {
	return ki.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ki *ki) WeekdaysNarrow() [][]byte {
	return ki.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ki *ki) WeekdayShort(weekday time.Weekday) []byte {
	return ki.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ki *ki) WeekdaysShort() [][]byte {
	return ki.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ki *ki) WeekdayWide(weekday time.Weekday) []byte {
	return ki.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ki *ki) WeekdaysWide() [][]byte {
	return ki.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ki' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ki' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ki *ki) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ki.currencies[currency]
	l := len(s) + len(ki.decimal) + len(ki.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ki.decimal) - 1; j >= 0; j-- {
				b = append(b, ki.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ki.group) - 1; j >= 0; j-- {
					b = append(b, ki.group[j])
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
		for j := len(ki.minus) - 1; j >= 0; j-- {
			b = append(b, ki.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ki.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ki'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ki.currencies[currency]
	l := len(s) + len(ki.decimal) + len(ki.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ki.decimal) - 1; j >= 0; j-- {
				b = append(b, ki.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ki.group) - 1; j >= 0; j-- {
					b = append(b, ki.group[j])
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

		for j := len(ki.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ki.currencyNegativePrefix[j])
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
			b = append(b, ki.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ki.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ki.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ki.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ki.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ki.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ki'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ki *ki) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ki.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ki.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
