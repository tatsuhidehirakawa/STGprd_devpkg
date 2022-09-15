package main 

import (
    "github.com/gin-gonic/gin"
    "bytes"                      // 追加
    "encoding/json"              // 追加
    "fmt"                        // 追加
    "net/http"                   // 追加
)

type Account_attribute struct {
    account_id        int       `json:"account_id"`
    update_date       time.Time `json:"update_date"`
    update_id         string    `json:"update_id"`
    pic_id_01         string    `json:"pic_id_01"`
    skil              string    `json:"skil"`
}

type Account_master struct {
    user_id           string    `json:"user_id"`
    account_name      string    `json:"account_name"`
    mail_addless      string    `json:"mail_addless"`
    password          string    `json:"password"`
    zip_code          string    `json:"zip_code"`
    barth_date        time.Time `json:"barth_date"`
    time_zone         string    `json:"time_zone"`
    language          string    `json:"language"`
    status            string    `json:"status"`
    registration_date string    `json: "registration_date"`
}

func main() {

    // *gin.EngineをReturnする関数を宣言
    router := gin.Default()

/*---GET処理-----+-----+-----+-----+-----+-----+-----+-----+---*/
    // jsonデータをGETさせる処理
    router.GET("/", func(c *gin.Context) {

    // GO構造体からjsonに変換する処理…<#001>
 
        // Table Name:"account_attribute"
        Account_attribute := new(txjsn)
        Account_attribute.account_id = "account_id"
        Account_attribute.update_date = "update_date"
        Account_attribute.update_id = "update_id"
        Account_attribute.pic_id_01 = "pic_id_01"
        Account_attribute.skil = "skil"
    
        // Table Name:"account_master"
        Account_master.user_id = "user_id"
        Account_master.account_name = "account_name"
        Account_master.mail_addless = "mail_addless"
        Account_master.password = "password"
        Account_master.zip_code = "zip_code"
        Account_master.barth_date = "barth_date"
        Account_master.time_zone = "time_zone"
        Account_master.language = "language"
        Account_master.status = "status"
        Account_master.registration_date = "registration_date"

        // encode json
        txjsn_json, _ := json.Marshal(txjsn)
        fmt.Printf("[+] %s\n", string(txjsn_json))

        // send json
        //// ポイント2, 3
        res, err := http.Post(URL, "application/json", bytes.NewBuffer(txjsn_json))
        defer res.Body.Close()

        if err != nil {
            fmt.Println("[!] " + err.Error())
        } else {
            fmt.Println("[*] " + res.Status)
        }
    }

    // DBから取り出した値をGO構造体に格納する処理…<#003>



    // 値を取り出すSQL文をDBに投げる処理…<#007>



    // gin.H{}でgoファイルの変数をHTMLテンプレートに渡す…<#005>
    c.JSON(200, gin.H{
            // Table Name:"account_attribute"
            "account_id": "account_id",
            "update_date": "update_date",
            "update_id": "update_id",
            "pic_id_01": "pic_id_01",
            "skil": "skil",
            // Table Name:"account_master"
            "user_id": "user_id",
            "account_name": "account_name",
            "mail_addless": "mail_addless",
            "password": "password",
            "zip_code": "zip_code",
            "barth_date": "barth_date",
            "time_zone": "time_zone",
            "language": "language",
            "status": "status",
            "registration_date": "registration_date",
        })
    })
