# Лабораторна робота №11
## Оцінка якості коду
##### Роботу виконав
- Кліщов Богдан
- КН-922б
##### Мета: Дослідити алгоритми визначення якості коду
##### Завдання:
Використовуючи код будь-якого додатку (неважливо, чи він створений вами, або взятий з github/gitlab):
-	рівень якості програмування
-	складність розуміння програми
-	трудомісткість кодування програми
-	цикломатичне число Мак-Кейба
-	метрика Чепіна
Метрики потрібно виконувати для:
-	source code
-	декомпільованого коду

##### Хід роботи:
Цикломатичне число Мак-Кейба розраховується по формулі:
- V(G)=e — n + 2p, 
* де e – кількість дуг, 
* n – кількість вершин, 
* p – число компонент зв'язності
Метрика Чепіна розраховується так:
- Q = a1 * P + a2 * M + a3 * C + a4 * T, 
* де a1, a2, a3, a4 - вагові коефіцієнти

Для тестування якості коду, складності, трудомісткості в мові програмування golang будемо використовувати пакет gocritic (https://github.com/go-critic/go-critic). Даний додаток аналізує код, знаходить синтаксичні помилки та строчки, які можна спростити.
Для отримання цикломатичного числа Мак-Кейба використовуємо додаток gocyclo (https://github.com/fzipp/gocyclo). 
Почнемо зі створення Make файлу, який буде запускати наші застосунки. Перевіряти будемо на створеному в попередній лабораторній роботі алгоритмі rsa
##### Код Make файлу:
    run:
      gocyclo ./
      gocritic check ./...

##### Результати аналізу та оцінки коду
Тепер запустимо програму, ввівши в консолі команду make. Результати gocyclo будуть такими (спочатку йде оцінка, потім назва пакету, назва функції, відносний шлях та строчка в коді):
    gocyclo ./
    12 rsa decrypt rsa/rsa_ext.go:234:1
    9 rsa pubKeyDecrypt rsa/rsa_ext.go:172:1
    9 rsa priKeyIO rsa/rsa_ext.go:136:1
    9 rsa pubKeyIO rsa/rsa_ext.go:100:1
    6 main main main.go:57:1
    5 rsa nonZeroRandomBytes rsa/rsa_ext.go:307:1
    5 rsa priKeyByte rsa/rsa_ext.go:79:1
    5 rsa pubKeyByte rsa/rsa_ext.go:58:1
    4 rsa priKeyEncrypt rsa/rsa_ext.go:202:1
    4 rsa getPriKey rsa/rsa_ext.go:42:1
    4 rsa getPubKey rsa/rsa_ext.go:25:1
    4 main applyPriEPubD main.go:97:1
    4 main applyPubEPriD main.go:81:1
    3 rsa modInverse rsa/rsa_ext.go:334:1
    3 rsa (*RSASecurity).PriKeyDECRYPT rsa/rsa.go:81:1
    3 rsa (*RSASecurity).PriKeyENCTYPT rsa/rsa.go:69:1
    3 rsa (*RSASecurity).PubKeyDECRYPT rsa/rsa.go:57:1
    3 rsa (*RSASecurity).PubKeyENCTYPT rsa/rsa.go:45:1
    3 rsa PriKeyDecrypt rsa/gorsa.go:49:1
    3 rsa PublicDecrypt rsa/gorsa.go:33:1
    2 rsa leftPad rsa/rsa_ext.go:324:1
    2 rsa copyWithLeftPad rsa/rsa_ext.go:299:1
    2 rsa (*RSASecurity).VerifySignSha256WithRsa rsa/rsa.go:150:1
    2 rsa (*RSASecurity).VerifySignSha1WithRsa rsa/rsa.go:140:1
    2 rsa (*RSASecurity).VerifySignMd5WithRsa rsa/rsa.go:127:1
    2 rsa PriKeyEncrypt rsa/gorsa.go:20:1
    2 rsa PublicEncrypt rsa/gorsa.go:7:1
    1 rsa encrypt rsa/rsa_ext.go:228:1
    1 rsa (*RSASecurity).SignSha256WithRsa rsa/rsa.go:116:1
    1 rsa (*RSASecurity).SignSha1WithRsa rsa/rsa.go:105:1
    1 rsa (*RSASecurity).SignMd5WithRsa rsa/rsa.go:94:1
    1 rsa (*RSASecurity).GetPublickey rsa/rsa.go:41:1
    1 rsa (*RSASecurity).GetPrivatekey rsa/rsa.go:37:1
    1 rsa (*RSASecurity).SetPrivateKey rsa/rsa.go:31:1
    1 rsa (*RSASecurity).SetPublicKey rsa/rsa.go:25:1
Даний застосунок знайшов всі функції та провів їх оцінку. Чим менше оцінка, тим краще розроблена функція. 

Результат роботи додатку gocritic показано нижче (спочатку йде назва та роздашування файлу, потім пояснення, що можна спростити):
    gocritic check ./...
    ./rsa/rsa_ext.go:61:3: assignOp: replace `k = k - 11` with `k -= 11`
    ./rsa/rsa_ext.go:82:3: assignOp: replace `k = k - 11` with `k -= 11`
    ./rsa/rsa_ext.go:103:3: assignOp: replace `k = k - 11` with `k -= 11`
    ./rsa/rsa_ext.go:139:3: assignOp: replace `k = k - 11` with `k -= 11`

##### Висновки
Досліджено алгоритми визначення якості коду
