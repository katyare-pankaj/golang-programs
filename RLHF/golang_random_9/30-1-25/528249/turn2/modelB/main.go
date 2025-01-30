func readOrWriteFile(filename string, doRead bool) error {
	// Open the file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	// Ensure the file is closed when the function returns
	defer file.Close()

	// Now you can read or write to the file as needed
	if doRead {
		// Simulate reading data from the file
		data := make([]byte, 1024)
		n, err := file.Read(data)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading file: %v", err)
		}
		fmt.Printf("Read %d bytes from the file:\n%s\n", n, data[:n])
	} else {
		// Simulate writing data to the file
		data := []byte("Hello, world!")
		n, err := file.Write(data)
		if err != nil {
			return fmt.Errorf("error writing file: %v", err)
		}
		fmt.Printf("Wrote %d bytes to the file:\n%s\n", n, data)
	}

	return nil
} 