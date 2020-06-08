package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	ScopeInt  []int
	Tuples    [][]int
}

type XModel struct {
	Vars    []XVar
	Cons    []XCon
	NumVars int
	NumCons int
	domsMap map[string]int
	varsMap map[string]int
	relsMap map[string]int
	consMap map[string]int
}

//func (xv XVar) BuildXVar(id int, name string, values []int) {
//
//}

func (xm *XModel) BuildXModel(xm2 *XModel2) {
	xm.NumVars = xm2.XVariables.NumVariables
	xm.NumCons = xm2.XConstraints.NumConstraints
	xm.Vars = make([]XVar, xm.NumVars)
	xm.Cons = make([]XCon, xm.NumCons)
	xm.domsMap = make(map[string]int)
	xm.varsMap = make(map[string]int)
	xm.relsMap = make(map[string]int)
	xm.consMap = make(map[string]int)
	// 初始化四个map
	for i, d := range xm2.XDomains.Domains {
		xm.domsMap[d.Name] = i
	}

	for i, v := range xm2.XVariables.Variables {
		xm.varsMap[v.Name] = i
		xm.Vars[i].Name = v.Name
		// 获取valueStr
		valueStr := xm2.XDomains.Domains[xm.domsMap[v.Name]].ValuesStr
		values := []int{}
		if strings.Contains(valueStr, "..") {
			//获取两部
			var lb, ub int
			fmt.Sscanf(valueStr, "%d..%d", &lb, &ub)
			for ii := lb; ii <= ub; ii++ {
				values = append(values, ii)
			}

		} else if !strings.Contains(valueStr, " ") {
			vs := strings.Split(valueStr, " ")
			for _, v := range vs {
				ii, _ := strconv.Atoi(v)
				values = append(values, ii)
			}
		}
		xm.Vars[i] = XVar{i, v.Name, make([]int, len(values))}
		copy(xm.Vars[i].Values, values)
	}

	for i, r := range xm2.XRelations.Relations {
		xm.relsMap[r.Name] = i
	}

	for i, c := range xm2.XConstraints.Constraints {
		xm.consMap[c.Name] = i
		name := c.Name
		relIndex := xm.relsMap[c.Reference]

		var semantics bool
		switch xm2.XRelations.Relations[relIndex].Semantics {
		case "supports":
			semantics = true
		case "conflicts":
			semantics = false
		default:
			fmt.Printf("Semantics error\n")
		}

		// arity, scope
		arity := xm2.XRelations.Relations[relIndex].Arity
		numTuples := xm2.XRelations.Relations[relIndex].NumTuples
		scope := make([]XVar, arity)
		scopeInt := make([]int, arity)
		scopeVarStr := strings.Split(xm2.XConstraints.Constraints[i].ScopeStr, " ")

		for ii, varStr := range scopeVarStr {
			vid := xm.varsMap[varStr]
			scope[ii] = xm.Vars[vid]
			scopeInt[ii] = vid
		}

		xm.Cons[i] = XCon{
			ID:        i,
			Name:      name,
			Semantics: semantics,
			Arity:     arity,
			Scope:     make([]XVar, arity),
			ScopeInt:  make([]int, arity),
			Tuples:    make([][]int, numTuples),
		}

		copy(xm.Cons[i].Scope, scope)
		copy(xm.Cons[i].ScopeInt, scopeInt)

		//xm.Cons[i].Tuples = make([][]int, numTuples)
		tuplesStr := strings.Split(xm2.XRelations.Relations[relIndex].TuplesStr, "|")

		for ii, s := range tuplesStr {
			tupleStr := strings.Split(s, " ")
			xm.Cons[i].Tuples[ii] = make([]int, arity)
			for jj, t := range tupleStr {
				tt, _ := strconv.Atoi(t)
				xm.Cons[i].Tuples[ii][jj] = tt
			}
		}

	}
}

func (xm *XModel) show() {
	//fmt.Println(xm.domsMap)
	//fmt.Println(xm.varsMap)
	fmt.Println("-------show model-------")
	for i, v := range xm.Vars {
		fmt.Println(i, v)
	}
	//fmt.Println(xm.relsMap)
	//fmt.Println(xm.consMap)
	for i, c := range xm.Cons {
		fmt.Println(i, c)
	}
}

func main() {

	file, err := os.Open("benchmarks/q4.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %xm2", err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %xm2", err)
		return
	}

	xm2 := XModel2{}
	err = xml.Unmarshal(data, &xm2)
	if err != nil {
		fmt.Printf("error: %xm2", err)
		return
	}

	fmt.Println(xm2)
	xm := XModel{}
	xm.BuildXModel(&xm2)
	//fmt.Println(xm)
	fmt.Println("---------------------------")
	xm.show()
	//fmt.Println(xm2.XConstraints)
	//fmt.Println(xm2.XConstraints.Constraints[0].Name)
	//fmt.Println(xm2.XDomains)

	//file, err := os.Open("servers.xml") // For read access.
	//if err != nil {
	//	fmt.Printf("error: %xm2", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error: %xm2", err)
	//	return
	//}
	//xm2 := SConfig{}
	//err = xml.Unmarshal(data, &xm2)
	//if err != nil {
	//	fmt.Printf("error: %xm2", err)
	//	return
	//}
	//
	//fmt.Println(xm2)
	//fmt.Println("SmtpServer : ", xm2.SmtpServer)
	//fmt.Println("SmtpPort : ", xm2.SmtpPort)
	//fmt.Println("Sender : ", xm2.Sender)
	//fmt.Println("SenderPasswd : ", xm2.SenderPasswd)
	//fmt.Println("Receivers.Flag : ", xm2.Receivers.Flag)
	//for i, element := range xm2.Receivers.User {
	//	fmt.Println(i, element)
	//}
}
