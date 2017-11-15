# Goartik-Cloud

**Goartik-Cloud** is Golang library for Samsung ARTIK Cloud.

## Todo
- Trails API
- Docs
- More example projects

## Installation

```sh
go get -u github.com/projekt86/goartik-cloud
```

## Examples

Set AccessToken to default APIClient
```go
client.DefaultAPIClient.AccessToken = "ACCESS_TOKEN"
```

Create New APIClient
```go
apiClient := &client.APIClient{
	AccessToken: "",
}
```

Create new Device API instance with default APIClient
```go
devicesAPIClient := api.NewDevicesAPI()
```

Create new Device API instance with custom APIClient
```go
apiClient := &client.APIClient{
	AccessToken: "ACCESS_TOKEN",
}
devicesAPIClient := api.NewDevicesAPIWithClient(apiClient)
```
Sample get request
```go
r, err := devicesAPIClient.GetDeviceStatus("8b40bdd25aac47a29bc67f866c3928d8", true, true)
if err != nil {
	fmt.Println(err.Error())
}
if err := r11.Error; err != nil {
	fmt.Println(*err)
} else {
	fmt.Println(*r11.Data)
}
```

## License
MIT licensed. See the LICENSE file for details.