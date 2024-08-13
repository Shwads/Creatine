package scriptParser

import (
	"os"
	"testing"
)

func TestSemAn(t *testing.T) {
	file, fileOpenErr := os.Open("ASTTest/requestScript.txt")
	if fileOpenErr != nil {
        t.Fatal("Failed to open file")
	}
	defer file.Close()

    // Get tokens from lexer
    tokens, lexerErr := lexScript(file)
    if lexerErr != nil {
        t.Fatal("Failed to lex file")
    }

	root := &ASTNode{
		Type:     Global,
		Children: make([]*ASTNode, 0),
	}

	i := 0

	treeErr := GlobalNode(&i, root, tokens)
	if treeErr != nil {
        t.Fatal("Failed to construct AST")
	}

    batches, semAnErr := semAn(root)
    if semAnErr != nil {
        t.Fatal("Failed to traverse AST")
    }

    // Check contents of batch 1
    if batch, ok := batches["this_is_a_batch"]; ok {
        if len(batch) != 3 {
            t.Fatal("batch 'this_is_a_batch' did not have expected length 3")
        }

        if batch[0] != "batch_contents1.yml" {
            t.Fatal("didn't find correct batch at index 0")
        }

        if batch[1] != "batch_contents2.yml" {
            t.Fatal("didn't find correct batch at index 1")
        }

        if batch[2] != "batch_contents3.yml" {
            t.Fatal("didn't find correct batch at index 2")
        }
    } else {
        t.Fatal("didn't find 'this_is_a_batch' in batches")
    }

    // Check contents of batch 2
    if batch, ok := batches["this_is_also_a_batch"]; ok {
        if len(batch) != 3 {
            t.Fatal("batch 'this_is_also_a_batch' did not have expected length 3")
        }

        if batch[0] != "batch_contents4.yml" {
            t.Fatal("didn't find correct batch at index 0")
        }

        if batch[1] != "batch_contents5.yml" {
            t.Fatal("didn't find correct batch at index 1")
        }

        if batch[2] != "batch_contents6.yml" {
            t.Fatal("didn't find correct batch at index 2")
        }
    } else {
        t.Fatal("didn't find 'this_is_a_batch' in batches")
    }
    
    if batch, ok := batches["this_is_an_empty_batch"]; ok {
        if len(batch) != 0 {
            t.Fatal("batch 'this_is_an_empty_batch' did not have expected length 0")
        }
    } else {
        t.Fatal("didn't find 'this_is_a_batch' in batches")
    }

    // Check the last batch
    if batch, ok := batches["last_batch"]; ok {
        if len(batch) != 6 {
            t.Fatal("batch 'this_is_also_a_batch' did not have expected length 3")
        }

        if batch[0] != "batch_contents1.yml" {
            t.Fatal("didn't find correct batch at index 0")
        }

        if batch[1] != "batch_contents2.yml" {
            t.Fatal("didn't find correct batch at index 1")
        }

        if batch[2] != "batch_contents3.yml" {
            t.Fatal("didn't find correct batch at index 2")
        }

        if batch[3] != "batch_contents4.yml" {
            t.Fatal("didn't find correct batch at index 3")
        }

        if batch[4] != "batch_contents5.yml" {
            t.Fatal("didn't find correct batch at index 2")
        }

        if batch[5] != "batch_contents6.yml" {
            t.Fatal("didn't find correct batch at index 2")
        }
    } else {
        t.Fatal("didn't find 'this_is_a_batch' in batches")
    }
}
