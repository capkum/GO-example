// 특정 디렉토리에서 파일리스트를 읽어와 디비에 입력하는 샘플코드
// getFiles()는 특정 디렉토리에서 파일목록을 읽어 []os.FileInfo로 리턴
// 샘플코드는 몽고디비의 json형식의 파일들을 mongoimport를 이용하여 입력하는 예
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type FileList struct {
	files []os.FileInfo
	err   error
}

// 임의 폴더안에 있는 파일명을 읽어온다
func getFiles(dirname string) *FileList {
	fileList := new(FileList)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fileList.err = err
	}
	fileList.files = files
	return fileList
}

// 읽어온  json 파일명을 몽고디비에 입력한다
// 디비는 mongodb를 예로 들었음
func insertDb(files *FileList, dirNm string, DbName string) {
	fmt.Println("======== start ==========")
	for _, value := range files.files {
		fileName := value.Name()
		collectionName := strings.Replace(fileName, ".json", "", -1)
		cmd := "mongoimport --db " + DbName + " --collection "
		cmd += collectionName + " --type json --file "
		cmd += dirNm + "/" + fileName

		_, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(fileName, " || ", collectionName)
	}
	fmt.Println("======== done ==========")
}

func main() {
	dbName := "DbName"
	dirName := "dirpath"
	files := getFiles(dirName)
	if files.err != nil {
		errorType := reflect.TypeOf(files.err)
		errorMsg := files.err.Error()
		fmt.Print(errorType)
		fmt.Println(": " + errorMsg)
		return
	}

	go insertDb(files, dirName, dbName)

	fmt.Scanln()
}
