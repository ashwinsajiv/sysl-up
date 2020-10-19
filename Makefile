SYSLGO_SYSL = sysl/bff.sysl
SYSLGO_PACKAGES = bff
SYSLGO_APP.bff = Bff
-include local.mk
include codegen.mk
all: copy
w = external_types_up.go
x = external_types_bff.go
y = gen/pkg/servers/bff/up
z = gen/pkg/servers/bff
copy:
	cp -rf $(w) $(y)
	cp -rf $(x) $(z)
