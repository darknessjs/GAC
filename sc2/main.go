package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
)

var key string = "df38f7wpmws82n64tcknhmwcqw3jfc9p"

func main() {

	http.HandleFunc("/", getApiData) //设置访问的路由
	err := http.ListenAndServe(":9001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//  reqest, _ = http.NewRequest("POST","http:/127.0.0.1/", bytes.NewBufferString(data.Encode()))
	//    respet1,_ := http.NewRequest("POST","http://127.0.0.1/",url.Values{"key":"Value"})
	//    reqest1.Header.Set("User-Agent","chrome 100")
	//    client.Do(reqest1)
}

func getApiData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	sendGetRes(r.URL.Path, w)
	//fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func sendGetRes(api string, w http.ResponseWriter) {
	locale := "zh_CN";
	apiKey := "dtwpk6g4rjjjqcypnjdtms69x5jjpmm4";
	var baseUrl string = "https://api.battlenet.com.cn"
	resultUrl := baseUrl + api + "?locale="+locale+"&apikey="+apiKey;
	fmt.Println(resultUrl)

	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", resultUrl, nil)

	response,_ := client.Do(reqest)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body);
		fmt.Println(bodystr)
		fmt.Fprintf(w, bodystr) //这个写入到w的是输出到客户端的
	} else {
		//fmt.Fprintf(w, "error") //这个写入到w的是输出到客户端的
	}
}