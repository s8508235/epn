package runnable

import (
	"bufio"
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/mi-dmp/epn/pkg/entity"
)

func EncryptPhoneNumber(
	logger *log.Logger,
	args entity.PhoneFileArgs,
) error {
	st := time.Now()
	if args.Input == args.Output {
		return errors.New("請讓輸入檔名和輸出檔名不同")
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

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) == 1 {
			sum := sha256.Sum256([]byte(record[0]))
			row := []string{fmt.Sprintf("%x", sum)}
			logger.Debugf("-- %s -- %s", record, row)
			if err := writer.Write(row); err != nil {
				return err
			}
		} else {
			return errors.New("請檢查輸入檔案是否只有一個欄位")
		}
	}
	logger.Debugf("花了 %f 秒", time.Since(st).Seconds())
	return nil
}
