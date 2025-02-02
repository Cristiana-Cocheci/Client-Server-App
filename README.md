# Client-Server-App
Distributed Systems project

Server is opened on localhost:8080, using tcp through the net package in Go. It can handle a maximum of n concurrent clients at once, which is declared in the env.txt file as ClientsNumber.

Clients can connect using telnet. They need to specify the type of command, the number of arrays sent as input and then the input. Here is an example:

```
> telnet localhost 8080
> ex6 1
> Hello WOrld
> ex1 1
> abc def ghi
Message from Server: adg, beh, cfi

```
For showcasing parallelism, I have written a predefined block of commands on 5 clients (in the example_requests folder). To run it the environment variable ReadFromFile needs to be set to true when running the application. 

Commands accepted by Server:

* ex1 : 
  * Input: an array of n strings of same length. 
  * Output: an array of n strings, where word at position i is made of the letters on the ith position in each input word.
  * Example: abc, def, ghi -> adg, beh, cfi
* ex3 : 
  * Input: an array of integers
  * Output: sum of the mirrored input integer
  * Example: 12, 32 -> 44
* ex5 : 
  * Input: an array of strings.
  * Output: returns transformed viable strings of binary numbers into decimal numbers. 
  * Example: aa 12 101 11 -> 5, 3
* ex6 : 
  * Input: a text
  * Output: the text encoded with Caesar's Cipher. The direction (left/right) and factor (k) of the encoding are chosen randomly and specified in the output.
  * Example: Hello WOrld -> Key: 21, Direction: LEFT => [Mjqqt BTwqi]
* ex7 :
  * Input: a string, where, before each letter, there is a number between 1 and 20, that specifies the number of times that letter appears in the actual string.
  * Output: decoded message
  * Example: 1G3o1L -> GoooL
* ex8 :
  * Input: an array of numbers
  * Output: from the numbers are selected the prime ones; it is returend the sum of the count of their digits
  * Example: 13, 7, 12 -> 3 (13 -> 2, 7 -> 1)
* ex12 :
  * Input: an array of integers
  * Output: each number's first digit is double, it is returned the sum of numbers
  * Example: 123, 12 -> 1235 (1123 + 12)
* exit :
  * server closes conection with client