package executioner

const (
	// ErrCommandNotFound defines the error when a command is not found
	ErrCommandNotFound = "Command not found"
	// ErrNoPermission defines the error when a user does not have sufficient permission
	ErrNoPermission = "You can't run this command"
	// ErrValuesOutOfBounds defines the error when a user gives more values to an arg then required
	ErrValuesOutOfBounds = "Values for argument '%s' are out of bounds"
	// ErrMissingRequiredArgs defines the error when a user fails to provide all required args for a command
	ErrMissingRequiredArgs = "You're missing required args: %s"
)
