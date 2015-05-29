package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

type HealthCheck struct {
    Msg string
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    h := &HealthCheck{"ok"}
    body := make(map[string]interface{})
    body["msg"] = h.Msg

    b, err := json.Marshal(body)
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}

func posReinforced(w http.ResponseWriter, r *http.Request) {
    /**
    * Steps
    * 1. Get time the POS was reinforced.
    * 2. Calculate time the POS will come out of reinforced.
    * 3. Return POS data and reinforced information.
    */
}

func InitHandlers() {
    http.HandleFunc("/health", healthCheckHandler)
    http.HandleFunc("/corp/account-balance", corpAccountBalanceHandler)
    http.HandleFunc("/corp/asset-list", corpAssetListHandler)
    http.HandleFunc("/corp/starbases/{id}", corpStarbaseDetailHandler)
    http.HandleFunc("/corp/starbases", corpStarbaseListHandler)
    http.HandleFunc("/pos/reinforced", posReinforced)
}

