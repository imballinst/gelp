last_revision := $(shell git rev-list --tags --max-count=1)
last_tag := $(shell git describe --tags $(last_revision))

# If the first argument is "run"...
ifeq (start,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

start:
	@echo $(last_revision) $(last_tag)
	@go run -ldflags "-X main.version=$(last_tag)" main.go $(RUN_ARGS)
# if err != nil {
# 	panic(err)
# }

# lastTag, err := ExecCommand("git describe --tags astRevision)
