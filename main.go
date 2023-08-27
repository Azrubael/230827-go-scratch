package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker           run <container> cmd args
// go run main.go   run             cmd args

func main() {
	fmt.Println("Hello, Azrubael!")

	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("WHAT???")
	}
}

func run() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Run()
	// must(cmd.Run())
}

func must(err error) {
	if err != nil {
		fmt.Printf("An Error")
		panic(err)
	}
}
