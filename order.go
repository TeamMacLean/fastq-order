package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Read struct {
	a string
	b string
	c string
	d string
}

func main() {

	//println(len(os.Args))

	if (len(os.Args) < 4) {
		println("command primaryFile secondaryFile output")
		os.Exit(1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]
	outputFile := os.Args[3]

	if (FileExists(outputFile)) {
		panic("Output file already exists")
	} else {
		TouchFile(outputFile)
	}

	println("processing the primary file...")
	fileOneReads := processFile(file1)
	println("the primary file has " + strconv.Itoa(len(fileOneReads)) + " read(s)")

	println("processing the secondary file...")
	fileTwoReads := processFile(file2)
	println("the secondary file has " + strconv.Itoa(len(fileTwoReads)) + " read(s)")

	println("creating output...")

	for _, element := range fileOneReads {
		name := element.a
		strippedName := name[:len(name) - 1];

		lookingFor := strippedName;

		if (strings.HasSuffix(name, "1")) {
			lookingFor = lookingFor + "2"
		} else {
			lookingFor = lookingFor + "1"
		}
		WriteLineToFile(outputFile, fileTwoReads[lookingFor])
	}

}

func TouchFile(filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func WriteLineToFile(fileName string, read Read) {
	f, err := os.OpenFile(fileName, os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(read.a + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(read.b + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(read.c + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(read.d + "\n"); err != nil {
		panic(err)
	}
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func processFile(file1 string) map[string]Read {

	inFile, _ := os.Open(file1)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	reads := make(map[string]Read)
	pos := 0;
	currentRead := Read{}

	for scanner.Scan() {
		text := scanner.Text()

		pos++;

		switch pos {
		case 1:// {}?
			currentRead = Read{}
			currentRead.a = text
		case 2:// {}?
			currentRead.b = text
		case 3:// {}?
			currentRead.c = text
		case 4:// {}?
			currentRead.d = text
			//println("good " + currentRead.a)
			if (!checkRead(currentRead)) {
				println("BAD!")
				os.Exit(1)
			}

			reads[currentRead.a] = currentRead
			pos = 0
		}
	}
	return reads;
}
func checkRead(read Read) bool {
	if (!strings.HasPrefix(read.a, "@")) {
		println("no @ " + read.a)
		return false
	}
	if (!strings.HasPrefix(read.c, "+")) {
		println("no + " + read.c)
		return false
	}

	return true;
}



