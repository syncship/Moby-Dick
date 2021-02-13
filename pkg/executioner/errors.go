package executioner

import "fmt"

// ErrCommandNotFound defines the error when a command is not found
type ErrCommandNotFound struct{}

func (e *ErrCommandNotFound) Error() string {
	return fmt.Sprintf("Command not found")
}

// ErrNoPermission defines the error when a user does not have sufficient permission
type ErrNoPermission struct{}

func (e *ErrNoPermission) Error() string {
	return fmt.Sprintf("You can't run this command")
}

// ErrValuesOutOfBounds defines the error when a user gives more values to an arg then required
type ErrValuesOutOfBounds struct {
	Arg string
}

func (e *ErrValuesOutOfBounds) Error() string {
	return fmt.Sprintf("Values for argument '%s' are out of bounds", e.Arg)
}

// ErrMissingRequiredArgs defines the error when a user fails to provide all required args for a command
type ErrMissingRequiredArgs struct {
	Args []string
}

func (e *ErrMissingRequiredArgs) Error() string {
	return fmt.Sprintf("You're missing required args: %s", e.Args)
}
