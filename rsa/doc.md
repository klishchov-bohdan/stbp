# Лабораторна робота №2
## Симетричне шифрування. Алгоритм AES
##### Роботу виконав
- Кліщов Богдан
- КН-922б
##### Мета: Дослідити принципи роботи симетричного шифрування на прикладі алгоритму AES
##### Завдання:
- Реалізувати алгоритм симетричного шифрування AES (будь-якої версії - 128 або 256).
- Довести коректність роботи реалізованого алгоритму шляхом порівняння результатів з існуючими реалізаціями (напр. сайтом-утилітою https://cryptii.com).

##### Розроблені функції шифрування та розшифровки:

    func (a *AES) EncryptCTR(in []byte, iv []byte) []byte {
        ivTmp := make([]byte, len(iv))
        copy(ivTmp, iv)
        plainTmp := make([]byte, len(in))
        copy(plainTmp, in)
        ivNumber := big.NewInt(0).SetBytes(iv)
        one := big.NewInt(1)
    
        i := 0
        for ; i < len(plainTmp)-a.len; i += a.len {
            a.encryptBlock(ivTmp, a.roundKeys)
            Xor(plainTmp[i:i+a.len], ivTmp)
            ivNumber.Add(ivNumber, one).FillBytes(ivTmp)
        }
        a.encryptBlock(ivTmp, a.roundKeys)
        Xor(plainTmp[i:], ivTmp)

        return plainTmp
    }
    
    func (a *AES) DecryptCTR(in []byte, iv []byte) []byte {
        ivTmp := make([]byte, len(iv))
        copy(ivTmp, iv)
        cipherTmp := make([]byte, len(in))
        copy(cipherTmp, in)
        ivNumber := big.NewInt(0).SetBytes(iv)
        one := big.NewInt(1)
    
        i := 0
        for ; i < len(cipherTmp)-a.len; i += a.len {
            a.encryptBlock(ivTmp, a.roundKeys)
            Xor(cipherTmp[i:i+a.len], ivTmp)
            ivNumber.Add(ivNumber, one).FillBytes(ivTmp)
        }
        a.encryptBlock(ivTmp, a.roundKeys)
        Xor(cipherTmp[i:], ivTmp)

        return cipherTmp
    }


##### Результати роботи програми
![results](./img/result1.png "Title")
![results](./img/result2.png "Title")

##### Висновки
Досліджено принципи роботи симетричного шифрування на прикладі алгоритму AES. Створено програму, яка реалізує алгоритм AES.
