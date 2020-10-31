package config

var ConfigFFmpegPath = "ffmpeg"

const ConfigServicePathBase = "miniTranscoder"
const ConfigServicePathWeb = ConfigServicePathBase + "/www"
const ConfigServicePathBin = ConfigServicePathBase + "/bin"

// db
const ConfigServiceSqliteDir = ConfigServicePathBase + "/db"
const ConfigServiceSqlitePath = ConfigServicePathBase + "/db/mt.sqlite"
const ConfigDatabaseUrl = ""

// data
const ConfigDataUploadPath = ConfigServicePathBase + "/data/upload"
const ConfigDataOutputPath = ConfigServicePathBase + "/data/output"