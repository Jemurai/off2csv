# OFF To CSV

Push off formatted findings to CSV.

## Running

1. Get your findings into OFF format (see [github.com/owasp/off](https://github.com/owasp/off))
1. `go get github.com/jemurai/off2csv`
1. `off2csv off-file.json`

## Releasing

off2csv works to follow golang best practices.  Therefore, when updating, we need to do the following:

- `go get -u` 
- `go mod tidy`
- `git commit -m "change with version"`
- `git tag v1.0.6`
- `git push origin v1.0.6`

Run the build.sh and get the different types of artifacts and include them in the release.

`build.sh github.com/jemurai/off2csv`
