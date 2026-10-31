# backend

## Environment Variables
```env
DATA_DIR=<dir> # Directory to store data (sqlite db and uploads)
FRONTEND_URL=<url to frontend>
BACKEND_URL=<url of the application>
DEFAULT_REDIRECT=<url to redirect to upon login>
```

## Development
> [!NOTE]
> This is a MUCH more technical description than for the frontend

### Running the backend without container
```bash
go run main.go
```
This will run it in development mode with a sqlite backend. Setting env vars will override the config file settings
