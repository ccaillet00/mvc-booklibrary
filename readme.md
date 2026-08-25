Use the App: 
# Local terminal app (default)
go run .

# Local HTTP server
go run . -mode http

# Both at once
go run . -mode both

# Via environment variable
APP_MODE=http go run .

Docker Integration: 

Dockerfile — sets ENV APP_MODE=http as default and adds EXPOSE 8080

compose.yaml — exposes port 8080 and reads APP_MODE from your environment with a fallback: APP_MODE=${APP_MODE:-http}

docker compose up                    # HTTP mode (default)
APP_MODE=both docker compose up      # CLI + HTTP (with tty/stdin already enabled)