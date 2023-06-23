package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"os"
)

func main() {
    url := fmt.Sprintf("https://api.buttercms.com/v2/pages/hugo_knowledge_base/?auth_token=<auth-token>")

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error making HTTP request: %s", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error fetching content: %d %s", resp.StatusCode, resp.Status)
        return
    }

    var data struct {
        Data json.RawMessage `json:"data"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Printf("Error decoding response: %s", err)
        return
    }

    
file, err := os.OpenFile("data/my_content_type.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
if err != nil {
    fmt.Printf("Error opening JSON file: %s", err)
    return
}
defer file.Close()

// Write the JSON data to the files
if err := json.NewEncoder(file).Encode(data); err != nil {
    fmt.Printf("Error writing JSON data to file: %s", err)
    return
}

    fmt.Printf("%s", data.Data)
}
