CC=go run
TARGET=app
SRC_DIR=src

rpn-calc:
	$(CC) $(SRC_DIR)/rpn-calculator/*.go
	$(TARGET)

list:
	$(CC) $(SRC_DIR)/linked-list/*.go
	$(TARGET)

img_gen:
	$(CC) $(SRC_DIR)/img_gen/*.go
	$(TARGET)

unsafe:
	$(CC) $(SRC_DIR)/unsafe_ptrs/*.go
	$(TARGET)

binary-read-write:
	$(CC) $(SRC_DIR)/to-binary/*.go
	$(TARGET)

get-collision-sha256:
	$(CC) $(SRC_DIR)/get-collision-sha256/*.go
	$(TARGET)
