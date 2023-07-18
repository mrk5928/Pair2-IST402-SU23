codebook = [[0b00, 0b01], [0b01, 0b10], [0b10, 0b11], [0b11, 0b00]]
messages = [0b01, 0b00, 0b10, 0b00]
cipherCBC = [0] * 4
cipherCFB = [0] * 4
iv = 0b10

def codesbookLookup(xor):
    for i in range(len(codebook)):
        if codebook[i][0] == xor:
            return codebook[i][1]
    return -1

def encryptBlockCBC(plaintext, iv):
    xor = plaintext ^ iv
    return codesbookLookup(xor)

def decryptBlockCBC(ciphertext, iv):
    xor = codesbookLookup(ciphertext) ^ iv
    return xor

def encryptBlockCFB(plaintext, iv):
    xor = plaintext ^ iv
    return xor

def decryptBlockCFB(ciphertext, iv):
    xor = ciphertext ^ iv
    return xor

# CBC Mode
xor = 0
lookupValue = codesbookLookup(iv)

# Display the original message
print("CBC Mode:")
for i in range(len(messages)):
    print("The plaintext value of a is", format(messages[i], "02b"))

# Encryption (CBC)
for i in range(len(messages)):
    xor = messages[i] ^ lookupValue
    lookupValue = codesbookLookup(xor)
    print("The ciphered value of a in CBC mode is", format(xor, "02b"))
    cipherCBC[i] = xor

# Decryption (CBC)
lookupValue = codesbookLookup(iv)
for i in range(len(cipherCBC)):
    xor = cipherCBC[i] ^ lookupValue
    lookupValue = codesbookLookup(cipherCBC[i])
    print("The plaintext value of a in CBC mode is", format(xor, "02b"))

print()

# CFB Mode
lookupValue = codesbookLookup(iv)

# Display the original message
print("CFB Mode:")
for i in range(len(messages)):
    print("The plaintext value of a is", format(messages[i], "02b"))

# Encryption (CFB)
for i in range(len(messages)):
    cipherCFB[i] = encryptBlockCFB(messages[i], lookupValue)
    lookupValue = cipherCFB[i]
    print("The ciphered value of a in CFB mode is", format(cipherCFB[i], "02b"))

# Decryption (CFB)
lookupValue = codesbookLookup(iv)
for i in range(len(cipherCFB)):
    plaintext = decryptBlockCFB(cipherCFB[i], lookupValue)
    lookupValue = cipherCFB[i]
    print("The plaintext value of a in CFB mode is", format(plaintext, "02b"))
