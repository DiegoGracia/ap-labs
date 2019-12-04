CC       = gcc
CFLAGS   = -Wall
LDFLAGS  = -pthread
OBJFILES = logger.o multiplier.o
TARGET   = multiplier

all: $(TARGET)

$(TARGET) : $(OBJFILES)
	$(CC) $(CFLAGS) -o $(TARGET) $(OBJFILES) $(LDFLAGS)

clean:
	rm -f $(OBJFILES) $(TARGET) *~

include ../common.mk
