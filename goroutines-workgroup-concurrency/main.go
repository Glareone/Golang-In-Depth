package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// TASK: Process 1M of files using 100 workers
// IN C# it's quite obvious how to do this using semaphoreSlim and awaitForeach
// in Go you dont have this functionality, but it's possible to emulate it.

// C# Code for the reference
//
//private readonly object _lock = new();
//private async Task Do()
//{
//  const int maxParallelTasks = 100;
//  var semaphore = new SemaphoreSlim(maxParallelTasks);
//  var tasks = new List<Task>();
//
//  await foreach (BlobItem blobItem in blobContainerClient.GetBlobsAsync())
//  {
//    Task insertBlobTask;
//
//    // Immediately start the asynchronous work, controlled by the semaphore
//	try
//	{
//		await semaphore.WaitAsync();
//		insertBlobTask = _retryPolicy.ExecuteAsync(async () =>
//		{
//			await InsertBlob(blobItem.Name);
//		});
//	}
//	finally
//	{
//		semaphore.Release();  // Release semaphore whether success or failure.
//	}
//
//	// prevent race condition, limited lock scope
//	lock (_lock)
//	{
//		tasks.Add(insertBlobTask);
//
//		// Periodically remove completed tasks from the list to manage memory for large sets of blobs
//		// Otherwise this array will become huge and consume a lot of memory
//		if (tasks.Count % 100 == 0)
//		{
//			tasks = tasks.Where(t => !t.IsCompleted).ToList();
//		}
//	}
//  }
//
//  await Task.WhenAll(tasks);
//}

func main() {
	numFiles := 100
	maxWorkers := 5

	// Semaphore using a buffered channel to limit concurrency
	semaphoreWorkerChan := make(chan int, maxWorkers)
	fileChan := make(chan string, numFiles)

	// Create a WaitGroup for the all files to control when all files processed
	// we don't need a separate Wait group for workers, the number of parallel workers controlled by semaphoreWorkerChan := make(chan int, maxWorkers)
	var allFilesSenderWaitGroup sync.WaitGroup
	// we wait when all files will be processed
	allFilesSenderWaitGroup.Add(numFiles)

	// Start worker goroutines
	for i := 0; i < numFiles; i++ {
		// acquire the lock in the Channel named semaphore for better understanding
		// in function I release it
		semaphoreWorkerChan <- i
		go fetchAndProcessEachFileIndividually(i, semaphoreWorkerChan, fileChan, &allFilesSenderWaitGroup)
	}

	// ... (process results from fileChan - optional) ...

	// Wait when all files finished their processing
	allFilesSenderWaitGroup.Wait()

	fmt.Println("All files processed!")
}

func fetchAndProcessEachFileIndividually(fileID int, semaphoreWorkerChan chan int, fileChan chan string, allFilesSenderWaitGroup *sync.WaitGroup) {
	// when file is fetched and processed - release 1 lock from the wait group
	defer allFilesSenderWaitGroup.Done()
	// once we get fileId from the semaphore - it's ready to get another task into work
	// using anonymous function. It's because we cant use defer directly like defer <- semaphoreWorkerChan.
	defer func() { <-semaphoreWorkerChan }()

	fileContent, err := fetchFile(fileID)

	if err != nil {
		fileChan <- fmt.Sprintf("fileID %d processed failed\n", fileID)
	}

	processFile(fileContent)
	fileChan <- fmt.Sprintf("fileID %d processed successfully\n", fileID)
	fmt.Printf("> %d FileID has been processed successfully.\nWorker now will be released using defer\n", fileID)
}

// Function to simulate fetching a file from a third-party service
func fetchFile(fileID int) (string, error) {
	// Simulate some network delay
	time.Sleep(time.Duration(rand.Int64N(1501)+200) * time.Millisecond) // 200-1700 ms
	return fmt.Sprintf("File content for ID %d", fileID), nil
}

// Function to simulate processing a file
func processFile(fileContent string) {
	// Simulate some processing time
	time.Sleep(time.Duration(rand.Int64N(501)+100) * time.Millisecond) // 100-600 ms
	fmt.Println("Processed:", fileContent)
}
