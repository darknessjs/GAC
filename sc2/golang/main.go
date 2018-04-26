package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/bitly/go-simplejson"
	"database/sql"
)

type sc2UserProfile struct {
	id int
	realm int
	name string
	displayName string
	clanName string
	clanTag string
	profilePath string
}
type sc2Characters struct {
	characters []sc2UserProfile
}

var key string = "df38f7wpmws82n64tcknhmwcqw3jfc9p"

func main() {

	http.HandleFunc("/", getApiData) //设置访问的路由
	http.HandleFunc("/saveAccessToken", saveUserAndGetData)
	err := http.ListenAndServe(":9001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//  reqest, _ = http.NewRequest("POST","http:/127.0.0.1/", bytes.NewBufferString(data.Encode()))
	//    respet1,_ := http.NewRequest("POST","http://127.0.0.1/",url.Values{"key":"Value"})
	//    reqest1.Header.Set("User-Agent","chrome 100")
	//    client.Do(reqest1)
}

func saveUserAndGetData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	var accessToken string = strings.Join(r.Form["accessToken"],"")

	requestURI := "https://api.battlenet.com.cn/sc2/profile/user?access_token=" + accessToken
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", requestURI, nil);
	response,_ := client.Do(reqest)
	if response.StatusCode == 200 {
		//var profile sc2Characters
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body);
		jsonProfile, _ := simplejson.NewJson([]byte(bodystr))
		savingUser := jsonProfile.Get("characters").GetIndex(0)
		id := savingUser.Get("id").MustInt()
		realm := savingUser.Get("realm").MustInt()
		displayName := savingUser.Get("displayName").MustString()
		name := savingUser.Get("name").MustString()
		clanName := savingUser.Get("clanName").MustString()
		clanTag := savingUser.Get("clanTag").MustString()
		profilePath := savingUser.Get("profilePath").MustString()
		//fmt.Print(id)
		//fmt.Print(realm)
		//fmt.Print(displayName)
		//fmt.Print(name)
		//fmt.Print(clanName)
		//fmt.Print(clanTag)
		//fmt.Print(profilePath)

		db, err := sql.Open("mysql", "root:Zirly123@/galaxycluster");
		if err != nil {
			panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
		}
		defer db.Close();
		insertSc2Profile, err := db.Prepare("INSERT INTO sc2_user VALUES(?,?,?,?,?,?,?)")
		insertSc2Profile.Exec(id,realm,name,displayName,clanName,clanTag,profilePath)

		fmt.Fprintf(w, bodystr) //这个写入到w的是输出到客户端的
	}
}

func getApiData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	if (r.URL.Path == "/favicon.ico") {
		fmt.Fprintf(w, "Hello astaxie!")
	} else {
		sendGetRes(r.URL.Path, w)
	}
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
		fmt.Println("抛出异常")
		w.WriteHeader(response.StatusCode);
		fmt.Fprintf(w, "2333");
		//fmt.Fprintf(w, "error") //这个写入到w的是输出到客户端的
	}
}