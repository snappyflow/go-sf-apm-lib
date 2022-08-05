package sfapmlib

// sf decryption key
const (
	EncryptedKey = "U25hcHB5RmxvdzEyMzQ1Ng=="
)

// tag values
const (
	ProjectName = "projectName"
	AppName     = "appName"
)

// config.yaml paths
const (
	WindowsConfigPath = "C:\\Program Files (x86)\\Sfagent\\config.yaml"
	LinuxConfigPath   = "/opt/sfagent/config.yaml"
)

// apm environment variables
const (
	GlobalLabels                    = "_tag_projectName=%s,_tag_appName=%s,_tag_profileId=%s"
	ElasticAPMServerURL             = "ELASTIC_APM_SERVER_URL"
	ElasticAPMGlobalLabels          = "ELASTIC_APM_GLOBAL_LABELS"
	ElasticAPMSpanFramesMinDuration = "ELASTIC_APM_SPAN_FRAMES_MIN_DURATION"
	ElasticAPMStackTraceLimit       = "ELASTIC_APM_STACK_TRACE_LIMIT"
	ElasticAPMVerifyServerCert      = "ELASTIC_APM_VERIFY_SERVER_CERT"
	SfProjectName                   = "SF_PROJECT_NAME"
	SfAppName                       = "SF_APP_NAME"
	SfProfileKey                    = "SF_PROFILE_KEY"
)

// apm environment variables default values
const (
	FramesMinDuration = "1ms"
	StackTraceLimit   = "2"
	VerifyServerCert  = "false"
)
