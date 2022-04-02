package data

import(
	"os"
	"time"
)


func dictIpadictAkBytes() ([]byte, error) {
	return _dictIpadictAk, nil
}

func dictIpadictAk() (*asset, error) {
	bytes, err := dictIpadictAkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ak", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
