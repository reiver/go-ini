package ini_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-ini"
)

func TestWriteINI(t *testing.T) {

	tests := []struct{
		Src any
		Expected string
	}{
		{
			Src: map[string]string(nil),
		},
		{
			Src: map[string]string{},
		},



		{
			Src: map[string]string{
				"apple":"ONCE",
			},
			Expected:
				`apple = ONCE` +"\n"+
				"",
		},
		{
			Src: map[string]string{
				"apple":"ONCE",
				"Banana":"twice",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`Banana = twice` +"\n"+
				"",
		},
		{
			Src: map[string]string{
				"apple":"ONCE",
				"Banana":"twice",
				"CHERRY":"THRICE",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`Banana = twice` +"\n"+
				`CHERRY = THRICE` +"\n"+
				"",
		},
		{
			Src: map[string]string{
				"apple":"ONCE",
				"Banana":"twice",
				"CHERRY":"THRICE",
				"dAtE":"fource",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`Banana = twice` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = fource` +"\n"+
				"",
		},



		{
			Src: map[string]string{
				"apple":"ONCE",
				"Banana":"twice",
				"bAnana":"2",
				"CHERRY":"THRICE",
				"dAtE":"fource",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`Banana = twice` +"\n"+
				`bAnana = 2` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = fource` +"\n"+
				"",
		},



		{
			Src: map[string]string{
				"apple":"ONCE",
				"bAnan":"2",
				"Banana":"twice",
				"CHERRY":"THRICE",
				"dAtE":"fource",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`bAnan = 2` +"\n"+
				`Banana = twice` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = fource` +"\n"+
				"",
		},









		{
			Src: map[string]string{
				"apple":"ONCE",
				"B":"L",
				"b":"l",
				"BA":"LO",
				"Ba":"Lo",
				"bA":"lO",
				"ba":"lo",
				"BAN":"LOW",
				"BAn":"LOw",
				"BaN":"LoW",
				"Ban":"Low",
				"bAN":"lOW",
				"bAn":"lOw",
				"baN":"loW",
				"ban":"low",
				"BaNaN":"up-down",
				"bAnAn":"down-up",
				"banana":"lower",
				"Banana":"Title",
				"BANANA":"CONST",
				"CHERRY":"THRICE",
				"dAtE":"fource",
			},
			Expected:
				`apple = ONCE` +"\n"+
				`B = L` +"\n"+
				`b = l` +"\n"+
				`BA = LO` +"\n"+
				`Ba = Lo` +"\n"+
				`bA = lO` +"\n"+
				`ba = lo` +"\n"+
				`BAN = LOW` +"\n"+
				`BAn = LOw` +"\n"+
				`BaN = LoW` +"\n"+
				`Ban = Low` +"\n"+
				`bAN = lOW` +"\n"+
				`bAn = lOw` +"\n"+
				`baN = loW` +"\n"+
				`ban = low` +"\n"+
				`BaNaN = up-down` +"\n"+
				`bAnAn = down-up` +"\n"+
				`BANANA = CONST` +"\n"+
				`Banana = Title` +"\n"+
				`banana = lower` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = fource` +"\n"+
				"",
		},
	}

	for testNumber, test := range tests {

		var dst strings.Builder
		err := ini.WriteINI(&dst, test.Src)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("SOURCE: (%T) %#v", test.Src, test.Src)
			continue
		}

		var actual string = dst.String()

		var expected string = test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual written-ini is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("SOURCE: (%T) %#v", test.Src, test.Src)
			continue
		}
	}
}
