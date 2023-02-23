// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	//todo: not handle Null type due to no time
	input, err := ioutil.ReadFile("input.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(input, &result)
	if err != nil {
		fmt.Println(err)
	}

	str := processJson(result)
	resultstr:=fmt.Sprintf("{%v}", str)

	fmt.Println(resultstr)

}

func processJson(inputvalue map[string]interface{}) string {
	resultstr := ""

	for key, value := range inputvalue {
		//omit empty key
		if key==""{
			continue
		}

		valuemap, _ := value.(map[string]interface{})

		str := getProcessedString(key, valuemap)

		if str != "" {
			if resultstr==""{
				resultstr=str
			}else {
				resultstr = fmt.Sprintf("%v,\n %v", resultstr, str)
			}
		}

	}
	return resultstr
}

func getProcessedString(key string, valuemap map[string]interface{}) string {
	str := ""

	if val, ok := valuemap["N"]; ok {
		str = processNumber(key, val)
		return str
	}

	if val, ok := valuemap["S"]; ok {
		str = processString(key, val)
		return str
	}

	if val, ok := valuemap["L"]; ok {
		str = processList(key, val)
		return str
	}

	if val, ok := valuemap["M"]; ok {
		str = processMap(key, val)
		return str
	}

	if val, ok := valuemap["BOOL"]; ok {
		str = processBool(key, val)
		return str
	}

	if _, ok := valuemap["NULL"]; ok {
		return ""
	}

	return str
}

func processMap(key string, val interface{}) string {

	resultstr := ""


	mapele, _ := val.(map[string]interface{})

	resultstr = processJson(mapele)

	if resultstr != "" {
		return fmt.Sprintf("\"%v\":{%v}", key, resultstr)
	}

	return ""
}

func processList(key string, val interface{}) string {
	//todo: remove the last comma in the resultstr
	resultstr := ""
	list, _ := val.([]interface{})

	for _, ele := range (list) {
		valuemapinlist, _ := ele.(map[string]interface{})
		str:=""
		if val, ok := valuemapinlist["N"]; ok {
			valnum, _ := val.(string)
			str = fmt.Sprintf("%v", valnum)
		}

		if val, ok := valuemapinlist["S"]; ok {
			valstr,_:=val.(string)
			if valstr!="" {
				str = fmt.Sprintf("\"%v\"", valstr)
			}
		}

		if str!="" {
			resultstr += fmt.Sprintf("%v,\n", str)
		}
	}

	if resultstr != "" {
		return fmt.Sprintf("\"%v\":[%v]", key, resultstr)
	}

	return ""
}

func processNumber(key string, val interface{}) string {
		resultstr := fmt.Sprintf("\"%v\": %v", key, val)
		return resultstr

}

func processString(key string, val interface{}) string {
		resultstr := fmt.Sprintf("\"%v\": \"%v\"", key, val)
		return resultstr
}

func processBool(key string, val interface{}) string {
	resultstr := fmt.Sprintf("\"%v\": %v", key, val)
	return resultstr

}
