package smn

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type smn struct {
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
	currencyPositivePrefix []byte
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

// New returns a new instance of translator for the 'smn' locale
func New() locales.Translator {
	return &smn{
		locale:                 "smn",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0x4b},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0x4b},
		monthsWide:             [][]uint8{[]uint8(nil), {0x75, 0xc4, 0x91, 0xc4, 0x91, 0xc3, 0xa2, 0x69, 0x76, 0x65, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x6b, 0x75, 0x6f, 0x76, 0xc3, 0xa2, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x6e, 0x6a, 0x75, 0x68, 0xc4, 0x8d, 0xc3, 0xa2, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x63, 0x75, 0xc3, 0xa1, 0xc5, 0x8b, 0x75, 0x69, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x76, 0x79, 0x65, 0x73, 0x69, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x6b, 0x65, 0x73, 0x69, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x73, 0x79, 0x65, 0x69, 0x6e, 0x69, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x70, 0x6f, 0x72, 0x67, 0x65, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0xc4, 0x8d, 0x6f, 0x68, 0xc4, 0x8d, 0xc3, 0xa2, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x72, 0x6f, 0x6f, 0x76, 0x76, 0xc3, 0xa2, 0x64, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x73, 0x6b, 0x61, 0x6d, 0x6d, 0xc3, 0xa2, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}, {0x6a, 0x75, 0x6f, 0x76, 0x6c, 0xc3, 0xa2, 0x6d, 0xc3, 0xa1, 0xc3, 0xa1, 0x6e, 0x75}},
		daysAbbreviated:        [][]uint8{{0x70, 0x61}, {0x76, 0x75}, {0x6d, 0x61}, {0x6b, 0x6f}, {0x74, 0x75}, {0x76, 0xc3, 0xa1}, {0x6c, 0xc3, 0xa1}},
		daysNarrow:             [][]uint8{{0x50}, {0x56}, {0x4d}, {0x4b}, {0x54}, {0x56}, {0x4c}},
		daysWide:               [][]uint8{{0x70, 0x61, 0x73, 0x65, 0x70, 0x65, 0x65, 0x69, 0x76, 0x69}, {0x76, 0x75, 0x6f, 0x73, 0x73, 0x61, 0x61, 0x72, 0x67, 0xc3, 0xa2}, {0x6d, 0x61, 0x6a, 0x65, 0x62, 0x61, 0x61, 0x72, 0x67, 0xc3, 0xa2}, {0x6b, 0x6f, 0x73, 0x6b, 0x6f, 0x68, 0x6f}, {0x74, 0x75, 0x6f, 0x72, 0xc3, 0xa2, 0x73, 0x74, 0x75, 0x76}, {0x76, 0xc3, 0xa1, 0x73, 0x74, 0x75, 0x70, 0x70, 0x65, 0x65, 0x69, 0x76, 0x69}, {0x6c, 0xc3, 0xa1, 0x76, 0x75, 0x72, 0x64, 0x75, 0x76}},
		timezones:              map[string][]uint8{"CDT": {0x43, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "AEST": {0x41, 0x45, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "BT": {0x42, 0x54}, "EST": {0x45, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CAT": {0x43, 0x41, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "EAT": {0x45, 0x41, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "COT": {0x43, 0x4f, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "IST": {0x49, 0x53, 0x54}, "CST": {0x43, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "MST": {0x4d, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "WIB": {0x57, 0x49, 0x42}, "MYT": {0x4d, 0x59, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "SGT": {0x53, 0x47, 0x54}, "SRT": {0x53, 0x52, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "GFT": {0x47, 0x46, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "CLST": {0x43, 0x4c, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (smn *smn) Locale() string {
	return smn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'smn'
func (smn *smn) PluralsCardinal() []locales.PluralRule {
	return smn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'smn'
func (smn *smn) PluralsOrdinal() []locales.PluralRule {
	return smn.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'smn'
func (smn *smn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'smn'
func (smn *smn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'smn'
func (smn *smn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (smn *smn) MonthAbbreviated(month time.Month) []byte {
	return smn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (smn *smn) MonthsAbbreviated() [][]byte {
	return smn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (smn *smn) MonthNarrow(month time.Month) []byte {
	return smn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (smn *smn) MonthsNarrow() [][]byte {
	return smn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (smn *smn) MonthWide(month time.Month) []byte {
	return smn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (smn *smn) MonthsWide() [][]byte {
	return smn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (smn *smn) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return smn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (smn *smn) WeekdaysAbbreviated() [][]byte {
	return smn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (smn *smn) WeekdayNarrow(weekday time.Weekday) []byte {
	return smn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (smn *smn) WeekdaysNarrow() [][]byte {
	return smn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (smn *smn) WeekdayShort(weekday time.Weekday) []byte {
	return smn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (smn *smn) WeekdaysShort() [][]byte {
	return smn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (smn *smn) WeekdayWide(weekday time.Weekday) []byte {
	return smn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (smn *smn) WeekdaysWide() [][]byte {
	return smn.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'smn' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'smn' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (smn *smn) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := smn.currencies[currency]
	l := len(s) + len(smn.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(smn.decimal) - 1; j >= 0; j-- {
				b = append(b, smn.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(smn.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, smn.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(smn.minus) - 1; j >= 0; j-- {
			b = append(b, smn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, smn.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'smn'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := smn.currencies[currency]
	l := len(s) + len(smn.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(smn.decimal) - 1; j >= 0; j-- {
				b = append(b, smn.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(smn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, smn.currencyNegativePrefix[j])
		}

		for j := len(smn.minus) - 1; j >= 0; j-- {
			b = append(b, smn.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(smn.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, smn.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, smn.currencyNegativeSuffix...)
	} else {

		b = append(b, smn.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'smn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (smn *smn) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
