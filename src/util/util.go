package util

import (
	"path/filepath"
	"log"
	"strconv"
    "reflect"
)

func HomePath() string {
	return AbsolutePath(".")
}

func AbsolutePath(relpath string) string {
	absolutePath, err := filepath.Abs(relpath)
	if err != nil {
		log.Println("current path error:", err)
	}
	return absolutePath
}

func HostAndPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}

// 再切片中查找对应元素的索引
func SliceIndex(slice interface{}, element interface{}) int {
    index := -1
    sv := reflect.ValueOf(slice)
    if sv.Kind() != reflect.Slice {
        return index
    }

    ev := reflect.ValueOf(element).Interface()
    length := sv.Len()

    for i := 0; i < length; i++ {
        iv := sv.Index(i).Interface()
        if reflect.DeepEqual(iv, ev) {
            return i
        }
    }
    return index
}


