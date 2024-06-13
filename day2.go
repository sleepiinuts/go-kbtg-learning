package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func sumNumber[K Number](ns []K) K {
	var sum K

	for _, n := range ns {
		sum += n
	}

	return sum
}

// marshal
// 1. {"address_no": "46/6", "road": "popular", "district": "pakkret", "province": "nonthaburi", "zip": "11120"}

// unmarshal
// 2. {"name": "john", "hobbies":["watch tv", "surf internet", "read book"]}

// marshal
// 3. {"first_name":"steve","last_name":"rogers","address":{"address_no": "46/6", "road": "popular", "district": "pakkret", "province": "nonthaburi", "zip": "11120"}}

type AddressJS struct {
	AddrNo   string `json:"address_no"`
	Road     string `json:"road"`
	District string `json:"district"`
	Province string `json:"province"`
	Zip      int    `json:"zip"`
}

type Person struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

type PersonInfo struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Addr      AddressJS `json:"address"`
}
