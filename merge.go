package jq

import (
	"encoding/json"
	"regexp"
)

// Merge two json, when a key exists in b1 and also exists in b2,
// if the value is a dict, then the value in b2 would override that in b1
// else if the value is a array, then the value in b2 would append to that in b1
func Merge(b1 []byte, b2 []byte) ([]byte, error) {
	re, _ := regexp.Compile(`(?s)^\s*\{.*\}\s*$|(?s)^\s*\[.*\]\s*$`)
	if !re.Match(b2) {
		return b2, nil
	}
	if !re.Match(b1) {
		return b1, nil
	}
	if matched, _ := regexp.Match(`(?s)^\s*{.*}\s*$`, b1); matched {
		// if dict
		var (
			j1  map[string]json.RawMessage
			j2  map[string]json.RawMessage
			err error
		)
		if err := json.Unmarshal(b1, &j1); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b2, &j2); err != nil {
			return nil, err
		}
		for k, v := range j2 {
			if _, ok := j1[k]; !ok {
				j1[k] = v
			} else {
				j1[k], err = Merge(j1[k], j2[k])
				if err != nil {
					return nil, err
				}
			}
		}
		return json.Marshal(j1)
	}
	// if array
	var (
		j1 []json.RawMessage
		j2 []json.RawMessage
	)
	if err := json.Unmarshal(b1, &j1); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b2, &j2); err != nil {
		return nil, err
	}
	j1 = append(j1, j2...)
	return json.Marshal(j1)

}
