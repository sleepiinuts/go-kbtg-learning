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

type Simple struct {
	Body struct {
		Service struct {
			XMLNS      string `xml:"xmlns,attr"`
			KK11238I01 struct {
				Header struct {
					FunctionName string `xml:"funcNm"`
					RequestID    string `xml:"rqUID"`
					RequestAppID string `xml:"rqAppId"`
					RequestDate  string `xml:"rqDt"`
					CorrID       string `xml:"corrID"`
					UserID       string `xml:"userId"`
					TerminalID   string `xml:"terminalId"`
					Language     string `xml:"userLangPref"`
					AuthUserID   string `xml:"authUserId"`
					AuthLevel    string `xml:"authLevel"`
					Error        struct {
						Message string `xml:"error"`
					} `xml:"errorVect"`
				} `xml:"Header"`
				Account struct {
					AccountID string `xml:"AcctId"`
				} `xml:"Account"`
			} `xml:"KK11238I01"`
		} `xml:"doService"`
	} `xml:"Body"`
}
