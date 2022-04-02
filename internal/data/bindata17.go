package data

import(
	"os"
	"time"
)


func dictIpadictArBytes() ([]byte, error) {
	return _dictIpadictAr, nil
}

func dictIpadictAr() (*asset, error) {
	bytes, err := dictIpadictArBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ar", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1602332411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
