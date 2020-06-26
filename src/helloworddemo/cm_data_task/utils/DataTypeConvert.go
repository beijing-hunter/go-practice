package utils

func SliceToMapString(datas []string) (result map[string]string) {

	if len(datas) == 0 {
		result = make(map[string]string, 1)
	} else {

		result = make(map[string]string, len(datas))

		for _, data := range datas {
			result[data] = data
		}
	}

	return
}

func SliceToMapInt(datas []int64) (result map[int64]int64) {

	if len(datas) == 0 {
		result = make(map[int64]int64, 1)
	} else {

		result = make(map[int64]int64, len(datas))

		for _, data := range datas {
			result[data] = data
		}
	}

	return
}
