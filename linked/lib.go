package linked

import "errors"

var ERROR_VALUE_TYPE error = errors.New("error type")

func Compare(a, b interface{}) (int, error) {
	// 大于为1 小于为-1 相等为0

	switch a := a.(type) {
	case int8:
		switch b := b.(type) {
		case int8:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int16:
		switch b := b.(type) {
		case int16:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int:
		switch b := b.(type) {
		case int:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int32:
		switch b := b.(type) {
		case int32:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int64:
		switch b := b.(type) {
		case int64:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}

	case uint8:
		switch b := b.(type) {
		case uint8:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint16:
		switch b := b.(type) {
		case uint16:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint:
		switch b := b.(type) {
		case uint:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint32:
		switch b := b.(type) {
		case uint32:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint64:
		switch b := b.(type) {
		case uint64:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case float64:
		switch b := b.(type) {
		case float64:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case float32:
		switch b := b.(type) {
		case float32:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case string:
		switch b := b.(type) {
		case string:
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	}
	return 0, ERROR_VALUE_TYPE

}
