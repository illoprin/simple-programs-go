CC=go build
TARGET=app.exe
SRC_DIR=src

rpn-calc:
	$(CC) -o $(TARGET) $(SRC_DIR)/rpn-calculator/*.go
	$(TARGET)

list:
	$(CC) -o $(TARGET) $(SRC_DIR)/linked-list/*.go
	$(TARGET)

img_gen:
	$(CC) -o $(TARGET) $(SRC_DIR)/img_gen/*.go
	$(TARGET)

unsafe:
	$(CC) -o $(TARGET) $(SRC_DIR)/unsafe_ptrs/*.go
	$(TARGET)
