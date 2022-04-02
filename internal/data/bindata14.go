package data

import(
	"os"
	"time"
)


func dictIpadictAoBytes() ([]byte, error) {
	return _dictIpadictAo, nil
}

func dictIpadictAo() (*asset, error) {
	bytes, err := dictIpadictAoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ao", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
