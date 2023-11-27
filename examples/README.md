# Running the Example

## Installation

First, install the required Go package using the following command:

```shell
go get github.com/jasonwwl/go-wework
```

## Set Environment Variables

To ensure the example runs correctly, you need to set the following environment variables. Replace `...` with your actual values.

### Internal Corporate Configuration

```shell
export WEWORK_CORP_ID="your-corp-id"
export WEWORK_SECRET="your-corp-secret"
export WEWORK_AGENT_ID="your-agent-id"
```

### Open Corporate Configuration

```shell
export WEWORK_PROVIDER_CORP_ID="provider-corp-id"
export WEWORK_PROVIDER_SECRET="provider-corp-secret"
export WEWORK_SUITE_ID="suite-id"
export WEWORK_SUITE_SECRET="suite-secret"
export WEWORK_SUITE_TOKEN="suite-token"
export WEWORK_SUITE_AES_KEY="suite-aes-key"
export WEWORK_AUTH_CORP_ID="authorized-corp-id"
```

## Run the Example

To run a specific example, use the following command, where `<target>` is the directory of the example you want to execute:

```shell
go run ./example/<target>
```

Ensure that all relevant environment variables are set correctly before running the example.
