package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"

	models "github.com/sleepiinuts/go-kbtg-learning/models"
)

type Address struct {
	HomeNo     string
	StreetName string
	Province   string
	Zipcode    int
}

func (a *Address) String() string {
	aValue := reflect.ValueOf(*a)
	aType := aValue.Type()
	i := aValue.NumField()

	str := "{\n"

	for idx := range i {
		str += fmt.Sprintf("   %s: %v\n", aType.Field(idx).Name, aValue.Field(idx))
	}

	return str + "}"
}

func main() {
	str := "golang is a simple programming language"
	i := 1024
	f64 := float64(100.35)
	b := true
	bb := []byte(`simple byte string`)

	fmt.Printf("%-20v -- type: %T\n", str, str)
	fmt.Printf("%-20v -- type: %T\n", i, i)
	fmt.Printf("%-20v -- type: %T\n", f64, f64)
	fmt.Printf("%-20v -- type: %T\n", b, b)
	fmt.Printf("%-20v -- type: %T\n", string(bb), bb)

	j := 10
	jj := float64(j)
	fmt.Printf("%v:%v\n", &j, &jj)

	addr := Address{
		HomeNo:     "home",
		StreetName: "street",
		Province:   "province",
		Zipcode:    112,
	}

	addr2 := Address{
		// HomeNo:     "home2",
		StreetName: "street2",
		Province:   "province2",
		Zipcode:    112,
	}

	// addr2 = addr
	fmt.Println("reassign: ", &addr)
	fmt.Println("reassign: ", &addr2)

	// fmt.Printf("addr: %v\n", addr)
	if jstr, err := xml.MarshalIndent(addr, "", "    "); err == nil {
		fmt.Println(string(jstr))
	}

	// fmt.Printf("%+v", addr)

	// fmt.Println(addr)
	fmt.Println("stringer: ", &addr)

	citizenID := []int{1, 1, 0, 3, 0, 0, 0, 0, 1, 9, 7, 8, 0}
	fmt.Println("my last 4 digits citizenID: ", citizenID[len(citizenID)-4:])

	cID := strings.Split(`1103000019780`, "")
	fmt.Println("my last 4 digits cID: ", cID[len(cID)-4:])

	printScore(100)

	printEvery5()

	fmt.Println("----")

	printEven()

	// defer func() {
	// 	str := recover()
	// 	fmt.Println("recover: ", str)
	// 	debug.PrintStack()
	// }()

	// panic("some message")

	c := models.Car{
		Name:  "sample name",
		Model: "sample model",
		Price: 123.45,
	}

	models.PrintDetails(c.Name, c.Model, c.Price)
	c.PrintDetaisWithCustomType()
	c.PrintDetailsWithCustomTypePretty()

	fmt.Println(models.PrintThousands(12345))

	e := models.Employee{
		FirstName: "FName",
		LastName:  "LName",
		Salary:    1234.56,
	}

	e.Print()
	e.PrintSalary()

	// start D2

	fmt.Println(sumNumber([]int{1, 2, 3, 4, 5}))
	fmt.Println(sumNumber([]float32{1.1, 1.2, 2.3}))

	addrJs := AddressJS{AddrNo: "46/6", Road: "popular", District: "pakkret", Province: "nonthaburi", Zip: 11120}
	byteJs, err := json.Marshal(addrJs)
	if err != nil {
		panic("err marshaljson")
	}

	fmt.Println(string(byteJs))

	personBt := []byte(`{"name": "john", "hobbies":["watch tv", "surf internet", "read book"]}`)
	var p1 Person

	err = json.Unmarshal(personBt, &p1)
	if err != nil {
		panic("err unmarshal json")
	}

	fmt.Printf("%+v\n", p1)

	pInf := PersonInfo{
		FirstName: "fname",
		LastName:  "lname",
		Addr:      addrJs,
	}

	pInfByte, err := json.MarshalIndent(pInf, "", "  ")
	if err != nil {
		panic("err marshal person info")
	}

	fmt.Println(string(pInfByte))

	strXML := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:v2="test.api.com/WS/KK11238I01/v2">
    <soapenv:Body>
        <doService xmlns="http://test.api.com/WS/KK11238I01/v2">
            <KK11238I01>
                <Header>
                    <funcNm>KK11238I01</funcNm>
                    <rqUID>123_20180608_1430_0000001</rqUID>
                    <rqAppId>123</rqAppId>
                    <rqDt>2018-06-08T14:34:09.47479164+07:00</rqDt>
                    <corrID>5543</corrID>
                    <userId>XXO20029</userId>
                    <terminalId>123</terminalId>
                    <userLangPref>TH</userLangPref>
                    <authUserId>XXO20029</authUserId>
                    <authLevel>1</authLevel>
                    <errorVect>
                        <error></error>
                    </errorVect>
                </Header>
                <Account>
                    <AcctId>12345677</AcctId>
                    <Concept1>CR To 12345677</Concept1>
                    <Concept2></Concept2>
                    <ExtAcctDt>2018-02-01</ExtAcctDt>
                    <FeeAmt>0.0</FeeAmt>
                    <ICA>123</ICA>
                    <OperationCode>01</OperationCode>
                    <OperationType>CR</OperationType>
                    <SubOperationCode>7677</SubOperationCode>
                    <SvcBranchId>2241</SvcBranchId>
                    <TrnAmt>300</TrnAmt>
                    <UseSvcBranchFlag>N</UseSvcBranchFlag>
                    <ValueDt>2018-02-01</ValueDt>
                    <AuthLevCnt></AuthLevCnt>
                    <TellerID></TellerID>
                    <TellerIDAuthNum></TellerIDAuthNum>
                    <TxnCode></TxnCode>
                    <TxnIndctcnt></TxnIndctcnt>                    
                </Account>
            </KK11238I01>
        </doService>
    </soapenv:Body>
</soapenv:Envelope>`

	var s Simple
	err = xml.Unmarshal([]byte(strXML), &s)
	if err != nil {
		panic("unmarshal XML error")
	}

	fmt.Println("funcNM: ", s.Body.Service.KK11238I01.Header.FunctionName)
	simpleJs, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(simpleJs))

	writeFile()
}

func printEvery5() {
	for i := 0; i <= 100; i += 5 {
		fmt.Printf("i: %d\n", i)
	}
}

func printEven() {
	for i := range 101 {
		if i%2 == 0 {
			fmt.Printf("i: %d\n", i)
		}
	}
}

func printScore(scr int) {
	switch {
	case scr > 90:
		fmt.Println("more than 90")
	case scr > 80:
		fmt.Println("more than 80")
	case scr > 70:
		fmt.Println("more than 70")
	case scr > 60:
		fmt.Println("more than 60")
	default:
		fmt.Println("else")
	}
}

func init() {
	fmt.Println("init function")
}
