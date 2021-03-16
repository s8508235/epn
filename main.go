package main

import (
	"errors"
	"io"
	"os"
	"regexp"
	"time"

	"bufio"
	"encoding/csv"

	"gitlab.com/mi-dmp/epn/pkg/entity"
	epnlog "gitlab.com/mi-dmp/epn/pkg/log"
	"gitlab.com/mi-dmp/epn/pkg/runnable"

	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:   appName,
	Short: appName + " 是個將輸入檔案每行輸出成 sha256 存成 csv 之執行檔",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{}

		worker, err := mermaid.NewMermaid(
			cmd,
			viper.New(),
			epnlog.NewLogger(),
			"EPN",
		)
		if err != nil {
			return err
		}
		return worker.Execute(runnable.EncryptPhoneNumber, initializers...)

	},
}

var check = &cobra.Command{
	Use:   "check",
	Short: "check 是個確認輸出檔案每行輸出為正確 sha256 格式之指令",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{}
		runnable := func(
			logger *log.Logger,
			args entity.CheckFileArgs,
		) error {
			st := time.Now()
			checkHandle, err := os.Open(args.Input)
			if err != nil {
				return err
			}
			reader := csv.NewReader(bufio.NewReader(checkHandle))
			re := regexp.MustCompile(`^[0-9a-f]{64}$`)
			for {
				record, err := reader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return err
				}
				if len(record) == 1 {
					if !re.Match([]byte(record[0])) {
						return errors.New("檔案格式有誤")
					}
				} else {
					return errors.New("請檢查輸入檔案是否只有一個欄位")
				}
			}
			logger.Debugf("花了 %f 秒", time.Since(st).Seconds())
			logger.Infoln("檔案沒有問題")
			return nil
		}
		worker, err := mermaid.NewMermaid(
			cmd,
			viper.New(),
			epnlog.NewLogger(),
			"EPN",
		)
		if err != nil {
			return err
		}
		return worker.Execute(runnable, initializers...)

	},
}

func init() {
	cmd.AddCommand(check)
	AddFileFlag(cmd)
	check.Flags().StringP("check_file", "c", "./output.csv", "The CSV file will be checked its format is SHA256.")
	cmd.PersistentFlags().String("log_level", "info", "")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
