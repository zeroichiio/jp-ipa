package data

import(
	"os"
	"time"
)


func dictIpadictAsBytes() ([]byte, error) {
	return _dictIpadictAs, nil
}

func dictIpadictAs() (*asset, error) {
	bytes, err := dictIpadictAsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.as", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
