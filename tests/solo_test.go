package tests

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

const (
	daemonName = "linkd"
	cliName    = "linkcli"
)

func TestMain(m *testing.M) {
	err := os.Chdir("..")

	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	makeCommand := exec.Command("make", "install")
	err = makeCommand.Run()
	if err != nil {
		fmt.Printf("build failed: %v", err)
		os.Exit(1)
	}

	//Ensure binaries
	gopath := os.Getenv("GOPATH")
	if _, err := os.Stat(strings.Join([]string{gopath, "bin", daemonName}, "/")); os.IsNotExist(err) {
		fmt.Printf("binary[%s] is not exist: %v", daemonName, err)
		os.Exit(1)
	}
	if _, err := os.Stat(strings.Join([]string{gopath, "bin", cliName}, "/")); os.IsNotExist(err) {
		fmt.Printf("binary[%s] is not exist: %v", cliName, err)
		os.Exit(1)
	}

	exitCode := m.Run()
	//Do something TearDown
	//

	os.Exit(exitCode)
}

func TestCliArgs(t *testing.T) {

	cmd := exec.Command(daemonName, "start")
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return
	}

	// Do other stuff while cmd runs in background:
	log.Println("Doing other stuff...")
	log.Print(cmd.Process.Pid)

	err := syscall.Kill(cmd.Process.Pid, syscall.SIGINT)
	if err != nil {
		panic(err)
	}

	// And when you need to wait for the command to finish:
	//if err := cmd.Wait(); err != nil {
	//	log.Printf("Cmd returned error: %v", err)
	//}

	//docker run --rm -v ~/workspace/github.com/kfangw/marshallchain/build:/linkd:Z --network marshalchain_localnet --ip 192.168.10.6 kfangw/linkdnode

	//actual := string(output)
	//fmt.Print(actual)

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		//{"no arguments", []string{}, "no-args.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//dir, err := os.Getwd()
			//if err != nil {
			//	t.Fatal(err)
			//}

			cmd := exec.Command(daemonName, tt.args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)

			//expected := loadFixture(t, tt.fixture)
			//
			//if !reflect.DeepEqual(actual, expected) {
			//	t.Fatalf("actual = %s, expected = %s", actual, expected)
			//}
			fmt.Print(actual)
		})
	}
}
