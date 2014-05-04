GODOC_REV=d8700e870053

WORKSPACE:=$(shell pwd)/_tmp
export GOPATH=${WORKSPACE}/sub
export PATH:=${GOPATH}/bin:${PATH}

_GOROOT:=${WORKSPACE}/go
GOREPO=code.google.com/p
APP=${GOPATH}/src/godoc

.PHONY: all
all: update build_doc

.PHONY: clean
clean:
	rm -rf ${APP}
	rm -rf ${_GOROOT}/*

build_doc: ${tr_status} ${APP}/doc/godoc.zip

tr_status:=${_GOROOT}/status/translate.html
tr_tool:=github.com/atotto/gophersjp-go-util/cmd/translate-status
${tr_status}: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	echo ${GOPATH}
	go get -u ${tr_tool}
	mkdir -p ${_GOROOT}/status
	translate-status -docroot=./ -goroot=${_GOROOT} -o=$@

.PHONY: update_doc
update_doc: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	cd ${_GOROOT};\
	rm -rf ${_GOROOT}/src/pkg/${GOREPO};\
	hg checkout tip -C;\
	hg --config extensions.purge= clean

	#cd ${GOPATH}/src; find code.google.com/p/go.tools -type d -not -path '*/.hg/*'\
	#	-exec mkdir -p '${_GOROOT}/src/pkg/{}' ';'
	#cd ${GOPATH}/src; find code.google.com/p/go.tools -type f -not -path '*/.hg/*'\
	#	-exec cp '{}' '${_GOROOT}/src/pkg/{}' ';'
	find */ -type d -not -path '.*' -not -path '_tmp/*' \
		-exec mkdir -p '${_GOROOT}/{}' ';'
	find */ -type f -not -path '.*' -not -path '_tmp/*' \
		-exec cp '{}' '${_GOROOT}/{}' ';'

${APP}/doc/godoc.zip: update_doc
	which zip || sudo apt-get install zip
	mkdir -p ${APP}/doc
	zip -q -r $@ ${_GOROOT}/*

.PHONY: run
run: update_doc
	GOPATH="";godoc -http=:6060 -play -goroot=${_GOROOT}

{APP}/.git: godep
	cd ${GOPATH}/src/${GOREPO}/go.tools; hg checkout ${GODOC_REV}

	mkdir -p ${APP}
	cd ${APP};\
	cp -r ${GOPATH}/src/${GOREPO}/go.tools/cmd/godoc/* ${APP};\
	git init;\
	git add -A;\
	git commit -a -m "Add godoc app";\
	echo 'web: godoc -http=:$$PORT -play -zip=doc/godoc.zip -goroot=${_GOROOT}' > Procfile;\
	git add -A;\
	git commit -a -m "Add Procfile";\
	godep save;\
	git add -A .;\
	git commit -a -m "Save godep"

deploy: build_doc {APP}/.git
	cd ${APP};\
	git remote add heroku-deploy ${HEROKU};\
	git push -f heroku-deploy master;\

update: ${_GOROOT} ${GOPATH}/src/${GOREPO}/go.tools
	cd ${_GOROOT}; hg pull; hg update tip
	cd ${GOPATH}/src/${GOREPO}/go.tools; hg pull; hg update tip

${_GOROOT}:
	hg clone -u tip https://code.google.com/p/go $@

${GOPATH}/src/${GOREPO}/go.tools:
	hg clone -u tip https://code.google.com/p/go.tools $@

godep:
	which godep ||go get -u github.com/kr/godep

