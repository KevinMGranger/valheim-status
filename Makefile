.PHONY: all

all: web parse

web: valwho-web/valwho-web.go
	cd valwho-web && go build -o ../web

parse: valwho-parse/valwho-parse.go
	cd valwho-parse && go build -o ../parse

clean:
	rm -f web parse

install: all
	mkdir -p "$(DESTDIR)/usr/libexec/valwho" "$(DESTDIR)/usr/bin/"
	install -m 0755 invocation logs who parse web "$(DESTDIR)/usr/libexec/valwho/"
	ln -sf "../libexec/valwho/who" "$(DESTDIR)/usr/bin/valwho"