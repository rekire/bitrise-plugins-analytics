package integration

import (
	"testing"

	"github.com/bitrise-io/bitrise-plugins-analytics/version"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

func Test_VersionTest(t *testing.T) {
	t.Log("version flag")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("")
		require.NoError(t, err)

		cmd := command.New(binPth, "--version")
		cmd.SetDir(tmpDir)
		out, err := cmd.RunAndReturnTrimmedCombinedOutput()
		require.NoError(t, err, out)
		require.Equal(t, version.VERSION, out)
	}

	t.Log("version flag")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("")
		require.NoError(t, err)

		cmd := command.New(binPth, "-v")
		cmd.SetDir(tmpDir)
		out, err := cmd.RunAndReturnTrimmedCombinedOutput()
		require.NoError(t, err, out)
		require.Equal(t, version.VERSION, out)
	}
}
