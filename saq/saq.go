package saq

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type saq struct {
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

// New returns a new instance of translator for the 'saq' locale
func New() locales.Translator {
	return &saq{
		locale:                 "saq",
		pluralsCardinal:        []locales.PluralRule{2, 6},
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
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4f, 0x62, 0x6f}, {0x57, 0x61, 0x61}, {0x4f, 0x6b, 0x75}, {0x4f, 0x6e, 0x67}, {0x49, 0x6d, 0x65}, {0x49, 0x6c, 0x65}, {0x53, 0x61, 0x70}, {0x49, 0x73, 0x69}, {0x53, 0x61, 0x61}, {0x54, 0x6f, 0x6d}, {0x54, 0x6f, 0x62}, {0x54, 0x6f, 0x77}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4f}, {0x57}, {0x4f}, {0x4f}, {0x49}, {0x49}, {0x53}, {0x49}, {0x53}, {0x54}, {0x54}, {0x54}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x6f, 0x62, 0x6f}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x77, 0x61, 0x61, 0x72, 0x65}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x6f, 0x6b, 0x75, 0x6e, 0x69}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x6f, 0x6e, 0x67, 0xe2, 0x80, 0x99, 0x77, 0x61, 0x6e}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x69, 0x6d, 0x65, 0x74}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x69, 0x6c, 0x65}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x70, 0x61}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x69, 0x73, 0x69, 0x65, 0x74}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x61, 0x6c}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x6d, 0x6f, 0x6e}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x6d, 0x6f, 0x6e, 0x20, 0x6f, 0x62, 0x6f}, {0x4c, 0x61, 0x70, 0x61, 0x20, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x6d, 0x6f, 0x6e, 0x20, 0x77, 0x61, 0x61, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x41, 0x72, 0x65}, {0x4b, 0x75, 0x6e}, {0x4f, 0x6e, 0x67}, {0x49, 0x6e, 0x65}, {0x49, 0x6c, 0x65}, {0x53, 0x61, 0x70}, {0x4b, 0x77, 0x65}},
		daysNarrow:             [][]uint8{{0x41}, {0x4b}, {0x4f}, {0x49}, {0x49}, {0x53}, {0x4b}},
		daysWide:               [][]uint8{{0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x61, 0x72, 0x65}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x6b, 0x75, 0x6e, 0x69}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x6f, 0x6e, 0x67, 0xe2, 0x80, 0x99, 0x77, 0x61, 0x6e}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x69, 0x6e, 0x65, 0x74}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x69, 0x6c, 0x65}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x73, 0x61, 0x70, 0x61}, {0x4d, 0x64, 0x65, 0x72, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x20, 0x6b, 0x77, 0x65}},
		periodsAbbreviated:     [][]uint8{{0x54, 0x65, 0x73, 0x69, 0x72, 0x61, 0x6e}, {0x54, 0x65, 0x69, 0x70, 0x61}},
		periodsWide:            [][]uint8{{0x54, 0x65, 0x73, 0x69, 0x72, 0x61, 0x6e}, {0x54, 0x65, 0x69, 0x70, 0x61}},
		erasAbbreviated:        [][]uint8{{0x4b, 0x4b}, {0x42, 0x4b}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x4b, 0x61, 0x62, 0x6c, 0x61, 0x20, 0x79, 0x61, 0x20, 0x43, 0x68, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x42, 0x61, 0x61, 0x64, 0x61, 0x20, 0x79, 0x61, 0x20, 0x43, 0x68, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:              map[string][]uint8{"HKT": {0x48, 0x4b, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "ART": {0x41, 0x52, 0x54}, "SRT": {0x53, 0x52, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "COT": {0x43, 0x4f, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "HAT": {0x48, 0x41, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "CDT": {0x43, 0x44, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "VET": {0x56, 0x45, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "CAT": {0x43, 0x41, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "BT": {0x42, 0x54}, "MST": {0x4d, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "ADT": {0x41, 0x44, 0x54}, "IST": {0x49, 0x53, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "WAT": {0x57, 0x41, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "CST": {0x43, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "GFT": {0x47, 0x46, 0x54}, "AST": {0x41, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "EST": {0x45, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "SGT": {0x53, 0x47, 0x54}},
	}
}

// Locale returns the current translators string locale
func (saq *saq) Locale() string {
	return saq.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'saq'
func (saq *saq) PluralsCardinal() []locales.PluralRule {
	return saq.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'saq'
func (saq *saq) PluralsOrdinal() []locales.PluralRule {
	return saq.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'saq'
func (saq *saq) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'saq'
func (saq *saq) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'saq'
func (saq *saq) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (saq *saq) MonthAbbreviated(month time.Month) []byte {
	return saq.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (saq *saq) MonthsAbbreviated() [][]byte {
	return saq.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (saq *saq) MonthNarrow(month time.Month) []byte {
	return saq.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (saq *saq) MonthsNarrow() [][]byte {
	return saq.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (saq *saq) MonthWide(month time.Month) []byte {
	return saq.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (saq *saq) MonthsWide() [][]byte {
	return saq.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (saq *saq) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return saq.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (saq *saq) WeekdaysAbbreviated() [][]byte {
	return saq.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (saq *saq) WeekdayNarrow(weekday time.Weekday) []byte {
	return saq.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (saq *saq) WeekdaysNarrow() [][]byte {
	return saq.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (saq *saq) WeekdayShort(weekday time.Weekday) []byte {
	return saq.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (saq *saq) WeekdaysShort() [][]byte {
	return saq.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (saq *saq) WeekdayWide(weekday time.Weekday) []byte {
	return saq.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (saq *saq) WeekdaysWide() [][]byte {
	return saq.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'saq' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'saq' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (saq *saq) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := saq.currencies[currency]
	l := len(s) + len(saq.decimal) + len(saq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(saq.decimal) - 1; j >= 0; j-- {
				b = append(b, saq.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(saq.group) - 1; j >= 0; j-- {
					b = append(b, saq.group[j])
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
		for j := len(saq.minus) - 1; j >= 0; j-- {
			b = append(b, saq.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, saq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'saq'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := saq.currencies[currency]
	l := len(s) + len(saq.decimal) + len(saq.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(saq.decimal) - 1; j >= 0; j-- {
				b = append(b, saq.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(saq.group) - 1; j >= 0; j-- {
					b = append(b, saq.group[j])
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

		for j := len(saq.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, saq.currencyNegativePrefix[j])
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
			b = append(b, saq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, saq.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, saq.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, saq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, saq.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, saq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'saq'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (saq *saq) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, saq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := saq.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
