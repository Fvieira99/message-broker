To run the message broker follow these steps: 

1 - First set up the rabbitmq environment on docker-compose file and start the container with:
```bash
docker compose up
```

2 - Start the payments consumer service on a separate terminal by running: 

```bash
go run ./payments
```

3 - Produce as many messages as you want, also on a separate terminal, with: 

```bash
go run ./orders
```

Then all the messages produced by the orders service must be consumed by the payments service.