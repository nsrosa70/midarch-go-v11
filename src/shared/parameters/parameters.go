package parameters

import "time"

// Dirs
//const DIR_BASE  = "/go/midarch-go"  // docker
const DIR_BASE = "/Users/nsr/Dropbox/go/midarch-go-v11"
const DIR_PLUGINS = DIR_BASE + "/src/plugins"
const DIR_CSP = DIR_BASE + "/src/cspspecs"
const DIR_SOURCE = DIR_BASE
const DIR_CONF = DIR_BASE + "/src/apps/confs"
const DIR_GO = "/usr/local/go/bin"
const DIR_FDR = "/Volumes/Macintosh HD/Applications/FDR4-2.app/Contents/MacOS"
const DIR_CSPARSER = DIR_BASE+"/src/verificationtools/cspdot/csparser"
const COMPONENTS_PATH = "components"
const CONNECTORS_PATH = "connectors"
const NAMINGCLIENTPROXY_PATH = "namingclientproxy"
const COMPONENTS_DIR = "components"
const CONNECTORS_DIR = "connectors"

const ADL_COMMENT = "//"

// Ports
const NAMING_PORT = 4040
const CALCULATOR_PORT = 2020
const FIBONACCI_PORT = 2030
const QUEUEING_PORT = 2040

//
const JAVA_COMMAND = "java"
const JAR_COMMAND = "-jar"

var SetOfPorts = map[string]int{
	"NAMING_PORT":     NAMING_PORT,
	"CALCULATOR_PORT": CALCULATOR_PORT,
	"FIBONACCI_PORT":  FIBONACCI_PORT,
	"QUEUEING_PORT":   QUEUEING_PORT}

const NO_CHANGE = 0
const REACTIVE_CHANGE = 1
const EVOLUTIVE_CHANGE = 2
const PROACTIVE_CHANGE = 3

const CHAN_BUFFER_SIZE = 1
const QUEUE_SIZE = 100

//const PLUGIN_BASE_NAME  = "calculatorinvoker"
//const PLUGIN_BASE_NAME = "sender"
const PLUGIN_BASE_NAME = "fibonacciinvoker"

const GRAPH_SIZE = 30

const MAX_NUMBER_OF_ACTIVE_CONSUMERS = 10
var IS_EVOLUTIVE  = false
var IS_CORRECTIVE = false
var IS_PROACTIVE  = false

var MONITOR_TIME time.Duration   // seconds
var INJECTION_TIME time.Duration // seconds
var REQUEST_TIME time.Duration   // milliseconds
var STRATEGY = 0             // 1 - no change 2 - change once 3 - change same plugin 4 - alternate plugins
var SAMPLE_SIZE = 0
var NAMING_HOST = ""
var QUEUEING_HOST = ""
var SHOW_PARAMETERS = true
var FIBONACCI_HOST = ""
