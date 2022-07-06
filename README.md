#### Module Proto
Sharing Proto file and go struct for Nano Project

##### Installation
```sh
go get github.com/fajarardiyanto/faltar/faltar-module-proto
```

##### Generate go struct from proto
```sh
# This is optimized for Go version 1.17+. Make sure your golang version is 1.17+.
# Don't generate directly
make go-proto
```
##### Folder structure
- go *(go struct generated from proto)*
- proto *(base proto folder)*
    - global *(can be used for global message, or service)*
    - modules *(base modules folder)*
        - module 1
        - module 2
    - services *(base services folder)*
        - service 1
        - service 2

##### Tips
for clean structure and clean code, don't create multiple message.
use import if possible
```proto
syntax = "proto3";

option go_package = "faltar-module-proto/go/modules/listener;listener";

package listener;
import "proto/global/global.proto";
import "proto/modules/monitor/monitor.proto";

service Listener {
	rpc Monitor(global.NoParameter) returns (stream monitor.MonitorResponse) {}
}
```

Before creating a new module or service, maybe it would be better to create a proto first in this repository. This repository can be used for "Golang" or "Javascript".
