package common

import (
	"github.com/ipfs/go-ipfs/Godeps/_workspace/src/gopkg.in/errgo.v1"
	"strings"
)

func MapGetKV(v map[string]interface{}, key string) (interface{}, error) {
	var ok bool
	var cursor interface{} = v
	parts := strings.Split(key, ".")
	for i, part := range parts {
		cursor, ok = cursor.(map[string]interface{})[part]
		if !ok {
			sofar := strings.Join(parts[:i], ".")
			return nil, errgo.Newf("%s key has no attributes", sofar)
		}
	}
	return cursor, nil
}

func MapSetKV(v map[string]interface{}, key string, value interface{}) error {
	var ok bool
	var mcursor map[string]interface{}
	var cursor interface{} = v

	parts := strings.Split(key, ".")
	for i, part := range parts {
		mcursor, ok = cursor.(map[string]interface{})
		if !ok {
			sofar := strings.Join(parts[:i], ".")
			return errgo.Newf("%s key is not a map", sofar)
		}

		// last part? set here
		if i == (len(parts) - 1) {
			mcursor[part] = value
			break
		}

		cursor, ok = mcursor[part]
		if !ok { // create map if this is empty
			mcursor[part] = map[string]interface{}{}
			cursor = mcursor[part]
		}
	}
	return nil
}
