package main

import (
	"time"

	"github.com/alexandrevicenzi/go-sse"
)

var events *sse.Server
var timer *time.Timer

