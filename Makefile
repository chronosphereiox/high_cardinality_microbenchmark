
ROOT_DIR := $(shell pwd)
TESTDATA_DIR := $(ROOT_DIR)/testdata
TESTDATA_PROM_DIR := $(TESTDATA_DIR)/prom
TESTDATA_M3_DIR := $(TESTDATA_DIR)/m3

clean:
	rm -rf $(TESTDATA_DIR)

setup:
	mkdir -p $(TESTDATA_DIR)

prom_generate: setup
	cd cmd/prom_generate && go run main/main.go -dir $(TESTDATA_PROM_DIR)

prom_benchindex: setup
	cd cmd/prom_benchindex && go run main/main.go -dir $(TESTDATA_PROM_DIR)/$(shell ls $(TESTDATA_PROM_DIR))

m3_generate: setup
	cd cmd/m3_generate && go run main/main.go -dir $(TESTDATA_M3_DIR)

m3_benchindex: setup
	cd cmd/m3_benchindex && go run main/main.go -dir $(TESTDATA_M3_DIR)
