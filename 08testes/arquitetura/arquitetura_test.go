package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()                   //Permite testar em pararelo
	if runtime.GOARCH == "arm64" { //para simular erro mudar para "amd64"
		t.Skip("Não funciona em arquitetura arm64")
	}
	//...
	t.Fail()
}
