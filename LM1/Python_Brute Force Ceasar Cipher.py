alphabet = 'abcdefghijklmnopqrstuvwxyz'
ALPHABET = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

def decrypt(n, ciphertext):
    """Decrypt the ciphertext and return the plaintext"""
    result = ''
    for l in ciphertext:
        try:
            if l.isupper():
                index = ALPHABET.index(l)
                i = (index - n) % 26
                result += ALPHABET[i]
            else:
                index = alphabet.index(l)
                i = (index - n) % 26
                result += alphabet[i]
        except ValueError:
            result += l
    return result


def brute_force(ciphertext):
    """Perform a brute force attack to find the key and plaintext message"""
    for key in range(1, 26):
        plaintext = decrypt(key, ciphertext)
        print(f"Key: {key}\tPlaintext: {plaintext}")


ciphertext = 'Ugew gnwj zwjw Oslkgf'
brute_force(ciphertext)
