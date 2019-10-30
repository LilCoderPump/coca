package call

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "../../language/java"
)

type JavaCallApp struct {

}

func (j *JavaCallApp) AnalysisPath(codeDir string) {
	files := (*JavaCallApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaCallApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaCallListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo := listener.getNodeInfo()
		bytes, _ := json.Marshal(nodeInfo)
		fmt.Println(string(bytes))
	}
}

func (j *JavaCallApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *JavaCallApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
