# Design specification

## Acceptance criteria
- [ ] service reads data from peripherals (e.g. some microcontrollers) according to internal scheduler
- [ ] service stores data in sqllite DB
- [ ] service sends data to Home Assistant
- [ ] all above (source of data, where it should be send, how often etc.) is based on the configuration stored as JSON file
- [ ] configuration should be maintainable through API
- [ ] service supports hot-reloading of the configuration
- [ ] service allows to verify if configuration works as expected
- [ ] service allows to read data read by the data source
- ...

## Comunication
Incoming communication goes thgough HTTP calls. 
Outcoming communication goes only to Home Assistant therefore it depends on how HA accepts data and will be clarified later.
Communication with data sources (e.g. microcontrollers) depends on the source itself, therefore will be defined through the configuration

## API

### POST /api/v1/configuration
Allows to add new configuration (one) for some data source
#### Body:
```json
{} // TBA
```
#### Returns:
HTTP 200
```json
{
    "id": UUID as string
}
```
HTTP 400
```json
{
    "msg": string
}
```

### POST /api/v1/configuration/reload
Allows to reload the configuration manually e.g. after manual modification of the configuration JSON file
#### Returns:
HTTP 204
```json
```
HTTP 400
```json
{
    "msg": string
}
```

### GET /api/v1/configuration/{id}
Allows to get configuration (one) by ID
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```json
{} // TBA, the same as JSON expected during inserting 
```
HTTP 404
```json
{
    "msg": string
}
```

### DELETE /api/v1/configuration/{id}
Allows to delete configuration (one) by ID
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 204
``` 
```
HTTP 404
```json
{
    "msg": string
}
```

### GET /api/v1/configurations
Allows to get configurations (all)
#### Returns:
HTTP 200
```json
[
    {} // TBA, the same as JSON expected during inserting 
]
```

### GET /api/v1/data-source/{id}/verify
Allows to check if it is possible to fetch data from data source
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```json
{} // TBA, the same as JSON expected during inserting 
```
HTTP 404
```json
{
    "msg": string
}
```

### GET /api/v1/data-source/{id}
Allows to get all read data from given data source
#### URL query
- `{id}` - ID of the configuration, valid UUID string
#### Returns:
HTTP 200
```json
[
    {} // object, depends on what configuration defines
]
```
HTTP 404
```json
{
    "msg": string
}
```