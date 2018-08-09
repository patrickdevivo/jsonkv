package jsonkv

import (
	"encoding/json"
	"strconv"
	"strings"
)

// KV represents a KV pair
type KV struct {
	Key   string
	Value string
}

func traverse(path, sep string, j interface{}) ([]*KV, error) {
	kvs := make([]*KV, 0)
	if sep == "" {
		sep = "/"
	}
	pathPre := ""
	if path != "" {
		pathPre = path + sep
	}

	switch j.(type) {
	case []interface{}:
		for sk, sv := range j.([]interface{}) {
			skvs, err := traverse(pathPre+strconv.Itoa(sk), sep, sv)
			if err != nil {
				return nil, err
			}
			kvs = append(kvs, skvs...)
		}
	case map[string]interface{}:
		for sk, sv := range j.(map[string]interface{}) {
			skvs, err := traverse(pathPre+sk, sep, sv)
			if err != nil {
				return nil, err
			}
			kvs = append(kvs, skvs...)
		}
	case float64:
		kvs = append(kvs, &KV{Key: path, Value: strconv.FormatFloat(j.(float64), 'f', -1, 64)})
	case bool:
		kvs = append(kvs, &KV{Key: path, Value: strconv.FormatBool(j.(bool))})
	case nil:
		kvs = append(kvs, &KV{Key: path, Value: ""})
	default:
		kvs = append(kvs, &KV{Key: path, Value: j.(string)})
	}

	return kvs, nil
}

// ToKVs takes a json byte array and returns a list of KV pairs
func ToKVs(jsonData []byte, sep string) ([]*KV, error) {
	var i interface{}
	err := json.Unmarshal(jsonData, &i)
	if err != nil {
		return nil, err
	}
	m := i.(map[string]interface{})

	return traverse("", sep, m)
}

// ToJSON converts a list of KVs to a JSON tree
func ToJSON(kvs []*KV, sep string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	for _, kv := range kvs {
		path := strings.Split(kv.Key, sep)
		var parent = m
		for s, segment := range path {
			if s == len(path)-1 {
				parent[segment] = string(kv.Value)
			} else {
				if parent[segment] == nil {
					parent[segment] = make(map[string]interface{})
				}
				switch parent[segment].(type) {
				case map[string]interface{}:
					parent = parent[segment].(map[string]interface{})
				default:
					delete(parent, segment)
				}
			}
		}
	}
	return m, nil
}
