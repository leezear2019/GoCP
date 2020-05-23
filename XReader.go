package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	ID     int
	Name   string
	Values []int
}

type XCon struct {
	ID        int
	Name      string
	Semantics bool
	Arity     int
	Scope     []XVar
	Tuples    [][]int
	ScopeInt  []int
}

type XModel struct {
	Vars    []XVar
	Cons    []XCon
	DomMap  map[string]int
	VarsMap map[string]int
	RelMap  map[string]int
	ConsMap map[string]int
}

func (xm XModel) BuildXModel(xm2 *XModel2) {
	// 初始化四个map
	for i, d := range xm2.XDomains.Domains {
		xm.DomMap[d.Name] = i
	}

	for i, v := range xm2.XVariables.Variables {
		xm.VarsMap[v.Name] = i
		// 获取valueStr
		valueStr := xm2.XDomains.Domains[xm.DomMap[v.Name]].ValuesStr
		values := []int{}
		if strings.Contains(valueStr, "..") {
			//获取两部
			var lb, ub int
			fmt.Sscanf(valueStr, "%d..%d", &lb, &ub)
			for ii := lb; ii <= ub; ii++ {
				values = append(values, ii)
			}

		} else if !strings.Contains(valueStr, " ") {

		}
		xm.Vars = append(xm.Vars, XVar{i, v.Name, make([]int, len(values))})
		copy(xm.Vars[i].Values, values)
	}

	for i, c := range xm2.XConstraints.Constraints {
		xm.VarsMap[c.Name] = i
	}

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
