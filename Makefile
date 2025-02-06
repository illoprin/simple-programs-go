CC=go build
TARGET=bin/app.exe
SRC_DIR=src

rpn-calc:
	$(CC) -o $(TARGET) $(SRC_DIR)/rpn-calculator/*