package config

import ( 
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Token struct {
	Token string 
}

var configFolder = "/cgptconfig/"
var configFile = "config.yml"

func GetApiKey() string {
	fmt.Print("Enter openai Apikey (You can get your from https://platform.openai.com/account/api-keys): ")
	var data string
	if _, err := fmt.Scan(&data); err != nil {
		panic(err)
	}
	return data
}

// saveToYAML saves the given data to a YAML file
func saveToYAML(filename string, token string) error {
	data := Token{Token: token}
	bytes, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	log.Println("Token saved to", filename)
	return nil
}

// readFromYAML reads data from a YAML file and unmarshals it into the given struct
func readFromYAML(filename string) (string, error) {
	data := Token{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if err := yaml.Unmarshal(bytes, &data); err != nil {
		return "", err
	}
	return data.Token, nil
}


func GetFolderAndConfig()( string, string) {
	var userConfigDir, _ = os.UserConfigDir()
	var folderName= filepath.Join(userConfigDir, configFolder)
	var fileName = filepath.Join(folderName,  configFile)
	return folderName, fileName
}


func SaveOrGetToken() string {
	folderName, fileName := GetFolderAndConfig()

	// make folder to store token
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		log.Println("Folder does not exist, creating folder ",folderName )
		if err := os.MkdirAll(folderName, 0755); err != nil {
			panic(err)
		}
	} 
	
	// get token if file exit otherwise save the tokem
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		token := GetApiKey()
		
		// save token
		saveToYAML(fileName, token)
		return token
	} else {
		token, err := readFromYAML(fileName)
		if err != nil  {
			panic(err)
		}
		return token 
	}

}






