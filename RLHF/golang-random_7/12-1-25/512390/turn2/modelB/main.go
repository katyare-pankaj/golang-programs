package main
import (
    "fmt"
    "sync"
    "time"
)

type DownloadStateMachine struct {
    currentState    string
    transitionCond *sync.Cond
    downloadFile    chan string
    wg             sync.WaitGroup
}
func newDownloadStateMachine() *DownloadStateMachine {
    dsm := &DownloadStateMachine{
        transitionCond: sync.NewCond(&sync.Mutex{}),
        downloadFile:    make(chan string),
    }
    return dsm
}

//Start the download protocol
func (dsm *DownloadStateMachine) start() {
    dsm.wg.Add(1)
    defer dsm.wg.Done()
    // transition from "STARTED" to "WAITING"
    dsm.transition("WAITING")
}

func (dsm *DownloadStateMachine) transition(nextState string) {
    dsm.transitionCond.L.Lock()
    defer dsm.transitionCond.L.Unlock()
    
    dsm.currentState = nextState
    dsm.transitionCond.Signal()
    fmt.Println("Transitioning to:", dsm.currentState)
}

func (dsm *DownloadStateMachine) runStateMachine() {
    go dsm.waitingState()
    go dsm.downloadingState()
    go dsm.verifyingState()
    go dsm.finishedState()
    go dsm.start()
    //Main thread wait till finish all state
    dsm.wg.Wait()
}

func (dsm *DownloadStateMachine) waitingState() {
    dsm.wg.Add(1)
    defer dsm.wg.Done()
    for {
        dsm.transitionCond.L.Lock()
        for dsm.currentState != "WAITING" {
            dsm.transitionCond.Wait()
        }
        fmt.Println("File available, Transitioning to: DOWNLOADING")
        dsm.transitionCond.L.Unlock()
        // Indicate that a file is ready for download
        dsm.downloadFile <- "file.txt"
    }
}

func (dsm *DownloadStateMachine) downloadingState() {
    dsm.wg.Add(1)
    defer dsm.wg.Done()
    for {
        dsm.transitionCond.L.Lock()
        for dsm.currentState != "DOWNLOADING" {
            dsm.transitionCond.Wait()
        }
        fileToDownload := <-dsm.downloadFile
        dsm.transitionCond.L.Unlock()
        
        // Simulate download duration
        fmt.Println("Downloading file:", fileToDownload)
        time.Sleep(3 * time.Second)
        
        // Transition to "VERIFYING" state
        dsm.transition("VERIFYING")
    }
}

func (dsm *DownloadStateMachine) verifyingState() {
    dsm.wg.Add(1)
    defer dsm.wg.Done()
    for {
        dsm.transitionCond.L.Lock()
        for dsm.currentState != "VERIFYING" {
            dsm.transitionCond.Wait()
        }
        dsm.transitionCond.L.Unlock()
       
        // Simulate verification duration
        fmt.Println("Verifying file")
        time.Sleep(2 * time.Second)

        // Simulate successful verification
        // Transition to "FINISHED" state
        dsm.transition("FINISHED")
    }
}
func (dsm *DownloadStateMachine) finishedState() {
    dsm.wg.Add(1)
    defer dsm.wg.Done()
    for {
        dsm.transitionCond.L.Lock()
        for dsm.currentState != "FINISHED" {
            dsm.transitionCond.Wait()
        }
        dsm.transitionCond.L.Unlock()
        fmt.Println("Download completed successfully.")
        // Transition back to "WAITING" state for the next download
        dsm.transition("WAITING")
    }