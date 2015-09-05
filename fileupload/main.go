package main

import (
    "github.com/astaxie/beego"

    "fmt"
    "path"
)

type UploadController struct {
    beego.Controller
}

func (this *UploadController) Post() {
    f, h, err := this.GetFile("fileToUpload")
    if err != nil {
        fmt.Println("getfile err ", err)
    }
    fmt.Println("filename:", h.Filename)
    
    tofilepath := "./static/files"
    tofilepath = path.Join(tofilepath, h.Filename)
    fmt.Println("tofilepath:",tofilepath)

    f.Close()
    err = this.SaveToFile("fileToUpload", tofilepath)
    if err != nil {
        fmt.Println("err:", err)
    }
    this.Redirect("/static/files", 302)
}

func main() {
    beego.Router("/upload", &UploadController{})
    beego.Run() 
}