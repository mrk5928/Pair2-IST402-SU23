codebook = [
    [0b00, 0b01],
    [0b01, 0b10],
    [0b10, 0b11],
    [0b11, 0b00]
]

plaintext = input("Enter the plaintext (binary): ")
message = [int(bit) for bit in plaintext]

message1 = []

iv = 0b10

def codebookLookup(xor):
    for i in range(4):
        if codebook[i][0] == xor:
            lookupValue = codebook[i][1]
            return lookupValue

for i in range(len(message)):
    a = f"{message[i]:b}"
    message1.append(a)

print("\nCBC encryption details:")
print(f"Plaintext: {message1}" )

stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream
    ciphertext.append(codebookLookup(xor))
    stream = ciphertext[i]
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

ciphertext.reverse()
message.reverse()

print("\nCBC decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    xor = codebookLookup(ciphertext[i])
    plaintext.append(xor ^ stream)
    stream = ciphertext[i]
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

plaintext.reverse()
original_plaintext = "".join([str(bit) for bit in plaintext])
print(f"\nOriginal message: {original_plaintext}")

print("\nCFB encryption details:")
print(f"Plaintext: {message1}")

stream = iv
ciphertext = []
for i in range(len(message)):
    xor = message[i] ^ stream
    ciphertext.append(codebookLookup(xor))
    stream = ciphertext[i] ^ iv
    print(f"The ciphered value of {message[i]:b} is {ciphertext[i]:b}")

ciphertext.reverse()
message.reverse()

print("\nCFB decryption details:")
print(f"Ciphertext: {ciphertext}")

stream = iv
plaintext = []
for i in range(len(ciphertext)):
    xor = codebookLookup(ciphertext[i])
    plaintext.append(xor ^ stream)
    stream = ciphertext[i] ^ plaintext[i]
    print(f"The deciphered value of {ciphertext[i]:b} is {plaintext[i]:b}")

plaintext.reverse()
original_plaintext = "".join([str(bit) for bit in plaintext])
print(f"\nOriginal message: {original_plaintext}")
