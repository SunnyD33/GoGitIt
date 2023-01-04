module GoGitIt/ggi

go 1.19

replace GoGitIt/pkg/utils => ../pkg/utils

replace GoGitIt/internal/auth => ../internal/auth

replace GoGitIt/internal/repos => ../internal/repos

replace GoGitIt/internal/search => ../internal/search

replace GoGitIt/internal/rate => ../internal/rate

replace GoItIt/internal/open => ../internal/open

replace GoItIt/internal/help => ../internal/help

require (
	GoGitIt/internal/auth v0.0.0-00010101000000-000000000000
	GoGitIt/internal/rate v0.0.0-00010101000000-000000000000
	GoGitIt/internal/repos v0.0.0-00010101000000-000000000000
	GoGitIt/internal/search v0.0.0-00010101000000-000000000000
	GoGitIt/pkg/utils v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
