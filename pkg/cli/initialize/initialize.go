package initialize

import (
	"fmt"
	"github.com/openfortra/fortra/pkg/schema"
	"github.com/openfortra/fortra/pkg/utils"
	"github.com/openfortra/fortra/internal/constants"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var (
	first, last, initial, suffix, email, ssan, phoneNumber, phoneType, passType, passOrigin, passID string
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
	userConfFile := utils.UserConfigFile()
	utils.ConfigFileDirExists()

	if _, err := os.Stat(userConfFile); !os.IsNotExist(err) {
		q := utils.CliQuestion("The user configuration file already exists. Overwrite it?")
		if q == "n" {
			return nil
		}
	}

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
	
	

	i := 0
	s := schema.SchemaInitializer()
	s.Employees[i].First = first
	s.Employees[i].Last = last
	s.Employees[i].Initial = initial
	s.Employees[i].Suffix = suffix
	s.Employees[i].Ssan = ssan
	s.Employees[i].Email = email
	j := 0
	for {
		phoneNumber = utils.CliReader("Phone Number", phoneNumber)
		phoneType = utils.CliReader("Phone Type", phoneType)
		if utils.StringInSlice(strings.ToLower(phoneType), constants.PhoneTypeList) != true {
			fmt.Printf("Invalid phone type! Valid phone types are %s:\n", constants.PhoneTypeList)
			phoneType = utils.CliReader("Phone Type", "")
		}
		s.Employees[i].Phones[j].Number = phoneNumber
		s.Employees[i].Phones[j].Type = phoneType
		q := utils.CliQuestion("Would you like to configure additional phone numbers?")
		if q == "n" {
			break
		}
		j = j + 1
		s.Employees[i].Phones = append(s.Employees[i].Phones, schema.Phone{})
		phoneNumber = ""
		phoneType = ""
	}
	j = 0
	for {
		passType = utils.CliReader("Travel Document Type", passType)
		if utils.StringInSlice(strings.ToLower(passType), constants.TravelTypeList) != true {
			fmt.Printf("Invalid travel document type! Valid document types are %s:\n", constants.TravelTypeList)
			passType = utils.CliReader("Travel Document Type", "")
		}
		passOrigin = utils.CliReader("Travel Document Origin", passOrigin)
		passID = utils.CliReader("Travel Document ID", passID)
		s.Employees[i].TravelsDocs[j].Type = passType
		s.Employees[i].TravelsDocs[j].Origin = passOrigin
		s.Employees[i].TravelsDocs[j].ID = passID
		q := utils.CliQuestion("Would you like to configure additional travel documents?")
		if q == "n" {
			break
		}
		j = j + 1
		s.Employees[i].TravelsDocs = append(s.Employees[i].TravelsDocs, schema.TravelDoc{})
		passType = ""
		passOrigin = ""
		passID = ""
	}

	utils.YamlWriter(&s, userConfFile)
	fmt.Printf("\nInitial Configuration successfully written! Happy travels...\n")

	return nil
}