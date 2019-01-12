package flaxid

import (
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Start at January 1st, 2015. Same as the reference python implementation
const EpochStart = 1420070400000
const TotalBits = 96
const TimestampBits = 40
const RandomBits = TotalBits - TimestampBits

// Modified Base 64 alphabet that preserves lexicographical ordering
const Base64Alphabet = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"_abcdefghijklmnopqrstuvwxyz"

func zeroPad(str string, length int) string {
	return strings.Repeat("0", length-len(str)) + str
}

func base64LexEncode(num *big.Int) string {
	bnum := zeroPad(num.Text(2), TotalBits)
	s := ""
	for i := 0; i <= TotalBits-6; i += 6 {
		x, _ := strconv.ParseInt(bnum[i:i+6], 2, 64)
		s += string(Base64Alphabet[x])
	}
	return s
}

func getFlaxIDNum(timestamp time.Time) *big.Int {
	ms := big.NewInt((timestamp.UnixNano() / 1e6) - EpochStart)
	ms.Lsh(ms, RandomBits)

	randomNum := big.NewInt(0)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum.Rand(rnd, big.NewInt(int64(math.Pow(2, RandomBits))))

	return ms.Add(ms, randomNum)
}

func New() string {
	return base64LexEncode(getFlaxIDNum(time.Now()))
}

func ForTimestamp(timestamp time.Time) string {
	return base64LexEncode(getFlaxIDNum(timestamp))
}
