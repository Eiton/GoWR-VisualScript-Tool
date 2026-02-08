# GOWR-Visualscript-Tool
This is a WIP tool translate the .visualscript file to pseudocode to better understand the game logic and file structure. Note that only a fraction of the nodes are translated and they are not guarenteed to be 100% correct.

## How To use
Install Golang

Clone the project ```git clone github.com/Eiton/GoWR-VisualScript-Tool```

Build the project ```go build -o visualscriptTool.exe ./src```

Read a .visualscript file ```./visualscriptTool.exe path/to/the/file > out.txt```