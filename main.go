package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"syscall"
)

var allFlag = flag.Bool("all", false, "print all information, in the following order, except omit -p and -i if unknown:")
var nodeNameFlag = flag.Bool("nodename", false, "print the network node hostname")
var kernelReleaseFlag = flag.Bool("kernel-release", false, "print the kernel release")
var kernelVersionFlag = flag.Bool("kernel-version", false, "print the kernel version")
var machineFlag = flag.Bool("machine", false, "print the machine hardware name")
var osFlag = flag.Bool("operating-system", false, "print the operating system")

func init()  {
	flag.BoolVar(allFlag, "a", false, "print all information, in the following order, except omit -p and -i if unknown:")
	flag.BoolVar(nodeNameFlag, "n", false, "print the network node hostname")
	flag.BoolVar(kernelReleaseFlag, "r", false, "print the kernel release")
	flag.BoolVar(kernelVersionFlag, "v", false, "print the kernel version")
	flag.BoolVar(machineFlag, "m", false, "print the machine hardware name")
	flag.BoolVar(osFlag, "o", false, "print the operating system")
}

func charsToString(ca []int8) string {
	s := make([]byte, len(ca))
	var lens int
	for ; lens < len(ca); lens++ {
		if ca[lens] == 0 {
			break
		}
		s[lens] = uint8(ca[lens])
	}
	return string(s[0:lens])
}

func generateOutput(output *string, input *string)  {
	if *output != "" {
		*output += " "
	}
	*output += *input
}

func main()  {
	var buf syscall.Utsname;
	var result string

	flag.Usage = func() {
		fmt.Println("Usage: uname [OPTION]...")
		fmt.Println("Print certain system information.  With no OPTION, same as -o.")
		fmt.Println()
		fmt.Println("  -a, --all                print all information, in the following order,")
		fmt.Println("                             except omit -p and -i if unknown:")
		fmt.Println("  -n, --nodename           print the network node hostname")
		fmt.Println("  -r, --kernel-release     print the kernel release")
		fmt.Println("  -v, --kernel-version     print the kernel version")
		fmt.Println("  -m, --machine            print the machine hardware name")
		fmt.Println("  -o, --operating-system   print the operating system")
	}

	flag.Parse()

	if len(os.Args) == 1 {
		*osFlag = true;
	}

	err := syscall.Uname(&buf)
	if err != nil {
		log.Fatal("Uname error\n")
	}

	if *allFlag {
		*nodeNameFlag = true;
		*kernelReleaseFlag = true;
		*kernelVersionFlag = true;
		*machineFlag = true;
		*osFlag = true;
	}

	if *nodeNameFlag {
		str := charsToString(buf.Nodename[:]);
		generateOutput(&result, &str)
	}

	if *kernelReleaseFlag {
		str := charsToString(buf.Release[:]);
		generateOutput(&result, &str)
	}

	if *kernelVersionFlag {
		str := charsToString(buf.Version[:]);
		generateOutput(&result, &str)
	}

	if *machineFlag {
		str := charsToString(buf.Machine[:]);
		generateOutput(&result, &str)
	}

	if *osFlag {
		str := charsToString(buf.Sysname[:]);
		generateOutput(&result, &str)
	}

	fmt.Println(result)
}