package data

import(
	"os"
	"time"
)


func dictIpadictAeBytes() ([]byte, error) {
	return _dictIpadictAe, nil
}

func dictIpadictAe() (*asset, error) {
	bytes, err := dictIpadictAeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ae", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
