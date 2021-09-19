package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
    bufSize   = 1024 * 8
)

//запись запроса GET в файл. Если имя файла задано пустой строкой,
//  то в качестве имени файла используется //последний фрагмент входного URL-адреса.
//   Eg: http://foo/baz.jar => baz.jar
func Wget(url, fileName string) {
    resp := getResponse(url)
    if fileName == "" {
        urlSplit := strings.Split(url, "/")
        fileName = urlSplit[len(urlSplit)-1]
    }
    writeToFile(fileName, resp)
}


// делается запрос GET возвращается ответ
func getResponse(url string) *http.Response {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(url)
    errorChecker(err)
    return resp
}


// запись овтета GET  файл
func writeToFile(fileName string, resp *http.Response) {
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
    errorChecker(err)
    defer file.Close()
    bufferedWriter := bufio.NewWriterSize(file, bufSize)
    errorChecker(err)
    _, err = io.Copy(bufferedWriter, resp.Body)
    errorChecker(err)
}

// проверка на ошибку и паника
func errorChecker(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	if len(os.Args) == 3 {
		Wget(os.Args[1], os.Args[2])
	} else if len(os.Args) == 2 {
		Wget(os.Args[1], "")
	}
}