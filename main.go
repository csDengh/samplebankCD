package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"sigs.k8s.io/yaml"
)

type AppSpec struct {
	Image         string   `json:"image"`
	Port          []int    `json:"port,omitempty"`
	Env           []AppEnv `json:"env,omitempty"`
	Replicas      int32    `json:"replicas"`
	EnableService bool     `json:"enable_service"`
	EnableIngress bool     `json:"enable_ingress,omitempty"`
}

type AppEnv struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MetaData struct {
	Name string `json:"name"`
}

type App struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	MetaData   `json:"metadata"`
	AppSpec    `json:"spec"`
}

func main() {
	var conf App

	yamlFile, err := ioutil.ReadFile("manifest/app.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = yaml.Unmarshal(yamlFile, &conf); err != nil {
		fmt.Println(err)
		return
	}
	name := os.Getenv("IMAGE")

	if name != "" {
		conf.AppSpec.Image = name
	}

	if out, err := yaml.Marshal(conf); err == nil {
		if err := ioutil.WriteFile("manifest/app.yaml", out, fs.ModeAppend); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

}
