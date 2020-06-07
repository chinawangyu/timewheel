package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"timewheel/Lib"
)

type HttpHandler struct {
}

//添加时间轮,存储到redis
func (http *HttpHandler) AddTimeWheel(w http.ResponseWriter, r *http.Request) {
	type BodyWheelTimeStruct struct {
		Runtime Lib.Time `json:"runtime"`
	}

	bodyWheelTimeStruct := &BodyWheelTimeStruct{}

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, bodyWheelTimeStruct); err != nil {
		log.Fatal(err.Error())
		return
	}

	//写文件-判断是加入时间轮还是加入数据文件

	fmt.Fprintf(w, string(body))
}
