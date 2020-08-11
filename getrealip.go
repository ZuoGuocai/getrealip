/*
Author: zuoguocai@126.com
*/

package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "net/http/httputil"
    "strings"
)




func main() {
    log.Println("Server Get Realip")
    http.HandleFunc("/", GetRealIP)
    http.ListenAndServe(":12345", nil)
}

func GetRealIP(w http.ResponseWriter, r *http.Request) {
    dump, _ := httputil.DumpRequest(r, false)
    log.Printf("%q\n", dump)
    head := `<h1 align="center" style="color:red;">Get Real IP</h1>`
    
    r1 := strings.Join([]string{"<h3 style='background-color:powderblue;'>","RemoteAddr:  ",r.RemoteAddr,"</h1>"},"")
    r2 := strings.Join([]string{"<h3 style='background-color:#DDA0DD;'>","X-Original-Forwarded-For:  ",r.Header.Get("X-Original-Forwarded-For"),"</h1>"},"")
    r3 := strings.Join([]string{"<h3 style='background-color:powderblue;'>","X-Forwarded-For:  ",r.Header.Get("X-Forwarded-For"),"</h1>"},"")
    r4 := strings.Join([]string{"<h3 style='background-color:#DDA0DD;'>","X-Real-Ip:  ",r.Header.Get("X-Real-Ip"),"</h1>"},"")
    html := strings.Join([]string{head,r1,r2,r3,r4},"")

    io.WriteString(w, fmt.Sprintf(html))
    return
}
