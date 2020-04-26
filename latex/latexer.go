package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)
//\usepackage{fullpage,latexsym,picinpar,amsmath,amsfonts}
func main() {
	content := 
`
\documentclass{article}
\usepackage{fullpage,latexsym,picinpar,amsmath,amsfonts}
\usepackage{graphicx}
\usepackage{array}
\newcolumntype{L}{>{\centering\arraybackslash}m{3cm}}
\usepackage{tabularx}
\begin{document}
\begin{center}
	\includegraphics[scale=0.25]{Cyber_525x438.png}\\
\end{center}	
\centerline{\large \bf REPORT}
\begin{center}
During our packet capture, we found the following details that might be insecure\\
 \begin{tabular}{||c c c c c||} 
 \hline 
 SrcIP & DstIP & SrcPort & DstPort & Finding \\ [0.5ex] 
 \hline\hline
`
	end :=
`
\end{tabular}
\end{center}
\end{document}
`
	// for loop to search file to add to latex
	toRead, readErr := ioutil.ReadFile("toRead.txt")

	if readErr != nil {
		fmt.Println("Error reading insecure packets from file", readErr)
		os.Exit(1)
	}

	f := func(c rune) bool {
		return c == '~'
	}

	splits := strings.FieldsFunc(string(toRead),f)

	var lineToAdd string = ""
	for i, c := range splits {
		fmt.Println("i",i)
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
		fmt.Println("linesToAdd",lineToAdd)

		toC := i+1

		for in := 1; in < 1000; in++ {
			if toC % (4+(5*in)) == 0 {
				lineToAdd += "\\multicolumn{1}{m{3cm}|}{"
			}
		}

		if toC == 4 {
			lineToAdd += "\\multicolumn{1}{m{3cm}|}{"
		}

		if toC % 5 == 0  {
			fmt.Println("GOT ALL 4")
			content += lineToAdd[0:len(lineToAdd)-3] 
			content += "}\\\\\n\\hline\n"
			lineToAdd = ""
		}
	}
	content = content[0:len(content)-1]

	content += end

	fmt.Println("To write:\n",content)
	writeErr := ioutil.WriteFile("report.tex", []byte(content), 0644)

	if writeErr != nil {
		fmt.Println("Error writing file")
		os.Exit(1)
	}

	cmd := exec.Command("/usr/bin/pdflatex","report.tex")
	cmdErr := cmd.Run()

	if cmdErr != nil {
		fmt.Println("Error running pdflatex on report.tex",cmdErr)
		//os.Exit(1)
	}
}