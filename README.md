# Oxd-go plugin

src/oxd-client - client source code

src/oxd-client-demo - simple communication demo

src/oxd-client-test - unit tests for communication


## Plugin usage
### Register Site Example

Lines:
1. Preparation of request and request parameters
2. Build Oxd Request (all request names are in constants.go)
3. Create Response

```go
    requestParams := utils.PrepareRegisterParams()
	request := client.BuildOxdRequest(constants.REGISTER_SITE,requestParams)
	var response transport.OxdResponse

```

1. Sending request and receiving response

```go
	client.Send(request,host,&response)
```

1. Create ResponseParameters
2. Fetching Parameters from response

```go
    var responseParams model.RegisterSiteResponseParams
	response.GetParams(&responseParams)
```