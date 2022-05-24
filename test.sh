mkdir -p testdata/subdir
echo old_content > testdata/subdir/foo
echo "Clearing test cache"
go test ./... -count 1
echo "Confirming that test is cached with no changes"
go test ./...
echo new_content > testdata/subdir/foo
echo "Running go test ./... after changing testdata/subdir/foo (should not pass)"
go test ./...
echo "Forcing a rerun (rightly fails)"
go test ./... -count 1
