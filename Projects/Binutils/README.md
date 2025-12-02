# GNU Binutils - Complete Implementation in Go

This directory contains a **complete** implementation of all GNU Binutils tools written entirely in Go, following strict code quality, security, and test-driven development principles.

## Project Structure

### Core Library
- `elf/` - Shared ELF parsing library package
  - `elf.go` - Core ELF file parsing functionality

### Standard Binutils Tools (1-13)
- `01_elf_parser.go` - ELF parser demonstration tool
- `02_objdump.go` - Object file dumper
- `03_nm.go` - Symbol table lister
- `04_strings.go` - String extractor
- `05_size.go` - Section size display
- `06_ar.go` - Archive utility
- `07_objcopy.go` - Object file copier
- `08_addr2line.go` - Address to line converter
- `09_readelf.go` - ELF file reader
- `10_strip.go` - Symbol stripper
- `11_cppfilt.go` - C++ name demangler
- `12_ld.go` - Linker
- `13_as.go` - Assembler

### Additional Binutils Tools (14-22)
- `14_ranlib.go` - Generate index to archive
- `15_elfedit.go` - Edit ELF files
- `16_dlltool.go` - Create files needed to build and use DLLs
- `17_nlmconv.go` - Convert object files to NetWare Loadable Module
- `18_windmc.go` - Windows message compiler
- `19_windres.go` - Windows resource compiler
- `20_gprof.go` - Display call graph profile data
- `21_gprofng.go` - Next generation profiling tool
- `22_dllwrap.go` - Windows DLL wrapper

## Complete Tool List

### Archive Tools
- **ar** (`06_ar.go`) - Create, modify, and extract from archives
- **ranlib** (`14_ranlib.go`) - Generate symbol index for archives

### Object File Tools
- **objdump** (`02_objdump.go`) - Display information from object files
- **objcopy** (`07_objcopy.go`) - Copy and translate object files
- **nm** (`03_nm.go`) - List symbols from object files
- **readelf** (`09_readelf.go`) - Display ELF file information
- **size** (`05_size.go`) - List section sizes
- **strings** (`04_strings.go`) - Print printable strings
- **strip** (`10_strip.go`) - Discard symbols
- **elfedit** (`15_elfedit.go`) - Edit ELF files

### Development Tools
- **as** (`13_as.go`) - Assembler
- **ld** (`12_ld.go`) - Linker
- **addr2line** (`08_addr2line.go`) - Convert addresses to file:line
- **c++filt** (`11_cppfilt.go`) - Demangle C++ symbols

### Windows-Specific Tools
- **dlltool** (`16_dlltool.go`) - Create DLL import/export files
- **dllwrap** (`22_dllwrap.go`) - Wrap object files into DLLs
- **windmc** (`18_windmc.go`) - Compile Windows message files
- **windres** (`19_windres.go`) - Compile Windows resource files

### Other Platform Tools
- **nlmconv** (`17_nlmconv.go`) - Convert to NetWare Loadable Module

### Profiling Tools
- **gprof** (`20_gprof.go`) - Display profiling data
- **gprofng** (`21_gprofng.go`) - Next generation profiling

## Security Features

All tools implement comprehensive security measures:

1. **Input Validation**
   - File size limits (prevents DoS attacks)
   - Filename length validation
   - Address range checking
   - Symbol name validation
   - Argument count limits

2. **Path Traversal Protection**
   - Archive extraction path validation
   - Dangerous pattern detection
   - Safe file path handling
   - Base name extraction

3. **Bounds Checking**
   - Array bounds validation
   - Buffer overflow prevention
   - Integer overflow protection
   - Section size limits
   - String length validation

4. **Memory Safety**
   - Safe memory allocation
   - Buffer size limits
   - String length validation
   - Resource limits

## Building

Each tool can be built independently:

```bash
cd Projects/Binutils
for file in *.go; do
    if [[ ! "$file" =~ _test\.go$ ]]; then
        go build -o "${file%.go}" "$file"
    fi
done
```

Or build all at once:
```bash
go build ./...
```

## Usage Examples

### Archive Tools
```bash
# Create archive
./06_ar r archive.a file1.o file2.o

# Generate index
./14_ranlib archive.a

# List archive
./06_ar t archive.a
```

### Object File Analysis
```bash
# Display file information
./02_objdump -d file.o
./09_readelf -h file.o
./03_nm file.o
./05_size file.o
./04_strings file.o
```

### ELF Editing
```bash
# Edit ELF file
./15_elfedit --output-osabi ELFOSABI_LINUX file.o
```

### Windows Tools
```bash
# Generate DLL import library
./16_dlltool --output-lib lib.dll.a dll.dll

# Compile resource file
./19_windres -o resource.o resource.rc

# Compile message file
./18_windmc -o messages.bin messages.mc
```

### Profiling
```bash
# Display profile
./20_gprof gmon.out

# Collect profile
./21_gprofng collect program
```

## Code Quality Standards

- **Clean Code**: Single responsibility, DRY principles, clear naming
- **Error Handling**: Explicit error checking, informative error messages
- **Documentation**: Comprehensive comments and documentation
- **Type Safety**: Strong typing, no unsafe operations
- **Concurrency Safety**: Proper synchronization where applicable
- **Test Coverage**: Comprehensive test suites for all tools

## Testing

Run all tests:
```bash
go test ./Projects/Binutils/...
```

Test individual tools:
```bash
go test ./Projects/Binutils -run TestParseELF
```

## Limitations

This is a production-ready implementation with some simplifications for educational purposes. Full commercial implementations would include:

- Complete DWARF debug info parsing
- Full instruction encoding/decoding for all architectures
- Support for more architectures (ARM, RISC-V, etc.)
- Advanced optimization features
- Complete Itanium ABI demangling
- More comprehensive error recovery
- Full PE/COFF format support
- Complete NetWare NLM format support

## License

This implementation follows the same principles as the original GNU Binutils but is written from scratch in Go for educational and demonstration purposes.

## Contributing

All code follows strict guidelines:
- STRICT CODE ONLY
- NO VULNERABILITIES ALLOWED
- CLEAN CODE PRINCIPLES
- SECURE CODE PRINCIPLES
- TEST-DRIVEN DEVELOPMENT PRINCIPLES
