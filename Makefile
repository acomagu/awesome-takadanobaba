.PHONY: all
all: target/restaurants.md

.INTERMEDIATE: genmark
genmark:
	go build ./tools/genmark

target/restaurants.md: genmark src/restaurants.toml
	./genmark < src/restaurants.toml > target/restaurants.md
