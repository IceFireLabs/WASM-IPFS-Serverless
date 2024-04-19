package confer

import (
	"errors"
	"fmt"

	utils "github.com/BlockCraftsman/WASM-IPFS-Serverless/utils"
	"gopkg.in/yaml.v2"
)

// parseYamlFromBytes parses YAML data from bytes.
func parseYamlFromBytes(originData []byte) (data confS, err error) {
	if len(originData) == 0 {
		err = errors.New("yaml source data is empty")
		return
	}

	err = yaml.Unmarshal(originData, &data)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal YAML: %v", err)
		return
	}

	return
}

// parseYamlFromFile parses YAML data from a file.
func parseYamlFromFile(filePath string) (confS, error) {
	fileData, err := utils.ReadFileData(filePath)
	if err != nil {
		return confS{}, err
	}

	return parseYamlFromBytes(fileData)
}
