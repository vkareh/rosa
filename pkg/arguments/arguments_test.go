package arguments

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("Client", func() {
	var (
		cmd      *cobra.Command
		childCmd *cobra.Command
	)

	Context("Region deprecation test", func() {
		BeforeEach(func() {
			cmd = &cobra.Command{
				Use:   "test",
				Short: "Test command used for testing deprecation",
				Long: "This command is used for testing the deprecation of the 'region' flag in " +
					"arguments.go - it is used for nothing else.",
			}
			childCmd = &cobra.Command{
				Use:   "child",
				Short: "Child command used for testing deprecation",
				Long: "This child command is used for testing the deprecation of the 'region' flag in " +
					"arguments.go - it is used for nothing else.",
			}
			cmd.AddCommand(childCmd)

			AddRegionFlag(cmd.PersistentFlags())
			AddDebugFlag(cmd.PersistentFlags())
		})
		It("Test deprecation of region flag", func() {
			MarkRegionDeprecated(cmd, []*cobra.Command{childCmd})
			regionFlag := cmd.PersistentFlags().Lookup("region")
			debugFlag := cmd.PersistentFlags().Lookup("debug")
			Expect(regionFlag.Deprecated).To(Equal(regionDeprecationMessage))
			Expect(debugFlag.Deprecated).To(Equal(""))
		})
	})

	Context("Test PreprocessUnknownFlagsWithId func", func() {
		BeforeEach(func() {
			cmd = &cobra.Command{
				Use:   "test",
				Short: "Test command used for testing non-positional args",
				Long: "This test command is being used specifically for testing non-positional args, " +
					"so we do not confuse users with hard rules for where, for example, the ID in " +
					"`rosa edit addon ID` must be. For example, we want to be able to do `rosa edit addon " +
					"-c test <ADDON_ID>` as well as `rosa edit addon <ADDON_ID> -c test`.",
				Args: func(cmd *cobra.Command, argv []string) error {

					return nil
				},
			}
			cmd.Flags().BoolP("help", "h", false, "")
			s := ""
			cmd.Flags().StringVarP(
				&s,
				"cluster",
				"c",
				"",
				"Name or ID of the cluster.",
			)
		})
		It("Returns without error", func() {
			err := PreprocessUnknownFlagsWithId(cmd, []string{"test", "-c", "test-cluster"})
			Expect(err).ToNot(HaveOccurred())
		})
		It("Returns error with no ID", func() {
			err := PreprocessUnknownFlagsWithId(cmd, []string{"-c", "test-cluster"})
			Expect(err).To(HaveOccurred())
			Expect(fmt.Sprint(err)).To(Equal("ID argument not found in list of arguments passed to command"))
		})
		It("Returns error with flag that has no value", func() {
			err := PreprocessUnknownFlagsWithId(cmd, []string{"test", "-c", "-c", "-c", "-c"})
			Expect(err).To(HaveOccurred())
			Expect(fmt.Sprint(err)).To(Equal("No value given for flag '-c'"))
		})
	})
})
