package handler
shv-ng
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

func handleInitialize(_ analysis.State, contents []byte) {
	// handling the request data
	var request lsp.InitializeRequest

	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("InitializeRequest failed: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("Connected to: %v %v",
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version),
	)

	// Response to the request
	msg := lsp.NewInitializeResponse(request.ID)
	writer := os.Stdout
	if err := WriteResponse(writer, msg); err != nil {
		logger.Error(fmt.Sprintf("EncodeMessage failed: %v", err))
	}
}
