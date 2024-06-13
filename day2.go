package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func sumComparable[K Number](ns []K) K {
	var sum K

	for _, n := range ns {
		sum += n
	}

	return sum
}
