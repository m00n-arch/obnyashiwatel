package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/m00n-arch/obnyashiwatel/internal/domain"
)

func main() {
	http.HandleFunc("/obnyash", obnyash)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

type request struct {
	Text string `json:"text"`
}

// {"text":"avc"}

func obnyash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var re request
	err = json.Unmarshal(body, &re)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(domain.RunObnyash(re.Text)))
	if err != nil {
		panic(err)
	}
}

func terminal() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	inputText, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println(domain.RunObnyash(inputText))
}
