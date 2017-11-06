package facade

import (
	"crypto"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/hash"
)

var (
	//Name is applicatin name
	Name = "hash"
	//Version is version number of application
	Version string
	//OS is OS name
	OS string
	//Arch is architecture name
	Arch string
)

var (
	cui        = gocli.NewUI()
	exitCode   = ExitNormal
	defaultAlg = hash.AlgoString(crypto.SHA256)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: Name + " [flags] [binary file]",
	RunE: func(cmd *cobra.Command, args []string) error {
		versionFlag, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if versionFlag {
			cui.OutputErr(Name)
			if len(Version) > 0 {
				cui.OutputErr(fmt.Sprintf(" v%s", Version))
			}
			if len(OS) > 0 && len(Arch) > 0 {
				cui.OutputErr(fmt.Sprintf(" (%s/%s)", OS, Arch))
			}
			cui.OutputErrln()
			return nil
		}
		listFlag, err := cmd.Flags().GetBool("list")
		if err != nil {
			return err
		}
		if listFlag {
			cui.Outputln(hash.FuncList())
			return nil
		}
		name, err := cmd.Flags().GetString("algo")
		if err != nil {
			return err
		}
		alg, err := hash.Algorithm(name)
		if err != nil {
			return err
		}
		compare, err := cmd.Flags().GetString("compare")
		if err != nil {
			return err
		}
		reader := cui.Reader()
		if len(args) > 0 {
			file, err2 := os.Open(args[0]) //args[0] is maybe file path
			if err2 != nil {
				return err2
			}
			defer file.Close()
			reader = file
		}

		v, err := hash.Value(reader, alg)
		if err != nil {
			return err
		}
		result := fmt.Sprintf("%x", v)

		if compare != "" {
			if strings.ToLower(compare) == result {
				cui.OutputErrln("matched")
				exitCode = ExitNormal
			} else {
				cui.OutputErrln("unmatched")
				exitCode = ExitAbnormal
			}
		} else {
			cui.Outputln(result)
			exitCode = ExitNormal
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ui *gocli.UI, args []string) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = ExitAbnormal
		}
	}()

	//execution
	cui = ui
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	if err := rootCmd.Execute(); err != nil {
		exit = ExitAbnormal
	} else {
		exit = exitCode
	}
	return
}

func init() {
	rootCmd.Flags().StringP("algo", "a", defaultAlg, "hash algorithm")
	rootCmd.Flags().StringP("compare", "c", "", "compare hash value")
	rootCmd.Flags().BoolP("list", "l", false, "listing hash functions")
	rootCmd.Flags().BoolP("version", "v", false, "version of "+Name)
}
