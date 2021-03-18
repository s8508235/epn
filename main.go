package main

import (
	"errors"
	"fmt"
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
	Short: appName + " is a executable that make input file content output to sha256 line by line",
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
	Short: "check will examine input file is a line-by-line sha256 file",
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
						return errors.New("incorrect file format")
					}
				} else {
					return errors.New("please make sure there is only one column in input file.")
				}
			}
			logger.Debugf("took %f seconds", time.Since(st).Seconds())
			logger.Infoln("file examine ok.")
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
	fmt.Printf("Version: %v\n", version)
	fmt.Printf("Date: %v\n", date)
	fmt.Printf("commit: %v\n", commit)
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
