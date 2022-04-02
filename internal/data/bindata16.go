package data

import(
	"os"
	"time"
)


func dictIpadictAqBytes() ([]byte, error) {
	return _dictIpadictAq, nil
}

func dictIpadictAq() (*asset, error) {
	bytes, err := dictIpadictAqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.aq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
