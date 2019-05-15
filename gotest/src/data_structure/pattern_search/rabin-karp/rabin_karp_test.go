package rabin_karp

import (
	"testing"
	"fmt"
)


func TestKmp(t *testing.T) {
	fmt.Println(search("ABABDABACDABABCABAB", "ABABCABAB"))
	fmt.Println(search("破釜破釜舟破釜破沉舟破釜破釜沉破釜破釜", "破釜破釜沉破釜破釜"))
	fmt.Println(search("AAAAAAAAAAAAAAAAAB", "AAAAB"))
	fmt.Println(search("ABABABCABABABCABABABC", "ABABAC"))
	txt := `中国人民银行（The People's Bank Of China，英文简称PBOC），简称央行，
是中华人民共和国的中央银行，中华人民共和国国务院组成部门。在国务院领导下，制定和执行货币政策，防范和化解金融风险，
维护金融稳定。1948年12月1日，在华北银行、北海银行、西北农民银行的基础上在河北省石家庄市合并组成中国人民银行。
1983年9月，国务院决定中国人民银行专门行使中国国家中央银行职能。1995年3月18日，第八届全国人民代表大会第三次会议通过
了《中华人民共和国中国人民银行法》，至此，中国人民银行作为中央银行以法律形式被确定下来。[1]中国人民银行根据
《中华人民共和国中国人民银行法》的规定，在国务院的领导下依法独立执行货币政策，
履行职责，开展业务，不受地方政府、社会团体和个人的干涉。`
	pattern := "中国人民银行"
	matched := search(txt, pattern)
	fmt.Println(matched)
	for _, v := range matched {
		fmt.Println(string([]rune(txt)[v:v+len([]rune(pattern))]))
	}
}