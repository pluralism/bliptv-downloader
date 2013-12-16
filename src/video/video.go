package video

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LinksDef struct {
	Type    string `xml:"url,attr"`
	Quality string `xml:"role,attr"`
	Size    int    `xml:"fileSize,attr"`
}

type Video struct {
	Title       string     `xml:"channel>item>title"`
	Show        string     `xml:"channel>item>show"`
	Language    string     `xml:"channel>item>language"`
	Description string     `xml:"channel>item>puredescription"`
	Duration    int        `xml:"channel>item>runtime"`
	Links       []LinksDef `xml:"channel>item>group>content"`
}

func (v Video) GetUserAnswer() int {
	var answer int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("Title:", v.Title)
	fmt.Println("Description:", v.Description)

	fmt.Println("Available qualities: ")

	for i := 0; i < len(v.Links); i++ {
		fmt.Println(i+1, "Quality:", v.Links[i].Quality, "\tSize:", toMB(v.Links[i].Size), "MB")
	}

	fmt.Printf("Which quality do you want to donwload?")
	_, err := fmt.Fscan(stdin, &answer)
	stdin.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for answer < 1 || answer > len(v.Links) {
		fmt.Printf("Invalid choice!")
		_, err := fmt.Fscan(stdin, &answer)
		stdin.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return answer
}

func (video Video) SetTitle(newTitle string) {
	if len(newTitle) != 0 {
		video.Title = strings.Title(newTitle)
	} else {
		fmt.Println("[!] The new title can't be empty!")
	}
}

func toMB(size int) int {
	return size / (1024 * 1024)
}
