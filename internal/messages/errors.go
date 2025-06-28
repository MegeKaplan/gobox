package messages

const (
	ErrNoGoModFound          = "No go.mod file found in the current directory. Would you like to create one?"
	ErrOperationAborted      = "Operation aborted by the user."
	ErrPackageInstallFailed  = "Failed to install the package: %s\n"
	ErrPackageSaveFailed     = "Failed to save the package information: %s"
	ErrProjectAlreadyExists  = "Project already exists. Please choose a different name or remove the existing directory."
	ErrPromptFailed          = "Prompt failed. Please try again."
	ErrGoModInitFailed       = "Failed to initialize go.mod file."
	ErrLoadingPackagesFailed = "Failed to load packages. Please try again."
	ErrRemovePackageFailed   = "Failed to remove the package: %s\n"
)
