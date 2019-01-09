package initialize

import (
	"fmt"
	"github.com/openfortra/fortra/pkg/schema"
	"github.com/openfortra/fortra/pkg/utils"
	"github.com/spf13/cobra"
	"io"
)

var (
	first, last, initial, suffix, email, ssan string
)

// FortraCmdInit gets all the compliance dependencies
func FortraCmdInit(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize new user configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunInit(out, cmd, args)
		},
	}
	cmd.Flags().StringVarP(&first, "first", "f", "", "First Name")
	cmd.Flags().StringVarP(&last, "last", "l", "", "Last Name")
	cmd.Flags().StringVarP(&initial, "initial", "i", "", "Middle Initial")
	cmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Suffix")
	cmd.Flags().StringVarP(&ssan, "ssan", "n", "", "Social Security Account Number")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email Address")
	cmd.Flags().SortFlags = false

	return cmd
}

// RunInit runs get when specified in cli
func RunInit(out io.Writer, cmd *cobra.Command, args []string) error {
	if cmd.Flags().NFlag() < 6 {
		fmt.Printf("\nInitializing new user. Hit [Enter] to Skip.\n")
		fmt.Printf("-----------------------\n")
	}

	first = utils.CliReader("First Name", first)
	last = utils.CliReader("Last Name", last)
	initial = utils.CliReader("Middle Initial", initial)
	suffix = utils.CliReader("Suffix", suffix)
	ssan = utils.CliReader("Social Security Account Number", ssan)
	email = utils.CliReader("Email Address", email)

	schema := schema.SchemaInitializer()
	schema.Employees[0].First = first
	schema.Employees[0].Last = last
	schema.Employees[0].Initial = initial
	schema.Employees[0].Suffix = suffix
	schema.Employees[0].Ssan = ssan
	schema.Employees[0].Email = email

	utils.YamlWriter(&schema, "test")
	fmt.Printf("\nInitial Configuration successfully written! Happy travels...\n")

	return nil
}