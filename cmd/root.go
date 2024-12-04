package cmd

import (
	"fmt"
	"log"
	"os"

	pwg "github.com/pwg/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgfile string

	rootCmd = &cobra.Command{
		Use:   "pwg",
		Short: "Password generator. Use to create cryptographically secure passwords.",
		Run:   runGenPwd,
	}
)

func init() {
	initConfig()

	rootCmd.Flags().IntVarP(&size, "size", "s", 0, "Size of new password in number of chars. Text passwords default to 4 words, random chars default to 32.")
	rootCmd.Flags().BoolVarP(&text, "text", "t", false, "Whether to create a random word passphrase.")
	rootCmd.Flags().BoolVarP(&dashes, "dashes", "d", false, "Whether to separate words by dashes with text passwords.")

	viper.BindPFlag("size", rootCmd.Flags().Lookup("size"))
	viper.BindPFlag("text", rootCmd.Flags().Lookup("text"))
	viper.BindPFlag("dashes", rootCmd.Flags().Lookup("dashes"))
}

func initConfig() {
	if cfgfile != "" {
		viper.SetConfigFile(cfgfile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func showPw(pw string) {
	fmt.Printf("New password:\n\n%s\n\n***This password will NOT be shown again. Save it before exiting!***\n\n", pw)
}

func runGenPwd(cmd *cobra.Command, args []string) {
	size, _ := cmd.Flags().GetInt("size")
	text, _ := cmd.Flags().GetBool("text")
	dashes, _ := cmd.Flags().GetBool("dashes")
	if text {
		if size == 0 {
			size = 4
		}
		pw, err := pwg.GenPhraseSecret(size, dashes)
		if err != nil {
			log.Fatal(err)
		}
		showPw(pw)
	} else {
		if size == 0 {
			size = 32
		}
		pw := pwg.GenSecret(size)
		showPw(pw)
	}
}
