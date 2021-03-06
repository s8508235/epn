package runnable

import (
	"bufio"
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/s8508235/epn/pkg/entity"
	log "github.com/sirupsen/logrus"
)

func EncryptPhoneNumber(
	logger *log.Logger,
	args entity.PhoneFileArgs,
) error {
	st := time.Now()
	if args.Input == args.Output {
		return errors.New("please let input and output file different")
	}
	inputHandle, err := os.Open(args.Input)
	if err != nil {
		return err
	}
	defer inputHandle.Close()
	outputHandle, err := os.Create(args.Output)
	if err != nil {
		return err
	}
	defer outputHandle.Close()
	reader := csv.NewReader(bufio.NewReader(inputHandle))
	reader.Comma = ','
	writer := csv.NewWriter(outputHandle)
	defer writer.Flush()

	re := regexp.MustCompile(`^09[0-9]{8}$`)
	failed := 0
	success := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) == 1 {
			b := []byte(record[0])
			if !re.Match(b) {
				logger.Debugln(fmt.Sprintf("%s is not a phone number", record[0]))
				failed++
				continue
			}
			sum := sha256.Sum256(b)
			row := []string{fmt.Sprintf("%x", sum)}
			logger.Debugf("-- %s -- %s", record, row)
			if err := writer.Write(row); err != nil {
				return err
			}
			success++
		} else {
			return errors.New("please make sure there is only one column in input file.")
		}
	}
	logger.Debugf("took %f seconds", time.Since(st).Seconds())
	logger.Infoln("file output completed", success, "success numbers", failed, "failed numbers")
	return nil
}
