# Crystal-Snowflake

A micro-backend for generate unique ids using snowflake algorithm.


## API Reference

#### Get Unique Id

```http
  GET /generate-id
```

The unique id will be received in the "X-Snowflake-Id" response header.


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`NODE_SIZE`: The node number, default is 1

`NODE_BITS`: Number of bits dedicated to the node, default is 10

`STEP_BITS`: Number of bits dedicated to the step, default is 12

`SERVER_PORT`: The server port, default is 8080

`SERVER_ADDRESS`: The server address, default is 0.0.0.0
## Execution

To execute this project run

```bash
  go run main.go
```

Or you can run the docker file with:
```bash
> docker build -f Dockerfile . -t image-name
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

