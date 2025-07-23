package file

// 文件

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
)

// 文件信息
type fileInfo struct {
	name string
	size int64
}

// WalkDir 递归获取目录下文件的名称和大小
func WalkDir(dirname string) ([]fileInfo, error) {
	op, err := filepath.Abs(dirname) // 获取目录的绝对路径
	if nil != err {
		return nil, err
	}
	files, err := os.ReadDir(op) // 获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		return nil, err
	}

	var fileInfos []fileInfo // 返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			fs, err := WalkDir(op + `/` + f.Name()) // 路径分隔符，linux 和 windows 不同
			if nil != err {
				return nil, err
			}
			fileInfos = append(fileInfos, fs...) // 将 slice 添加到 slice
		} else {
			info, err := f.Info()
			if nil != err {
				return nil, err
			}
			fi := fileInfo{op + `/` + f.Name(), info.Size()}
			fileInfos = append(fileInfos, fi) // slice 中添加成员
		}
	}
	return fileInfos, nil
}

// DirSize 获取目录下所有文件大小
func DirSize(dirname string) string {
	var (
		ss       int64
		files, _ = WalkDir(dirname)
	)
	for _, n := range files {
		ss += n.size
	}
	return FileSize(ss)
}

func MergeAbs(path string, fileName ...string) string {
	var paths = []string{gfile.RealPath(path)}
	paths = append(paths, fileName...)
	return gfile.Join(paths...)
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(data int64) string {
	var factor float64 = 1024
	res := float64(data)
	for _, unit := range []string{"", "K", "M", "G", "T", "P"} {
		if res < factor {
			return fmt.Sprintf("%.2f%sB", res, unit)
		}
		res /= factor
	}
	return fmt.Sprintf("%.2f%sB", res, "P")
}
