.PHONY: backend frontend clean

BACKEND_SRC := $(shell find backend -name "*.go")

all: backend frontend

build/system: $(BACKEND_SRC)
	@echo "[+] Building backend..."
	mkdir -p build
	cd backend && go build -o ../build/system .
	@echo "[✔] Backend build finished"

backend:
	@if [ -f build/system ]; then \
		if [ -z "$$(find backend -name '*.go' -newer build/system)" ]; then \
			echo "[✔] backend is up-to-date, no build needed"; \
			exit 0; \
		fi; \
	fi; \
	$(MAKE) build/system

frontend:
	@echo "[+] Building frontend..."
	@echo "[✔] Frontend build finished"

clean:
	rm -rf build