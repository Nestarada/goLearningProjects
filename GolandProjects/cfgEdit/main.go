package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type cfgStruct struct {
	Service string `json:"service"`
	Data    []struct {
		OnLoad bool   `json:"onload"`
		Prior  string `json:"prior"`
	} `json:"data"`
}

var cfgs map[string]struct{}

func cfgListGetter() {
	for i := range cfgs {
		delete(cfgs, i)
	}
	filepath.Walk("config", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			cfgs[strings.TrimRight(info.Name(), ".json")] = struct{}{}
		}
		return nil
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		cfgListGetter()
		name := strings.TrimLeft(r.URL.Path, "/config/")
		if _, exists := cfgs[name]; exists {
			err := os.Remove("config/" + name + ".json")
			if err != nil {
				fmt.Errorf("error deleting file %v", err)
			}
			fmt.Fprintf(w, "Config succesfully deleted")
		} else {
			fmt.Fprintf(w, "Config %s not found", name)
		}
	case http.MethodGet:
		cfgListGetter()
		if r.URL.Path == "/config/List" {
			fmt.Fprintf(w, "List of cfgs:\n")
			for i := range cfgs {
				fmt.Fprintf(w, "%s\n", i)
			}
		} else {
			name := strings.TrimLeft(r.URL.Path, "/config/")
			if _, exists := cfgs[name]; exists {
				b, _ := ioutil.ReadFile(strings.TrimLeft(r.URL.Path+".json", "/"))
				fmt.Fprintf(w, "%s", b)
			} else {
				fmt.Fprintf(w, "Config %s not found", name)
			}
		}
	case http.MethodPost:
		cfgListGetter()
		config := cfgStruct{}
		reqSize, _ := strconv.Atoi(r.Header.Get("Content-Length"))
		buff := make([]byte, reqSize)
		_, err := r.Body.Read(buff)
		if err != nil {
			fmt.Errorf("error reading request %v", nil)
		}
		err = json.Unmarshal(buff, &config)
		if err != nil {
			fmt.Errorf("error parsing request body %v", err)
		}
		if _, exists := cfgs[config.Service]; !exists {
			err := os.WriteFile("config/"+config.Service+".json", buff, 0666)
			if err != nil {
				fmt.Errorf("error creating config %v", err)
			}
			fmt.Fprintf(w, "Config was created")
		} else {
			err := os.Rename("config/"+config.Service+".json", "config/"+config.Service+"_old.json")
			if err != nil {
				fmt.Errorf("error creating config %v", err)
			}
			err = os.WriteFile("config/"+config.Service+".json", buff, 0666)
			if err != nil {
				fmt.Errorf("error creating config %v", err)
			}
			fmt.Fprintf(w, "Config was updated, old config saved as %s", config.Service+"_old.json")
		}
	}
}

func main() {
	cfgs = make(map[string]struct{})
	cfgListGetter()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		fmt.Errorf("error while listening for requests %v", err)
	}
}
