DICT_FILES = $(filter-out dict/user.dict.utf8, $(wildcard dict/*.utf8))
DICT_H_FILES = $(addprefix deps/cppjieba/dict/, $(notdir $(patsubst %.utf8,%.h, $(DICT_FILES))))

build: $(DICT_H_FILES)
	go build -v -a

deps/cppjieba/dict/%.h: dict/%.utf8
	xxd -i $< > $@

clean:
	rm -f $(DICT_H_FILES)

.PHONY: build clean
