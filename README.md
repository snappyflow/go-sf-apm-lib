# go-sf-apm-lib

This package collects the profile key, project name, and app name to send the trace data to SnappyFlow. These are collected automatically from the environment variables set by the user.
The environment variables to be set are:
- **SF_PROJECT_NAME**: specify your project name here
- **SF_APP_NAME**: specify your app name here
- **SF_PROFILE_KEY**: specify the snappyflow key here

If these environment variables are not set, the values are alternatively fetched from the sfagent's config.yaml file.

## Getting started

- **Pre-requisite**

    - Run below command to download or update the sf-go-apm-lib package in your current project.
        ```bash
        go get github.com/snappyflow/go-sf-apm-lib
        ```

- **Example**

```go
import _ "github.com/snappyflow/go-sf-apm-lib"

func main() {
    // rest of the application code
}
```

This package needs to be imported as a blank import solely for its initialization only. 