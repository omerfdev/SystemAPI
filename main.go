package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type SystemInfo struct {
    Hostname string `json:"hostname"`
    OS       string `json:"os"`
    Arch     string `json:"architecture"`
}

func main() {
    http.HandleFunc("/systeminfo", handleSystemInfo)

    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

func handleSystemInfo(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    systemInfo := SystemInfo{
        Hostname: hostname,
        OS:       getOS(),
        Arch:     getArchitecture(),
    }

    jsonResponse, err := json.Marshal(systemInfo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func getOS() string {
    return os.Getenv("OS")
}

func getArchitecture() string {
    return os.Getenv("PROCESSOR_ARCHITECTURE")
}
