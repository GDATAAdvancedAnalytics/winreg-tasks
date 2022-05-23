winreg-tasks
============

This repository contains structure definitions and some tooling for the BLOBs found in the TaskCache registry key on Windows (`HKLM\Microsoft\Windows NT\CurrentVersion\Schedule\TaskCache`).

To successfully run the tool, it must be launched from a command window with elevated privileges. Otherwise, the tool cannot read the registry keys which are needed to display the requested information.

Examples:
```powershell
# display the available commands:
.\winreg-tasks.exe [-h|--help]

# iterates all tasks and prints a list of Actions, Triggers,
# and the DynamicInfo of all tasks registered to the system
.\winreg-tasks.exe parseall
# same as before but only prints errors; most useful when you
# changed something and want to see if anything broke
.\winreg-tasks.exe parseall -q

# get a more detailed dump of the actions of a task:
.\winreg-tasks.exe actions '{00000000-1111-2222-3333-444444444444}'
# you can pass the path alternatively (leading backslash required!):
.\winreg-tasks.exe actions '\My Task'

# get the triggers of a given task:
.\winreg-tasks.exe triggers '{00000000-1111-2222-3333-444444444444}'
# you can pass the path alternatively (leading backslash required!):
.\winreg-tasks.exe triggers '\My Task'

# get the DynamicInfo of a given task:
.\winreg-tasks.exe dynamicinfo '{00000000-1111-2222-3333-444444444444}'
# you can pass the path alternatively (leading backslash required!):
.\winreg-tasks.exe dynamicinfo '\My Task'
```

All commands (except `parseall`) support the `-d` or `--dump` flag which prints the data read from the registry key as a 16 bytes wide hex dump. I found it much more easy to work with these dumps than exporting a value of a key with regedit and then converting the `hex:00,11,22,...` notation to something more readable.

Generate
========

If you want to re-generate the source files, just use the `generate.sh` script. If you need another language, please adapt the script to your needs.

**Note**: the current release (v0.9) version of kaitai does not support UTF16 strings in Golang. You need to compile the upstream version from Github. Since I don't know anything about scala, the steps below might not be how it's supposed to be done. It works, however, so you might want stick with the following commands if you're not familiar with scala either:
```bash
# install sbt and scala 2.X (important: it MUST be scala 2.X; the lastest release 3.X is incompatible); I use sdkman and would everyone else encourage to do so as well
sdk install sbt
sdk install scala 2.13.8

# clone kaitai-struct-compiler
git clone https://github.com/kaitai-io/kaitai_struct_compiler ~/projects/kaitai
cd ~/projects/kaitai

# generate the staging package (this is where I'm pretty sure there is a much more efficient way to do this):
sbt stage

# cd into package sources
git clone https://github.com/gdataadvancedanalytics/winreg-tasks ~/projects/winreg-tasks
cd ~/projects/winreg-tasks

# run generate script
KAITAI_COMPILER=~/projects/kaitai/jvm/target/universal/stage/bin/kaitai-struct-compiler ./generate.sh
```


Build
=====
If you did not change anything and just want to use the tool, simply run the build script; the `winreg-tasks.exe` is written to the `out` folder. Just make sure, you have a working installation of Golang 1.18 (or later) and then run:
```bash
./build.sh
```
Or, if you are on a Windows platform, you might just want to install the package from source:
```powershell
go install github.com/gdataadvancedanalytics/winreg-tasks/golang/cmd@latest
```

Using the Generated Code
========================
At least for golang, using the generated code is as simple as importing this repository as a package. The Golang files located in `./golang/cmd/` may serve as examples on how to use this package.

Minimum example:
```golang
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/golang/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"golang.org/x/sys/windows/registry"
)

func main() {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Schedule\TaskCache\Tasks\{B75AF762-3C5C-4C74-ADB1-B99F98FDE0E5}`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("cannot open task key: %v", err)
		return
	}
	defer key.Close()

	dynamicInfoRaw, _, err := key.GetBinaryValue("DynamicInfo")
	if err != nil {
		fmt.Printf("cannot get dynamic info for task: %v", err)
		return
	}

	dynamicInfo := generated.NewDynamicInfo()
	if err = dynamicInfo.Read(kaitai.NewStream(bytes.NewReader(dynamicInfoRaw)), dynamicInfo, dynamicInfo); err != nil {
		fmt.Printf("cannot parse dynamic info: %v", err)
		return
	}

	lastErrorCode := dynamicInfo.LastErrorCode
	log.Printf("Last Error Code: 0x%08x", lastErrorCode)
}
```


Licensing
=========
Copyright 2022 G DATA Advanced Analytics GmbH

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
