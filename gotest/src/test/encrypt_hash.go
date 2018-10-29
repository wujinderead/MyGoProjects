package test

import (
	"os"
	"errors"
	"hash"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os/exec"
	"crypto/sha1"
	"crypto/md5"
)

const (
	size224 = 224
	size256 = 256
	size384 = 384
	size512 = 512
	size1 = 160
	md5size = 128
)

func ShaSumFromFile(filename string, size int) ([]byte, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("open file error")
	}
	defer reader.Close()
	//get hasher
	var hasher = func(size int) hash.Hash {
		switch size {
		case size224:
			return sha256.New224()
		case size256:
			return sha256.New()
		case size384:
			return sha512.New384()
		case size512:
			return sha512.New()
		case 1:
			return sha1.New()
		default:
			return nil
		}
	}(size)
	if hasher == nil {
		return nil, errors.New(fmt.Sprintf("unsupported sha size: %d", size))
	}
	//get data and feed to hasher
	data := make([]byte, 500)  // read 500 bytes every time
	for {
		number, err := reader.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, errors.New("error when read")
			}
		}
		hasher.Write(data[:number])
	}
	return hasher.Sum([]byte{}), nil
}

func ShaSumFromCmd(filename string, size int) (string, error) {
	var cmd *exec.Cmd
	switch size {
	case size224:
	case size256:
	case size384:
	case size512:
	case 1:
	default:
		return "", errors.New(fmt.Sprintf("unsupported sha size: %d", size))
	}
	cmd = exec.Command(fmt.Sprintf("sha%dsum", size), filename)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", errors.New("get stdout error")
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		return "", errors.New("execution error")
	}
	// 读取输出结果
	if size == 1 {
		size = size1
	}
	result := make([]byte, size/4)
	_, err = stdout.Read(result)
	if err != nil {
		return "", errors.New("read result error")
	}
	return string(result), nil
}

func Md5SumFromFile(filename string) ([]byte, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("open file error")
	}
	defer reader.Close()
	//get hasher
	var hasher = md5.New()
	//get data and feed to hasher
	data := make([]byte, 500)  // read 500 bytes every time
	for {
		number, err := reader.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, errors.New("error when read")
			}
		}
		hasher.Write(data[:number])
	}
	return hasher.Sum([]byte{}), nil
}

func Md5SumFromCmd(filename string) (string, error) {
	var cmd *exec.Cmd
	cmd = exec.Command("md5sum", filename)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", errors.New("get stdout error")
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		return "", errors.New("execution error")
	}
	// 读取输出结果
	size := md5size
	result := make([]byte, size/4)
	_, err = stdout.Read(result)
	if err != nil {
		return "", errors.New("read result error")
	}
	return string(result), nil
}
