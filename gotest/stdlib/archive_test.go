package stdlib

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestZip(t *testing.T) {
	reader, err := zip.OpenReader("/usr/local/go/lib/time/zoneinfo.zip")
	defer func() {
		_ = reader.Close()
	}()
	if err != nil {
		fmt.Println("open reader err:", err.Error())
		return
	}
	for _, file := range reader.File {
		fmt.Println("file name:", file.Name)
		fmt.Println("file mthd:", file.Method)
		fmt.Println("file isdr:", file.FileInfo().IsDir())
		fmt.Println("file mode:", file.FileInfo().Mode())
		fmt.Println("file cprs:", file.CompressedSize64)
		fmt.Println("file ucpr:", file.UncompressedSize64)
		if file.Name == "Africa/Casablanca" {
			rf, err := file.Open()
			if err != nil {
				fmt.Println("open file err:", err.Error())
				_ = rf.Close()
			}
			byter, err := ioutil.ReadAll(rf)
			if err != nil {
				fmt.Println("read file err:", err.Error())
				_ = rf.Close()
			}
			fmt.Println("content:", string(byter))
			location, err := time.LoadLocationFromTZData("Morocco/Casablanca", byter)
			if err != nil {
				fmt.Println("load location err:", err.Error())
				_ = rf.Close()
			}
			fmt.Println("location:", location.String())
			_ = rf.Close()
		}
		fmt.Println()
	}
}

func TestTar(t *testing.T) {
	file, err := os.Open("/home/xzy/images/mysql.tar")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("file info err:", err.Error())
		return
	}
	header, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println("header err:", err.Error())
		return
	}
	fmt.Println("name:", header.Name)
	fmt.Println("mode:", header.Mode)
	fmt.Println("size:", header.Size)
	fmt.Println("accs:", header.AccessTime)
	fmt.Println("chgn:", header.ChangeTime)
	fmt.Println("modt:", header.ModTime)
	fmt.Println("type:", string(header.Typeflag))
	fmt.Println("majd:", header.Devmajor)
	fmt.Println("mind:", header.Devminor)
	fmt.Println("fmt :", header.Format)
	fmt.Println("gid :", header.Gid)
	fmt.Println("gnam:", header.Gname)
	fmt.Println("uid :", header.Uid)
	fmt.Println("unam:", header.Uname)
	fmt.Println("link:", header.Linkname)
	fmt.Println("pax :", header.PAXRecords)
	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("next err:", err.Error())
			break
		}
		fmt.Println("name:", header.Name)
		fmt.Println("mode:", header.Mode)
		fmt.Println("size:", header.Size)
		if strings.Contains(header.Name, "json") {
			byter, err := ioutil.ReadAll(reader)
			if err != nil {
				fmt.Println("read err:", err)
				continue
			}
			fmt.Println(header.Name, ":")
			fmt.Println(string(byter))
		}
	}
}

func TestTarWrite(t *testing.T) {
	file, err := os.OpenFile("/tmp/test.tar", os.O_WRONLY|os.O_CREATE, os.FileMode(0644))
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}
	writer := tar.NewWriter(file)
	defer func() {
		_ = writer.Close()
	}()
	_ = writer.WriteHeader(&tar.Header{Name: "poem1.txt", Mode: 0666, Size: int64(len(poem))})
	_, _ = writer.Write([]byte(poem))
	_ = writer.WriteHeader(&tar.Header{Name: "somedir", Mode: 0666, Size: 0, Typeflag: tar.TypeDir})
	_, _ = writer.Write([]byte(poem))
	_ = writer.WriteHeader(&tar.Header{Name: "somedir/poem2.txt", Mode: 0666, Size: int64(len(traditional))})
	_, _ = writer.Write([]byte(traditional))
}

func TestTarRead(t *testing.T) {
	file, err := os.Open("/tmp/test.tar")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("file info err:", err.Error())
		return
	}
	header, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println("header err:", err.Error())
		return
	}
	fmt.Println("name:", header.Name)
	fmt.Println("mode:", header.Mode)
	fmt.Println("size:", header.Size)
	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("next err:", err.Error())
			break
		}
		fmt.Println("name:", header.Name)
		fmt.Println("mode:", header.Mode)
		fmt.Println("size:", header.Size)
	}
}

func TestTarGz(t *testing.T) {
	file, err := os.Open("/tmp/ttt.tar.gz")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}
	greader, err := gzip.NewReader(file)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		fmt.Println("gzip reader err:", err.Error())
		return
	}
	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, greader)
	if err != nil {
		fmt.Println("copy err:", err)
		return
	}
	fmt.Println("copied:", n)
	fmt.Println("size:", buf.Len())
	treader := tar.NewReader(buf)
	for {
		header, err := treader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("next err:", err.Error())
			break
		}
		fmt.Println("name:", header.Name)
		fmt.Println("mode:", header.Mode, header.FileInfo().Mode(), uint32(header.FileInfo().Mode().Perm()))
		fmt.Println("size:", header.Size)
	}
}
