package ipfs

import (
	"errors"
	"os"

	"github.com/ipfs/go-cid"
	carstorage "github.com/ipld/go-car/v2/storage"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/storage"
)

func ExtractCarFile(carfilePath string, outputDir string) (extractedFiles int, err error) {
	var store storage.ReadableStorage
	var roots []cid.Cid

	carFile, err := os.Open(carfilePath)
	if err != nil {
		return 0, err
	}

	defer carFile.Close()

	store, err = carstorage.OpenReadable(carFile)
	if err != nil {
		return 0, err
	}

	roots = store.(carstorage.ReadableCar).Roots()

	ls := cidlink.DefaultLinkSystem()
	ls.TrustedStorage = true
	ls.SetReadStorage(store)

	for _, root := range roots {
		count, err := extractRoot(&ls, root, outputDir, nil)
		if err != nil {
			return 0, err
		}
		extractedFiles += count
	}

	if extractedFiles == 0 {
		return 0, errors.New("no files extracted")
	}

	return extractedFiles, nil
}
