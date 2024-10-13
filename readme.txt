# Покращений MVP для зашифрованих комунікацій

Цей проект реалізує покращений MVP для зашифрованих комунікацій між броньованими підрозділами та командними центрами з додатковими заходами безпеки.

## Впроваджені заходи безпеки

1. **TLS з'єднання**: 
   - Використовується HTTPS (WSS для WebSocket) для шифрування всього трафіку між клієнтом і сервером.
   - Генерується самопідписаний сертифікат для TLS (у реальному сценарії слід використовувати сертифікати, підписані довіреним центром сертифікації).

2. **Безпечний обмін ключами**:
   - Реалізовано протокол обміну ключами Diffie-Hellman на еліптичних кривих (ECDH).
   - Кожна сесія використовує унікальний ключ, отриманий з спільного секрету ECDH.

3. **Шифрування повідомлень**:
   - Використовується AES-256 у режимі CFB для шифрування вмісту повідомлень.
   - Кожне повідомлення шифрується унікальним сесійним ключем.

4. **Цифрові підписи**:
   - Додано заглушки для підписання та перевірки підписів повідомлень (потрібна повна реалізація).
   - Це забезпечить цілісність повідомлень та автентифікацію відправника.

5. **Захист від атак повторного відтворення**:
   - Кожне повідомлення містить часову мітку, яку можна використовувати для відхилення старих повідомлень.

## Наступні кроки для покращення безпеки

1. **Повна реалізація цифрових підписів**:
   - Використовуйте алгоритм ECDSA для створення та перевірки підписів.

2. **Управління ключами**:
   - Реалізуйте безпечне зберігання та ротацію ключів.

3. **Автентифікація клієнтів**:
   - Впровадьте систему автентифікації для перевірки особистості кожного клієнта.

4. **Захист від DoS-атак**:
   - Додайте обмеження на кількість з'єднань та повідомлень від одного клієнта.

5. **Аудит та логування**:
   - Впровадьте систему аудиту для відстеження важливих подій безпеки.

6. **Оновлення TLS**:
   - Використовуйте сертифікати, підписані довіреним центром сертифікації.
   - Налаштуйте правильні параметри TLS для максимальної безпеки.

7. **Тестування на проникнення**:
   - Проведіть ретельне тестування на проникнення для виявлення можливих вразливостей.

## Запуск проекту

1. Переконайтеся, що у вас встановлено Go версії 1.21 або новіше.
2. Клонуйте репозиторій та перейдіть до директорії проекту.
3. Виконайте `go mod tidy` для завантаження залежностей.
4. Запустіть сервер: `go run main.go --mode server`
5. В окремому терміналі запустіть клієнт: `go rum main.go --mode client`


Server-Client Architecture:
The program is designed as a client-server application using WebSockets for real-time, bidirectional communication.
Encryption:
It uses AES encryption in GCM mode to secure messages between clients and the server.
Startup:

When you run the program, main.go checks the command-line flag to determine whether to start in server or client mode.


Server Mode:

The server starts and listens on port 8080 for WebSocket connections.
It uses TLS for secure connections.
When a client connects, the server performs a key exchange (currently a placeholder function).
The server maintains a list of connected clients.
It receives messages from clients, timestamps them, and broadcasts to all connected clients.


Client Mode:

The client connects to the server using a WebSocket over TLS.
It performs a key exchange with the server.
The client runs two goroutines:
a. One for reading user input and sending messages.
b. Another for receiving and displaying messages from the server.


Message Flow:

User types a message in the client.
The client encrypts the message using the session key.
The encrypted message is sent to the server.
The server receives the message and broadcasts it to all clients.
Each client receives the encrypted message and decrypts it using their session key.
The decrypted message is displayed to the user.


Security Features:

TLS is used for the WebSocket connection.
Messages are encrypted end-to-end using AES-GCM.
Each message includes a timestamp to prevent replay attacks.


Key Exchange:

Currently, this is a placeholder function returning a fixed key.
In a real implementation, this would involve a secure key exchange protocol like Diffie-Hellman.


Error Handling:

The program includes basic error handling for network issues, encryption/decryption problems, etc.


Concurrency:

The server uses goroutines to handle multiple clients concurrently.
Mutex is used to safely manage the shared clients map.



To use the program:

Start the server in one terminal.
Start one or more clients in separate terminals.
Type messages in any client terminal.
See the encrypted messages broadcasted to all clients and decrypted for display.

This design allows for secure, real-time communication between multiple clients through a central server. 
The use of WebSockets enables instant message delivery, while the encryption ensures the privacy of the communications.

