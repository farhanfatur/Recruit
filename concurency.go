package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const endpointProvinsi = "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
const endpointMuseumByKodeProp = "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_prop="

var concurrentLimitFlag int
var outputFlag string

var dataMuseumCsvHeader = []string{"alamat_jalan", "bangunan", "bujur", "desa_kelurahan", "kabupaten_kota", "kecamatan", "kode_pengelolaan", "koleksi", "lintang", "luas_tanah", "museum_id", "nama", "pengelola", "propinsi", "sdm", "standar", "status_kepemilikan", "sumber_dana", "tahun_berdiri", "tipe"}

type M map[string]interface{}

func grab(endpoint string, result interface{}) error {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	buf, _ := ioutil.ReadAll(response.Body)
	str := strings.Replace(string(buf), "\ufeff", "", -1)

	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		return err
	}

	return nil
}

func groupMuseumByKabupatenKota(data []M) map[string][]M {
	dataMuseumByKabupatenKota := make(map[string][]M, 0)

	for _, row := range data {
		key := row["kabupaten_kota"].(string)
		rows := make([]M, 0)

		if v, ok := dataMuseumByKabupatenKota[key]; ok {
			rows = v
		}
		rows = append(rows, row)
		dataMuseumByKabupatenKota[key] = rows
	}

	return dataMuseumByKabupatenKota
}

func saveToCsv(kabupatenKota string, rows []M) error {
	file, err := os.Create(kabupatenKota + ".csv")
	if err != nil {
		return nil
	}
	defer file.Close()

	Writer := csv.NewWriter(file)
	defer Writer.Flush()

	for i, row := range rows {
		if i == 0 {
			Writer.Write(dataMuseumCsvHeader)
		} else {
			csvContent := make([]string, 0)
			for _, key := range dataMuseumCsvHeader {
				if val, ok := row[key]; ok {
					csvContent = append(csvContent, fmt.Sprintf("%v", val))
				} else {
					csvContent = append(csvContent, "")
				}
			}
			Writer.Write(csvContent)
		}
	}

	return nil
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)

		endPoint := endpointMuseumByKodeProp + strings.TrimSpace(j)
		dataMuseum := struct {
			Data []M
		}{}
		err := grab(endPoint, &dataMuseum)
		if err != nil {
			wg.Done()
			continue
		}

		dataMuseumByKabupatenKota := groupMuseumByKabupatenKota(dataMuseum.Data)
		for key, rows := range dataMuseumByKabupatenKota {
			saveToCsv(filepath.Join(outputFlag, key), rows)
		}
		fmt.Println("worker", id, "finished job", j)
		wg.Done()
	}
}

func makeSureDirectoryExists(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			log.Fatal(merr)
		}
	}
}

func main() {
	flag.IntVar(&concurrentLimitFlag, "concurrent_limit", 0, "")
	flag.StringVar(&outputFlag, "output", "", "")

	flag.Parse()
	makeSureDirectoryExists(outputFlag)
	dataProvinsi := struct {
		Data []M
	}{}
	err := grab(endpointProvinsi, &dataProvinsi)
	if err != nil {
		log.Fatal(err)
	}
	jobs := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(len(dataProvinsi.Data))

	for w := 1; w <= concurrentLimitFlag; w++ {
		go worker(w, jobs, wg)
	}

	for _, each := range dataProvinsi.Data {
		jobs <- each["kode_wilayah"].(string)
	}

	close(jobs)
	wg.Wait()
}
