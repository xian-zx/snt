// 全角字符unicode编码从65281~65374 （十六进制 0xFF01 ~ 0xFF5E）
// 半角字符unicode编码从33~126 （十六进制 0x21~ 0x7E）
// 全角字符 - 半角字符 = 65248
// 空格比较特殊，全角为 12288（0x3000），半角为 32（0x20）
package quanban

const (
	BanStr  = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	QuanStr = "！＂＃＄％＆＇（）＊＋，－．／０１２３４５６７８９：；＜＝＞？＠ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ［＼］＾＿｀ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ｛｜｝～"
)


// 全角字符 => 半角字符
func Q2B(s string) string {
	rs := []rune(s)
	for i, r := range rs {
		if r >= 65281 && r <= 65374 {
			rs[i] = r - 65248
		} else if r == 12288 {
			rs[i] = 32
		}
	}
	return string(rs)
}

// 半角字符 => 全角字符
func B2Q(s string) string {
	rs := []rune(s)
	for i, r := range rs {
		if r >= 33 && r <= 126 {
			rs[i] = r + 65248
		} else if r == 32 {
			rs[i] = 12288
		}
	}
	return string(rs)
}