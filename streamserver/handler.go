package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//接收到的视频流处理
func StreamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vid := ps.ByName("vid")
	videoLink := VIDEO_DIR + vid //视频链接
	video, err := os.Open(videoLink)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")  //视频解析方式
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

//上传文件处理
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body= http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)//限制上传视频大小
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "File is too big")//code:400
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Internal error")//code:500
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError,"Internal error")//code:500
	}

	fn := p.ByName("vid")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)//设置权限666
	if err != nil {
		log.Printf("Write file error:%v", err)
		SendErrorResponse(w, http.StatusInternalServerError,"Internal error")//code:500
	}

	SendNormalResponse(w, http.StatusCreated, "Uploaded successfully")//code:201
	// 返回正确信息
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

