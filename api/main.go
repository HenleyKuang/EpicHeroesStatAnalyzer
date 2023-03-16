package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"main/dmgformula"
	"main/imageutils"
	"main/parser"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/otiai10/gosseract/v2"
)

var digitsClient *gosseract.Client
var alphabetClient *gosseract.Client

func init() {
	digitsClient = gosseract.NewClient()
	digitsClient.SetTessdataPrefix("./traineddata/")
	digitsClient.SetLanguage("digitsall_layer")
	digitsClient.SetWhitelist("0123456789.")
	// defer digitsClient.Close()

	alphabetClient = gosseract.NewClient()
	alphabetClient.SetTessdataPrefix("./traineddata/")
	alphabetClient.SetLanguage("eng")
	alphabetClient.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ")
	// defer alphabetClient.Close()
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

func heroAnalysis(w http.ResponseWriter, r *http.Request) {
	// SomeValue := r.FormValue("SomeValue")
	// if SomeValue == "" {
	// 	w.Header().Set("Invalid Request", "Request is invalid")
	// 	w.WriteHeader(400)
	// } else {
	// 1. Get Image Obj.
	bytes := readFileToBytes()
	if bytes == nil {
		w.Header().Set("Error", "bytes are nil")
		w.WriteHeader(400)
	}
	// fmt.Println(bytes)
	imgObj, err := imageutils.BytesToImageObject(bytes)
	if err != nil {
		fmt.Println("[BytesToImageObject] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	imgObj = imageutils.ImageObjToGrayScale(imgObj)
	// 2. Get Hero Name.
	croppedHeroNameImgObj, err := imageutils.CropToHeroName(imgObj)
	if err != nil {
		fmt.Println("[CropToHeroName] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	croppedHeroNameImgBytes, err := imageutils.ImageObjectToBytes(croppedHeroNameImgObj)
	if err != nil {
		fmt.Println("[ImageObjectToBytes for HeroName] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	heroName, err := parser.HeroNameFromBytes(alphabetClient, croppedHeroNameImgBytes)
	if err != nil {
		fmt.Println("[HeroNameFromBytes] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	// 3. Get Main Stats.
	croppedMainStatsImgObj, err := imageutils.CropToMainStats(imgObj)
	if err != nil {
		fmt.Println("[CropToMainStats] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	croppedMainStatsImgBytes, err := imageutils.ImageObjectToBytes(croppedMainStatsImgObj)
	if err != nil {
		fmt.Println("[ImageObjectToBytes for Main Stats] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	mainStats, err := parser.MainStatsFromBytes(digitsClient, croppedMainStatsImgBytes)
	if err != nil {
		fmt.Println("[MainStatsFromBytes] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	// 4. Get Percentage Stats.
	croppedPercentageStatsImgObj, err := imageutils.CropToPercentageStats(imgObj)
	if err != nil {
		fmt.Println("[CropToPercentageStats] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	croppedPercentageImgBytes, err := imageutils.ImageObjectToBytes(croppedPercentageStatsImgObj)
	if err != nil {
		fmt.Println("[ImageObjectToBytes for Percentage Stats] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	percentageStats, err := parser.PercentageStatsFromBytes(digitsClient, croppedPercentageImgBytes)
	if err != nil {
		fmt.Println("[PercentageStatsFromBytes] ", err)
		w.Header().Set("Error", err.Error())
		w.WriteHeader(400)
	}
	allStats := mergeMaps(mainStats, percentageStats)
	dmgMap := map[string]int{
		"Basic Atk DMG":             dmgformula.BasicAtkDmg(allStats["ATK"], 0, allStats["Broken Armor"]),
		"Basic Atk DMG with Crit":   dmgformula.BasicAtkCritDmg(allStats["ATK"], 0, allStats["Broken Armor"], allStats["Crit"], allStats["Crit DMG"]),
		"Passive Atk Dmg":           dmgformula.PassiveAtkDmg(allStats["ATK"], 0, allStats["Broken Armor"], 0),
		"Passive Atk Dmg with Crit": dmgformula.PassiveAtkCritDmg(allStats["ATK"], 0, allStats["Broken Armor"], 0, allStats["Crit"], allStats["Crit DMG"]),
		"Skill Atk Dmg":             dmgformula.SkillAtkDmg(allStats["ATK"], 0, allStats["Broken Armor"], allStats["Skill DMG"], 0),
		"Skill Atk Dmg with Crit":   dmgformula.SkillAtkCritDmg(allStats["ATK"], 0, allStats["Broken Armor"], allStats["Skill DMG"], 0, allStats["Crit"], allStats["Crit DMG"]),
	}
	responseMap := map[string]interface{}{
		"Hero":          heroName,
		"Stats":         allStats,
		"Estimated Dmg": dmgMap,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseMap)
	// w.WriteHeader(200)
	// }
	fmt.Println("Endpoint Hit: heroAnalysis")
}

// HandleRequests sets up the api endpoints.
func HandleRequests(port string) {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.Path("/heroanalysis").
		// Queries("playerID", "{playerID:[0-9]+}",
		// 	"teamID", "{teamID:[0-9]+}",
		// 	"gameID", "{gameID:[0-9]+}",
		// 	"statType", "{statType:[A-Z0-9]+}").
		HandlerFunc(heroAnalysis)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func readFileToBytes() []byte {
	// Read the entire file into a byte slice
	imgPath := "./playground/data/toko_stats.jpg"
	fileObj, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("os.Open ", err)
		return nil
	}
	defer fileObj.Close()

	fileInfo, _ := fileObj.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)
	// read file into bytes
	buffer := bufio.NewReader(fileObj)
	_, err = buffer.Read(bytes)
	if err != nil {
		fmt.Println("[buffer.Read] ", err)
		return nil
	}
	return bytes
}

func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
