package main

import (
	"bufio"
	"fmt"
	"os"
	"url"
)

func main() {
	var userURL string
	stdin := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a valid URL:")
	_, err := fmt.Fscan(stdin, &userURL)
	stdin.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for len(userURL) == 0 {
		fmt.Println("[!] URL can't be empty!")
		_, err := fmt.Fscan(stdin, &userURL)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	response := url.GetPageSource(userURL)
	videoUser := url.GetVideoDetails(response)
	videoNumber := videoUser.GetUserAnswer()
	url.DownloadFile(videoUser.Links[videoNumber-1].Type, videoUser)
}
