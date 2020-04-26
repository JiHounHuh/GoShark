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
/*
\begin{center}
 \begin{tabular}{||c c c c||} 
 \hline
 SrcIP & SrcPort & DstIP & DstPort \\ [0.5ex] 
 \hline\hline
 1 & 6 & 87837 & 787 \\ 
 \hline
 2 & 7 & 78 & 5415 \\
 \hline
 3 & 545 & 778 & 7507 \\
 \hline
 4 & 545 & 18744 & 7560 \\
 \hline
 5 & 88 & 788 & 6344 \\ [1ex] 
 \hline
\end{tabular}
\end{center}
*/

	// for loop to search file to add to latex
	toRead, readErr := ioutil.ReadFile("toRead")

	if readErr != nil {
		fmt.Println("Error reading insecure packets from file", readErr)
		os.Exit(1)
	}

	splits := strings.Fields(string(toRead))

	var lineToAdd string = ""
	for i, c := range splits {
		fmt.Println("i",i)
		lineToAdd += c // dstIP, srcIP, srcPort, dstPort
		lineToAdd += " & "
		fmt.Println("linesToAdd",lineToAdd)

		toC := i+1

		if toC % 5 == 0  {
			fmt.Println("GOT ALL 4")
			content += lineToAdd[0:len(lineToAdd)-3] 
			content += " \\\\\n\\hline\n"
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