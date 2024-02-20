package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		INI string
		Expected map[string]string
	}{
		{
			INI:
				"one=1",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\n",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\r",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
			},
		},
		{
			INI:
				"one=1" +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
			},
		},



		{
			INI:
				"one=1" +"\n"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1" +"\n\r"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1" +"\r"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1" +"\r\n"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1" +"\u0085"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1" +"\u2028"+
				"two=22",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},



		{
			INI:
				"one=1"  +"\n"+
				"two=22" +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\n\r"+
				"two=22" +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\r"+
				"two=22" +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\r\n"+
				"two=22" +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\u0085"+
				"two=22" +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\u2028"+
				"two=22" +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},



		{
			INI:
				"one=1"  +"\n"+
				"two=22" +"\n"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\n\r"+
				"two=22" +"\n\r"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\r"+
				"two=22" +"\r"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\r\n"+
				"two=22" +"\r\n"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\u0085"+
				"two=22" +"\u0085"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"  +"\u2028"+
				"two=22" +"\u2028"+
				"[three]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},



		{
			INI:
				"one=1"   +"\n"+
				"two=22"  +"\n"+
				"[three]" +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"   +"\n\r"+
				"two=22"  +"\n\r"+
				"[three]" +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"   +"\r"+
				"two=22"  +"\r"+
				"[three]" +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"   +"\r\n"+
				"two=22"  +"\r\n"+
				"[three]" +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"   +"\u0085"+
				"two=22"  +"\u0085"+
				"[three]" +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},
		{
			INI:
				"one=1"   +"\u2028"+
				"two=22"  +"\u2028"+
				"[three]" +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
			},
		},



		{
			INI:
				"one=1"   +"\n"+
				"two=22"  +"\n"+
				"[three]" +"\n"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"   +"\n\r"+
				"two=22"  +"\n\r"+
				"[three]" +"\n\r"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"   +"\r"+
				"two=22"  +"\r"+
				"[three]" +"\r"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"   +"\r\n"+
				"two=22"  +"\r\n"+
				"[three]" +"\r\n"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"   +"\u0085"+
				"two=22"  +"\u0085"+
				"[three]" +"\u0085"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"   +"\u2028"+
				"two=22"  +"\u2028"+
				"[three]" +"\u2028"+
				"apple=red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},



		{
			INI:
				"one=1"     +"\n"+
				"two=22"    +"\n"+
				"[three]"   +"\n"+
				"apple=red" +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"     +"\n\r"+
				"two=22"    +"\n\r"+
				"[three]"   +"\n\r"+
				"apple=red" +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"     +"\r"+
				"two=22"    +"\r"+
				"[three]"   +"\r"+
				"apple=red" +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"     +"\r\n"+
				"two=22"    +"\r\n"+
				"[three]"   +"\r\n"+
				"apple=red" +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"     +"\u0085"+
				"two=22"    +"\u0085"+
				"[three]"   +"\u0085"+
				"apple=red" +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},
		{
			INI:
				"one=1"     +"\u2028"+
				"two=22"    +"\u2028"+
				"[three]"   +"\u2028"+
				"apple=red" +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
			},
		},



		{
			INI:
				"one=1"     +"\n"+
				"two=22"    +"\n"+
				"[three]"   +"\n"+
				"apple=red" +"\n"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"     +"\n\r"+
				"two=22"    +"\n\r"+
				"[three]"   +"\n\r"+
				"apple=red" +"\n\r"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"     +"\r"+
				"two=22"    +"\r"+
				"[three]"   +"\r"+
				"apple=red" +"\r"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"     +"\r\n"+
				"two=22"    +"\r\n"+
				"[three]"   +"\r\n"+
				"apple=red" +"\r\n"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"     +"\u0085"+
				"two=22"    +"\u0085"+
				"[three]"   +"\u0085"+
				"apple=red" +"\u0085"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"     +"\u2028"+
				"two=22"    +"\u2028"+
				"[three]"   +"\u2028"+
				"apple=red" +"\u2028"+
				"banana = yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n"+
				""                +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r"+
				""                +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r"+
				""                +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n"+
				""                +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085"+
				""                +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028"+
				""                +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n"+
				""                +"\n"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r"+
				""                +"\n\r"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r"+
				""                +"\r"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n"+
				""                +"\r\n"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085"+
				""                +"\u0085"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028"+
				""                +"\u2028"+
				"cherry: red",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n"+
				""                +"\n"+
				"cherry: red"     +"\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r"+
				""                +"\n\r"+
				"cherry: red"     +"\n\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r"+
				""                +"\r"+
				"cherry: red"     +"\r",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n"+
				""                +"\r\n"+
				"cherry: red"     +"\r\n",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085"+
				""                +"\u0085"+
				"cherry: red"     +"\u0085",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028"+
				""                +"\u2028"+
				"cherry: red"     +"\u2028",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
			},
		},



		{
			INI:
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"BANANA = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n"+
				""                +"\n"+
				"Cherry: red"     +"\n"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"BANANA = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r"+
				""                +"\n\r"+
				"Cherry: red"     +"\n\r"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"BANANA = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r"+
				""                +"\r"+
				"Cherry: red"     +"\r"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"BANANA = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n"+
				""                +"\r\n"+
				"Cherry: red"     +"\r\n"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"BANANA = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085"+
				""                +"\u0085"+
				"Cherry: red"     +"\u0085"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"BANANA = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028"+
				""                +"\u2028"+
				"Cherry: red"     +"\u2028"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},



		{
			INI:
				";INI/1"          +"\n"+
				""                +"\n"+
				"one=1"           +"\n"+
				"two=22"          +"\n"+
				"[three]"         +"\n"+
				"apple=red"       +"\n"+
				"banana = yellow" +"\n"+
				""                +"\n"+
				"[four]"          +"\n"+
				""                +"\n"+
				"Cherry: red"     +"\n"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				";INI/1"          +"\n\r"+
				""                +"\n\r"+
				"one=1"           +"\n\r"+
				"two=22"          +"\n\r"+
				"[three]"         +"\n\r"+
				"apple=red"       +"\n\r"+
				"banana = yellow" +"\n\r"+
				""                +"\n\r"+
				"[four]"          +"\n\r"+
				""                +"\n\r"+
				"Cherry: red"     +"\n\r"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				";INI/1"          +"\r"+
				""                +"\r"+
				"one=1"           +"\r"+
				"two=22"          +"\r"+
				"[three]"         +"\r"+
				"apple=red"       +"\r"+
				"banana = yellow" +"\r"+
				""                +"\r"+
				"[four]"          +"\r"+
				""                +"\r"+
				"Cherry: red"     +"\r"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				";INI/1"          +"\r\n"+
				""                +"\r\n"+
				"one=1"           +"\r\n"+
				"two=22"          +"\r\n"+
				"[three]"         +"\r\n"+
				"apple=red"       +"\r\n"+
				"banana = yellow" +"\r\n"+
				""                +"\r\n"+
				"[four]"          +"\r\n"+
				""                +"\r\n"+
				"Cherry: red"     +"\r\n"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				";INI/1"          +"\u0085"+
				""                +"\u0085"+
				"one=1"           +"\u0085"+
				"two=22"          +"\u0085"+
				"[three]"         +"\u0085"+
				"apple=red"       +"\u0085"+
				"banana = yellow" +"\u0085"+
				""                +"\u0085"+
				"[four]"          +"\u0085"+
				""                +"\u0085"+
				"Cherry: red"     +"\u0085"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
		{
			INI:
				";INI/1"          +"\u2028"+
				""                +"\u2028"+
				"one=1"           +"\u2028"+
				"two=22"          +"\u2028"+
				"[three]"         +"\u2028"+
				"apple=red"       +"\u2028"+
				"banana = yellow" +"\u2028"+
				""                +"\u2028"+
				"[four]"          +"\u2028"+
				""                +"\u2028"+
				"Cherry: red"     +"\u2028"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},









		{
			INI:
				";INI/1"            +"\u2028"+
				""                  +"\u2028"+
				"one=1"             +"\u2028"+
				"two=22"            +"\u2028"+
				"[three]"           +"\u2028"+
				"\tapple=red"       +"\u2028"+
				"\tbanana = yellow" +"\u2028"+
				""                  +"\u2028"+
				"[four]"            +"\u2028"+
				""                  +"\u2028"+
				"Cherry: red"       +"\u2028"+
				"date : red, yellow",
			Expected: map[string]string{
				`one`:`1`,
				`two`:`22`,
				`three.apple`:`red`,
				`three.banana`:`yellow`,
				`four.cherry`: `red`,
				`four.date`: `red, yellow`,
			},
		},
	}

	for testNumber, test := range tests {
		var actual map[string]string = map[string]string{}

		var p []byte = []byte(test.INI)

		err := ini.Unmarshal(p, &actual)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("INI: %s", test.INI)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual decoded result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("INI: %s", test.INI)
				continue
			}
		}
	}
}
