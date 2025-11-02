package ecacpluginsdk

import (
	"net/rpc"
	"os"
)

type RPCRequest struct {
	Params map[string]any
}

type RPCResponse struct {
	Result string
	Error  string
}

type Runner interface {
	Run(params map[string]any) (string, error)
}

func Serve(r Runner) {
	rpc.RegisterName("Plugin", &adapter{runner: r})
	rpc.ServeConn(&stdio{})
}

type adapter struct{ runner Runner }

func (a *adapter) Run(req RPCRequest, resp *RPCResponse) error {
	result, err := a.runner.Run(req.Params)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Result = result
	return nil
}

type stdio struct{}

func (s *stdio) Read(b []byte) (int, error)  { return os.Stdin.Read(b) }
func (s *stdio) Write(b []byte) (int, error) { return os.Stdout.Write(b) }
func (s *stdio) Close() error                { return nil }
