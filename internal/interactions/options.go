package interactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/m00n-arch/obnyashiwatel/internal/domain"
)

func Obnyash(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		obnyashPost(w, r)
	case http.MethodOptions:
		obnyashOptions(w, r)
	}
}

func obnyashPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	type request struct {
		Text string `json:"text"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var re request
	err = json.Unmarshal(body, &re)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write([]byte(domain.RunObnyash(re.Text)))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func obnyashOptions(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}
