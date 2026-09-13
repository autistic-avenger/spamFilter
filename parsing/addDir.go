package parsing

import (
	"io/fs"
	"path/filepath"
)

func AddDir(path string, mainTokenMap map[string]int, TokenCount *int) {
	
	filepath.WalkDir("./data/enron1", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		err = AddTokens(path, mainTokenMap)
		if err != nil {
			panic(err)
		}

		for tk := range mainTokenMap {
			*TokenCount += mainTokenMap[tk]
		}
		return nil
	})
}