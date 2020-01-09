var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// Cancel traversal when input is detected.
go func() {
	os.Stdin.Read(make([]byte, 1))
	close(done)
}()

