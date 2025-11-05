# Working with Files

- directories
- file path helpers
- files
- using file system interfaces

## Directories
os.ReadDir - read directories
takes path as string parameter
returns slice of os.DirEntry value which represents a file or directory


to get info about just one file, we can use "os.Stat"
if the file exists we will get valid info, else there will be an error

info, err := os.Stat("path/to/a/file.txt")
if err != nil {
  ...
}

WalkDir function shall be used to walk all subdirectories
- provide a function argument

Go will skip few special directories

## filepath helpers
Ext
FromSlash
Base
...etc

## Files

## FS interface

interfaces
File
MapFile

## embedding files
