package main

package main  
import (  
    "fmt"
    "log"
    "strings"
    "github.com/nlpodyssey/gotokenizers"
    "github.com/dgraph-io/dgo/v2"
    "github.com/dgraph-io/dgo/v2/protos/api"
    "github.com/sajari/word2vec"
)

func initWord2VecModel() *word2vec.Model {
    // Load the pre-trained Word2Vec model (you can download it from various sources)
    model, err := word2vec.Load("path/to/your/word2vec.bin")
    if err != nil {
        log.Fatalf("Error loading Word2Vec model: %v", err)
    }
    return model
}

func compressText(text string, model *word2vec.Model) string {
    tokenizer := gotokenizers.NewSentenceTokenizer()
    sentences := tokenizer.Tokenize(text)
    var compressedText []string
    
    for _, sentence := range sentences {
        tokens := strings.Fields(sentence)
        var compressedSentence []string
        for _, token := range tokens {
            // Get the nearest words to the token in the Word2Vec model
            nearestWords, err := model.NearestWords(token, 10)
            if err != nil {
                log.Printf("Error finding nearest words for %s: %v", token, err)
                continue
            }
            // Use the first nearest word as the compressed representation
            compressedSentence = append(compressedSentence, nearestWords[0])
        }
        compressedText = append(compressedText, strings.Join(compressedSentence, " "))
    }
    return strings.Join(compressedText, ". ")
}

func storeWordEmbeddings(model *word2vec.Model) {
    // Connect to the Dgraph database
    dg, err := dgo.NewDgraphClient(api.NewDgraphClientConfig().WithHost("localhost:9080"))
    if err != nil {
        log.Fatalf("Error connecting to Dgraph: %v", err)
    }
    defer dg.Close()

    // Create a schema for the word embeddings
    err = dg.Alter(`
        create type WordEmbedding {
            word: string @index(term) .
            embedding: [float] .
        }
    `)
    if err != nil {
        log.Fatalf("Error creating schema: %v", err)
    }

    // Set up a mutation to insert word embeddings
    mu := new(api.Mutation)
    nquads := make([]string, 0)
    for _, word := range model.Vocab {
        embedding := model.WordVector(word)
        nquad := fmt.Sprintf(`
            <_:word-%s> <word> "%s" .
            <_:word-%s> <embedding> %v .
        `, word, word, word, embedding)
        nquads = append(nquads, nquad)
    }
    mu.SetNquads(nquads)

    // Commit the mutation to store the word embeddings in the Dgraph database
    txn := dg.NewTxn()
    if err := txn.Mutate(mu); err != nil {
        log.Fatalf("Error mutating: %v", err)
    }
    if err := txn.Commit(); err != nil {
        log.Fatalf("Error committing: %v", err)
    }
}

func main() {
    // Initialize the Word2Vec model
    model := initWord2VecModel()
    
    // Example text for compression
    text := "This is a sample text for compression using word embeddings."
    
    // Compress the text using Word2Vec
    compressedText := compressText(text, model)
    fmt.Println("Compressed Text:", compressedText)
    
    // (Optional) Store the word embeddings in a Graph Database
    storeWordEmbeddings(model)
}
