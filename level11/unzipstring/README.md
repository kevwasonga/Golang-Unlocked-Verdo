Unzipstring

Instructions


Write a function called Unzipstring that takes a string that will be a kind of code, and your function will have to decrypt it and return a new string with the output. It works as follows: The string will be formed by a number followed by a letter, and the purpose is to print this letter the number of times that is requested by the number that precedes it.

"3w" ==> www

"2m3e" ==> mmee


    The letter after each number must be between a and z or A and Z.
    The number before the letter must be between 0 and 9.
    You cannot have two numbers or two letters in a row.
    If the input string does not respect the format, return Invalid Input.

Expected Function


func Unzipstring(s string) string {
}
