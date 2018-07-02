GO := go
GOBUILD := GOOS=linux $(GO) build
GOTEST := $(GO) test

HANDLER_DIR := handler/
HANDLER_ONE := one
HANDLER_TWO := two

define build_handler
		cd $(HANDLER_DIR)$(1) && \
		$(GOBUILD) -o ../../$(1) && \
		cd ../../ && \
		zip $(1).zip ./$(1) && \
		rm $(1)
endef

build:
		$(call build_handler,$(HANDLER_ONE))
		$(call build_handler,$(HANDLER_TWO))
		mv *.zip pkg

clean:
		rm -f pkg/*

dep:
		dep ensure