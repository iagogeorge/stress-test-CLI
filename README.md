
# Load Test CLI

## Objective

Create a CLI system in Go to perform load tests on a web service. The user should provide the service URL, the total number of requests, and the number of concurrent calls.

The system should generate a report with specific information after the tests are executed.

## CLI Parameters

- `--url`: URL of the service to be tested.
- `--requests`: Total number of requests.
- `--concurrency`: Number of simultaneous calls.

## Test Execution

- Perform HTTP requests to the specified URL.
- Distribute requests according to the defined concurrency level.
- Ensure that the total number of requests is met.

## Report Generation

Generate a report at the end of the tests containing:
- Total time spent in execution.
- Total number of requests made.
- Number of requests with HTTP 200 status.
- Distribution of other HTTP status codes (such as 404, 500, etc.).

## Application Execution

You can use this application by making a call via Docker. For example:
```bash
docker run <your docker image> --url=http://google.com --requests=1000 --concurrency=10
```


## Building and Running the Project

### Build and Run the Load Test CLI

Navigate to the directory where the `Dockerfile` for the CLI is located and run:

```bash
docker build -t load-tester .
docker run load-tester --url=http://google.com --requests=100 --concurrency=10
```
```
