package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	mainPart1()
}

func mainPart1() {
	answer := run("input")
	// answer := run("input_real")
	println(answer)

}

func run(filename string) int {
	a := tj.FileToSlice(filename)

	for _, fileLine := range a {
		var l line
		l.load(fileLine)
		println("---------------------------------")
	}
	return 0
}

// func asBits(val uint64) []uint64 {
// 	bits := []uint64{}
// 	for i := 0; i < 24; i++ {
// 		bits = append([]uint64{val & 0x1}, bits...)
// 		// or
// 		// bits = append(bits, val & 0x1)
// 		// depending on the order you want
// 		val = val >> 1
// 	}
// 	return bits
// }

func asBits(bytes []byte) []bool {
	bools := []bool{}
	for _, b := range bytes {
		// println("Looking at byte: ", b)
		for i := 7; i >= 0; i-- {
			// for i := 0; i < 8; i++ {
			// println("Mask:", 0x1<<i)
			bools = append(bools, (b&(0x1<<i)) > 0)
		}
	}
	return bools
}

func bitsToInt(bits []bool, start, length int) int {
	result := 0
	for i, bit := range bits[start : start+length] {
		if bit {
			result += int(math.Pow(2, float64(length-1-i)))
		}
	}
	return result
}

type line struct {
	hex string
	v   int
	t   int
	l   int
}

func printBits(bits []bool) {
	for _, b := range bits {
		if b {
			print("1")
		} else {
			print("0")
		}
	}
	println()
}

func (l *line) load(line string) {
	l.hex = line
	bytes := []byte{}
	for i := 0; i < len(line); i += 2 {
		i, err := strconv.ParseUint(string(l.hex[i])+string(l.hex[i+1]), 16, 32)
		if err != nil {
			fmt.Printf("%s", err)
		}
		bytes = append(bytes, byte(i))
	}
	bits := asBits(bytes)
	result, _ := l.parsePacket(bits, math.MaxInt)
	println("Result: ", result[0])
}

func (l *line) parsePacket(bits []bool, packetLimit int) ([]int, int) {
	currentBit := 0
	pv := 0
	pt := 0
	totalVersion := 0
	values := []int{}
	subPackets := 0
	for {
		if subPackets >= packetLimit {
			return values, currentBit
		}
		subPackets++
		pv = bitsToInt(bits, currentBit, 3)
		totalVersion += pv
		currentBit += 3
		pt = bitsToInt(bits, currentBit, 3)
		// println("Type: ", pt)
		currentBit += 3
		switch pt {
		case 4:
			// literal value packet
			value, bitsConsumed := l.parseAsLiteral(bits[currentBit:])
			currentBit += bitsConsumed
			values = append(values, value)
		default:
			// Other types of packet
			l.l = bitsToInt(bits, currentBit, 1)
			currentBit += 1
			totalLengthInBits := 0
			numberOfSubPackets := 0
			subValues := []int{}
			bitCount := 0
			if l.l == 0 {
				totalLengthInBits = bitsToInt(bits, currentBit, 15)
				currentBit += 15
				subValues, _ = l.parsePacket(bits[currentBit:currentBit+totalLengthInBits], math.MaxInt)
				currentBit += totalLengthInBits
			} else {
				numberOfSubPackets = bitsToInt(bits, currentBit, 11)
				currentBit += 11
				subValues, bitCount = l.parsePacket(bits[currentBit:], numberOfSubPackets)
				currentBit += bitCount
			}
			result := 0
			switch pt {
			case 0:
				// sum
				for _, val := range subValues {
					result += val
				}
			case 1:
				// product
				result = 1
				for _, val := range subValues {
					result *= val
				}
			case 2:
				// minimum
				result = math.MaxInt
				for _, val := range subValues {
					if val < result {
						result = val
					}
				}
			case 3:
				// maximum
				result = math.MinInt
				for _, val := range subValues {
					if val > result {
						result = val
					}
				}
			case 5:
				// greater than
				if subValues[0] > subValues[1] {
					result = 1
				}
			case 6:
				// less than
				if subValues[0] < subValues[1] {
					result = 1
				}
			case 7:
				//equal to
				if subValues[0] == subValues[1] {
					result = 1
				}
			}
			values = append(values, result)
		}
		// println(l.hex)
		// println("version: ", l.v, "type:", l.t, "lengthid:", l.l)
		if currentBit+7 >= len(bits) {
			break
		}
	}
	// println("Total version: ", totalVersion)
	return values, currentBit
}

func (l *line) parseAsLiteral(bits []bool) (int, int) {
	combined := []bool{}
	i := 0
	for {
		chunk := bits[i+1 : i+1+4]
		combined = append(combined, chunk...)
		if !bits[i] {
			break
		}
		i += 5
	}
	return bitsToInt(combined, 0, len(combined)), i + 5
}
