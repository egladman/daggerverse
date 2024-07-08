GIT ?= git

# Find all dagger modules in the git root
rwildcard = $(wildcard $1$2) $(foreach d,$(wildcard $1*),$(dir $(call rwildcard,$d/,$2)))
DAGGER_MODULES ?= $(patsubst %/,%,$(call rwildcard,*,dagger.json))

# Semantic version of the dagger module (i.e., 1.0.0 1.2.3)
ifeq ($(VERSION),)
$(error version not specified)
endif

# Dagger modules are independently versioned by prefixing the tag with the subpath
#   https://docs.dagger.io/manuals/developer/publish-modules/#semantic-versioning

.PHONY: $(DAGGER_MODULES)
$(DAGGER_MODULES): GIT_REMOTE ?= origin
$(DAGGER_MODULES):
	$(GIT) tag $@/v$(VERSION)
	$(GIT) push $(GIT_REMOTE) $@/v$(VERSION)
