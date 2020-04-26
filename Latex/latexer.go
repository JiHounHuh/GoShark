package Latex

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)
//\usepackage{fullpage,latexsym,picinpar,amsmath,amsfonts}
func MakeReport() {
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
		return c == '~' || c == '\n'
	}

	splits := strings.FieldsFunc(string(toRead),f)

	var lineToAdd string = ""
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

		if i % ((5*rowCount)-2) == 0 && i != 0 {
			lineToAdd += "\\multicolumn{1}{m{3cm}|}{"
		}

		if i % ((5 * rowCount) - 1) == 0 && i != 0 {
			content += strings.Replace(lineToAdd[0:len(lineToAdd)-3],"_","\\_",-1)
			content += "}\\\\\n\\hline\n"
			lineToAdd = ""
			rowCount += 1
		}
	}
	content = content[0:len(content)-1]

	content += end

	writeErr := ioutil.WriteFile("report.tex", []byte(content), 0644)

	if writeErr != nil {
		fmt.Println("Error writing file")
		os.Exit(1)
	}

	cmd := exec.Command("/usr/bin/pdflatex","report.tex")
	cmdErr := cmd.Run()

	if cmdErr != nil {
		fmt.Println("Error running pdflatex on report.tex",cmdErr)
		os.Exit(1)
	}
}
