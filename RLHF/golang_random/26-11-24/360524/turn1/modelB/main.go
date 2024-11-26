
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// UserRole defines the different roles for users
type UserRole string

const (
	// AdminRole represents the "admin" user role
	AdminRole UserRole = "admin"
	// UserRole represents the "user" user role
	UserRole UserRole = "user"
)

// UserFlags represents the command-line flags available to users
type UserFlags struct {
	Name string
	Role UserRole
}

// AdminFlags represents the additional command-line flags available to admins
type AdminFlags struct {
	Database string
}

var (
	// userFlags contains the user-specific flags
	userFlags UserFlags
	// adminFlags contains the admin-specific flags
	adminFlags AdminFlags
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "My simple application",
		Run:   runFunc,
	}

	// Initialize the flags
	initUserFlags(rootCmd)
	initAdminFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// initUserFlags initializes the user-specific flags
func initUserFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&userFlags.Name, "name", "n", "", "User name")
	cmd.Flags().VarP((*userRoleFlag)(&userFlags.Role), "role", "r", "User role (admin or user)")
}

// initAdminFlags initializes the admin-specific flags
func initAdminFlags(cmd *cobra.Command) {
	adminCmd := &cobra.Command{
		Use:   "admin",
		Short: "Admin commands",
		Run:   adminRunFunc,
	}
	adminCmd.Flags().StringVarP(&adminFlags.Database, "database", "d", "", "Database name")

	cmd.AddCommand(adminCmd)
}

// userRoleFlag is a custom flag type for validating user roles
type userRoleFlag UserRole

func (f *userRoleFlag) String() string {
	return string(*f)
}

func (f *userRoleFlag) Set(s string) error {
	// Validate the role flag
	switch s {
	case string(AdminRole), string(UserRole):
		*f = userRoleFlag(s)
		return nil
	default:
		return fmt.Errorf("invalid role: %s, must be 'admin' or 'user'", s)
	}
}

func (f *userRoleFlag) Type() string {
	return "userRole"
}

// runFunc is the main function called when the root command is executed
func runFunc(cmd *cobra.Command, args []string) {
	if err := validateUserFlags(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Name: %s\nUser Role: %s\n", userFlags.Name, userFlags.Role)
}

// adminRunFunc is the function called when the admin command is executed
func adminRunFunc(cmd *cobra.Command, args []string) {
	if err := validateAdminFlags(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Admin Database: %s\n", adminFlags.Database)
}

func validateUserFlags() error {
	// Add any specific validation logic for user flags here
	if userFlags.Name == "" {
		return fmt.Errorf("user name is required")
	}
	if userFlags.Role == "" {
		return fmt.Errorf("user role is required")
	}
	return nil
}