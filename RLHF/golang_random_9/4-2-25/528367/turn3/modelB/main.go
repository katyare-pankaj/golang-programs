package main  
  
import (  
    "fmt"
    "io/ioutil"
    "os"
)

// FileManager defines an interface for file operations
type FileManager interface {  
    Read(filename string) (string, error)
    Write(filename string, data string) error
    Delete(filename string) error
    FileExists(filename string) (bool, error)
}  

// DiskFileManager implements the FileManager interface for disk storage  
type DiskFileManager struct{}

func (dfm *DiskFileManager) Read(filename string) (string, error) {  
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}  

func (dfm *DiskFileManager) Write(filename string, data string) error {  
    return ioutil.WriteFile(filename, []byte(data), 0644)
}  

func (dfm *DiskFileManager) Delete(filename string) error {  
    return os.Remove(filename)
}

func (dfm *DiskFileManager) FileExists(filename string) (bool, error) {
    _, err := os.Stat(filename)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

// InMemoryFileManager implements the FileManager interface for in-memory storage  
type InMemoryFileManager struct {  
    files map[string]string
}  

func NewInMemoryFileManager() *InMemoryFileManager {  
    return &InMemoryFileManager{files: make(map[string]string)}
}  

func (imfm *InMemoryFileManager) Read(filename string) (string, error) {  
    data, exists := imfm.files[filename]
    if !exists {
        return "", fmt.Errorf("file not found: %s", filename)
    }
    return data, nil
}  

func (imfm *InMemoryFileManager) Write(filename string, data string) error {  
    imfm.files[filename] = data
    return nil
}  

func (imfm *InMemoryFileManager) Delete(filename string) error {  
    delete(imfm.files, filename)
    return nil
}

func (imfm *InMemoryFileManager) FileExists(filename string) (bool, error) {
	_, exists := imfm.files[filename]
	return exists, nil
}

func main() {
	filename := "sample.txt"
    data := "Hello, World!"

    //Using DiskFileManager
    fmt.Println("Disk File Manager:")
    diskManager := &DiskFileManager{}
    
    exists, err := diskManager.FileExists(filename)
    if err != nil {
        fmt.Println("Error checking file existence:", err)
    } else {
        fmt.Println("File exists:", exists)
    }
    
    if err := diskManager.Write(filename, data); err != nil {
        fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
        return
    }
    
    exists, err = diskManager.FileExists(filename)
    if err != nil {
        fmt.Println("Error checking file existence after write:", err)
    } else {
        fmt.Println("File exists after write:", exists)
    }
    
    readContent, err := diskManager.Read(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
    } else {
        fmt.Println("Content read:", readContent)
    }
    
    if err := diskManager.Delete(filename); err != nil {
        fmt.Fprintf(os.Stderr, "Delete error: %v\n", err)
    } else {
        fmt.Println("File deleted successfully.")
    }
    
    exists, err = diskManager.FileExists(filename)
    if err != nil {