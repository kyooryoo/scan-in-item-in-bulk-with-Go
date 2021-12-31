package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var library string
	var circ_desk string
	var in_house_use string
	var barcode string
	var file string
	flag.StringVar(&library, "l", "MAIN", "Specify library code. Default is MAIN.")
	flag.StringVar(&circ_desk, "p", "DEFAULT_CIRC_DESK", "Specify circulation desk code. Default is DEFAULT_CIRC_DESK.")
	flag.StringVar(&in_house_use, "i", "false", "Specify if register in house use. Default is false.")
	flag.StringVar(&barcode, "b", "NOTEXIST", "Specify the barcode to scan. Default is NOTEXIST.")
	flag.StringVar(&file, "f", "./barcodes.txt", "Specify the file path for barcode list. Default is ./barcodes.txt")
	flag.Parse()

	apikey_value := os.Getenv("APIKEY")
	apikey := fmt.Sprintf("apikey %s", apikey_value)
	url_base := "https://api-cn.hosted.exlibrisgroup.com.cn/almaws/v1"

	if barcode != "NOTEXIST" {
		scanitem(barcode, url_base, apikey, library, circ_desk, in_house_use)
	} else {
		barcodes, err := os.Open("./barcodes.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer barcodes.Close()

		scanner := bufio.NewScanner(barcodes)
		for scanner.Scan() {
			barcode := fmt.Sprint(scanner.Text())
			scanitem(barcode, url_base, apikey, library, circ_desk, in_house_use)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func scanitem(barcode string, url_base string, apikey string, library string, circ_desk string, in_house_use string) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Timeout: timeout}
	url_barcode := fmt.Sprintf("%s/items?item_barcode=%s", url_base, barcode)
	get_req, _ := http.NewRequest("GET", url_barcode, nil)
	get_req.Header.Set("accept", "application/json")
	get_req.Header.Set("authorization", apikey)
	get_res, err := client.Do(get_req)
	if err != nil {
		log.Fatalln(err)
	}
	defer get_res.Body.Close()

	var result map[string]map[string]interface{}
	json.NewDecoder(get_res.Body).Decode(&result)
	mms_id := result["bib_data"]["mms_id"]
	holding_id := result["holding_data"]["holding_id"]
	item_id := result["item_data"]["pid"]
	base_item := fmt.Sprintf("%s/bibs/%s/holdings/%s/items/%s", url_base, mms_id, holding_id, item_id)
	params := fmt.Sprintf("op=scan&library=%s&circ_desk=%s&register_in_house_use=%s", library, circ_desk, in_house_use)

	url_item := fmt.Sprintf("%s?%s", base_item, params)
	post_req, _ := http.NewRequest("POST", url_item, nil)
	post_req.Header.Set("accept", "application/json")
	post_req.Header.Set("authorization", apikey)

	post_res, err := client.Do(post_req)
	if err != nil {
		log.Fatalln(err)
	}
	defer post_res.Body.Close()
	if post_res.StatusCode == 200 {
		fmt.Printf("%s is processed!\n", barcode)
	} else {
		fmt.Printf("Check item %s!\n", barcode)
	}
}
