package main

import (
	"fmt"
	"github.com/gitchander/heuristic/math/crygo"
)

func main() {

	if err := BlockExample(); err != nil {
		fmt.Println(err.Error())
	}

	if err := StreamExample(); err != nil {
		fmt.Println(err.Error())
	}
}

func BlockExample() error {

	key := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	}
	fmt.Printf("key: [ % x ]\n", key)

	table := crygo.NewTableDefault()

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	b1 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	b2 := make([]byte, crygo.BlockSize)
	b3 := make([]byte, crygo.BlockSize)

	blockCipher.Encrypt(b2, b1)
	fmt.Printf("[ % x ]\n", b2)

	blockCipher.Decrypt(b3, b2)
	fmt.Printf("[ % x ]\n", b3)

	return nil
}

func StreamExample() error {

	key := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	}

	syn := []byte{
		0xF1, 0x09, 0xAC, 0x11, 0x73, 0xB8, 0x04, 0x13,
	}

	table := crygo.NewTableDefault()

	blockCipher, err := crygo.NewBlockCipher(table, key)
	if err != nil {
		return err
	}

	streamCipher, err := crygo.NewStreamCipher(blockCipher, syn)
	if err != nil {
		return err
	}

	b1 := []byte("中國是中國，台灣和新加坡的官方語言。世界各地的講它超過13十億人")
	b2 := make([]byte, len(b1))
	b3 := make([]byte, len(b1))

	streamCipher.XORKeyStream(b2, b1)

	fmt.Printf("b1: [ % x ]\n", b1)
	fmt.Printf("b2: [ % x ]\n", b2)

	streamCipher, err = crygo.NewStreamCipher(blockCipher, syn)
	if err != nil {
		return err
	}

	streamCipher.XORKeyStream(b3[:5], b2[:5])
	streamCipher.XORKeyStream(b3[5:9], b2[5:9])
	streamCipher.XORKeyStream(b3[9:17], b2[9:17])
	streamCipher.XORKeyStream(b3[17:], b2[17:])

	fmt.Printf("b3: [ % x ]\n", b3)
	fmt.Println(string(b3))

	return nil
}
