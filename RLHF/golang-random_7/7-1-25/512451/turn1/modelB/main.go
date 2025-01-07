import Foundation

// Function to count the number of lines in a file
func countLines(in file: String, completion: @escaping (Int) -> Void) {
    DispatchQueue.global().async {
        let lineCount = try! String(contentsOfFile: file).components(separatedBy: .newlines).filter { !$0.isEmpty }.count
        DispatchQueue.main.async {
            completion(lineCount)
        }
    }
}

// Function to count the number of words in a file
func countWords(in file: String, completion: @escaping (Int) -> Void) {
    DispatchQueue.global().async {
        let wordCount = try! String(contentsOfFile: file).components(separatedBy: .whitespacesAndNewlines).filter { !$0.isEmpty }.count
        DispatchQueue.main.async {
            completion(wordCount)
        }
    }
}

// Function to check if a file contains a specific phrase
func containsPhrase(_ file: String, _ phrase: String, completion: @escaping (Bool) -> Void) {
    DispatchQueue.global().async {
        let content = try! String(contentsOfFile: file)
        let contains = content.range(of: phrase) != nil
        DispatchQueue.main.async {
            completion(contains)
        }
    }
}

// Function to print the summary after all tasks are complete
func printSummary(lineCount: Int, wordCount: Int, containsPhrase: Bool) {
    print("File Summary:")
    print("Line Count: \(lineCount)")
    print("Word Count: \(wordCount)")
    print("Contains Phrase: \(containsPhrase)")
}

func processFile(_ file: String, _ phrase: String) {
    var waitGroup = DispatchGroup()
    
    var lineCount: Int = 0
    var wordCount: Int = 0
    var containsPhrase: Bool = false
    
    // Independent tasks
    waitGroup.enter()
    countLines(in: file) { count in
        lineCount = count
        waitGroup.leave()
    }
    
    waitGroup.enter()
    countWords(in: file) { count in
        wordCount = count
        waitGroup.leave()
    }
    
    waitGroup.enter()
    containsPhrase(file, phrase) { contains in
        containsPhrase = contains
        waitGroup.leave()
    }
    
    // Dependent task
    waitGroup.notify(queue: .main) {
        printSummary(lineCount: lineCount, wordCount: wordCount, containsPhrase: containsPhrase)
    }
}

// Example usage:
let filePath = "/path/to/your/file.txt"
let phraseToFind = "some specific phrase"
processFile(filePath, phraseToFind)