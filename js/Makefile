SHELL := /bin/bash -euo pipefail
export PATH := $(V23_ROOT)/release/go/bin:node_modules/.bin:$(V23_ROOT)/third_party/cout/node/bin:$(PATH)

# This target causes any target files to be deleted if the target task fails.
# This is especially useful for browserify, which creates files even if it
# fails.
.DELETE_ON_ERROR:

# Default browserify options: use sourcemaps.
BROWSERIFY_OPTS := --debug
# Names that should not be mangled by minification.
RESERVED_NAMES := 'context,ctx,callback,cb,$$stream,serverCall'
# Don't mangle RESERVED_NAMES, and screw ie8.
MANGLE_OPTS := --mangle [--except $(RESERVED_NAMES) --screw_ie8 ]
# Don't remove unused variables from function arguments, which could mess up signatures.
# Also don't evaulate constant expressions, since we rely on them to conditionally require modules only in node.
COMPRESS_OPTS := --compress [ --no-unused --no-evaluate ]

# Browserify and extract sourcemap, but do not minify.
define BROWSERIFY
	mkdir -p $(dir $2)
	browserify $1 $(BROWSERIFY_OPTS) | exorcist $2.map > $2
endef

# Browserify, minify, and extract sourcemap.
define BROWSERIFY-MIN
	mkdir -p $(dir $2)
	browserify $1 $(BROWSERIFY_OPTS) --g [ uglifyify $(MANGLE_OPTS) $(COMPRESS_OPTS) ] | exorcist $2.map > $2
endef

node_modules: package.json
	npm prune
	npm install
	# Link Vanadium from V23_ROOT.
	rm -rf ./node_modules/vanadium
	cd "$(V23_ROOT)/release/javascript/core" && npm link
	npm link vanadium
	touch node_modules

src/gen-vdl: $(shell find ../go/src/ -name "*.vdl")
	v23 run vdl generate --lang=javascript --js-out-dir=src/gen-vdl v.io/x/media_sharing

build/bundle.js: src/index.js $(shell find src/ -name "*.js") src/gen-vdl node_modules
	mkdir -p build
ifndef NOMINIFY
	$(call BROWSERIFY,$<,$@)
else
	$(call BROWSERIFY-MIN,$<,$@)
endif

build/index.html: public/index.html
	mkdir -p build
	cp $< $@

build/style.css: public/style.css
	mkdir -p build
	cp $< $@

build: build/bundle.js build/index.html build/style.css

serve: build
	static build

.PHONY: lint
lint: node_modules
	jshint .

.PHONY: clean
clean:
	rm -rf \
		build \
		node_modules
