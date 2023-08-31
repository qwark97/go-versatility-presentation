# Design specification

## Acceptance criteria
- [x] service reads data from peripherals (e.g. some microcontrollers) according to internal scheduler
- [x] service stores data in memory
- [ ] service sends data to Home Assistant
- [x] all above (source of data, where it should be send, how often etc.) is based on the configuration stored as JSON file
- [x] configuration should be maintainable through API
- [x] service supports hot-reloading of the configuration
- [x] service allows to verify if configuration works as expected
- [x] service allows to read data read by the data source
- ...

## Comunication
Incoming communication goes through HTTP calls. 
Outcoming communication goes only to Home Assistant therefore it depends on how HA accepts data and will be clarified later.
Communication with data sources (e.g. microcontrollers) goes through HTTP calls

## API

### POST /api/v1/configuration
Allows to add new configuration (one) for some data source
#### Body:
```json
{
    "method": string,
    "addr": string,
    "frequency": string,
    "description": string,
    "unit": string
}
```
#### Returns:
HTTP 201
```json
```
HTTP 400
When passed `id` is invalid `UUID`
```json
```
HTTP 500
When failed to add configuration
```json
```
### POST /api/v1/configuration/reload
Allows to reload the configuration manually e.g. after manual modification of the configuration JSON file
#### Returns:
HTTP 204
```json
```
HTTP 500
When failed to reload configuration
```json
```

### GET /api/v1/configuration/{id}
Allows to get configuration (one) by ID
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```json
{
    "id": UUID,
    "method": string,
    "addr": string,
    "frequency": string,
    "description": string,
    "unit": string
}
```
HTTP 400
When passed `id` is invalid `UUID`
```json
```
HTTP 404
When did not find configuration by given `id`
```json
```
HTTP 500
When failed to get configuration by `id`
```json
```

### DELETE /api/v1/configuration/{id}
Allows to delete configuration (one) by ID. If ID not exists, response will be valid
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 204
```json
```
HTTP 400
When passed `id` is invalid `UUID`
```json
```
HTTP 500
When failed to delete configuration by `id`
```json
```

### GET /api/v1/configurations
Allows to get configurations (all)
#### Returns:
HTTP 200
```json
[
    {
        "id": UUID,
        "method": string,
        "addr": string,
        "frequency": string,
        "description": string,
        "unit": string
    }
]
```
HTTP 500
When failed to get configurations
```json
```

### GET /api/v1/data-source/{id}/verify
Allows to check if it is possible to fetch data from data source. Success means that endpoint (`addr`) was reachable and returned something
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```json
{
    "success": bool
}
```
HTTP 400
When passed `id` is invalid `UUID`
```json
```
HTTP 500
When failed to verify configuration by `id`
```json
```

### GET /api/v1/data-source/{id}
Allows to get all read data from given data source
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```text
"<readData><unit> (<description>)"
```
HTTP 400
When passed `id` is invalid `UUID`
```json
```
HTTP 500
When failed to read last reading of given periferal
```json
```