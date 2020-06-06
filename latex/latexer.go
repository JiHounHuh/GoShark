package latex

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var templates, parseErr = template.ParseFiles("./reportTemplate.tex")

func CompileReport(filename string) error {
	if parseErr != nil {
		fmt.Println("Cannot parse ./reportTemplate.tex, please check if it exists and you have suffienct permissions to access it.")
		return parseErr
	}

	cmd := exec.Command("/usr/bin/pdflatex", filename)
	cmdErr := cmd.Run()

	if cmdErr != nil {
		fmt.Println("Error running pdflatex on report.tex", cmdErr)
		return cmdErr
	}

	return nil
}

func MakeReport(filename string) error {
	outputName := "report.tex"
	toReadBytes, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		fmt.Println("Error reading insecure packets from file", readErr)
		return readErr
	}

	f := func(c rune) bool { return c == '~' || c == '\n' }
	splits := strings.FieldsFunc(string(toReadBytes), f)

	var lineToAdd string = ""
	var content string = ""
	rowCount := 1
	for i, c := range splits {
		if len(c) == 1 {
			if int(c[0]) < 32 {
				continue
			} else {
				lineToAdd += c
			}
		} else {
			lineToAdd += c
		}
		lineToAdd += " & "

		if i%((5*rowCount)-2) == 0 && i != 0 {
			lineToAdd += "\\multicolumn{1}{m{8.5cm}|}{"
		}

		if i%((5*rowCount)-1) == 0 && i != 0 {
			// we need to escape underscores in the latex
			content += strings.Replace(lineToAdd[0:len(lineToAdd)-3], "_", "\\_", -1)
			content += "}\\\\\n\\hline\n"
			lineToAdd = ""
			rowCount += 1
		}
	}

	if len(content) != 0 {
		content = content[0 : len(content)-1]
	} else {
		content = string(toReadBytes)
	}

	reportTex, creatErr := os.Create(outputName)
	if creatErr != nil {
		fmt.Println("Error creating .tex file")
		return creatErr
	}

	// write our report.tex to the writer
	writeErr := templates.Execute(reportTex, content)

	fmt.Println(outputName, content, filename)

	if writeErr != nil {
		fmt.Println("Error writing file, %s", writeErr)
		return writeErr
	}

	reportTex.Close()
	compErr := CompileReport(outputName)

	if compErr != nil {
		fmt.Println("Error with CompileReport()")
		return compErr
	}

	return nil
}
