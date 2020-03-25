package basicsample

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"` // todo: ``符号在这里的作用是什么？字段必须以大写字母开头吗
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} //todo: 为什么用slice来形容数组、map？json.Marshal() 为什么同样的方法，返回值的个数不一样？
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}} //todo:为什么本行末尾的两个}需要放在一起，将最后一个}换行会显示错误
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a", "b"]}`)
	var dat map[string]interface{}                    //todo:为什么用这种结构存放反序列化后的值？interface{}是什么结构
	if err := json.Unmarshal(byt, &dat); err != nil { //todo: 这是什么判断语句
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1, "fruits":["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // todo:这样不是更安全吗，为什么有未处理的错误，难道是需要接收一下错误，再明确不处理？
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) // todo: 没明白这个小片段的代码，os.Stdout是什么？
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
