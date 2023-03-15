package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/imageutils"
	"main/parser"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gorilla/mux"
)

var nbaClient *connection.Client

var notFound = ttlcache.ErrNotFound
var cache *ttlcache.Cache

func init() {
	digitsClient := gosseract.NewClient()
	digitsClient.SetTessdataPrefix("./traineddata/")
	digitsClient.SetLanguage("digitsall_layer")
	digitsClient.SetWhitelist("0123456789.")
	defer digitsClient.Close()

	alphabetClient := gosseract.NewClient()
	alphabetClient.SetTessdataPrefix("./traineddata/")
	alphabetClient.SetLanguage("eng")
	alphabetClient.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ")
	defer alphabetClient.Close()

	cache = ttlcache.NewCache()
	cache.SetTTL(time.Duration(3 * 24 * time.Hour)) // 3 Days TTL
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	metrics := cache.GetMetrics()
	fmt.Fprintf(w, fmt.Sprintf("H: %d, R: %d, M: %d", metrics.Hits, metrics.Retrievals, metrics.Misses))
	fmt.Println("Endpoint Hit: homePage")
}

func heroDmgCalculation(w http.ResponseWriter, r *http.Request) {
	playerID := r.FormValue("playerID") // e.g. "1629001" for Melton
	teamID := r.FormValue("teamID")     // e.g. "1610612763" for MEM
	gameID := r.FormValue("gameID")     // e.g. "0022100040" for MEM vs LAL on 10/24/2021
	statType := r.FormValue("statType") // e.g. "STL"
	if playerID == "" || teamID == "" || gameID == "" || statType == "" {
		w.Header().Set("Invalid Request", "Request is invalid")
		w.WriteHeader(400)
	} else {
		playerVideoResults, _ := nbaClient.GetPlayerVideos("2021-22", gameID, teamID, playerID, statType)
		json.NewEncoder(w).Encode(playerVideoResults)
		w.WriteHeader(200)
	}
	fmt.Println("Endpoint Hit: playergamevideos")
}

func nbaGames(w http.ResponseWriter, r *http.Request) {
	gameDate := r.FormValue("gameDate") // e.g. "10-29-2021"
	if gameDate == "" {
		w.Header().Set("Invalid Request", "Request is invalid")
		w.WriteHeader(400)
	} else {
		gameResults, _ := nbaClient.GetGames(gameDate)
		json.NewEncoder(w).Encode(gameResults)
		w.WriteHeader(200)
	}
	fmt.Println("Endpoint Hit: nbagames")
}

func playerVideos(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	playerName := r.FormValue("playerName")             // e.g. "De'Anthony Melton"
	teamAbbreviation := r.FormValue("teamAbbreviation") // e.g. "MEM"
	gameDate := r.FormValue("gameDate")                 // e.g. "10-28-2021"
	statType := r.FormValue("statType")                 // e.g. "STL"
	if playerName == "" || teamAbbreviation == "" || gameDate == "" || statType == "" {
		w.Header().Set("Invalid Request", "Request is invalid")
		w.WriteHeader(400)
	} else {
		key := fmt.Sprintf("%s:%s:%s:%s", playerName, teamAbbreviation, gameDate, statType)
		var gameResults []*connection.PlayerVideoResult
		if val, err := cache.Get(key); err != notFound {
			fmt.Printf("[playervideos] CacheHit: %s\n", key)
			json.Unmarshal(val.([]byte), &gameResults)
		} else {
			gameResults = service.GetVideos(nbaClient, playerName, teamAbbreviation, gameDate, statType)
			if len(gameResults) > 0 {
				// Sometimes, nbaapi lags and returns 0 results if game isn't over/just ended. Only cache results if more than 1 is returned.
				b, _ := json.Marshal(gameResults)
				cache.Set(key, b)
			}
		}
		json.NewEncoder(w).Encode(gameResults)
		w.WriteHeader(200)
	}
	fmt.Println("Endpoint Hit: playervideos")
}

// HandleRequests sets up the api endpoints.
func HandleRequests(port string) {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/playerindex", playerIndex)
	myRouter.Path("/playergamevideos").
		Queries("playerID", "{playerID:[0-9]+}",
			"teamID", "{teamID:[0-9]+}",
			"gameID", "{gameID:[0-9]+}",
			"statType", "{statType:[A-Z0-9]+}").
		HandlerFunc(playerGameVideos)
	myRouter.Path("/playergamevideos").HandlerFunc(playerGameVideos)
	myRouter.Path("/nbagames").
		Queries("gameDate", "{gameDate:\\d{4}-\\d{2}-\\d{2}").
		HandlerFunc(nbaGames)
	myRouter.Path("/nbagames").HandlerFunc(nbaGames)
	myRouter.Path("/playervideos").
		Queries("playerName", "{playerName:.+}",
			"teamAbbreviation", "{teamAbbreviation:.+}",
			"gameDate", "{gameDate:\\d{4}-\\d{2}-\\d{2}",
			"statType", "{statType:[A-Z0-9]+}").
		HandlerFunc(playerVideos)
	myRouter.Path("/playervideos").HandlerFunc(playerVideos)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}
Footer
© 2023 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
