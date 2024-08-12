
# Banner File Checksum Verifier

## Overview

This project is designed to verify the integrity of banner files using SHA-256 checksums. It computes the checksum of a specified banner file and compares it against a known checksum to ensure that the file has not been altered. This is useful for validating the authenticity of banner files before they are used in your application.

## Features

- **Checksum Verification**: Computes the SHA-256 checksum of a banner file and compares it with a known checksum.
- **Support for Multiple Files**: Can verify multiple banner files using a map of file paths and checksums.
- **Error Handling**: Provides feedback on whether the checksum is valid or if an error occurred during verification.

## File Structure

```
.
├── banners
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── main.go
└── README.md
```

- `banners/standard.txt`: The standard banner file.
- `banners/shadow.txt`: The shadow banner file.
- `banners/thinkertoy.txt`: The thinkertoy banner file.
- `main.go`: The main Go program that verifies the checksums.
- `README.md`: This file.

## Usage

### Prerequisites

- Go 1.16 or higher installed on your machine.
- Your banner files should be placed in the `banners/` directory.

### Running the Program

1. Clone the repository and navigate to the project directory.


2. Update the `main.go` file with the correct file paths and known checksums if they differ.

3. Run the program:

   ```bash
   go run main.go
   ```

### Example Output

If the checksums match:

```bash
Checksum for banners/standard.txt is valid.
Checksum for banners/shadow.txt is valid.
Checksum for banners/thinkertoy.txt is valid.
```

If a checksum does not match:

```bash
Checksum for banners/standard.txt is invalid.
```

If there is an error reading a file:

```bash
Error verifying banners/standard.txt: open banners/standard.txt: no such file or directory
```

## Modifying the Program

### Verifying a Single Banner File

If you only have one banner file, you can simplify the program by directly specifying the file path and checksum without using a map. See the simplified version in the `main.go` file.

### Adding New Banner Files

To add new banner files:

1. Place the new banner file in the `banners/` directory.
2. Compute the SHA-256 checksum of the new file using the method described below.
3. Update the `knownChecksums` map in `main.go` with the new file path and its checksum.

### Computing the SHA-256 Checksum

To compute the SHA-256 checksum of a file, you can use a tool like `shasum`:

```bash
shasum -a 256 banners/newbanner.txt
```

Copy the output checksum and add it to the `knownChecksums` map in `main.go`.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributions

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For any questions or support, please contact [your email address].

---

This `README.md` provides a clear and concise guide to your project, including its purpose, usage, and how to modify it. You can customize it further based on your specific needs and preferences.