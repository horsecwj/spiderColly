package graphQl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func main() {
	song := make(map[string]string)
	song["username"] = "******"
	song["password"] = "******"

	//	```
	//  {
	//    flashLoans(first: 10, orderBy: timestamp, orderDirection: desc) {
	//      id
	//      reserve {
	//        name
	//        symbol
	//      }
	//      amount,
	//      target,
	//      timestamp
	//    }
	//  }
	//```

	bytesData, _ := json.Marshal(song)

	res, err := http.Post("http://xxxxxx.com",
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	fmt.Println(*str)

}
