package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"bytes"
	"mime/multipart"
	"io/ioutil"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm() //解析参数 默认不解析
	fmt.Println("path",r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello world")
}
func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.Method)
	if r.Method == "GET"{
		t,_ := template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	}else{
		r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}

}

func upload(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method == "GET"{
		crutime  := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t,_ := template.ParseFiles("upload.gtpl")
		t.Execute(w,token)

	}else {
		r.ParseMultipartForm( 32 << 20)
		file,handle,err := r.FormFile("uploadfile")
		if err != nil{
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w,"%v",handle.Header)

		f,err := os.OpenFile("./test/"+handle.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil{
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f,file)

	}
	
}

func postFile(filename string ,targetUrl string) error  {
	bodyBuf := &bytes.Buffer{}
	bodyWriter  := multipart.NewWriter(bodyBuf)

	fileWriter,err  := bodyWriter.CreateFormFile("uploadfile",filename)

	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh,err := os.Open(filename)

	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	defer fh.Close()

	_,err = io.Copy(fileWriter,fh)

	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()

	resp,err := http.Post(targetUrl,contentType,bodyBuf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	resp_body,err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

func postfile1(w http.ResponseWriter,r *http.Request)  {
	targetUrl := "http://localhost:9090/upload"

	filename := "./wite.txt"

	postFile(filename,targetUrl)

}



func main()  {
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/postfile",postfile1)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}