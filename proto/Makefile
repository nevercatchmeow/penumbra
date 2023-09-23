# CUSTOMIZE: Change the following variables to suit your needs
CODE_DIR := generated/code
DOC_DIR := generated/doc
TMP_DIR := github.com/nevercatchmeow/penumbra/proto

# DON'T EDIT BELOW THIS LINE
DIRS_TO_REMOVE := $(shell path_to_remove=$(TMP_DIR); while [ "$$path_to_remove" != "." ] ; do  echo -n "$$path_to_remove " ; path_to_remove=`dirname $$path_to_remove`;  done)

.PHONY: protoc
protoc:
	mkdir -p $(CODE_DIR) && mkdir -p $(DOC_DIR) && \
	protoc -I ./protos -I ./pb --go_out=./ --validate_out="lang=go:." --doc_out=./$(DOC_DIR) --doc_opt=html,index.html pb/*.proto && \
	rm -rf $(CODE_DIR)/* && mv $(TMP_DIR)/* $(CODE_DIR) && \
	for DIR in $(DIRS_TO_REMOVE) ; do rm -rf $$DIR ; done
