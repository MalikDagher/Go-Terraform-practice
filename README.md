# Go & Terraform Learning Playground

Two starter projects. **Open a NEW terminal first** so the freshly installed
`go` and `terraform` commands are on your PATH.

## Go (`go-playground/`)

```powershell
cd C:\Users\malik\learning\go-playground
go mod init playground   # one time: creates go.mod
go run .                 # build + run main.go
```

Other useful commands:
- `go build` — compile to an `.exe`
- `go fmt ./...` — auto-format your code
- `go test ./...` — run tests

## Terraform (`terraform-playground/`)

```powershell
cd C:\Users\malik\learning\terraform-playground
terraform init      # one time: downloads the "local" provider
terraform plan      # preview what will change
terraform apply     # create hello.txt (type "yes" to confirm)
terraform destroy   # remove what you created
```

Try overriding the variable:
```powershell
terraform apply -var="greeting=Terraform is fun"
```

## Recommended VS Code extensions
- **Go** (golang.go) — IntelliSense, debugging, formatting
- **HashiCorp Terraform** (hashicorp.terraform) — syntax + the language server
