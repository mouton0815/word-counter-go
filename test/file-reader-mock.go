package test

type FileReaderMock struct {
    result []string
}

func (r *FileReaderMock) Read(path string) {
    r.result = append(r.result, path)
}

func (r *FileReaderMock) Close() {
}

func NewFileReaderMock(capacity int) FileReaderMock {
    return FileReaderMock{result: make([]string, 0, capacity)}
}
