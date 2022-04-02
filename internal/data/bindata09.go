package data

import(
	"os"
	"time"
)


func dictIpadictAjBytes() ([]byte, error) {
	return _dictIpadictAj, nil
}

func dictIpadictAj() (*asset, error) {
	bytes, err := dictIpadictAjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.aj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
