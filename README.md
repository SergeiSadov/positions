This was initially a test task that I have decided to rewrite. It consists of two endpoints:

#### /summary:

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "domain": {
      "type": "string"
    },
    "positions_count": {
      "type": "integer"
    }
  },
  "required": [
    "domain",
    "positions_count"
  ]
}
```

#### /positions

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "domain": {
      "type": "string"
    },
    "positions": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "keyword": {
            "type": "string"
          },
          "position": {
            "type": "integer"
          },
          "url": {
            "type": "string"
          },
          "volume": {
            "type": "integer"
          },
          "results": {
            "type": "integer"
          },
          "cpc": {
            "type": "number"
          },
          "updated": {
            "type": "string",
            "format": "date"
          }
        }
      }
    }
  },
  "required": [
    "domain",
    "positions"
  ]
}
```

### Run:

Prerequisite: installed docker, kubernetes (minikube), helm

```sh
cp config.example.json config.json
sh build.sh
sh build_ek.sh
```

### Logging:

Wait until elk cluster is running  (check with command ``kubectl get pods``)

Go to: http://localhost:5601/app/kibana#/management/kibana/index_pattern?_g=()
Add patterns:
error*
info*

### Services:
- kibana: localhost:5601

### Endpoints:

- http://localhost:8000/positions
- http://localhost:8000/summary