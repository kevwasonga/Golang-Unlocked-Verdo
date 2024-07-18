

```markdown
# Search and Replace

This program takes 3 arguments and performs a search and replace operation on the first argument (a string). It replaces all occurrences of the character specified by the second argument with the character specified by the third argument.

## Usage

To run the program, use the following command:

```sh
go run . "<string>" "<old_char>" "<new_char>"
```

### Parameters

1. `<string>`: The input string where the replacement will occur.
2. `<old_char>`: The character to be replaced.
3. `<new_char>`: The character to replace the old character with.

### Example

```sh
$ go run . "hello world" "o" "a"
hella warld
```

In this example, all occurrences of 'o' in the string "hello world" are replaced with 'a', resulting in "hella warld".

## Edge Cases

- If the number of arguments is not exactly 3, the program does nothing and exits.
- If the second argument (old_char) is not found in the first argument (string), the original string is returned unchanged.

### More Examples

```sh
$ go run . "hella there" "a" "o"
hello there

$ go run . "hallo thara" "a" "e"
hello there

$ go run . "abcd" "z" "l"
abcd

$ go run . "something" "a" "o" "b" "c"
# No output because the number of arguments is not exactly 3
```

## Implementation Details

The program is implemented in Go and uses the `z01` package for printing characters. The main logic is contained in the `replace` function which iterates over each character in the input string and performs the replacement.

### Main Function

The main function:
1. Retrieves the input arguments.
2. Converts the second and third arguments to runes (characters).
3. Calls the `replace` function to perform the replacement.
4. Prints the resulting string.

### Replace Function

The `replace` function:
1. Iterates over each character in the input string.
2. Replaces occurrences of the old character with the new character.
3. Constructs and returns the modified string.

## Dependencies

This project depends on the `z01` package for printing characters. Make sure to install the package before running the program.

## License

This project is licensed under the MIT License.
```

This `README.md` provides an overview of the program, usage instructions, examples, and implementation details. It should help users understand how to use and extend the program.