package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "bufio"
    "encoding/json"
    "regexp" 
    "mux"
)

type object struct {
    AOR string `json:"addressOfRecord"`
    TID string `json:"tenantId"`
    URI string `json:"uri"`
    Cont  string `json:"contact"`
    PH string `json:"path"`
    SRC string `json:"source"`
    TGT string `json:"target"`
    UA string `json:"userAgent"`
    RUA string `json:"rawUserAgent"`
    CT string `json:"created"`
    LID string `json:"lineId"`
}

func main() {
	r := mux.NewRouter()
    
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := "Welcome to Golang"
        fmt.Fprintf(w, "%s\n",content)
        
    
	})
    
    
	r.HandleFunc("/aor/", func(w http.ResponseWriter, r *http.Request) {
	
        file, err := os.Open("regs.json")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {             // internally, it advances token based on sperator
            data := scanner.Text()
            match, _ := regexp.MatchString("{\"addressOfRecord\".*?\"}", data )
            if match == true {
                theJson := data
                var obj object
                json.Unmarshal([]byte(theJson), &obj)

                fmt.Fprintf(w, "%s\n",obj.AOR)
                }
            }
	})
    

    r.HandleFunc("/aor/{id:[a-zA-Z0-9_.-]*$}", func(w http.ResponseWriter, r *http.Request) {
        file, err := os.Open("regs.json")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        id := mux.Vars(r)["id"]
        confMap := map[string]string{}
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {             // internally, it advances token based on sperator
            data := scanner.Text()
            match, _ := regexp.MatchString("{\"addressOfRecord\".*?\"}", data )
            if match == true {
                theJson := data
                var obj object
                json.Unmarshal([]byte(theJson), &obj)
                confMap[obj.AOR] =  data
                }

            }
        if v, ok := confMap[id]; ok {
            fmt.Fprintf(w, "%s", v)
        }else {
                fmt.Fprintf(w, "404 page not found\n")
            }
	})
	http.ListenAndServe(":80", r)
}

