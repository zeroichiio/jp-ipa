package data

import(
	"os"
	"time"
)


func dictIpadictApBytes() ([]byte, error) {
	return _dictIpadictAp, nil
}

func dictIpadictAp() (*asset, error) {
	bytes, err := dictIpadictApBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ap", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
