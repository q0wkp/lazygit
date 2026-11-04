package file

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var StageFailureShowsRealStatus = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "When staging a file fails, the file is shown with its real status rather than as staged",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("first commit")
		shell.CreateFile("file", "content")

		// makes git add fail, while git status still works
		shell.CreateFile(".git/index.lock", "")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsFocused().
			Lines(
				Equals("?? file").IsSelected(),
			).
			PressPrimaryAction().
			Tap(func() {
				t.ExpectPopup().Alert().Title(Equals("Error")).Content(Contains("index.lock")).Confirm()
			}).
			Lines(
				Equals("?? file").IsSelected(),
			)
	},
})
