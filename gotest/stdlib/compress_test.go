package stdlib

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"testing"
)

var (
	poem = `春江潮水连海平，海上明月共潮生。
滟滟随波千万里，何处春江无月明。
江流宛转绕芳甸，月照花林皆似霰。
空里流霜不觉飞，汀上白沙看不见。
江天一色无纤尘，皎皎空中孤月轮。
江畔何人初见月，江月何年初照人。
人生代代无穷已，江月年年只相似。
不知江月待何人，但见长江送流水。
白云一片去悠悠，青枫浦上不胜愁。
谁家今夜扁舟子，何处相思明月楼。
可怜楼上月徘徊，应照离人妆镜台。
玉户帘中卷不去，捣衣砧上拂还来。
此时相望不相闻，愿逐月华流照君。
鸿雁长飞光不度，鱼龙潜跃水成文。
昨夜闲潭梦落花，可怜春半不还家。
江水流春去欲尽，江潭落月复西斜。
斜月沉沉藏海雾，碣石潇湘无限路。
不知乘月几人归，落月摇情满江树。`
	traditional = `春江潮水連海平，海上明月共潮生。
灩灩隨波千萬裏，何處春江無月明。
江流宛轉繞芳甸，月照花林皆似霰。
空裏流霜不覺飛，汀上白沙看不見。
江天一色無纖塵，皎皎空中孤月輪。
江畔何人初見月，江月何年初照人。
人生代代無窮已，江月年年只相似。
不知江月待何人，但見長江送流水。
白雲一片去悠悠，青楓浦上不勝愁。
誰家今夜扁舟子，何處相思明月樓。
可憐樓上月裴回，應照離人妝鏡臺。
玉戶簾中卷不去，擣衣砧上拂還來。
此時相望不相聞，願逐月華流照君。
鴻雁長飛光不度，魚龍潛躍水成文。
昨夜閑潭夢落花，可憐春半不還家。
江水流春去欲盡，江潭落月複西斜。
斜月沉沉藏海霧，碣石瀟湘無限路。
不知乘月幾人歸，落月搖情滿江樹。`
)

func TestZlib(t *testing.T) {
	// call writer.Flush() before call writer.Close(). otherwise it will cause EOF error when read
	{
		buffer := new(bytes.Buffer)
		writer := zlib.NewWriter(buffer)
		_, _ = writer.Write([]byte(poem))
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Flush()
		_ = writer.Close() // cause no EOF when read
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := zlib.NewReader(buffer)
		uncompressed := make([]byte, 2000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		fmt.Println("========== ==========")
		fmt.Println()
	}
	{
		buffer := new(bytes.Buffer)
		writer := zlib.NewWriter(buffer)
		_, _ = writer.Write([]byte(poem))
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Close() // cause EOF when read
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := zlib.NewReader(buffer)
		uncompressed := make([]byte, 2000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		fmt.Println("========== ==========")
		fmt.Println()
	}
	// when data was flushed separately
	{
		buffer := new(bytes.Buffer)
		writer := zlib.NewWriter(buffer)
		_, _ = writer.Write([]byte(poem))
		_ = writer.Flush()
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Close() // cause EOF when read
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := zlib.NewReader(buffer)
		uncompressed := make([]byte, 1000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))
		fmt.Println(string(uncompressed[:n]) == poem)

		n, err = reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))
		fmt.Println(string(uncompressed[:n]) == traditional)
		fmt.Println("========== ==========")
		fmt.Println()
	}
	{
		buffer := new(bytes.Buffer)
		writer := zlib.NewWriter(buffer)
		_, _ = writer.Write([]byte(poem))
		_ = writer.Flush()
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Flush()
		_ = writer.Close() // does not cause EOF when read
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := zlib.NewReader(buffer)
		uncompressed := make([]byte, 1000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))
		fmt.Println(string(uncompressed[:n]) == poem)

		n, err = reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))
		fmt.Println(string(uncompressed[:n]) == traditional)

		fmt.Println("========== ==========")
		fmt.Println()
	}
}

func TestGzip(t *testing.T) {
	{
		buffer := new(bytes.Buffer)
		writer, err := gzip.NewWriterLevel(buffer, gzip.BestCompression)
		if err != nil {
			fmt.Println("create writer err:", err.Error())
			return
		}
		_, _ = writer.Write([]byte(poem))
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Flush()
		_ = writer.Close()
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := gzip.NewReader(buffer)
		uncompressed := make([]byte, 1000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		n, err = reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		fmt.Println("========== ==========")
		fmt.Println()
	}
	{
		buffer := new(bytes.Buffer)
		writer, err := gzip.NewWriterLevel(buffer, gzip.BestSpeed)
		if err != nil {
			fmt.Println("create writer err:", err.Error())
			return
		}
		_, _ = writer.Write([]byte(poem))
		_, _ = writer.Write([]byte(traditional))
		_ = writer.Flush()
		_ = writer.Close()
		fmt.Println("uncompressed:", len([]byte(poem))+len([]byte(traditional)))
		compressed := buffer.Bytes()
		fmt.Println("compressed  :", len(compressed))

		reader, _ := gzip.NewReader(buffer)
		uncompressed := make([]byte, 1000)
		n, err := reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		n, err = reader.Read(uncompressed)
		fmt.Println("read err:", err)
		fmt.Println("uncompressed:", string(uncompressed[:n]))

		fmt.Println("========== ==========")
		fmt.Println()
	}
}
