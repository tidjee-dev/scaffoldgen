package app

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	completionShell string
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate the autocompletion script for the specified shell",
	Long: `To load completions:

Bash:
  $ source <(scaffoldgen completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ scaffoldgen completion bash > /etc/bash_completion.d/scaffoldgen
  # macOS:
  $ scaffoldgen completion bash > /usr/local/etc/bash_completion.d/scaffoldgen

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ scaffoldgen completion zsh > "${fpath[1]}/_scaffoldgen"

  # You will need to start a new shell for this setup to take effect.

fish:
  $ scaffoldgen completion fish | source

  # To load completions for each session, execute once:
  $ scaffoldgen completion fish > ~/.config/fish/completions/scaffoldgen.fish

PowerShell:
  PS> scaffoldgen completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> scaffoldgen completion powershell > scaffoldgen.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
