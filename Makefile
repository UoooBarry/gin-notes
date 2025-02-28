dev:
	go run . & \
	pnpm build:css &
	pnpm watch:js
