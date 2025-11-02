package ecacpluginsdk

import (
	"net/rpc"
	"os"
)

// Det här interface måste pluginen implementera
type Runner interface {
	Run(params map[string]any) (string, error)
}

// Startar RPC-server i pluginprocessen
func Serve(r Runner) {
	rpc.RegisterName("Plugin", r)
	rpc.ServeConn(&stdio{})
}

type stdio struct{}

func (s *stdio) Read(b []byte) (int, error)  { return os.Stdin.Read(b) }
func (s *stdio) Write(b []byte) (int, error) { return os.Stdout.Write(b) }
func (s *stdio) Close() error                { return nil }
