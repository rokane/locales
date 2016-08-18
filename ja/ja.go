package ja

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ja struct {
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

// New returns a new instance of translator for the 'ja' locale
func New() locales.Translator {
	return &ja{
		locale:                 "ja",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0xe5, 0x85, 0x83}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xef, 0xbf, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x31, 0xe6, 0x9c, 0x88}, {0x32, 0xe6, 0x9c, 0x88}, {0x33, 0xe6, 0x9c, 0x88}, {0x34, 0xe6, 0x9c, 0x88}, {0x35, 0xe6, 0x9c, 0x88}, {0x36, 0xe6, 0x9c, 0x88}, {0x37, 0xe6, 0x9c, 0x88}, {0x38, 0xe6, 0x9c, 0x88}, {0x39, 0xe6, 0x9c, 0x88}, {0x31, 0x30, 0xe6, 0x9c, 0x88}, {0x31, 0x31, 0xe6, 0x9c, 0x88}, {0x31, 0x32, 0xe6, 0x9c, 0x88}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}, {0x37}, {0x38}, {0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x31, 0xe6, 0x9c, 0x88}, {0x32, 0xe6, 0x9c, 0x88}, {0x33, 0xe6, 0x9c, 0x88}, {0x34, 0xe6, 0x9c, 0x88}, {0x35, 0xe6, 0x9c, 0x88}, {0x36, 0xe6, 0x9c, 0x88}, {0x37, 0xe6, 0x9c, 0x88}, {0x38, 0xe6, 0x9c, 0x88}, {0x39, 0xe6, 0x9c, 0x88}, {0x31, 0x30, 0xe6, 0x9c, 0x88}, {0x31, 0x31, 0xe6, 0x9c, 0x88}, {0x31, 0x32, 0xe6, 0x9c, 0x88}},
		daysAbbreviated:        [][]uint8{{0xe6, 0x97, 0xa5}, {0xe6, 0x9c, 0x88}, {0xe7, 0x81, 0xab}, {0xe6, 0xb0, 0xb4}, {0xe6, 0x9c, 0xa8}, {0xe9, 0x87, 0x91}, {0xe5, 0x9c, 0x9f}},
		daysNarrow:             [][]uint8{{0xe6, 0x97, 0xa5}, {0xe6, 0x9c, 0x88}, {0xe7, 0x81, 0xab}, {0xe6, 0xb0, 0xb4}, {0xe6, 0x9c, 0xa8}, {0xe9, 0x87, 0x91}, {0xe5, 0x9c, 0x9f}},
		daysShort:              [][]uint8{{0xe6, 0x97, 0xa5}, {0xe6, 0x9c, 0x88}, {0xe7, 0x81, 0xab}, {0xe6, 0xb0, 0xb4}, {0xe6, 0x9c, 0xa8}, {0xe9, 0x87, 0x91}, {0xe5, 0x9c, 0x9f}},
		daysWide:               [][]uint8{{0xe6, 0x97, 0xa5, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe6, 0x9c, 0x88, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe7, 0x81, 0xab, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe6, 0xb0, 0xb4, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe6, 0x9c, 0xa8, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe9, 0x87, 0x91, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}, {0xe5, 0x9c, 0x9f, 0xe6, 0x9b, 0x9c, 0xe6, 0x97, 0xa5}},
		periodsAbbreviated:     [][]uint8{{0xe5, 0x8d, 0x88, 0xe5, 0x89, 0x8d}, {0xe5, 0x8d, 0x88, 0xe5, 0xbe, 0x8c}},
		periodsNarrow:          [][]uint8{{0xe5, 0x8d, 0x88, 0xe5, 0x89, 0x8d}, {0xe5, 0x8d, 0x88, 0xe5, 0xbe, 0x8c}},
		periodsWide:            [][]uint8{{0xe5, 0x8d, 0x88, 0xe5, 0x89, 0x8d}, {0xe5, 0x8d, 0x88, 0xe5, 0xbe, 0x8c}},
		erasAbbreviated:        [][]uint8{{0xe7, 0xb4, 0x80, 0xe5, 0x85, 0x83, 0xe5, 0x89, 0x8d}, {0xe8, 0xa5, 0xbf, 0xe6, 0x9a, 0xa6}},
		erasNarrow:             [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		erasWide:               [][]uint8{{0xe7, 0xb4, 0x80, 0xe5, 0x85, 0x83, 0xe5, 0x89, 0x8d}, {0xe8, 0xa5, 0xbf, 0xe6, 0x9a, 0xa6}},
		timezones:              map[string][]uint8{"ACDT": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AEST": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe6, 0x9d, 0xb1, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "SAST": {0xe5, 0x8d, 0x97, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x95, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "WART": {0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xbc, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x81, 0xe3, 0x83, 0xb3, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "MDT": {0xe3, 0x83, 0x9e, 0xe3, 0x82, 0xab, 0xe3, 0x82, 0xaa, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "ADT": {0xe5, 0xa4, 0xa7, 0xe8, 0xa5, 0xbf, 0xe6, 0xb4, 0x8b, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "HKST": {0xe9, 0xa6, 0x99, 0xe6, 0xb8, 0xaf, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "CHAST": {0xe3, 0x83, 0x81, 0xe3, 0x83, 0xa3, 0xe3, 0x82, 0xbf, 0xe3, 0x83, 0xa0, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "WARST": {0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xbc, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x81, 0xe3, 0x83, 0xb3, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "PST": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe5, 0xa4, 0xaa, 0xe5, 0xb9, 0xb3, 0xe6, 0xb4, 0x8b, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "MEZ": {0xe4, 0xb8, 0xad, 0xe5, 0xa4, 0xae, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "∅∅∅": {0xe3, 0x82, 0xa2, 0xe3, 0x82, 0xaf, 0xe3, 0x83, 0xac, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AWST": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "WIB": {0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0x8d, 0xe3, 0x82, 0xb7, 0xe3, 0x82, 0xa2, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "EST": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0x9d, 0xb1, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "UYST": {0xe3, 0x82, 0xa6, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xb0, 0xe3, 0x82, 0xa2, 0xe3, 0x82, 0xa4, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "LHST": {0xe3, 0x83, 0xad, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0x8f, 0xe3, 0x82, 0xa6, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "WEZ": {0xe8, 0xa5, 0xbf, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "PDT": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe5, 0xa4, 0xaa, 0xe5, 0xb9, 0xb3, 0xe6, 0xb4, 0x8b, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "NZDT": {0xe3, 0x83, 0x8b, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "GYT": {0xe3, 0x82, 0xac, 0xe3, 0x82, 0xa4, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x8a, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AKST": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa9, 0xe3, 0x82, 0xb9, 0xe3, 0x82, 0xab, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "BOT": {0xe3, 0x83, 0x9c, 0xe3, 0x83, 0xaa, 0xe3, 0x83, 0x93, 0xe3, 0x82, 0xa2, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "EDT": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0x9d, 0xb1, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "WITA": {0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0x8d, 0xe3, 0x82, 0xb7, 0xe3, 0x82, 0xa2, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "WIT": {0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0x8d, 0xe3, 0x82, 0xb7, 0xe3, 0x82, 0xa2, 0xe6, 0x9d, 0xb1, 0xe9, 0x83, 0xa8, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AWDT": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "LHDT": {0xe3, 0x83, 0xad, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0x8f, 0xe3, 0x82, 0xa6, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "SGT": {0xe3, 0x82, 0xb7, 0xe3, 0x83, 0xb3, 0xe3, 0x82, 0xac, 0xe3, 0x83, 0x9d, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xab, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "MST": {0xe3, 0x83, 0x9e, 0xe3, 0x82, 0xab, 0xe3, 0x82, 0xaa, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "ACWDT": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe4, 0xb8, 0xad, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "TMT": {0xe3, 0x83, 0x88, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xaf, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0x8b, 0xe3, 0x82, 0xb9, 0xe3, 0x82, 0xbf, 0xe3, 0x83, 0xb3, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "HNT": {0xe3, 0x83, 0x8b, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x95, 0xe3, 0x82, 0xa1, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "EAT": {0xe6, 0x9d, 0xb1, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x95, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "COST": {0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x93, 0xe3, 0x82, 0xa2, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "ChST": {0xe3, 0x83, 0x81, 0xe3, 0x83, 0xa3, 0xe3, 0x83, 0xa2, 0xe3, 0x83, 0xad, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "CHADT": {0xe3, 0x83, 0x81, 0xe3, 0x83, 0xa3, 0xe3, 0x82, 0xbf, 0xe3, 0x83, 0xa0, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AKDT": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa9, 0xe3, 0x82, 0xb9, 0xe3, 0x82, 0xab, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "GFT": {0xe4, 0xbb, 0x8f, 0xe9, 0xa0, 0x98, 0xe3, 0x82, 0xae, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x8a, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "HKT": {0xe9, 0xa6, 0x99, 0xe6, 0xb8, 0xaf, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "ECT": {0xe3, 0x82, 0xa8, 0xe3, 0x82, 0xaf, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0xab, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "CST": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "NZST": {0xe3, 0x83, 0x8b, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "ARST": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xbc, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x81, 0xe3, 0x83, 0xb3, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "CLST": {0xe3, 0x83, 0x81, 0xe3, 0x83, 0xaa, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "JDT": {0xe6, 0x97, 0xa5, 0xe6, 0x9c, 0xac, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "HAT": {0xe3, 0x83, 0x8b, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x95, 0xe3, 0x82, 0xa1, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "GMT": {0xe3, 0x82, 0xb0, 0xe3, 0x83, 0xaa, 0xe3, 0x83, 0x8b, 0xe3, 0x83, 0x83, 0xe3, 0x82, 0xb8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "WESZ": {0xe8, 0xa5, 0xbf, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "COT": {0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x93, 0xe3, 0x82, 0xa2, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "ACWST": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe4, 0xb8, 0xad, 0xe8, 0xa5, 0xbf, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "AEDT": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe6, 0x9d, 0xb1, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "WAST": {0xe8, 0xa5, 0xbf, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x95, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "ACST": {0xe3, 0x82, 0xaa, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb9, 0xe3, 0x83, 0x88, 0xe3, 0x83, 0xa9, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xa2, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "SRT": {0xe3, 0x82, 0xb9, 0xe3, 0x83, 0xaa, 0xe3, 0x83, 0x8a, 0xe3, 0x83, 0xa0, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "VET": {0xe3, 0x83, 0x99, 0xe3, 0x83, 0x8d, 0xe3, 0x82, 0xba, 0xe3, 0x82, 0xa8, 0xe3, 0x83, 0xa9, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "OESZ": {0xe6, 0x9d, 0xb1, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "AST": {0xe5, 0xa4, 0xa7, 0xe8, 0xa5, 0xbf, 0xe6, 0xb4, 0x8b, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "IST": {0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x89, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "HADT": {0xe3, 0x83, 0x8f, 0xe3, 0x83, 0xaf, 0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xbb, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xaa, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb7, 0xe3, 0x83, 0xa3, 0xe3, 0x83, 0xb3, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "JST": {0xe6, 0x97, 0xa5, 0xe6, 0x9c, 0xac, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "BT": {0xe3, 0x83, 0x96, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xbf, 0xe3, 0x83, 0xb3, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "CLT": {0xe3, 0x83, 0x81, 0xe3, 0x83, 0xaa, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "ART": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xbc, 0xe3, 0x83, 0xb3, 0xe3, 0x83, 0x81, 0xe3, 0x83, 0xb3, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "CDT": {0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe4, 0xb8, 0xad, 0xe9, 0x83, 0xa8, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "UYT": {0xe3, 0x82, 0xa6, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xb0, 0xe3, 0x82, 0xa2, 0xe3, 0x82, 0xa4, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "TMST": {0xe3, 0x83, 0x88, 0xe3, 0x83, 0xab, 0xe3, 0x82, 0xaf, 0xe3, 0x83, 0xa1, 0xe3, 0x83, 0x8b, 0xe3, 0x82, 0xb9, 0xe3, 0x82, 0xbf, 0xe3, 0x83, 0xb3, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "MESZ": {0xe4, 0xb8, 0xad, 0xe5, 0xa4, 0xae, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe5, 0xa4, 0x8f, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "HAST": {0xe3, 0x83, 0x8f, 0xe3, 0x83, 0xaf, 0xe3, 0x82, 0xa4, 0xe3, 0x83, 0xbb, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0xaa, 0xe3, 0x83, 0xa5, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb7, 0xe3, 0x83, 0xa3, 0xe3, 0x83, 0xb3, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "MYT": {0xe3, 0x83, 0x9e, 0xe3, 0x83, 0xac, 0xe3, 0x83, 0xbc, 0xe3, 0x82, 0xb7, 0xe3, 0x82, 0xa2, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "OEZ": {0xe6, 0x9d, 0xb1, 0xe3, 0x83, 0xa8, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0xad, 0xe3, 0x83, 0x83, 0xe3, 0x83, 0x91, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}, "CAT": {0xe4, 0xb8, 0xad, 0xe5, 0xa4, 0xae, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x95, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0x99, 0x82, 0xe9, 0x96, 0x93}, "WAT": {0xe8, 0xa5, 0xbf, 0xe3, 0x82, 0xa2, 0xe3, 0x83, 0x95, 0xe3, 0x83, 0xaa, 0xe3, 0x82, 0xab, 0xe6, 0xa8, 0x99, 0xe6, 0xba, 0x96, 0xe6, 0x99, 0x82}},
	}
}

// Locale returns the current translators string locale
func (ja *ja) Locale() string {
	return ja.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ja'
func (ja *ja) PluralsCardinal() []locales.PluralRule {
	return ja.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ja'
func (ja *ja) PluralsOrdinal() []locales.PluralRule {
	return ja.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ja'
func (ja *ja) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ja'
func (ja *ja) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ja'
func (ja *ja) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ja *ja) MonthAbbreviated(month time.Month) []byte {
	return ja.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ja *ja) MonthsAbbreviated() [][]byte {
	return ja.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ja *ja) MonthNarrow(month time.Month) []byte {
	return ja.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ja *ja) MonthsNarrow() [][]byte {
	return ja.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ja *ja) MonthWide(month time.Month) []byte {
	return ja.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ja *ja) MonthsWide() [][]byte {
	return ja.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ja *ja) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ja.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ja *ja) WeekdaysAbbreviated() [][]byte {
	return ja.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ja *ja) WeekdayNarrow(weekday time.Weekday) []byte {
	return ja.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ja *ja) WeekdaysNarrow() [][]byte {
	return ja.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ja *ja) WeekdayShort(weekday time.Weekday) []byte {
	return ja.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ja *ja) WeekdaysShort() [][]byte {
	return ja.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ja *ja) WeekdayWide(weekday time.Weekday) []byte {
	return ja.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ja *ja) WeekdaysWide() [][]byte {
	return ja.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ja' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ja.decimal) + len(ja.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ja' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ja *ja) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ja.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ja.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ja.currencies[currency]
	l := len(s) + len(ja.decimal) + len(ja.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
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
		b = append(b, ja.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ja.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ja'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ja.currencies[currency]
	l := len(s) + len(ja.decimal) + len(ja.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ja.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ja.group[0])
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

		for j := len(ja.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ja.currencyNegativePrefix[j])
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
			b = append(b, ja.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ja.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtDateMedium(t time.Time) []byte {

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

// FmtDateLong returns the long date representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)
	b = append(b, ja.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ja.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ja'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ja *ja) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0xe6, 0x99, 0x82}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe5, 0x88, 0x86}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe7, 0xa7, 0x92, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ja.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
