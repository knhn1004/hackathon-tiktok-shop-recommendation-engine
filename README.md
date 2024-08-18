

## Development
Start
```
docker-compose -f docker-compose.dev.yml up --build --force-recreate
```

Logs
```
docker-compose -f docker-compose.dev.yml logs -f frontend
```

Down
```
docker-compose -f docker-compose.dev.yml down
```

db:seed (important to see items in API endpoint)
```
docker-compose -f docker-compose.dev.yml exec api-node make seed
```