package stdlib

import (
	"testing"
	"fmt"
	"bufio"
	"bytes"
)

func TestBufReader(t *testing.T) {
	buffer := bytes.NewBufferString("赵客缦胡缨，吴钩霜雪明，\n" +
		"银鞍照白马，飒沓如流星。\n" +
		"十步杀一人，千里不留行，\n" +
		"事了拂衣去，深藏身与名。\n" +
		"闲过信陵饮，脱剑膝前横，\n" +
		"将炙啖朱亥，持觞劝侯嬴。\n" +
		"三杯吐然诺，五岳倒为轻，\n" +
		"眼花耳热后，意气素霓生。\n" +
		"救赵挥金槌，邯郸先震惊，\n" +
		"千秋二壮士，烜赫大梁城。\n" +
		"纵死侠骨香，不惭世上英，\n" +
		"谁能书阁下，白首太玄经。")
	reader := bufio.NewReader(buffer)

	byter := make([]byte, 6)
	reader.Read(byter)  // read some bytes to a []byte
	fmt.Println("read bytes: ", string(byter))

	str, _ := reader.ReadString('\n')
	fmt.Println("read string: ", str)   // ReadString has '\n' at end

	byter, isPrefix, _ := reader.ReadLine()
	fmt.Println("read line: ", string(byter), isPrefix) // ReadLine has not '\n' at end

	n, _ := reader.Discard(36)  // skip some bytes
	fmt.Println("skipped: ", n)

	abyte, _ := reader.ReadByte() // read a byte
	fmt.Println("read byte: ", abyte, int('\n'))

	byter, _ = reader.Peek(6)  // peek some bytes, not moving the pointer
	fmt.Println("peeked: ", string(byter))

	runer, size, _ := reader.ReadRune() // read a rune
	fmt.Println("read rune: ", string(runer), " ,size: ", size)

	byter, _ = reader.ReadBytes('\n')  // read bytes by delim
	fmt.Println("read bytes: ", string(byter))

	byter, _ = reader.ReadSlice('\n')  // return a slice point to the bytes in buffer, invalid after next read
	fmt.Println("read slice: ", string(byter))  // the bytes can be read

	_ = reader.UnreadByte()  // pointer go back one byte
	abyte, _ = reader.ReadByte()
	fmt.Println("unread byte: ", abyte)

	runer, size, _ = reader.ReadRune() // read a rune
	fmt.Println("read rune: ", string(runer), " ,size: ", size)
	_ = reader.UnreadRune()  // pointer go back a rune
	runer, size, _ = reader.ReadRune() // read a rune
	fmt.Println("unread rune: ", string(runer), " ,size: ", size)

	// the buufer size, the buffered bytes size
	fmt.Printf("size: %d, buffered: %d\n", reader.Size(), reader.Buffered())

}