package jq_test

import (
	"testing"

	"github.com/zgljl2012/jq"
)

func TestMerge(t *testing.T) {
	testcases := map[string]struct {
		s1     string
		s2     string
		expect string
	}{
		"simple": {
			s1:     `{"hello1":"world1"}`,
			s2:     `{"hello2":"world2"}`,
			expect: `{"hello1":"world1","hello2":"world2"}`,
		},
		"override": {
			s1:     `{"hello1":"world1"}`,
			s2:     `{"hello1":"world2"}`,
			expect: `{"hello1":"world2"}`,
		},
		"multi level": {
			s1:     `{"a":{"b":1}}`,
			s2:     `{"a":{"c":2}}`,
			expect: `{"a":{"b":1,"c":2}}`,
		},
		"multi level with override": {
			s1:     `{"a":{"b":1}}`,
			s2:     `{"a":{"b":2,"c":2}}`,
			expect: `{"a":{"b":2,"c":2}}`,
		},
		"multi level more": {
			s1:     `{"a":{"b":{"c":1,"d":2}},"b":1}`,
			s2:     `{"a":{"b":{"c":2},"c":2}}`,
			expect: `{"a":{"b":{"c":2,"d":2},"c":2},"b":1}`,
		},
		"array": {
			s1:     `{"a":[0, 1]}`,
			s2:     `{"a":[3,4]}`,
			expect: `{"a":[0,1,3,4]}`,
		},
		"array dict": {
			s1:     `{"a":[{"a":1},{"b":1}]}`,
			s2:     `{"a":[{"c":1}]}`,
			expect: `{"a":[{"a":1},{"b":1},{"c":1}]}`,
		},
		"complex case": {
			s1:     `{"a":{"b":{"c":{"d":{"e":1}}}}}`,
			s2:     `{"a":{"b":{"c":{"f":{"h":2}}}}}`,
			expect: `{"a":{"b":{"c":{"d":{"e":1},"f":{"h":2}}}}}`,
		},
	}
	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			if target, err := jq.Merge([]byte(testcase.s1), []byte(testcase.s2)); err != nil {
				t.Error(err)
			} else {
				if string(target) != testcase.expect {
					t.Errorf("expect %s, but got %s", testcase.expect, string(target))
				}
			}
		})
	}
}
