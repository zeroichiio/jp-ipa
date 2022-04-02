package data

import(
	"os"
	"time"
)


func dictIpadictAuBytes() ([]byte, error) {
	return _dictIpadictAu, nil
}

func dictIpadictAu() (*asset, error) {
	bytes, err := dictIpadictAuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.au", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
