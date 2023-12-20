build:
	bazel build //...
build/%: 
	bazel build //cmd/$*

run/%: 
	bazel run //cmd/$*

test:
	bazel test //...

clean:
	bazel clean

deps-update:
	bazel run //:gazelle -- update-repos -from_file=go.mod