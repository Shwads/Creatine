package scriptParser

import (
	"log"
	"os"
)

// The main branch of our script reader.
// Calls the lexer and parser and extracts information from the script file.
func ParseScript(executeAll bool, name string) error {
	file, fileOpenErr := os.Open(name)
	if fileOpenErr != nil {
		log.Printf("Encountered error: %s when opening script file: %s\n", fileOpenErr, name)
		return fileOpenErr
	}
	defer file.Close()

    // Get tokens from lexer
    tokens, lexerErr := lexScript(file)
    if lexerErr != nil {
        return lexerErr
    }

	//printTokens(tokens)

    // Our tree constructor expects a root node upon which to build the AST
	root := &ASTNode{
		Type:     Global,
		Children: make([]*ASTNode, 0),
	}

	i := 0

	treeErr := GlobalNode(&i, root, tokens)
	if treeErr != nil {
		return treeErr
	}

	root.printTree(0)

    batches, semAnErr := semAn(root)
    if semAnErr != nil {
        return semAnErr
    }

    printBatches(batches)

	return nil
}
