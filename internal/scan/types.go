package scan

type FileInfo struct{
    Path string
    Size int64 // bytes
}

// Result associates a SHA-256 hash - list of files with that hash
type Result map[string][]FileInfo
