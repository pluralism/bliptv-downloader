package url

import (
	"encoding/xml"
	"errors"
	"files"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"video"
)

func getID(url string) string {
	return regexp.MustCompile(`\d+$`).FindString(url)
}

func IsValidURL(url string) bool {
	// [^/] matches anything that is not a forward slash (/).
	pattern := regexp.MustCompile(`(http://)?blip.tv/[^/]+/[^/]+\d+`)
	return pattern.MatchString(url)
}

func GetDetailsURL(url string) (string, error) {
	const BASE_URL string = "http://blip.tv/rss/flash/"

	if IsValidURL(url) {
		return string(BASE_URL + getID(url)), nil
	}
	return "Error", errors.New("[!] Invalid video ID")
}

func GetPageSource(url string) string {
	URL, errURL := GetDetailsURL(url)
	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("[!] An error occurred while downloading the source code!")
		fmt.Println(err)
		os.Exit(1)
	}

	if errURL != nil {
		fmt.Println(errURL)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[!] An error occurred while reading the body!")
		os.Exit(1)
	}

	bodyString := string(body)
	strings.Replace(bodyString, "&lt;", "<", -1)
	strings.Replace(bodyString, "&gt;", ">", -1)

	return bodyString
}

func DownloadFile(url string, v video.Video) {
	var qual map[string]string
	qual = make(map[string]string)
	qual["video/mp4"] = "mp4"
	qual["video/quicktime"] = "mov"
	qual["video/msvideo"] = "avi"
	qual["video/x-msvideo"] = "avi"

	tempURL := url + "?showplayer=1" //Or any number
	response, err := http.Get(tempURL)

	if err != nil {
		fmt.Println("[!] An error occurred while trying to download the file!")
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	var fileName string
	contentType := response.Header.Get("Content-Type")
	value, ok := qual[contentType]

	if ok != false {
		tempDir := "C:\\BLIPTV-DOWNLOADS\\"

		if err != nil {
			fmt.Println("[!] Could not get the current username!")
			os.Exit(1)
		}

		if !files.FileExists(tempDir + v.Show) {
			os.Mkdir(tempDir, 0700)
			os.Mkdir(tempDir+v.Show, 0700)
			fmt.Println("[*] Created directory \"", tempDir+v.Show, "\"")
		}
		fileName = tempDir + v.Show + "\\" + files.GetFileName(v.Title) + "." + value
	} else {
		fmt.Println("[!] The program could not recognize the format of the video!")
		os.Exit(1)
	}

	write, err := os.Create(fileName)
	defer write.Close()

	fmt.Println("[*] Downloading the file... It can take a while...")
	writeFile, err := io.Copy(write, response.Body)

	if err != nil {
		fmt.Println("[!] An error occurred while downloading the file!")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("[*] Copied", writeFile, "bytes!")
	os.Exit(1)
}

func GetVideoDetails(source string) video.Video {
	v1 := video.Video{}
	err := xml.Unmarshal([]byte(source), &v1)
	if err != nil {
		fmt.Println(err)
		return video.Video{}
	}

	return v1
}
