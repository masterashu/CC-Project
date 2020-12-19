package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func help() {
	fmt.Printf("Use:\n")
	fmt.Printf("%s download [image-name]", os.Args[0])
	fmt.Printf("%s run [image-name]\n", os.Args[0])
	fmt.Printf("\t\t\t Use to download and run a image")
	fmt.Printf("%s start [image-name]\n", os.Args[0])
	fmt.Printf("\t\t\t Use to start pre-downloaded image")
	fmt.Printf("%s download [image-name]\n", os.Args[0])
	fmt.Printf("\t\t\t Use to download a image")
	fmt.Printf("%s help\n", os.Args[0])
	fmt.Printf("\t\t\t Print the help message")
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Error: No Arguments Specified\n")
		help()
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case "start":
			os.Mkdir(".temp", 0777)
			os.Chdir(".temp")
			run()
			break
		case "download":
			os.Mkdir(".temp", 0777)
			os.Chdir(".temp")
			download()
			break
		case "run":
			os.Mkdir(".temp", 0777)
			os.Chdir(".temp")
			download()
			run()
			break
		case "help":
			help()
			break
		// For internal use
		case "initialize":
			child()
			break
		default:
			fmt.Printf("Error: Invalid Argument. Use %s help\n", os.Args[0])
			os.Exit(1)
		}
	}
}

func checkImageArg() string {
	if len(os.Args) == 2 {
		fmt.Printf("Error: No Image name specified.\n")
	}
	return os.Args[2]
}

func makeRootFS(image string) {
	os.Remove("image.tar")
	os.RemoveAll(image)
	cmd := exec.Command("/bin/bash", "-c", "docker export $(docker create "+image+") -o image.tar")
	os.Mkdir(image, 0777)
	must(cmd.Run())
	newcmd := exec.Command("/bin/bash", "-c", "tar -xf image.tar -C "+image)
	must(newcmd.Run())
	os.Remove("image.tar")
}

func download() {
	image := checkImageArg()
	makeRootFS(image)
}

func run() {
	image := checkImageArg()
	cmd := exec.Command("/proc/self/exe", "initialize", image)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:  syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Credential:  &syscall.Credential{Uid: 0, Gid: 0},
		UidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getuid(), Size: 1}},
		GidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getgid(), Size: 1}},
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Inside Container, Running %s as pid %d\n", os.Args[2], os.Getpid())

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	must(syscall.Chroot(os.Args[2]))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
	fmt.Printf("Exited from container\n")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
