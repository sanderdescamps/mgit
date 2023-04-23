package helper

import (
	"fmt"
)

// func MergeDir(a map[string]any, b map[string]any) {
// 	for bk, bv := range b {
// 		if _, ok := a[bk]; !ok {
// 			a[bk] = bv
// 		} else if {

// 		}
// 	}
// }

// func Merge[M ~map[K]V, K comparable, V any](maps ...M) M {
// 	merged := make(M)
// 	for _, m := range maps {
// 		for key, val := range m {
// 			merged[key] = val
// 		}
// 	}

// 	return merged
// }

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func MergeDir(maps ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, m := range maps {
		for key, val := range m {
			switch v := val.(type) {
			case bool:
				merged[key] = v
			case int:
				merged[key] = v
			case string:
				merged[key] = v
			case float32, float64:
				merged[key] = v.(float64)
			case map[string]interface{}:
				if merged[key] != nil {
					merged[key] = MergeDir(merged[key].(map[string]interface{}), val.(map[string]interface{}))
				} else {
					merged[key] = val.(map[string]interface{})
				}

			case []string:
				if _, ok := merged[key]; ok {
					merged[key] = append(merged[key].([]string), val.([]string)...)
				} else {
					merged[key] = val.([]string)
				}
			default:
				fmt.Printf("Unsupported value; Can not merge maps; key=%s value=%v(%T)\n", key, val, val)
			}
		}
	}
	return merged
}
