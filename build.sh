XC_OS=${XC_OS:-$(go env GOOS)}
XC_ARCH=${XC_ARCH:-$(go env GOARCH)}

gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "build/{{.OS}}_{{.Arch}}/packer-{{.Dir}}" \
    ./...
