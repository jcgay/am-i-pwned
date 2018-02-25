package pwned

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestThatPasswordIsPwned(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(passwords("625DC21EF05F6AD4DDF47C5F203837AA32C:15320")) // sha1sum toto (uppercase) without 5 first chars
	}))
	defer ts.Close()
	url = ts.URL

	result, _ := CheckPassword("toto")

	if result != 15320 {
		t.Error("Expecting [15320], got: ", result)
	}
}

func TestThatPasswordIsNotPwnedWhenNoResponseFromTheAPIs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	url = ts.URL

	result, _ := CheckPassword("toto")

	if result != 0 {
		t.Error("Expecting [0], got: ", result)
	}
}

func TestThatPasswordIsNotPwnedWhenShaWithTheSamePrefixArePresent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(passwords(""))
	}))
	defer ts.Close()
	url = ts.URL

	result, _ := CheckPassword("toto")

	if result != 0 {
		t.Error("Expecting [0], got: ", result)
	}
}

func TestThatPasswordIsReadFromArgs(t *testing.T) {
	result := SelectPassword([]string{"toto"}, strings.NewReader("titi"))

	if result != "toto" {
		t.Error("Expecting [toto], got: ", result)
	}
}

func TestThatPasswordIsReadFromReader(t *testing.T) {
	result := SelectPassword([]string{}, strings.NewReader("titi"))

	if result != "titi" {
		t.Error("Expecting [titi], got: ", result)
	}
}

func passwords(shaPasswd string) []byte {
	resp := `0008A6A4913DCE9EA2BEF992928EA06211F:825
0040F96632FCBA25C3061346E5F47396A6B:1
01737C05FCB6D893A9F99FC73FC48BF3EF7:3
01B9511761BE1FBA925F6975E9DA6784100:3
03874ED288CDCC2EA00108B034C3000FE4F:2
03E3D034578F78A1E3625774C2BF2CB48ED:1
043288CF004EFBFDE0D15DDD0B50DB337F5:19
0483D534FF4C5A1C6AE7E81D9838CE1B34A:2
048DF07E35DFF0873B8040BAF2E718C0A64:2
04EACCD3B9F6CBB950F49DDA681C3DF341B:2
0511561A7167F4398F853D26D3601348F56:2
05817AE51A573DC932B684FC89D798013CD:1
071BA726E95BB46B804BBA2ABE86A1DC9A5:2
094B33D99DA9CBDE4F79426A6A670195448:8
09A47ECBEC4CBFF5D2FCE32D237D566E234:1
0ADB932107C7E608242E79FE535C1B3D648:1
0B26401E5F21F4441149DA6F344EEC696A8:2
0CE5DEE980524B3EA859774C0D9EF0D12CF:265
0CEE040EEAAB2A86B32ED766CB14A9C726A:103
0D876EA172607E19F8568CCF455C7EA36E1:1
0DE5776B81F24941708A9B324A4B5C8BF15:2
0E49A19AE1EA749193965454E6392034369:10
0EAA6A6BE508F86687B4ADDACB36D4C9A2A:2
0ECD1155ABF95C201FB2E97DC55FB4A7B14:1
0EEAAF0D378900D798CC7004D20177195F7:1
0F6E13BB39FD232385D2B3CACCE1B684E79:6
10F12B2650D5132FC30FE6FFC5F6D1A5D7D:2
1155951432A8229D96C4E983383DD176D5E:1
11E4DF3590358A72FFE503ABCF347E61BA9:2
1219586938607A44E939CC0F045D14E8DF0:4
123251CCEFE165A104890B2EECD754C8246:4
125858E827ECB3BA2EB8647D917AC458977:3
128264D249E627B016048925BC98590FAB5:3
1321D475C331A44CB627A9FDD00E294DC00:5
` +
		shaPasswd +
		`
144F1DBC1E8286A1308367227C617F1448F:2
1476366181421DF6A5B2F5A23B0920D52FB:2
14CFF81BD2564F4A65F642DAFC3FAC321A7:2
14F1F7D89C86C9A6E055242CB87D7A81832:1
15769CC9CA70A668178D8321C7EB0A6A001:3
1598E6CC54B336396D92A4323130ABB2BAB:1
163DF9D1F2AB2B337BBF6F1DED009A990E0:1
166318D06BD93F13B21CBB9177C485D0C95:1
1697A2AADC2AE54ABB38946939DAD174FC0:5
16DB56C7FE1951992800B09B7208A5791AD:6
16E492256D5DF1E98A4C47501F3B9A7A7F4:1
17F1829D8D6BDF62DF32EEF6F93BD8B072A:2
182530C099B002A845D72123BB94659E868:3
19ACDCF1284EF78BCC83B498D3A8828C5A4:2
1A5E31D4BA08537F2DED056F37593B2DC4F:5
1A82CA6EB62E71B2D3EEA03451FB8A80A9F:2
1AAAB3A09B9EED976C50F89317D054EF927:1
1B4D298A834A16AC4D74A91C7645049F73B:9`
	bytes := []byte(strings.Replace(resp, "\n", "\r\n", -1))
	return bytes
}
