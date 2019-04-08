package map2value

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func MapToUInt8(m map[string]interface{}, key string) (uint8, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(uint8); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToUInt16(m map[string]interface{}, key string) (uint16, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(uint16); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToUInt32(m map[string]interface{}, key string) (uint32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(uint32); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToUInt64(m map[string]interface{}, key string) (uint64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(uint64); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToInt(m map[string]interface{}, key string) (int, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(int); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToInt8(m map[string]interface{}, key string) (int8, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(int8); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToInt16(m map[string]interface{}, key string) (int16, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(int16); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToInt32(m map[string]interface{}, key string) (int32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(int32); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToInt64(m map[string]interface{}, key string) (int64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(int64); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}

}

func MapToFloat32(m map[string]interface{}, key string) (float32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(float32); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}
}

func MapToFloat64(m map[string]interface{}, key string) (float64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(float64); ok {
			return v, nil
		} else {
			return 0, errors.New("mismatched Type")
		}
	} else {
		return 0, errors.New("Key is not exist")
	}
}

func MapToString(m map[string]interface{}, key string) (string, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(string); ok {
			return v, nil
		} else {
			return "", errors.New("mismatched Type")
		}
	} else {
		return "", errors.New("Key is not exist")
	}

}

func MapToTime(m map[string]interface{}, key string) (time.Time, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(time.Time); ok {
			return v, nil
		} else {
			return time.Now(), errors.New("mismatched Type")
		}
	} else {
		return time.Now(), errors.New("Key is not exist")
	}
}

func MapToList(m map[string]interface{}, key string) ([]interface{}, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]interface{}); ok {
			return v, nil
		} else {
			var a []interface{}
			return a, errors.New("mismatched Type")
		}
	} else {
		var a []interface{}
		return a, errors.New("Key is not exist")
	}
}

func MapToStringList(m map[string]interface{}, key string) ([]string, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]string); ok {
			return v, nil
		} else {
			var a []string
			return a, errors.New("mismatched Type")
		}
	} else {
		var a []string
		return a, errors.New("Key is not exist")
	}
}

func MapToMapStringInterface(m map[string]interface{}, key string) (map[string]interface{}, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.(map[string]interface{}); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToIndentJSON(m interface{}) (string, error) {
	recordString, err := json.MarshalIndent(m, "", "\t")
	if err == nil {
		return fmt.Sprintf("RecordMap: %s", recordString), nil
	} else {
		return fmt.Sprintf("RecordMap: %+v", m), err
	}
}

func MapToFloat32List(m map[string]interface{}, key string) ([]float32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]float32); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToFloat64List(m map[string]interface{}, key string) ([]float64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]float64); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToInt8List(m map[string]interface{}, key string) ([]int8, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]int8); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToInt16List(m map[string]interface{}, key string) ([]int16, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]int16); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}
func MapToInt32List(m map[string]interface{}, key string) ([]int32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]int32); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToInt64List(m map[string]interface{}, key string) ([]int64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]int64); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToUint8List(m map[string]interface{}, key string) ([]uint8, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]uint8); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToUint16List(m map[string]interface{}, key string) ([]uint16, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]uint16); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}
func MapToUint32List(m map[string]interface{}, key string) ([]uint32, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]uint32); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToUint64List(m map[string]interface{}, key string) ([]uint64, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]uint64); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}

func MapToBytes(m map[string]interface{}, key string) ([]byte, error) {
	if _v, valid := m[key]; valid {
		if v, ok := _v.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.New("mismatched Type")
		}
	} else {
		return nil, errors.New("Key is not exist")
	}
}
