package executioner

import "fmt"

// ErrCommandNotFound defines the error when a command is not found
type ErrCommandNotFound struct{}

func (e *ErrCommandNotFound) Error() string {
	return fmt.Sprintf("Não encontrei esse comando.")
}

// ErrNoPermission defines the error when a user does not have sufficient permission
type ErrNoPermission struct{}

func (e *ErrNoPermission) Error() string {
	return fmt.Sprintf("Você não tem permissão pra usar esse comando.")
}

// ErrValuesOutOfBounds defines the error when a user gives more values to an arg then required
type ErrValuesOutOfBounds struct {
	Arg string
}

func (e *ErrValuesOutOfBounds) Error() string {
	return fmt.Sprintf("Valores pro argumento '%s' estão fora dos limites.", e.Arg)
}

// ErrMissingRequiredArgs defines the error when a user fails to provide all required args for a command
type ErrMissingRequiredArgs struct {
	Args []string
}

func (e *ErrMissingRequiredArgs) Error() string {
	return fmt.Sprintf("Você esqueceu de passar argumentos obrigatórios: %s", e.Args)
}
