# Cipher-Decoder

## Stage 1/3: The protocol

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


## Stage 2/3: A secret that we share

### 1) Description
In this stage, we need to figure out what secret number Alice and Bob (that's you) are going to use to communicate. We can reveal it through a series of computations. First, we need to find
**B**. Recall from Stage 1 that this is produced by Bob picking a random natural number
**b**, where **1 < b < p**, and performing the following calculation with the publicly-known
**g** and **p**:

**B = g^b mod p**

For example, let's say **g = 3, b = 7, p = 23**.

This is a relatively simple calculation to perform:

**g^b mod p = 3^7 mod23 = 2187 mod 23 = 2.**

In theory, this algorithm may work - a one-step computation seems easy enough, right? But what if
**g** and **b** both are 100? Then, **g^b** is a number with more than two hundred digits! This won't be memory-efficient to compute and store (it might even cause an integer-overflow error, depending on the variable type you're using!). In real cryptography, the numbers are significantly more than 100, so we're going to need a more efficient algorithm.

Through the mathematics of "Modular Exponentiation" we can find an alternative algorithm which avoids dealing with gigantic numbers. To perform this procedure, start off with
**c = 1** (this is just a starter number for the calculation, which is always chosen as 1), then iteratively perform the following computation for successive values of c ("c-primes" - here, these are just used to store the result of the calculation so we can carry it forward into the next calculation):

1) c′ = (c⋅g)modp
2) c′′ = (c′⋅g)modp
3) c′′′ = (c′′⋅g)modp
4) ...

If you perform this calculation **b**-times, the c-prime you end up with is guaranteed to equal the result for
**B** you would've got from the previous memory-intensive calculation! Now you don't have to worry about those huge numbers! (You can reference the above linked Wikipedia article for a more detailed explanation if you want)

If **g = 3, b = 7, p = 23**, then again starting off with
**c = 1** we can work out:

1) (1⋅3)mod23 = 3
2) (3⋅3)mod23 = 9
3) (9⋅3)mod23 = 4
4) (4⋅3)mod23 = 12
5) (12⋅3)mod23 = 13
6) (13⋅3)mod23 = 16
7) (16⋅3)mod23 = 2

As you can see, the **b**-th (7th in this case) calculation reveals that
**B=2**, which as promised is the same as the earlier, direct method of calculating it. Notice that while this involved more steps, the numbers stayed small. This will improve our program's memory efficiency because it won't have to deal with computing and storing giant numbers.

Now that Bob (you) has their **B**, you're just lacking one number (Alice's
**A**) to compute the shared secret. Alice will send it to you.

Compute the shared secret **S = A^b modp**, then send
**B** to Alice, so she can compute the shared secret too.

### 2) Objectives

At this stage, your program should:

- Take **p** and **g** from the input (same as in Stage 1);
- Print OK;
- Compute **B**;
- Take **A** from the input;
- Compute the shared secret (you're going to use this in Stage 3 to encrypt messages);
- Print ```B is [B]```.
- Print ```A is [A]```. (this won't be kept for the next stage; it's just for validating this stage)
- Print ```S is [S]```. (this won't be kept for the next stage; it's just for validating this stage)

### 3) Examples

For the purposes of this stage, the tests are acting as "Alice" and your program responds as "Bob". The example below shows how your program should work.
The greater-than symbol followed by a space (```> ```) represents the tests' input. Note that it's not part of the input.

#### Example 1: a = 21, b = 15

    > g is 245 and p is 999
    OK
    > A is 413
    B is 179
    A is 413
    S is 512

#### Example 2: a = 21, b = 15; different g and p

    > g is 28 and p is 644
    OK
    > A is 336
    B is 364
    A is 336
    S is 224

## Stage 3/3: Encrypt and decrypt

### 1) Description

Let's head onto encryption. In this stage, messages will be encrypted using the **Caesar cipher**. The Caesar cipher is a very simple algorithm, not well suited to any kind of real-world application outside of teaching purposes. Each letter of the alphabet is cyclically shifted down the alphabet by a certain number of positions. We can encrypt only letters; other characters will remain as they are. For example:

If the shared secret number is **4**, then we shift each letter down by four positions:

Plain	a	b	c	d	e	f	g	h	i	j	k	l	m	n	o	p	q	r	s	t	u	v	w	x	y	z

Cipher	e	f	g	h	i	j	k	l	m	n	o	p	q	r	s	t	u	v	w	x	y	z	a	b	c	d

Plaintext: ```This is a message.```

Ciphertext: ```Xlmw mw e qiwweki.```

In this stage, you need to encrypt the message ```Will you marry me?``` and send it to Alice. If her answer is ```Yeah, okay!```, encrypt ```Great!``` and send it back. If Alice's answer is ```Let's be friends.```, encrypt and output ```What a pity!``` Otherwise, don't print anything.

### 2) Objectives

At this stage, your program should:

- Take **p** and **g** from the input;
- Print ```OK```;
- Compute **B**;
- Take **A** from the input;
- Compute the shared secret;
- Print ```B is [B]```;
- Encrypt the message ```Will you marry me?``` with the Caesar cipher and print the result;
- Receive Alice's encrypted answer and decrypt it. If Alice's answer is ```Yeah, okay!```, encrypt ```Great!``` and print it. If Alice's answer is ```Let's be friends.```, encrypt ```What a pity!``` and output it. If the answer is anything else, don't print anything.

### 3) Examples

For the purposes of this stage, the tests are acting as "Alice" and your program responds as "Bob". The example below shows how your program should work.
The greater-than symbol followed by a space (```> ```) represents the tests' input. Note that it's not part of the input.

#### Example 1: success (Bob has chosen **b = 7** here, resulting in **s = 565**)

    > g is 245 and p is 999
    OK
    > A is 232
    B is 236
    Pbee rhn ftkkr fx?  // Will you marry me?
    > Rxta, hdtr!  // Yeah, okay!
    Zkxtm!  // Great!

#### Example 2: better luck next time (Bob has chosen **b = 7** here, resulting in **s = 565**)

    > g is 245 and p is 999
    OK
    > A is 232
    B is 236
    Pbee rhn ftkkr fx?  // Will you marry me?
    > Exm'l ux ykbxgwl.  // Let's be friends.
    Patm t ibmr!  // What a pity!
