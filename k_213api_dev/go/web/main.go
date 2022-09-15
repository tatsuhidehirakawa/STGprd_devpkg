package main 

import (
    "github.com/gin-gonic/gin"
    "bytes"                      // 追加
    "encoding/json"              // 追加
    "fmt"                        // 追加
    "net/http"                   // 追加
)

/* models/models.goへ移植
type txjsn struct {
    // Table Name:"account_attribute"
    account_id        int         `json: "account_id"`
    update_date       timestamp   `json: "update_date"`
    update_id         string      `json: "update_id"`
    pic_id_01         string      `json: "pic_id_01"`
    skil              string      `json: "skil"`
    // Table Name:"account_master"
    user_id           string      `json: "user_id"`
    account_name      string      `json: "account_name"`
    mail_addless      string      `json: "mail_addless"`
    password          string      `json: "password"`
    zip_code          string      `json: "zip_code"`
    barth_date        timestamp   `json: "barth_date"`
    time_zone         string      `json: "time_zone"`
    language          string      `json: "language"`
    status            string      `json: "status"`
    registration_date string `json: "registration_date"`
}
*/

func main() {

    // *gin.EngineをReturnする関数を宣言
    router := gin.Default()

/*---GET処理-----+-----+-----+-----+-----+-----+-----+-----+---*/
    // jsonデータをGETさせる処理
    router.GET("/", func(c *gin.Context) {

    // GO構造体からjsonに変換する処理…<2>
 
    func main() {
        // Table Name:"account_attribute"
        txjsn := new(txjsn)
        txjsn.account_id = "account_id"
        txjsn.update_date = "update_date"
        txjsn.update_id = "update_id"
        txjsn.pic_id_01 = "pic_id_01"
        txjsn.skil = "skil"
        // Table Name:"account_master"
        txjsn.user_id = "user_id"
        txjsn.account_name = "account_name"
        txjsn.mail_addless = "mail_addless"
        txjsn.password = "password"
        txjsn.zip_code = "zip_code"
        txjsn.barth_date = "barth_date"
        txjsn.time_zone = "time_zone"
        txjsn.language = "language"
        txjsn.status = "status"
        txjsn.registration_date = "registration_date"

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

    // DBから取り出した値をGO構造体に格納する処理…<3>



    // 値を取り出すSQL文をDBに投げる処理…<7>



    // gin.H{}でgoファイルの変数をHTMLテンプレートに渡す…<5>
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
/*---POST処理----+-----+-----+-----+-----+-----+-----+-----+---*/
    // jsonデータをPOSTさせる処理
    router.POST("/", func(c *gin.Context) {

    // jsonをGO構造体に変換する処理…<1>



    // GO構造体から取り出した値をDBに格納する処理…<4>



    // 値を取り出すSQL文をDBに投げる処理…<8>



    // gin.H{}でHTMLテンプレートからgoファイルの変数を受け取る…<6>
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
/*---+-----+-----+-----+-----+-----+-----+-----+-----+-----+---*/
    // サーバを起動する処理
    router.Run(":3000")
}
