package commandline

import (
	"fmt"
	jjson "github.com/JIIL07/jcloud/pkg/json"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/JIIL07/jcloud/internal/storage"
	"github.com/JIIL07/jcloud/pkg/cookies"
	"github.com/JIIL07/jcloud/pkg/ctx"
)

func HandleCmdExec(w http.ResponseWriter, r *http.Request) {
	var ok bool
	s, ok = jctx.FromContext[*storage.Storage](r.Context(), "storage")
	if !ok {
		http.Error(w, "Storage not found", http.StatusInternalServerError)
		return
	}

	store, err := cookies.Store.Get(r, "admin")
	if err != nil {
		jjson.RespondWithError(w, err)
		return
	}

	if store.IsNew {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized")) // nolint:errcheck
		return
	}

	var req jjson.Request
	req.Command = r.URL.Query().Get("command")

	output, err := ExecuteCommand(req.Command)
	if err != nil {
		jjson.RespondWithError(w, err)
		return
	}

	jjson.RespondWithJSON(w, output)

}

func ExecuteCommand(command string) (string, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %w", err)
	}

	return string(output), nil
}
