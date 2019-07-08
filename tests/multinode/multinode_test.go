package multinode

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	daemonName = "linkd"
	cliName    = "linkcli"
)

func TestMain(m *testing.M) {
	err := os.Chdir("../..")

	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	//err = exec.Command("rm", "-rf", "./build").Run()
	//if err != nil {
	//	fmt.Printf("remove build failed: %v", err)
	//	os.Exit(1)
	//}
	//
	//makeCommand := exec.Command("make")
	//err = makeCommand.Run()
	//if err != nil {
	//	fmt.Printf("build failed: %v", err)
	//	os.Exit(1)
	//}

	exitCode := m.Run()

	//Do something TearDown
	//err = exec.Command("rm", "-rf", "./build").Run()
	//if err != nil {
	//	fmt.Printf("remove build failed: %v", err)
	//	os.Exit(1)
	//}

	os.Exit(exitCode)
}

func TestCliArgs(t *testing.T) {
	makeCommand := exec.Command("make", "build")
	out, err := makeCommand.Output()
	if err != nil {
		fmt.Printf("build failed: %v\n", err)
		fmt.Println(string(out))
		t.Error(err)
	}
}
