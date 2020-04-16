# shis

## what's this
shis is a command-line tool.  
Fast and quickly search of bash history.

The rough function is as follows.  
```
$ history | grep [keyword]
```

## How to use this
This cmd application was developed by go-lang.  
So you need to get or generate a binary to run.

release page is [here](https://github.com/shogysd/shis/releases)

### How to make binary
Please execute this command on macOS or Linux with go-lang installed.  
( If you are using another operating system (e.g. Windows)? Sorry, this application is not supported. )  

#### get source files
Download the source files from [https://github.com/shogysd/shis/archive/master.zip](https://github.com/shogysd/shis/archive/master.zip).  
( If you want to use git command, please type this command. )
```
// download by git
$ git clone https://github.com/shogysd/shis.git
```

#### build source files
Go to the project root directory and execute the command.
```
$ make build
```
If the build is successful, the binary is generated here -> `/bin/[Arch]/shis` .

### How to run it
plwase type this!
```
$ shis [search target]
// e.g. $ shis ls
```
When you type this command, get the help message.
```
$ shis -h
```

## ATTENTION
To use all features, please make the following settings.

- `PROMPT_COMMAND='history -a'`
  - This command line tool gets history from file.  
  Therefore, to perform an accurate history search, you must write the history to a file each time you run the command.
- `HISTTIMEFORMAT='[Arbitrary set value]'`
  - Time stamps must be saved for accurate time to appear in search results.
