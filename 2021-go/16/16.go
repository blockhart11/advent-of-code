package _16

import (
	"fmt"
	"math"
	"strings"
)

func hexToBinMap() map[string]string {
	return map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
}

func hexToBinStr(hex string) string {
	var result string
	hexMap := hexToBinMap()
	for _, char := range hex {
		result += hexMap[string(char)]
	}
	return result
}

// Lifted from day 03
func binToDec(in string) int {
	l := len(in)
	result := 0
	for i, v := range in {
		if v == '1' {
			exp := float64(l - i - 1)
			result += int(math.Pow(2, exp))
		}
	}
	return result
}

type packet struct {
	version         int
	typeId          int
	lengthTypeId    int
	subPacketLength int
	subPacketCount  int
	literal         int
	subpackets      []packet
}

const noLimit int = -1

func processInput(inputHex string) (int, int) {
	var versionCount, evalResult int
	fmt.Printf("\n\n*** DECODING MESSAGE ***\n\n")
	packets, _ := decode(hexToBinStr(inputHex), noLimit)
	for _, p := range packets {
		versionCount += versionSum(p)
	}
	evalResult = eval(packets[0])
	fmt.Printf("Sum of versions: %d\n", versionCount)
	fmt.Printf("Evaluation of result: %d\n\n", evalResult)
	return versionCount, evalResult
}

func versionSum(p packet) int {
	result := p.version
	for _, subPacket := range p.subpackets {
		result += versionSum(subPacket)
	}
	return result
}

func read(input string, length int, reader *int) string {
	result := input[*reader : *reader+length]
	*reader += length
	return result
}

// Could probably use a "reader" instead. Learn how, maybe?
func decode(inputBin string, limit int) ([]packet, int) {
	readerIdx := 0
	if !strings.Contains(inputBin, "1") {
		// only zeros left, EOF
		return nil, readerIdx
	}

	var result []packet
	var count int

	for {
		fmt.Println("Decoding next packet...")
		pNew := packet{}

		// version
		pNew.version = binToDec(read(inputBin, 3, &readerIdx))
		fmt.Println("Version", pNew.version)

		// type id
		pNew.typeId = binToDec(read(inputBin, 3, &readerIdx))
		fmt.Println("Type ID", pNew.typeId)

		if pNew.typeId == 4 {
			// read literal values
			var literalBin string
			for {
				// leading bit
				var final bool
				switch leadBit := read(inputBin, 1, &readerIdx); leadBit {
				case "0":
					final = true
				}
				literalBin += read(inputBin, 4, &readerIdx)
				if final {
					break
				}
			}
			pNew.literal = binToDec(literalBin)
			fmt.Println("Literal", pNew.literal)
		} else {
			// operator
			pNew.literal = -1
			pNew.lengthTypeId = binToDec(read(inputBin, 1, &readerIdx))
			switch pNew.lengthTypeId {
			case 0:
				// 15 bits saying how long the subpackets are
				pNew.subPacketLength = binToDec(read(inputBin, 15, &readerIdx))
				fmt.Printf("Decoding operator packet with %d subpacket bits...\n", pNew.subPacketLength)
				pNew.subpackets, _ = decode(read(inputBin, pNew.subPacketLength, &readerIdx), noLimit)
			case 1:
				// 11 bits saying how many sub-packets there are
				pNew.subPacketCount = binToDec(read(inputBin, 11, &readerIdx))
				fmt.Printf("Decoding operator packet with %d subpackets...\n", pNew.subPacketCount)

				var readCount int
				pNew.subpackets, readCount = decode(inputBin[readerIdx:], pNew.subPacketCount)
				readerIdx += readCount
			default:
				// error
				fmt.Errorf("operator packet length type ID isn't 0 or 1")
				return nil, readerIdx
			}
		}
		result = append(result, pNew)

		// reached packet limit?
		count++
		if count == limit {
			fmt.Println("Hit subpacket limit:", count)
			break
		}
		// End of file?
		if !strings.Contains(inputBin[readerIdx:], "1") {
			// only zeros left, EOF
			fmt.Println("Reached EOF at position", readerIdx)
			break
		}
	}

	return result, readerIdx
}

func eval(p packet) int {
	var result int
	switch p.typeId {
	case 0:
		// sum
		for _, sub := range p.subpackets {
			result += eval(sub)
		}
	case 1:
		// product
		result = 1
		for _, sub := range p.subpackets {
			result *= eval(sub)
		}
	case 2:
		// max
		min := -1
		for _, sub := range p.subpackets {
			next := eval(sub)
			if next < min || min == -1 {
				min = next
			}
		}
		result = min
	case 3:
		// max
		max := -1
		for _, sub := range p.subpackets {
			next := eval(sub)
			if next > max || max == -1 {
				max = next
			}
		}
		result = max
	case 4:
		// literal value
		result = p.literal
	case 5:
		// gt
		if eval(p.subpackets[0]) > eval(p.subpackets[1]) {
			result = 1
		} else {
			result = 0
		}
	case 6:
		// lt
		if eval(p.subpackets[0]) < eval(p.subpackets[1]) {
			result = 1
		} else {
			result = 0
		}
	case 7:
		// equal (return 1 if true, 0 otherwise. always 2 subpackets)
		if eval(p.subpackets[0]) == eval(p.subpackets[1]) {
			result = 1
		} else {
			result = 0
		}
	default:
		fmt.Errorf("unknown type ID")
		return -1
	}
	return result
}
