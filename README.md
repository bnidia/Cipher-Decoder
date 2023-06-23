# Cipher-Decoder

## The protocol

### 1) Description

Cryptography is a good tool to keep your communication secret. One of the main issues of cryptography is how to send an encryption key to another person over an unsecured communication channel. One solution to this problem is what's called the **Diffie-Hellman key exchange protocol**.

In this project, we will tell you the story of Alice and Bob. They are fictional characters commonly used as placeholders in conversations on cryptographic systems and protocols. You will act as Bob, a guy who has an urge to communicate with Alice.

The simplest implementation of the protocol is the following:

1. Alice and Bob agree to use the modulus **p = 23** and base **g = 5**. These numbers are no secret and are picked at random;
2. Alice chooses the secret integer **a = 4**, computes **A**, and sends it to Bob:
A = g^a mod p = 5^4 mod23 = 4

In plain language, this says that **A** is computed by raising
**g** to the exponent **a** and taking the remainder from division by
**p** (i.e., the number "mod" **p**).
3. Similarly, Bob chooses the secret integer **b = 3**, computes
**B**, and sends it to Alice:
B = g^b mod p = 5^3 mod23 = 10
4. Alice computes the resultant shared secret **S**
on her end, using the **B** from Bob, her own secret **a**, and the shared
S = B^a mod p = 10^4 mod23 = 18
5. Bob computes the shared secret on his end, using Alice's
**A**, his secret **b**, and the shared **p**
(notice that it's the same result):
S = A^b mod p = 4^3 mod23 = 18
6. Alice and Bob now share a secret number (18). They can use the shared secret to send encrypted messages to each other.

Look at the math formula below. This is the main idea of the Diffie-Hellman protocol:

(x^a)^b = (x^b)^a

Let's now rewrite it with the variables we used in our example above:

(g^a)^b = (g^b)^a

Now, add the modulus:

(g^a mod p)^b mod p = (g^b mod p)^a mod p

If something's not clear in the Description, you can read this Wikipedia article that explains the protocol in more detail.

In this stage, Alice will tell you **g** and **p**. You are going to need them in later stages.

### 2) Objectives

In this stage, your program should:
- Read a string from the input;
- Get **p** and **g** from the string;
- Print **p** and **g**.

### 3) Example

For the purposes of this stage, the tests are acting as "Alice" and your program responds as "Bob". The example below shows how your program should work.
The greater-than symbol followed by a space (> ) represents the tests' input. Note that it's not part of the input.

#### Example 1:
    > g is 245 and p is 999
    g = 245 and p = 999
