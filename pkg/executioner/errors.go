package executioner

import "fmt"

/* TODO: remove this comment
Porque acho que assim fica melhor em alguns casos
Apesar que parece mais "sujo" os errors ficam strong typed

1. vc pode fazer type check para verificar o tipo de erro quanto as string consts nao
	if e, ok := err.(ErrCommandNotFound); ok {
		// é command not found
	}
2. format no error: um dia mais tarde vc pode nem lembrar que tem format e usar um
	return executioner.ErrMissingRequiredArgs
*/

// ErrCommandNotFound defines the error when a command is not found
type ErrCommandNotFound struct {
	Command string
}

func (e ErrCommandNotFound) Error() string {
	return fmt.Sprintf("Command '%q' not found", e.Command)
}

// ErrNoPermission defines the error when a user does not have sufficient permission
type ErrNoPermission struct{}

func (ErrNoPermission) Error() string {
	return "You can't run this command"
}

// ErrValuesOutOfBounds defines the error when a user gives more values to an arg then required
type ErrValuesOutOfBounds struct {
	ArgName string
}

func (e ErrValuesOutOfBounds) Error() string {
	return fmt.Sprintf("Values for argument '%s' are out of bounds", e.ArgName)
}

// ErrMissingRequiredArgs defines the error when a user fails to provide all required args for a command
type ErrMissingRequiredArgs struct {
	Args []string
}

func (e ErrMissingRequiredArgs) Error() string {
	return fmt.Sprintf("You're missing required args: %s", e.Args)
}
