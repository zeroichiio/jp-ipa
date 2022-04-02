package data

import(
	"os"
	"time"
)


func dictIpadictAmBytes() ([]byte, error) {
	return _dictIpadictAm, nil
}

func dictIpadictAm() (*asset, error) {
	bytes, err := dictIpadictAmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.am", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
