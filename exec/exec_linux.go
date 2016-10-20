package exec

import (
	"os/exec"
	"github.com/johnlion/btsite/config"
)

func run( which string ){
	exec.Command( "/bin/sh" , "-c", "title", config.FULL_NAME  ).Start()
}