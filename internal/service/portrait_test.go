package service

import "testing"

type PortraitAuthTest struct {
	in  PortraitAuthRequest
	out string
}

var PortraitAuthTests = []PortraitAuthTest{
	{in: PortraitAuthRequest{UserName: "谢理", CertNumber: "511381199108190691", Photo: ""},
		out: "0XXX"},
}

func TestPortraitAuth(t *testing.T) {
	for _, test := range PortraitAuthTests {
		if reuslt, err := PortraitAuth(&test.in); err != nil {
			t.Errorf("Portrait auth failed")
		} else {
			if reuslt != test.out {
				t.Errorf("Portrait auth expects %s, but got %s", test.out, reuslt)
			}
		}
	}

}
