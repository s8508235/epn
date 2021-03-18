package runnable

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"gitlab.com/mi-dmp/epn/pkg/entity"
	epnlog "gitlab.com/mi-dmp/epn/pkg/log"
)

var faker = rand.New(rand.NewSource(time.Now().Unix()))

func setup() {
	rand.Seed(time.Now().Unix())
}
func BenchmarkEncryptPhoneNumber_N_100000(b *testing.B) {
	setup()
	benchmarkEncryptPhoneNumber(b, 100000)
}

func BenchmarkEncryptPhoneNumber_N_1000000(b *testing.B) {
	setup()
	benchmarkEncryptPhoneNumber(b, 1000000)
}
func BenchmarkEncryptPhoneNumber_N_1500000(b *testing.B) {
	setup()
	benchmarkEncryptPhoneNumber(b, 1500000)
}

func BenchmarkEncryptPhoneNumber_N_5000000(b *testing.B) {
	setup()
	benchmarkEncryptPhoneNumber(b, 5000000)
}

func BenchmarkEncryptPhoneNumber_N_10000000(b *testing.B) {
	setup()
	benchmarkEncryptPhoneNumber(b, 10000000)
}

// Generate random ASCII digit
func randDigit(r *rand.Rand) rune {
	return rune(byte(r.Intn(10)) + '0')
}

func replaceWithNumbers(r *rand.Rand, str string) string {
	const hashtag = '#'
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(randDigit(r))
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(r.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

func GenerateRandPhoneNumber() string {
	return fmt.Sprintf("09%s", replaceWithNumbers(faker, "########"))
}

func benchmarkEncryptPhoneNumber(
	b *testing.B,
	numOfRecord int,
) {
	logger := epnlog.NewLogger()
	st := time.Now()
	f, err := os.Create("test.csv")
	if err != nil {
		b.Error(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	for i := 0; i < numOfRecord; i++ {
		pn := GenerateRandPhoneNumber()
		record := []string{pn}
		if err := writer.Write(record); err != nil {
			b.Error(err)
		}
	}
	writer.Flush()
	logger.Infof("%f sec to generate file", time.Since(st).Seconds())

	err = EncryptPhoneNumber(logger, entity.PhoneFileArgs{Input: "test.csv", Output: "test_output.csv"})
	if err != nil {
		b.Error(err)
	}
	logger.Infof("done encrypt took %f sec", time.Since(st).Seconds())
}
