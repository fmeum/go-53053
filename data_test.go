package data_test

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestDataDep(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	os.RemoveAll(tmp)
	require.NoError(t, err)
	testdata, err := filepath.Abs("testdata")
	require.NoError(t, err)
	os.Symlink(testdata, tmp)
	cmd := exec.Command("bash", "-c", "cat "+filepath.Join(tmp, "subdir", "foo"))
	out, err := cmd.Output()
	require.NoError(t, err)
	require.Equal(t, "old_content\n", string(out))
}
