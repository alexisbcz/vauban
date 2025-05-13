setup:
	@echo "🔧 Installing Reflex for hot reload"
	go install github.com/cespare/reflex@latest
	@echo "📦 Installing dependencies..."
	go mod tidy

css:
	npx @tailwindcss/cli -i ./pkg/views/styles.css -o ./pkg/public/assets/styles.css --watch

dev:
	reflex --start-service -- sh -c 'go run ./cmd/http'