# gmvault-search
A quick script to search in [gmvault](https://github.com/gaubert/gmvault) meta files for emails that contain given searchterms in their subject.

*Tested on gmvault backups generated with gmvault v1.8.1-beta.*

## Installation
```golang
go get github.com/samwierema/gmvault-search
```

## Usage
The script accepts two (required) parameters: filepath, and searchterm. E.g.
```golang
gmvault-search /path/to/gmvault/dir/ "search words"
```
The search will return a list of meta files, and the email subjects that belong to them. E.g.
```golang
Filename:  /path/to/gmvault/dir/1234567890.meta
Subject:   Email with search words in the subject
---
```
If you want to see the actual email, look up the corresponding ```eml.gz``` file, unzip, and enjoy!

## Author
* [Sam Wierema](http://wiere.ma)
