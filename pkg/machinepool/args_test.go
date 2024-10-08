package machinepool

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("NewMachinepoolArgsFunction test", func() {
	var (
		cmd                    *cobra.Command
		invalidMachinepoolName = "-c"
		validMachinepoolName   = "valid"
	)

	BeforeEach(func() {
		cmd = &cobra.Command{}
		cmd.Flags().String(machinepoolFlagName, "", "Machine pool of the cluster to target")
	})

	Context("When machinepool is a flag", func() {
		It("Returns an error if the machinepool identifier is invalid", func() {
			cmd.Flags().Set(machinepoolFlagName, invalidMachinepoolName)

			validateArgs := NewMachinepoolArgsFunction(true)
			err := validateArgs(cmd, []string{})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(ErrInvalidMachinePoolIdentifier.Error()))
		})

		It("Does not return an error for a valid machinepool identifier", func() {
			cmd.Flags().Set(machinepoolFlagName, validMachinepoolName)

			validateArgs := NewMachinepoolArgsFunction(true)
			err := validateArgs(cmd, []string{})

			Expect(err).NotTo(HaveOccurred())
		})

		It("Does not return an error if machinepool flag is not set", func() {
			validateArgs := NewMachinepoolArgsFunction(true)
			err := validateArgs(cmd, []string{})

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("When machinepool is an argument", func() {
		It("Returns an error if the machinepool identifier is missing", func() {
			validateArgs := NewMachinepoolArgsFunction(false)
			err := validateArgs(cmd, []string{})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(ErrMissingMachinePoolIdentifier.Error()))
		})

		It("Returns an error if the machinepool identifier is invalid", func() {
			cmd.Flags().Set(machinepoolFlagName, invalidMachinepoolName)

			validateArgs := NewMachinepoolArgsFunction(false)
			err := validateArgs(cmd, []string{invalidMachinepoolName})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(ErrInvalidMachinePoolIdentifier.Error()))
		})

		It("Does not return an error for a valid machinepool identifier", func() {
			validateArgs := NewMachinepoolArgsFunction(false)
			err := validateArgs(cmd, []string{validMachinepoolName})

			Expect(err).NotTo(HaveOccurred())
		})
	})

})
