codebook = [
    [0b00, 0b01],  # Codebook entry: 00 -> 01
    [0b01, 0b10],  # Codebook entry: 01 -> 10
    [0b10, 0b11],  # Codebook entry: 10 -> 11
    [0b11, 0b00]   # Codebook entry: 11 -> 00
]

plaintext = input("Enter the plaintext (binary): ")
message = [int(bit) for bit in plaintext]

message1 = []

iv = 0b10  # Initialization Vector

def codebookLookup(xor):
    # Look up the codebook for the given XOR value
    for i in range(4):
        if codebook[i][0] == xor:
            lookupValue = codebook[i][1]
            return lookupValue

# Convert plaintext to a list of binary strings
for i in range(len(message)):
    a = f"{message[i]:b}"
    message1.append(a)

print("\nCBC encryption details:")
print(f"Plaintext: {message1}" )

stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream  # XOR the message with the stream
    ciphertext.append(codebookLookup(xor))  # Encrypt the XOR result using the codebook
    stream = ciphertext[i]  # Update the stream with the current ciphertext
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

ciphertext.reverse()
message.reverse()

print("\nCBC decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    xor = codebookLookup(ciphertext[i])  # Decrypt the ciphertext using the codebook
    plaintext.append(xor ^ stream)  # XOR the decrypted value with the stream
    stream = ciphertext[i]  # Update the stream with the current ciphertext
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

plaintext.reverse()
original_plaintext = "".join([str(bit) for bit in plaintext])
print(f"\nOriginal message: {original_plaintext}")

print("\nCFB encryption details:")
print(f"Plaintext: {message1}")

stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream  # XOR the message with the stream
    ciphertext.append(codebookLookup(xor))  # Encrypt the XOR result using the codebook
    stream = ciphertext[i] ^ iv  # Update the stream with the XOR result of ciphertext and IV
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

ciphertext.reverse()
message.reverse()

print("\nCFB decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    xor = codebookLookup(ciphertext[i])  # Decrypt the ciphertext using the codebook
    plaintext.append(xor ^ stream)  # XOR the decrypted value with the stream
    stream = ciphertext[i] ^ plaintext[i]  # Update the stream with the XOR result of ciphertext and plaintext
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

plaintext.reverse()
original_plaintext = "".join([str(bit) for bit in plaintext])
print(f"\nOriginal message: {original_plaintext}")
