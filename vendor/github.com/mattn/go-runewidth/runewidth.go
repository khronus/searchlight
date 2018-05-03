package runewidth

var EastAsianWidth = IsEastAsian()
var DefaultCondition = &Condition{EastAsianWidth}

type interval struct {
	first rune
	last  rune
}

var combining = []interval{
	{0x0300, 0x036F}, {0x0483, 0x0486}, {0x0488, 0x0489},
	{0x0591, 0x05BD}, {0x05BF, 0x05BF}, {0x05C1, 0x05C2},
	{0x05C4, 0x05C5}, {0x05C7, 0x05C7}, {0x0600, 0x0603},
	{0x0610, 0x0615}, {0x064B, 0x065E}, {0x0670, 0x0670},
	{0x06D6, 0x06E4}, {0x06E7, 0x06E8}, {0x06EA, 0x06ED},
	{0x070F, 0x070F}, {0x0711, 0x0711}, {0x0730, 0x074A},
	{0x07A6, 0x07B0}, {0x07EB, 0x07F3}, {0x0901, 0x0902},
	{0x093C, 0x093C}, {0x0941, 0x0948}, {0x094D, 0x094D},
	{0x0951, 0x0954}, {0x0962, 0x0963}, {0x0981, 0x0981},
	{0x09BC, 0x09BC}, {0x09C1, 0x09C4}, {0x09CD, 0x09CD},
	{0x09E2, 0x09E3}, {0x0A01, 0x0A02}, {0x0A3C, 0x0A3C},
	{0x0A41, 0x0A42}, {0x0A47, 0x0A48}, {0x0A4B, 0x0A4D},
	{0x0A70, 0x0A71}, {0x0A81, 0x0A82}, {0x0ABC, 0x0ABC},
	{0x0AC1, 0x0AC5}, {0x0AC7, 0x0AC8}, {0x0ACD, 0x0ACD},
	{0x0AE2, 0x0AE3}, {0x0B01, 0x0B01}, {0x0B3C, 0x0B3C},
	{0x0B3F, 0x0B3F}, {0x0B41, 0x0B43}, {0x0B4D, 0x0B4D},
	{0x0B56, 0x0B56}, {0x0B82, 0x0B82}, {0x0BC0, 0x0BC0},
	{0x0BCD, 0x0BCD}, {0x0C3E, 0x0C40}, {0x0C46, 0x0C48},
	{0x0C4A, 0x0C4D}, {0x0C55, 0x0C56}, {0x0CBC, 0x0CBC},
	{0x0CBF, 0x0CBF}, {0x0CC6, 0x0CC6}, {0x0CCC, 0x0CCD},
	{0x0CE2, 0x0CE3}, {0x0D41, 0x0D43}, {0x0D4D, 0x0D4D},
	{0x0DCA, 0x0DCA}, {0x0DD2, 0x0DD4}, {0x0DD6, 0x0DD6},
	{0x0E31, 0x0E31}, {0x0E34, 0x0E3A}, {0x0E47, 0x0E4E},
	{0x0EB1, 0x0EB1}, {0x0EB4, 0x0EB9}, {0x0EBB, 0x0EBC},
	{0x0EC8, 0x0ECD}, {0x0F18, 0x0F19}, {0x0F35, 0x0F35},
	{0x0F37, 0x0F37}, {0x0F39, 0x0F39}, {0x0F71, 0x0F7E},
	{0x0F80, 0x0F84}, {0x0F86, 0x0F87}, {0x0F90, 0x0F97},
	{0x0F99, 0x0FBC}, {0x0FC6, 0x0FC6}, {0x102D, 0x1030},
	{0x1032, 0x1032}, {0x1036, 0x1037}, {0x1039, 0x1039},
	{0x1058, 0x1059}, {0x1160, 0x11FF}, {0x135F, 0x135F},
	{0x1712, 0x1714}, {0x1732, 0x1734}, {0x1752, 0x1753},
	{0x1772, 0x1773}, {0x17B4, 0x17B5}, {0x17B7, 0x17BD},
	{0x17C6, 0x17C6}, {0x17C9, 0x17D3}, {0x17DD, 0x17DD},
	{0x180B, 0x180D}, {0x18A9, 0x18A9}, {0x1920, 0x1922},
	{0x1927, 0x1928}, {0x1932, 0x1932}, {0x1939, 0x193B},
	{0x1A17, 0x1A18}, {0x1B00, 0x1B03}, {0x1B34, 0x1B34},
	{0x1B36, 0x1B3A}, {0x1B3C, 0x1B3C}, {0x1B42, 0x1B42},
	{0x1B6B, 0x1B73}, {0x1DC0, 0x1DCA}, {0x1DFE, 0x1DFF},
	{0x200B, 0x200F}, {0x202A, 0x202E}, {0x2060, 0x2063},
	{0x206A, 0x206F}, {0x20D0, 0x20EF}, {0x302A, 0x302F},
	{0x3099, 0x309A}, {0xA806, 0xA806}, {0xA80B, 0xA80B},
	{0xA825, 0xA826}, {0xFB1E, 0xFB1E}, {0xFE00, 0xFE0F},
	{0xFE20, 0xFE23}, {0xFEFF, 0xFEFF}, {0xFFF9, 0xFFFB},
	{0x10A01, 0x10A03}, {0x10A05, 0x10A06}, {0x10A0C, 0x10A0F},
	{0x10A38, 0x10A3A}, {0x10A3F, 0x10A3F}, {0x1D167, 0x1D169},
	{0x1D173, 0x1D182}, {0x1D185, 0x1D18B}, {0x1D1AA, 0x1D1AD},
	{0x1D242, 0x1D244}, {0xE0001, 0xE0001}, {0xE0020, 0xE007F},
	{0xE0100, 0xE01EF},
}

type ctype int

const (
	narrow ctype = iota
	ambiguous
	wide
	halfwidth
	fullwidth
	neutral
)

type intervalType struct {
	first rune
	last  rune
	ctype ctype
}

var ctypes = []intervalType{
	{0x0020, 0x007E, narrow},
	{0x00A1, 0x00A1, ambiguous},
	{0x00A2, 0x00A3, narrow},
	{0x00A4, 0x00A4, ambiguous},
	{0x00A5, 0x00A6, narrow},
	{0x00A7, 0x00A8, ambiguous},
	{0x00AA, 0x00AA, ambiguous},
	{0x00AC, 0x00AC, narrow},
	{0x00AD, 0x00AE, ambiguous},
	{0x00AF, 0x00AF, narrow},
	{0x00B0, 0x00B4, ambiguous},
	{0x00B6, 0x00BA, ambiguous},
	{0x00BC, 0x00BF, ambiguous},
	{0x00C6, 0x00C6, ambiguous},
	{0x00D0, 0x00D0, ambiguous},
	{0x00D7, 0x00D8, ambiguous},
	{0x00DE, 0x00E1, ambiguous},
	{0x00E6, 0x00E6, ambiguous},
	{0x00E8, 0x00EA, ambiguous},
	{0x00EC, 0x00ED, ambiguous},
	{0x00F0, 0x00F0, ambiguous},
	{0x00F2, 0x00F3, ambiguous},
	{0x00F7, 0x00FA, ambiguous},
	{0x00FC, 0x00FC, ambiguous},
	{0x00FE, 0x00FE, ambiguous},
	{0x0101, 0x0101, ambiguous},
	{0x0111, 0x0111, ambiguous},
	{0x0113, 0x0113, ambiguous},
	{0x011B, 0x011B, ambiguous},
	{0x0126, 0x0127, ambiguous},
	{0x012B, 0x012B, ambiguous},
	{0x0131, 0x0133, ambiguous},
	{0x0138, 0x0138, ambiguous},
	{0x013F, 0x0142, ambiguous},
	{0x0144, 0x0144, ambiguous},
	{0x0148, 0x014B, ambiguous},
	{0x014D, 0x014D, ambiguous},
	{0x0152, 0x0153, ambiguous},
	{0x0166, 0x0167, ambiguous},
	{0x016B, 0x016B, ambiguous},
	{0x01CE, 0x01CE, ambiguous},
	{0x01D0, 0x01D0, ambiguous},
	{0x01D2, 0x01D2, ambiguous},
	{0x01D4, 0x01D4, ambiguous},
	{0x01D6, 0x01D6, ambiguous},
	{0x01D8, 0x01D8, ambiguous},
	{0x01DA, 0x01DA, ambiguous},
	{0x01DC, 0x01DC, ambiguous},
	{0x0251, 0x0251, ambiguous},
	{0x0261, 0x0261, ambiguous},
	{0x02C4, 0x02C4, ambiguous},
	{0x02C7, 0x02C7, ambiguous},
	{0x02C9, 0x02CB, ambiguous},
	{0x02CD, 0x02CD, ambiguous},
	{0x02D0, 0x02D0, ambiguous},
	{0x02D8, 0x02DB, ambiguous},
	{0x02DD, 0x02DD, ambiguous},
	{0x02DF, 0x02DF, ambiguous},
	{0x0300, 0x036F, ambiguous},
	{0x0391, 0x03A2, ambiguous},
	{0x03A3, 0x03A9, ambiguous},
	{0x03B1, 0x03C1, ambiguous},
	{0x03C3, 0x03C9, ambiguous},
	{0x0401, 0x0401, ambiguous},
	{0x0410, 0x044F, ambiguous},
	{0x0451, 0x0451, ambiguous},
	{0x1100, 0x115F, wide},
	{0x2010, 0x2010, ambiguous},
	{0x2013, 0x2016, ambiguous},
	{0x2018, 0x2019, ambiguous},
	{0x201C, 0x201D, ambiguous},
	{0x2020, 0x2022, ambiguous},
	{0x2024, 0x2027, ambiguous},
	{0x2030, 0x2030, ambiguous},
	{0x2032, 0x2033, ambiguous},
	{0x2035, 0x2035, ambiguous},
	{0x203B, 0x203B, ambiguous},
	{0x203E, 0x203E, ambiguous},
	{0x2074, 0x2074, ambiguous},
	{0x207F, 0x207F, ambiguous},
	{0x2081, 0x2084, ambiguous},
	{0x20A9, 0x20A9, halfwidth},
	{0x20AC, 0x20AC, ambiguous},
	{0x2103, 0x2103, ambiguous},
	{0x2105, 0x2105, ambiguous},
	{0x2109, 0x2109, ambiguous},
	{0x2113, 0x2113, ambiguous},
	{0x2116, 0x2116, ambiguous},
	{0x2121, 0x2122, ambiguous},
	{0x2126, 0x2126, ambiguous},
	{0x212B, 0x212B, ambiguous},
	{0x2153, 0x2154, ambiguous},
	{0x215B, 0x215E, ambiguous},
	{0x2160, 0x216B, ambiguous},
	{0x2170, 0x2179, ambiguous},
	{0x2189, 0x218A, ambiguous},
	{0x2190, 0x2199, ambiguous},
	{0x21B8, 0x21B9, ambiguous},
	{0x21D2, 0x21D2, ambiguous},
	{0x21D4, 0x21D4, ambiguous},
	{0x21E7, 0x21E7, ambiguous},
	{0x2200, 0x2200, ambiguous},
	{0x2202, 0x2203, ambiguous},
	{0x2207, 0x2208, ambiguous},
	{0x220B, 0x220B, ambiguous},
	{0x220F, 0x220F, ambiguous},
	{0x2211, 0x2211, ambiguous},
	{0x2215, 0x2215, ambiguous},
	{0x221A, 0x221A, ambiguous},
	{0x221D, 0x2220, ambiguous},
	{0x2223, 0x2223, ambiguous},
	{0x2225, 0x2225, ambiguous},
	{0x2227, 0x222C, ambiguous},
	{0x222E, 0x222E, ambiguous},
	{0x2234, 0x2237, ambiguous},
	{0x223C, 0x223D, ambiguous},
	{0x2248, 0x2248, ambiguous},
	{0x224C, 0x224C, ambiguous},
	{0x2252, 0x2252, ambiguous},
	{0x2260, 0x2261, ambiguous},
	{0x2264, 0x2267, ambiguous},
	{0x226A, 0x226B, ambiguous},
	{0x226E, 0x226F, ambiguous},
	{0x2282, 0x2283, ambiguous},
	{0x2286, 0x2287, ambiguous},
	{0x2295, 0x2295, ambiguous},
	{0x2299, 0x2299, ambiguous},
	{0x22A5, 0x22A5, ambiguous},
	{0x22BF, 0x22BF, ambiguous},
	{0x2312, 0x2312, ambiguous},
	{0x2329, 0x232A, wide},
	{0x2460, 0x24E9, ambiguous},
	{0x24EB, 0x254B, ambiguous},
	{0x2550, 0x2573, ambiguous},
	{0x2580, 0x258F, ambiguous},
	{0x2592, 0x2595, ambiguous},
	{0x25A0, 0x25A1, ambiguous},
	{0x25A3, 0x25A9, ambiguous},
	{0x25B2, 0x25B3, ambiguous},
	{0x25B6, 0x25B7, ambiguous},
	{0x25BC, 0x25BD, ambiguous},
	{0x25C0, 0x25C1, ambiguous},
	{0x25C6, 0x25C8, ambiguous},
	{0x25CB, 0x25CB, ambiguous},
	{0x25CE, 0x25D1, ambiguous},
	{0x25E2, 0x25E5, ambiguous},
	{0x25EF, 0x25EF, ambiguous},
	{0x2605, 0x2606, ambiguous},
	{0x2609, 0x2609, ambiguous},
	{0x260E, 0x260F, ambiguous},
	{0x2614, 0x2615, ambiguous},
	{0x261C, 0x261C, ambiguous},
	{0x261E, 0x261E, ambiguous},
	{0x2640, 0x2640, ambiguous},
	{0x2642, 0x2642, ambiguous},
	{0x2660, 0x2661, ambiguous},
	{0x2663, 0x2665, ambiguous},
	{0x2667, 0x266A, ambiguous},
	{0x266C, 0x266D, ambiguous},
	{0x266F, 0x266F, ambiguous},
	{0x269E, 0x269F, ambiguous},
	{0x26BE, 0x26BF, ambiguous},
	{0x26C4, 0x26CD, ambiguous},
	{0x26CF, 0x26E1, ambiguous},
	{0x26E3, 0x26E3, ambiguous},
	{0x26E8, 0x26FF, ambiguous},
	{0x273D, 0x273D, ambiguous},
	{0x2757, 0x2757, ambiguous},
	{0x2776, 0x277F, ambiguous},
	{0x27E6, 0x27ED, narrow},
	{0x2985, 0x2986, narrow},
	{0x2B55, 0x2B59, ambiguous},
	{0x2E80, 0x2E9A, wide},
	{0x2E9B, 0x2EF4, wide},
	{0x2F00, 0x2FD6, wide},
	{0x2FF0, 0x2FFC, wide},
	{0x3000, 0x3000, fullwidth},
	{0x3001, 0x303E, wide},
	{0x3041, 0x3097, wide},
	{0x3099, 0x3100, wide},
	{0x3105, 0x312E, wide},
	{0x3131, 0x318F, wide},
	{0x3190, 0x31BB, wide},
	{0x31C0, 0x31E4, wide},
	{0x31F0, 0x321F, wide},
	{0x3220, 0x3247, wide},
	{0x3248, 0x324F, ambiguous},
	{0x3250, 0x32FF, wide},
	{0x3300, 0x4DBF, wide},
	{0x4E00, 0xA48D, wide},
	{0xA490, 0xA4C7, wide},
	{0xA960, 0xA97D, wide},
	{0xAC00, 0xD7A4, wide},
	{0xE000, 0xF8FF, ambiguous},
	{0xF900, 0xFAFF, wide},
	{0xFE00, 0xFE0F, ambiguous},
	{0xFE10, 0xFE1A, wide},
	{0xFE30, 0xFE53, wide},
	{0xFE54, 0xFE67, wide},
	{0xFE68, 0xFE6C, wide},
	{0xFF01, 0xFF60, fullwidth},
	{0xFF61, 0xFFBF, halfwidth},
	{0xFFC2, 0xFFC8, halfwidth},
	{0xFFCA, 0xFFD0, halfwidth},
	{0xFFD2, 0xFFD8, halfwidth},
	{0xFFDA, 0xFFDD, halfwidth},
	{0xFFE0, 0xFFE7, fullwidth},
	{0xFFE8, 0xFFEF, halfwidth},
	{0xFFFD, 0xFFFE, ambiguous},
	{0x1B000, 0x1B002, wide},
	{0x1F100, 0x1F10A, ambiguous},
	{0x1F110, 0x1F12D, ambiguous},
	{0x1F130, 0x1F169, ambiguous},
	{0x1F170, 0x1F19B, ambiguous},
	{0x1F200, 0x1F203, wide},
	{0x1F210, 0x1F23B, wide},
	{0x1F240, 0x1F249, wide},
	{0x1F250, 0x1F252, wide},
	{0x20000, 0x2FFFE, wide},
	{0x30000, 0x3FFFE, wide},
	{0xE0100, 0xE01F0, ambiguous},
	{0xF0000, 0xFFFFD, ambiguous},
	{0x100000, 0x10FFFE, ambiguous},
}

type Condition struct {
	EastAsianWidth bool
}

func NewCondition() *Condition {
	return &Condition{EastAsianWidth}
}

// RuneWidth returns the number of cells in r.
// See http://www.unicode.org/reports/tr11/
func (c *Condition) RuneWidth(r rune) int {
	if r == 0 {
		return 0
	}
	if r < 32 || (r >= 0x7f && r < 0xa0) {
		return 1
	}
	for _, iv := range combining {
		if iv.first <= r && r <= iv.last {
			return 0
		}
	}

	if c.EastAsianWidth && IsAmbiguousWidth(r) {
		return 2
	}

	if r >= 0x1100 &&
		(r <= 0x115f || r == 0x2329 || r == 0x232a ||
			(r >= 0x2e80 && r <= 0xa4cf && r != 0x303f) ||
			(r >= 0xac00 && r <= 0xd7a3) ||
			(r >= 0xf900 && r <= 0xfaff) ||
			(r >= 0xfe30 && r <= 0xfe6f) ||
			(r >= 0xff00 && r <= 0xff60) ||
			(r >= 0xffe0 && r <= 0xffe6) ||
			(r >= 0x20000 && r <= 0x2fffd) ||
			(r >= 0x30000 && r <= 0x3fffd)) {
		return 2
	}
	return 1
}

func (c *Condition) StringWidth(s string) (width int) {
	for _, r := range []rune(s) {
		width += c.RuneWidth(r)
	}
	return width
}

func (c *Condition) Truncate(s string, w int, tail string) string {
	if c.StringWidth(s) <= w {
		return s
	}
	r := []rune(s)
	tw := c.StringWidth(tail)
	w -= tw
	width := 0
	i := 0
	for ; i < len(r); i++ {
		cw := c.RuneWidth(r[i])
		if width+cw > w {
			break
		}
		width += cw
	}
	return string(r[0:i]) + tail
}

func (c *Condition) Wrap(s string, w int) string {
	width := 0
	out := ""
	for _, r := range []rune(s) {
		cw := RuneWidth(r)
		if r == '\n' {
			out += string(r)
			width = 0
			continue
		} else if width+cw > w {
			out += "\n"
			width = 0
			out += string(r)
			width += cw
			continue
		}
		out += string(r)
		width += cw
	}
	return out
}

func (c *Condition) FillLeft(s string, w int) string {
	width := c.StringWidth(s)
	count := w - width
	if count > 0 {
		b := make([]byte, count)
		for i := range b {
			b[i] = ' '
		}
		return string(b) + s
	}
	return s
}

func (c *Condition) FillRight(s string, w int) string {
	width := c.StringWidth(s)
	count := w - width
	if count > 0 {
		b := make([]byte, count)
		for i := range b {
			b[i] = ' '
		}
		return s + string(b)
	}
	return s
}

// RuneWidth returns the number of cells in r.
// See http://www.unicode.org/reports/tr11/
func RuneWidth(r rune) int {
	return DefaultCondition.RuneWidth(r)
}

func ct(r rune) ctype {
	for _, iv := range ctypes {
		if iv.first <= r && r <= iv.last {
			return iv.ctype
		}
	}
	return neutral
}

// IsAmbiguousWidth returns whether is ambiguous width or not.
func IsAmbiguousWidth(r rune) bool {
	return ct(r) == ambiguous
}

// IsAmbiguousWidth returns whether is ambiguous width or not.
func IsNeutralWidth(r rune) bool {
	return ct(r) == neutral
}

func StringWidth(s string) (width int) {
	return DefaultCondition.StringWidth(s)
}

func Truncate(s string, w int, tail string) string {
	return DefaultCondition.Truncate(s, w, tail)
}

func Wrap(s string, w int) string {
	return DefaultCondition.Wrap(s, w)
}

func FillLeft(s string, w int) string {
	return DefaultCondition.FillLeft(s, w)
}

func FillRight(s string, w int) string {
	return DefaultCondition.FillRight(s, w)
}
