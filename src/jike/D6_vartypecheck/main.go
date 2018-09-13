package main

import "fmt"

func check_accert(c interface{}) {
	if v, ok := interface{}(c).(string); ok {
		fmt.Printf("[Info]v is string,value is %s.\n", v)
	} else if v, ok := interface{}(c).([]string); ok {
		fmt.Printf("[Info]v is []string,value is %v\n", v)
	} else {
		fmt.Println("[Error]Unkonw type\n")
	}
}

func check_switch(c interface{}) {
	switch v := c.(type) {
	case string:
		fmt.Printf("[Info]v is string,value is %s.\n", v)
	case []string:
		fmt.Printf("[Info]v is []string,value is %v\n", v)
	case map[int]string:
		fmt.Printf("[Info]v is map[int]string,value is %v\n", v)
	default:
		fmt.Println("[Error]Unkonw type\n")

	}
}
func main() {
	var container1 = []string{"zero", "one", "two"}
	var container2 = map[int]string{0: "zero", 1: "one", 2: "two"}
	block := "function"
	check_accert(block)
	check_accert(container1)
	check_accert(container2)
	check_switch(block)
	check_switch(container2)
}

//方法3：用reflect
//  switch t := diskSpecificationConfarray.(type) {
//  case targetTypeXXX:
//      log.Printf("[]openapi_model.DiskSpecificationConf: %v\n", t)
//  default:
//      var r = reflect.TypeOf(resp.Elements)
//      log.Printf("Other:%v\n", r)
//  }
//
//   var structXXX
//  util.FormatResponse(<interface{}>XXX, targetTypeXXX)
//
// 将interface强制转换成已知类型
//func FormatResponse(data interface{}, view interface{}) error {
//    b, err := json.Marshal(data)
//    if err != nil {
//        return err
//    }
//    decoder := json.NewDecoder(bytes.NewReader(b))
//    decoder.UseNumber()
//    err = decoder.Decode(view)
//    if err != nil {
//        return err
//    }
//    return nil
//}
