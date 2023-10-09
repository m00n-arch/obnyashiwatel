package main

import (
	"log"
	"net/http"

	"github.com/m00n-arch/obnyashiwatel/internal/interactions"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/obnyash", interactions.Obnyash)
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":80", nil))
}

// {"text":"avc"}

// func terminal() {
//
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter text: ")
// 	inputText, err := reader.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println(domain.RunObnyash(inputText))
// }
