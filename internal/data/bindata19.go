package data

import(
	"os"
	"time"
)


func dictIpadictAtBytes() ([]byte, error) {
	return _dictIpadictAt, nil
}

func dictIpadictAt() (*asset, error) {
	bytes, err := dictIpadictAtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.at", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
