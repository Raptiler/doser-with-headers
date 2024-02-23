# doser-with-headers

`doser-with-headers` is a versatile command-line tool designed for conducting load testing on web applications. This tool allows you to easily send high volumes of GET or POST requests to a specified URL, making it an invaluable resource for developers looking to test the robustness and scalability of their web applications under heavy load. Additionally, it supports custom HTTP headers, enabling more complex testing scenarios, such as API testing with authentication tokens or custom content types.

## Features

- **Custom HTTP Headers:** Specify custom headers for your requests to test different scenarios and configurations.
- **GET and POST Requests:** Choose between sending GET or POST requests depending on your testing requirements.
- **Payload Support:** Include a data payload with your POST requests to simulate form submissions or API requests with body content.
- **Concurrency Control:** Adjust the number of concurrent threads to simulate varying levels of load and stress on your application.
- **Simplicity and Flexibility:** Easy to use with a simple command-line interface, allowing for quick adjustments and testing in various environments.

## Usage

To use `doser-with-headers`, you can specify various command-line options to tailor your load testing. Below are the options available:

- `-H string`: Specify custom headers for your requests. Usage: `-H 'Key: Value;Key2: Value2'`
- `-d string`: Specify data payload for POST request.
- `-g string`: Specify a GET request. Usage: `-g '<url>'`
- `-p string`: Specify a POST request. Usage: `-p '<url>'`
- `-t int`: Specify the number of threads to be used for sending requests concurrently. The default is 500.

### Examples

**Sending Requests**

```shell
./doser -g 'http://example.com' -t 100
./doser -p 'http://example.com/api/submit' -d '{"key":"value"}' -H 'Content-Type: application/json' -t 200
./doser -g 'http://example.com' -H 'Authorization: Bearer token;Custom-Header: value' -t 150
```
### Installation
```shell
git clone https://github.com/Raptiler/doser-with-headers.git
cd doser-with-headers
go build
```

