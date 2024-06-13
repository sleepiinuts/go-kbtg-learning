package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

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

type Travel struct {
	Attr string `xml:"Body>doService xmlns,attr"`
}

type Product struct {
	Name       string
	Price      float64
	LaunchDate CustomDate
	Date       time.Time
}

type CustomDate struct {
	time.Time
}

const dateFormat = "2006-01-02"

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// write your code here

	str := string(data)

	t, err := time.Parse(dateFormat, str[1:len(str)-1]) // --> need to remove quote first!
	if err != nil {
		return err
	}

	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// write your code here
	return json.Marshal(cd.Time.Format(dateFormat))

	// fmt.Println("byte: ", []byte(fmt.Sprintf("%q", cd.Time.Format(dateFormat))))
	// cmp, _ := json.Marshal(cd.Time.Format(dateFormat))
	// fmt.Println("json: ", cmp)
	// return []byte(cd.Time.Format(dateFormat)), nil --> require quoted text to work!!
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
type Meta struct {
	Pagination struct {
		Total int `json:"total"`
		Pages int `json:"pages"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"pagination"`
}

type Response struct {
	Code int `json:"code"`
	Meta `json:"meta"`
	Data []User
}

func HttpRequest() {
	req, err := http.NewRequest(http.MethodGet, "https://gorest.co.in/public-api/users", nil)
	if err != nil {
		fmt.Println("err connecting: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err response: ", err)
	}

	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("err decoder: ", err)
	}

	fmt.Printf("%+v\n", response.Meta)

	for _, d := range response.Data {
		fmt.Printf("%+v\n", d)
	}

}

func writeFile() {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := range 100 {
		file.Write([]byte(fmt.Sprintf("%d\n", i)))
	}
}
