GODOC_REV=11bd89d4e1e9a84d4d66c88b7eccfecd66e8f6c9
GO_VER=go1.3

WORKSPACE:=$(shell pwd)/_tmp
export GOPATH=${WORKSPACE}/sub
export PATH:=${GOPATH}/bin:${PATH}

_GOROOT:=${WORKSPACE}/go
GOREPO=code.google.com/p
APP=${GOPATH}/src/godoc
APP_CONTENTS=${APP}/doc/godoc.zip

.PHONY: all
all: update build_doc

.PHONY: clean
clean:
	rm -rf ${APP}
	rm -rf ${_GOROOT}/*

tr_status:=${_GOROOT}/status/translate.html
tr_tool:=github.com/atotto/gophersjp-go-util/cmd/translate-status
${tr_status}: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	go get -u ${tr_tool}
	mkdir -p ${_GOROOT}/status
	translate-status -docroot=$(shell pwd) -goroot=${_GOROOT} -o=$@

build_doc: update_doc ${tr_status} ${APP_CONTENTS}

.PHONY: update_doc
update_doc: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	#cd ${GOPATH}/src; find code.google.com/p/go.tools -type d -not -path '*/.hg/*'\
	#	-exec mkdir -p '${_GOROOT}/src/{}' ';'
	#cd ${GOPATH}/src; find code.google.com/p/go.tools -type f -not -path '*/.hg/*'\
	#	-exec cp '{}' '${_GOROOT}/src/{}' ';'
	find */ -type d -not -path '.*' -not -path '_tmp/*' \
		-exec mkdir -p '${_GOROOT}/{}' ';'
	find */ -type f -not -path '.*' -not -path '_tmp/*' \
		-exec cp '{}' '${_GOROOT}/{}' ';'

${APP_CONTENTS}: update_doc
	which zip || sudo apt-get install zip
	mkdir -p ${APP}/doc
	cd ${_GOROOT}/../;\
	zip -q -r $@ go/*

public:
	cp _robots.txt ${_GOROOT}/robots.txt
	cd ${_GOROOT}/../;\
	zip -f -r ${APP_CONTENTS} go/*

.PHONY: run
run: update_doc
	GOPATH="";godoc -http=:6060 -play -goroot=${_GOROOT}

.PHONY: godep
godep:
	which godep ||go get -u github.com/kr/godep

${APP}/godeps: godep
	cd ${GOPATH}/src/${GOREPO}/go.tools; hg checkout ${GODOC_REV}

	mkdir -p ${APP}
	cd ${APP};\
	cp -r ${GOPATH}/src/${GOREPO}/go.tools/cmd/godoc/* ${APP};\
	git init;\
	git add -A;\
	git commit -a -m "Add godoc app";\
	echo 'web: godoc -http=:$$PORT -play -zip=doc/godoc.zip -goroot=/go' > Procfile;\
	git add -A;\
	git commit -a -m "Add Procfile";\
	godep save;\
	git add -A .;\
	git commit -a -m "Save godep";\
	cat Godeps/Godeps.json | jq -a '.GoVersion = "${GO_VER}"' > Godeps.json;\
	mv Godeps.json Godeps/Godeps.json;\
	git commit -a -m "go version ${GO_VER}"


deploy: ${APP_CONTENTS} ${APP}/godeps
	cd ${APP};\
	git remote add heroku-deploy ${HEROKU};\
	git push -f heroku-deploy master;\

update: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	cd ${_GOROOT};\
	rm -rf ${_GOROOT}/src/${GOREPO};\
	hg checkout tip -C;\
	hg --config extensions.purge= clean;\
	hg pull; hg update tip
	cd ${GOPATH}/src/${GOREPO}/go.tools; hg pull; hg update tip

${_GOROOT}:
	hg clone -u tip https://code.google.com/p/go $@

${GOPATH}/src/${GOREPO}/go.tools:
	hg clone -u tip https://code.google.com/p/go.tools $@

