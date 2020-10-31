package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// create dir if not exist
func CreatePath(path string) error {

	exist, _ := PathExists(path)
	if exist == false {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func PathLastName(path string) string {

	slist := strings.Split(path, "/")

	if len(slist) > 0 {
		return slist[len(slist)-1]
	}

	return path
}

func GetDirFilenames(path string) ([]string, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	return names, nil
}

func CheckFileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteFile(filePath string, content string) error {
	data := []byte(content)

	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return err
}

func WriteBytesToFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return err
}

func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// only remove dir's content, not remove dir itself
func RemoveDirContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}

// dstDir not exist, cp A B  --> will create B
// dstDir exit, cp A B --> A will be created in B, such as: B/A
func CopyDir(srcDir string, dstDir string) error {
	command := exec.Command("cp", "-rf", srcDir, dstDir)
	return command.Run()
}

// if dstDir exist, remove it, then cp A B  --> will create B
func CopyDirForce(srcDir string, dstDir string) error {
	RemoveDir(dstDir)
	return CopyDir(srcDir, dstDir)
}

func CopyFile(sourceFilePath string, destinationFilePath string, force bool) error {

	// check source file exist
	exist := CheckFileExist(sourceFilePath)
	if !exist {
		return errors.New("Not exist source file")
	}

	// check destination file exist
	exist = CheckFileExist(destinationFilePath)
	if exist && !force {
		return errors.New("Exist destination file")
	} else {
		// copy
		data, err := ioutil.ReadFile(sourceFilePath)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(destinationFilePath, data, 0644)
		if err != nil {
			return err
		}
	}


	return nil
}

func LinkDirSoft(srcDir string, dstDir string) error {
	// remove dstDir if it exist
	RemoveDir(dstDir)

	command := exec.Command("ln", "-sf", srcDir, dstDir)
	return command.Run()
}

func LinkFileSoft(src string, dst string) error {

	command := exec.Command("ln", "-sf", src, dst)
	return command.Run()
}

func FileSize(filePath string) int64 {
	if CheckFileExist(filePath) {
		stat, err := os.Stat(filePath)
		if err == nil {
			return stat.Size()
		}
	}
	return 0
}