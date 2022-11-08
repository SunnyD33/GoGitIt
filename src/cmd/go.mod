module GoGitIt/cmd

go 1.19

replace GoGitIt/pkg/utils => ../pkg/utils

replace GoGitIt/internal/auth => ../internal/auth

require (
	GoGitIt/internal/auth v0.0.0-00010101000000-000000000000
	GoGitIt/pkg/utils v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.4.0 // indirect
