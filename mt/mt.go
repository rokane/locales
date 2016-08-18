package mt

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type mt struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	decimal            []byte
	group              []byte
	minus              []byte
	percent            []byte
	perMille           []byte
	timeSeparator      []byte
	inifinity          []byte
	currencies         [][]byte // idx = enum of currency code
	monthsAbbreviated  [][]byte
	monthsNarrow       [][]byte
	monthsWide         [][]byte
	daysAbbreviated    [][]byte
	daysNarrow         [][]byte
	daysShort          [][]byte
	daysWide           [][]byte
	periodsAbbreviated [][]byte
	periodsNarrow      [][]byte
	periodsShort       [][]byte
	periodsWide        [][]byte
	erasAbbreviated    [][]byte
	erasNarrow         [][]byte
	erasWide           [][]byte
	timezones          map[string][]byte
}

// New returns a new instance of translator for the 'mt' locale
func New() locales.Translator {
	return &mt{
		locale:             "mt",
		pluralsCardinal:    []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:     nil,
		decimal:            []byte{0x2e},
		group:              []byte{0x2c},
		minus:              []byte{0x2d},
		percent:            []byte{0x25},
		perMille:           []byte{0xe2, 0x80, 0xb0},
		timeSeparator:      []byte{0x3a},
		inifinity:          []byte{0xe2, 0x88, 0x9e},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x72, 0x61}, {0x4d, 0x61, 0x72}, {0x41, 0x70, 0x72}, {0x4d, 0x65, 0x6a}, {0xc4, 0xa0, 0x75, 0x6e}, {0x4c, 0x75, 0x6c}, {0x41, 0x77, 0x77}, {0x53, 0x65, 0x74}, {0x4f, 0x74, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x69, 0xc4, 0x8b}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0xc4, 0xa0}, {0x4c}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:         [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x6e, 0x61, 0x72}, {0x46, 0x72, 0x61, 0x72}, {0x4d, 0x61, 0x72, 0x7a, 0x75}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x65, 0x6a, 0x6a, 0x75}, {0xc4, 0xa0, 0x75, 0x6e, 0x6a, 0x75}, {0x4c, 0x75, 0x6c, 0x6a, 0x75}, {0x41, 0x77, 0x77, 0x69, 0x73, 0x73, 0x75}, {0x53, 0x65, 0x74, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x75}, {0x4f, 0x74, 0x74, 0x75, 0x62, 0x72, 0x75}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x75}, {0x44, 0x69, 0xc4, 0x8b, 0x65, 0x6d, 0x62, 0x72, 0x75}},
		daysAbbreviated:    [][]uint8{{0xc4, 0xa6, 0x61, 0x64}, {0x54, 0x6e, 0x65}, {0x54, 0x6c, 0x69}, {0x45, 0x72, 0x62}, {0xc4, 0xa6, 0x61, 0x6d}, {0xc4, 0xa0, 0x69, 0x6d}, {0x53, 0x69, 0x62}},
		daysNarrow:         [][]uint8{{0xc4, 0xa6}, {0x54}, {0x54}, {0x45}, {0xc4, 0xa6}, {0xc4, 0xa0}, {0x53}},
		daysWide:           [][]uint8{{0x49, 0x6c, 0x2d, 0xc4, 0xa6, 0x61, 0x64, 0x64}, {0x49, 0x74, 0x2d, 0x54, 0x6e, 0x65, 0x6a, 0x6e}, {0x49, 0x74, 0x2d, 0x54, 0x6c, 0x69, 0x65, 0x74, 0x61}, {0x4c, 0x2d, 0x45, 0x72, 0x62, 0x67, 0xc4, 0xa7, 0x61}, {0x49, 0x6c, 0x2d, 0xc4, 0xa6, 0x61, 0x6d, 0x69, 0x73}, {0x49, 0x6c, 0x2d, 0xc4, 0xa0, 0x69, 0x6d, 0x67, 0xc4, 0xa7, 0x61}, {0x49, 0x73, 0x2d, 0x53, 0x69, 0x62, 0x74}},
		periodsAbbreviated: [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:        [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:    [][]uint8{{0x51, 0x4b}, {0x57, 0x4b}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0x51, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75}, {0x57, 0x61, 0x72, 0x61, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75}},
		timezones:          map[string][]uint8{"MST": {0x4d, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HKT": {0x48, 0x4b, 0x54}, "IST": {0x49, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "GFT": {0x47, 0x46, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "BT": {0x42, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WIB": {0x57, 0x49, 0x42}, "ARST": {0x41, 0x52, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "EAT": {0x45, 0x41, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "MESZ": {0xc4, 0xa6, 0x69, 0x6e, 0x20, 0xc4, 0x8a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x20, 0x45, 0x77, 0x72, 0x6f, 0x70, 0x65, 0x77, 0x20, 0x74, 0x61, 0x73, 0x2d, 0x53, 0x61, 0x6a, 0x66}, "ADT": {0x41, 0x44, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "MEZ": {0xc4, 0xa6, 0x69, 0x6e, 0x20, 0xc4, 0x8a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x20, 0x45, 0x77, 0x72, 0x6f, 0x70, 0x65, 0x77, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64}, "JST": {0x4a, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "COT": {0x43, 0x4f, 0x54}, "CAT": {0x43, 0x41, 0x54}, "WAT": {0x57, 0x41, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "EST": {0x45, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "CDT": {0x43, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "ACST": {0x41, 0x43, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "AEST": {0x41, 0x45, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "EDT": {0x45, 0x44, 0x54}, "PDT": {0x50, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "SRT": {0x53, 0x52, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "AST": {0x41, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "ART": {0x41, 0x52, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "CST": {0x43, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (mt *mt) Locale() string {
	return mt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mt'
func (mt *mt) PluralsCardinal() []locales.PluralRule {
	return mt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mt'
func (mt *mt) PluralsOrdinal() []locales.PluralRule {
	return mt.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mt'
func (mt *mt) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)

	if n == 1 {
		return locales.PluralRuleOne
	} else if (n == 0) || (nMod100 >= 2 && nMod100 <= 10) {
		return locales.PluralRuleFew
	} else if nMod100 >= 11 && nMod100 <= 19 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mt'
func (mt *mt) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mt'
func (mt *mt) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mt *mt) MonthAbbreviated(month time.Month) []byte {
	return mt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mt *mt) MonthsAbbreviated() [][]byte {
	return mt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mt *mt) MonthNarrow(month time.Month) []byte {
	return mt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mt *mt) MonthsNarrow() [][]byte {
	return mt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mt *mt) MonthWide(month time.Month) []byte {
	return mt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mt *mt) MonthsWide() [][]byte {
	return mt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mt *mt) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return mt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mt *mt) WeekdaysAbbreviated() [][]byte {
	return mt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mt *mt) WeekdayNarrow(weekday time.Weekday) []byte {
	return mt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mt *mt) WeekdaysNarrow() [][]byte {
	return mt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mt *mt) WeekdayShort(weekday time.Weekday) []byte {
	return mt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mt *mt) WeekdaysShort() [][]byte {
	return mt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mt *mt) WeekdayWide(weekday time.Weekday) []byte {
	return mt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mt *mt) WeekdaysWide() [][]byte {
	return mt.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(mt.decimal) + len(mt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mt.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mt *mt) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(mt.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mt.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mt.currencies[currency]
	l := len(s) + len(mt.decimal) + len(mt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mt.group[0])
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
		b = append(b, mt.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mt'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mt.currencies[currency]
	l := len(s) + len(mt.decimal) + len(mt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mt.group[0])
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

		b = append(b, mt.minus[0])

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
			b = append(b, mt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mt.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x74, 0x61, 0x27, 0xe2, 0x80, 0x99, 0x20}...)
	b = append(b, mt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, mt.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x74, 0x61, 0x27, 0xe2, 0x80, 0x99, 0x20}...)
	b = append(b, mt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'mt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mt *mt) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
