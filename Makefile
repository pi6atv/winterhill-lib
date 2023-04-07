GIT_COMMIT := $(shell git describe --tags | tr -d v)
GIT_DIRTY := $(if $(shell git status --porcelain | grep -qE .),~$(shell git rev-parse --short HEAD))
VERSION := $(GIT_COMMIT)$(GIT_DIRTY)
ARCH ?= amd64
PACKAGE := winterhill-web

.PHONY: all
all: clean package

test:
	go vet ./...
	go test -race ./...

$(PACKAGE): app/winterhill-web/dist test
	CGO_ENABLED=1 go build -o $(PACKAGE) app/$(PACKAGE)/main.go

app/winterhill-web/dist:
	cd web && npm install && npm run build && mv dist ../app/winterhill-web

.PHONY: clean
clean:
	rm -rf app/winterhill-web/dist $(PACKAGE) build

.PHONY: package
package: $(PACKAGE) app/winterhill-web/dist
	mkdir -p build build/opt/$(PACKAGE)/
	cp $(PACKAGE) build/opt/$(PACKAGE)
#	cp nginx-site.conf build/etc/nginx/sites-enabled/$(PACKAGE).conf
#	cp nginx-proxy.conf build/etc/nginx/snippets/$(PACKAGE)-proxy.conf
	cp systemd.service build/$(PACKAGE).service
#	cp config/$(PACKAGE).yaml build/opt/$(PACKAGE)/$(PACKAGE).yaml
#	cp -a web/dist  build/opt/$(PACKAGE)/web/apip
#	cp grafana-dashboard.json build/var/lib/grafana/dashboards/$(PACKAGE)-dashboard.json
#	cp prometheus.yaml build/etc/prometheus/targets/$(PACKAGE).yaml
	cd build && \
		fpm -s dir -t deb -n $(PACKAGE) -v "$(VERSION)" \
			--deb-systemd $(PACKAGE).service \
			--deb-systemd-enable --deb-systemd-auto-start --deb-systemd-restart-after-upgrade \
			-a $(ARCH) -m "Wim Fournier <debian@fournier.nl>" \
			.
