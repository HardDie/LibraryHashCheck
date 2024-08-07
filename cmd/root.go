/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/HardDie/ghashdeep/internal/crawler"
	"github.com/HardDie/ghashdeep/internal/validators"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "ghashdeep",
	Short: "This utility will help you easily calculate or check previously calculated hash sums of the entire library recursively with a single command",
}

func Execute(v string) {
	Version = v
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("algorithm", "a", "md5", "The hashing algorithm you prefer to use. Possible algorithms: md5, sha1, sha224, sha256, sha384, sha512, xxhash, blake3")
}

func chooseHashAlg(cmd *cobra.Command) (crawler.HashMethod, error) {
	alg, _ := cmd.Flags().GetString("algorithm")
	if alg == "" {
		alg = "md5"
	}
	switch alg {
	case "md5":
		return validators.NewMd5(), nil
	case "sha1":
		return validators.NewSha1(), nil
	case "sha224":
		return validators.NewSha224(), nil
	case "sha256":
		return validators.NewSha256(), nil
	case "sha384":
		return validators.NewSha384(), nil
	case "sha512":
		return validators.NewSha512(), nil
	case "xxhash":
		return validators.NewXxhash(), nil
	case "blake3":
		return validators.NewBlake3(), nil
	}
	return nil, fmt.Errorf("unknown flag --alg value %q", alg)
}
