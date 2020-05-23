package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//
//type SConfig struct {
//	XMLName      xml.Name   `xml:"config"`     // 指定最外层的标签为config
//	SmtpServer   string     `xml:"smtpServer"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
//	SmtpPort     int        `xml:"smtpPort"`
//	Sender       string     `xml:"sender"`
//	SenderPasswd string     `xml:"senderPasswd"`
//	Receivers    SReceivers `xml:"receivers"` // 读取receivers标签下的内容，以结构方式获取
//}
//
//type SReceivers struct {
//	Flag string   `xml:"flag,attr"` // 读取flag属性
//	User []string `xml:"user"`      // 读取user数组
//}
//

type XDomain2 struct {
	//XMLName   xml.Name `xml:"domain"`
	Name      string `xml:"name,attr"`
	NumValues int    `xml:"nbValues,attr"`
	ValuesStr string `xml:",innerxml"`
}

type XDomains2 struct {
	NumDomains int        `xml:"nbDomains,attr"`
	Domains    []XDomain2 `xml:"domain"`
}
type XVariable2 struct {
	Name       string `xml:"name,attr"`
	DomainName string `xml:"domain,attr"`
}
type XVariables2 struct {
	NumVariables int          `xml:"nbVariables,attr"`
	Variables    []XVariable2 `xml:"variable"`
}
type XRelation2 struct {
	Name      string `xml:"name,attr"`
	Arity     int    `xml:"arity,attr"`
	NumTuples int    `xml:"nbTuples,attr"`
	Semantics string `xml:"semantics,attr"`
	TuplesStr string `xml:",innerxml"`
}
type XRelations2 struct {
	NumRelations int          `xml:"nbRelations,attr"`
	Relations    []XRelation2 `xml:"relation"`
}

type XConstraint2 struct {
	Name      string `xml:"name,attr"`
	Arity     int    `xml:"arity,attr"`
	ScopeStr  string `xml:"scope,attr"`
	Reference string `xml:"reference,attr"`
}
type XConstraints2 struct {
	NumConstraints int            `xml:"nbConstraints,attr"`
	Constraints    []XConstraint2 `xml:"constraint"`
}

type XModel2 struct {
	XMLName      xml.Name      `xml:"instance"`
	XDomains     XDomains2     `xml:"domains"`
	XVariables   XVariables2   `xml:"variables"`
	XRelations   XRelations2   `xml:"relations"`
	XConstraints XConstraints2 `xml:"constraints"`
}

type XVar struct {
}

type XCon struct {
}

type XModel struct {
}

func main() {

	file, err := os.Open("benchmarks/q4.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := XModel2{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println(v.XConstraints)
	fmt.Println(v.XConstraints.Constraints[0].Name)
	//fmt.Println(v.XDomains)

	//file, err := os.Open("servers.xml") // For read access.
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//v := SConfig{}
	//err = xml.Unmarshal(data, &v)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//
	//fmt.Println(v)
	//fmt.Println("SmtpServer : ", v.SmtpServer)
	//fmt.Println("SmtpPort : ", v.SmtpPort)
	//fmt.Println("Sender : ", v.Sender)
	//fmt.Println("SenderPasswd : ", v.SenderPasswd)
	//fmt.Println("Receivers.Flag : ", v.Receivers.Flag)
	//for i, element := range v.Receivers.User {
	//	fmt.Println(i, element)
	//}
}
