package ipfs

import "os"

func GetDATAFromIPFSCID(ipfsC *IPFSClient, cid string) (D [][]byte, err error) {
	data, err := ipfsC.GetDataFromCID(cid)
	if err != nil {
		return nil, err
	}

	fcar, err := os.CreateTemp("", cid+"Car")

	if err != nil {
		return nil, err
	}

	defer os.Remove(fcar.Name())

	_, err = fcar.Write(data)
	if err != nil {
		return nil, err
	}
	fcar.Close()

	wasmdname, err := os.MkdirTemp("", cid+"CarExtractOutputDir")
	if err != nil {
		return nil, err
	}

	defer os.RemoveAll(wasmdname)

	_, err = ExtractCarFile(fcar.Name(), wasmdname)

	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(wasmdname)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			//Get the wasm directory name and splice it
			wasmPath := wasmdname + "/" + e.Name()

			//Read the content of wasm file as bytes
			wasmdata, err := os.ReadFile(wasmPath)
			if err != nil {
				return nil, err
			}
			D = append(D, wasmdata)
		}

	}

	return
}
