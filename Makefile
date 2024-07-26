live/templ:
	templ generate --watch --proxy="http://localhost:2000" --open-browser=false -v

live/server:
	go run github.com/air-verse/air@v1.52.3 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--misc.clean_on_exit true

live/tailwind:
	npx tailwindcss -i ./input.css -o ./assets/styles.css --minify --watch

live/sync_assets:
	go run github.com/air-verse/air@v1.52.3 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

live: 
	make -j5 live/templ live/server live/tailwind live/sync_assets

