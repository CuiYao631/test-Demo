package main

import (
	"bufio"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//测试

	ctx := context.Background()
	//firebaseApp
	opt := option.WithCredentialsFile("/Users/xiaocui/dev-env-374915-firebase-adminsdk-2skau-0308936127.json") //"/Users/xiaocui/dev-env-374915-firebase-adminsdk-2skau-0308936127.json")
	config := &firebase.Config{ProjectID: "dev-env-374915"}                                                    //"dev-env-374915"}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	//firebase firestoreClient
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting firestore client: %v\n", err)
	}

	out := make([][]string, 0)

	fi, err := os.Open("map.txt")
	if err != nil {
		log.Fatal(err)
	}
	fo, err := os.Create("map-out.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(fi)
	for {
		a, _, c := r.ReadLine()
		if c != nil {
			break
		}
		instruct := strings.Split(string(a), ",")
		if instruct[7] == "A" {
			lat, _ := strconv.ParseFloat(instruct[8], 64)
			lng, _ := strconv.ParseFloat(instruct[10], 64)
			outLat, outLng := GeoLocation(lat, instruct[9], lng, instruct[11])
			//outLats, _ := fmt.Printf("%.5f", outLat)
			//outLngs, _ := fmt.Printf("%.5f", outLng)

			//log.Println(strconv.FormatFloat(outLat, 'f', 5, 32))
			//log.Println(strconv.FormatFloat(outLng, 'f', 5, 32))
			latStr := strconv.FormatFloat(outLat, 'f', 5, 32)
			lngStr := strconv.FormatFloat(outLng, 'f', 5, 32)
			latStrInstruct := strings.Split(latStr, ".")
			lngStrInstruct := strings.Split(lngStr, ".")

			doc := make(map[string]interface{})
			doc["bike_imei"] = "869731053931051"
			doc["lat_int_part"] = latStrInstruct[0]
			doc["lat_decimal_part"] = latStrInstruct[1]
			doc["lng_int_part"] = lngStrInstruct[0]
			doc["lng_decimal_part"] = lngStrInstruct[1]
			doc["satellite_num"] = "0"
			doc["loc_precision"] = "0"
			doc["created_at"] = time.Now().Format(time.RFC3339)
			_, _, err := firestoreClient.Collection("BikeGeoLog").Add(ctx, doc)
			if err != nil {

			}

			//out = append(out, []string{strconv.FormatFloat(outLat, 'f', 5, 32), strconv.FormatFloat(outLng, 'f', 5, 32)})
		}
		//out = append(out, in)

	}

	for _, v := range out {
		fmt.Fprintln(fo, v[0], v[1])
	}
	err = fo.Close()
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(out)
	//
	//latReq := 3354.2101
	//lngReq := 15112.7983
	//lat, lng := GeoLocation(latReq, "S", lngReq, "E")
	//fmt.Printf("%.5f", lng)
	//log.Println("")
	//log.Println(lat, lng)
}
func GeoLocation(latReq float64, ns string, lngReq float64, ew string) (lat, lng float64) {

	latStr := fmt.Sprintf("%.4f", latReq)
	latParts := strings.Split(latStr, ".")

	latIntPart, _ := strconv.Atoi(latParts[0][:2])
	latFracPart, _ := strconv.ParseFloat(latParts[0][2:]+"."+latParts[1], 64)
	lat = float64(latIntPart) + latFracPart/60
	if ns == "S" {
		lat = -lat
	}

	lngStr := fmt.Sprintf("%.4f", lngReq)
	lngParts := strings.Split(lngStr, ".")

	lngIntPart, _ := strconv.Atoi(lngParts[0][:3])
	lngFracPart, _ := strconv.ParseFloat(lngParts[0][3:]+"."+lngParts[1], 64)
	lng = float64(lngIntPart) + lngFracPart/60
	if ew == "W" {
		lng = -lng
	}
	return
}
