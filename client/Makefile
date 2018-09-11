CSC		= /cygdrive/c/windows/microsoft.net/framework/v4.0.30319/csc.exe
TARGET	= WaifuShareClient.exe
SRC		= main.cs

CSC_FLAGS	= /r:System.Net.Http.dll /nologo 

DEBUG_FLAGS		= 
RELEASE_FLAGS	= 

.PHONY: debug
debug: CSC_FLAGS+=$(DEBUG_FLAGS)
debug: all

.PHONY: release
release: CSC_FLAGS+=$(RELEASE_FLAGS)
release: all

all: $(TARGET)
$(TARGET): $(SRC)
	$(CSC) $(CSC_FLAGS) /out:$(TARGET) $(SRC) | iconv -f cp932 -t utf-8

.PHONY: clean
clean:
	rm $(TARGET)

