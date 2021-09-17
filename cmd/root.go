package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	encode bool
	decode bool
)

var rootCmd = &cobra.Command{
	Use:   "b64",
	Short: "Encode string to base64 or decode base64 to string",
	Example: `	encode: b64 -e "hello guys"
	decode: b64 -d aGVsbG8gZ3V5cw==`,
	Run: func(cmd *cobra.Command, args []string) {
		if encode {
			str := base64.StdEncoding.EncodeToString([]byte(args[0]))

			fmt.Print(str)
		} else if decode {
			str, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Print(string(str))
		} else {
			fmt.Printf("%s\n\n", cmd.Short)
			fmt.Print(cmd.UsageString())
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&encode, "encode", "e", false, "Encode a string to Base64")
	rootCmd.Flags().BoolVarP(&decode, "decode", "d", false, "Decode Base64 to string")
}
