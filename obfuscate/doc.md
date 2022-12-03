# Лабораторна робота №12
## Дослідження засобів обфускації програмного забезпечення
#### Роботу виконав
- Кліщов Богдан
- КН-922б
#### Мета: Дослідити існуючі утиліти Обфускації програмного забезпечення
#### Завдання:
Для обраної мови програмування слід обфускувати будь-який проект.
Слід використовувати декілька утиліт для обфускації. Мінімальна кількість обфускаторів - 4. У звіті навести порівняльну характеристику обраних обфускаторів, що складатиметься з:
- інформації щодо використовуваних методів обфускації
- плюсів та мінусів кожного з обфускаторів.
Навести опис використання кожного з обраних обфускаторів.

#### Хід роботи:
В даній лабораторній роботі було використано такі засоби обфускації коду:
- garble (https://github.com/burrowers/garble)
- obfuscator (https://github.com/stgleb/obfuscator)
- mumbojumbo (https://github.com/jeromer/mumbojumbo)
- poltergeist (https://github.com/Shell-Company/poltergeist)

Проект для обфускації був обраний з попередньої лабораторної роботи - sha256. 

#### Код Make файлу:
      run:
         garble -debugdir ./garble build obfuscate/sha256
         var=$(<./sha256/sha256.go) && mumbojumbo -s="$var" -p="sha256" | goimports > ./mumbojumbo/sha256.go
         go run main.go
         poltergeist -encode -file ./sha256/sha256.go > polter.go

#### garble
Особливість даного методу заключається в зміні назв функцій та змінних на незрозумілі символи
Частина обфускованого коду виглядає так:

      func Vw2WoqqX(kdNdo9nm []byte) [32]byte {
         kdNdo9nm =  /*line m0EOGLTU.go:1*/ ys7QvSCZ(kdNdo9nm)

         goAW401K, fCqjYtt6, vzMizMAy, itIP2ByV, ri8wxafH, tjZVztDa, x3btOA2l, oaOEcnee :=  /*line _AkaK2iS.go:1*/ uint32(0x6a09e667),  /*line GIMB_0TR.go:1*/ uint32(0xbb67ae85),  /*line QPCjLSkd.go:1*/ uint32(0x3c6ef372),  /*line XOjPr79k.go:1*/ uint32(0xa54ff53a),
             /*line KUZi3DCC.go:1*/ uint32(0x510e527f),  /*line wiUQxUVy.go:1*/ uint32(0x9b05688c),  /*line kDo1pdPS.go:1*/ uint32(0x1f83d9ab),  /*line uMEuq84M.go:1*/ uint32(0x5be0cd19)

         for iMByIZp0 := 0; iMByIZp0 <  /*line ZYPU5dBd.go:1*/ len(kdNdo9nm); iMByIZp0 += chunkSize {
            var iBmv8zrT [64]uint32
            for idzQKtZu := 0; idzQKtZu*4 < chunkSize; idzQKtZu++ {
               iBmv8zrT[idzQKtZu] =  /*line h8egaGOj.go:1*/ binary.BigEndian.Uint32(kdNdo9nm[iMByIZp0+idzQKtZu*4 : iMByIZp0+(idzQKtZu+1)*4])
            }

            for idzQKtZu := 16; idzQKtZu < 64; idzQKtZu++ {
               iAMkUhcH :=  /*line fhxicu9z.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-15], -7) ^  /*line Glen5Btc.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-15], -18) ^ (iBmv8zrT[idzQKtZu-15] >> 3)
               kw0mLN4V :=  /*line gnct4evx.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-2], -17) ^  /*line BFfzqrS9.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-2], -19) ^ (iBmv8zrT[idzQKtZu-2] >> 10)
               iBmv8zrT[idzQKtZu] = iBmv8zrT[idzQKtZu-16] + iAMkUhcH + iBmv8zrT[idzQKtZu-7] + kw0mLN4V
            }

            jkaWZWr3, ytpv7Fq6, cM51f7as, c3DFzJxK, iaqIONzz, quIJzx4t, yYB7n2Kl, vB7RJcOO := goAW401K, fCqjYtt6, vzMizMAy, itIP2ByV, ri8wxafH, tjZVztDa, x3btOA2l, oaOEcnee
            for idzQKtZu := 0; idzQKtZu < 64; idzQKtZu++ {
               WmfGM9K_ :=  /*line XJFHiT98.go:1*/ bits.RotateLeft32(iaqIONzz, -6) ^  /*line DwIDJowJ.go:1*/ bits.RotateLeft32(iaqIONzz, -11) ^  /*line bNfPfBNz.go:1*/ bits.RotateLeft32(iaqIONzz, -25)
               cujd6iWL := (iaqIONzz & quIJzx4t) ^ ((^iaqIONzz) & yYB7n2Kl)
               khYsNgsw := vB7RJcOO + WmfGM9K_ + cujd6iWL + Le5rqThl[idzQKtZu] + iBmv8zrT[idzQKtZu]
               EUnrRCmx :=  /*line qCIdlOWW.go:1*/ bits.RotateLeft32(jkaWZWr3, -2) ^  /*line Lfc8aG7N.go:1*/ bits.RotateLeft32(jkaWZWr3, -13) ^  /*line j5lxZczU.go:1*/ bits.RotateLeft32(jkaWZWr3, -22)
               fAkQYphG := (jkaWZWr3 & ytpv7Fq6) ^ (jkaWZWr3 & cM51f7as) ^ (ytpv7Fq6 & cM51f7as)
               euX3kGs9 := EUnrRCmx + fAkQYphG
               vB7RJcOO = yYB7n2Kl
               yYB7n2Kl = quIJzx4t
               quIJzx4t = iaqIONzz
               iaqIONzz = c3DFzJxK + khYsNgsw
               c3DFzJxK = cM51f7as
               cM51f7as = ytpv7Fq6
               ytpv7Fq6 = jkaWZWr3
               jkaWZWr3 = khYsNgsw + euX3kGs9
            }
            goAW401K += jkaWZWr3
            fCqjYtt6 += ytpv7Fq6
            vzMizMAy += cM51f7as
            itIP2ByV += c3DFzJxK
            ri8wxafH += iaqIONzz
            tjZVztDa += quIJzx4t
            x3btOA2l += yYB7n2Kl
            oaOEcnee += vB7RJcOO
         }

         c_ehY6dh := [32]byte{}
          /*line dRdKybWc.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[:4], goAW401K)
          /*line w2QaDzJZ.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[4:8], fCqjYtt6)
          /*line mCkuiB7d.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[8:12], vzMizMAy)
          /*line B66KSLjd.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[12:16], itIP2ByV)
          /*line E_GhpZ3y.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[16:20], ri8wxafH)
          /*line La1RqG_M.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[20:24], tjZVztDa)
          /*line lTcP8Yzs.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[24:28], x3btOA2l)
          /*line UK_J7VtK.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[28:], oaOEcnee)

         return c_ehY6dh
      }

#### obfuscator
Обфускація за допомогої даного пакету виконується так:

      func main() {
         fileStr, _ := ioutil.ReadFile("./sha256/sha256.go")
         result, err := obfuscator.Obfuscate(fileStr)

         if err != nil {
            log.Fatal(err)
         }
         create, err := os.Create("./obfuscator/sha256.go")
         if err != nil {
            log.Fatal(err)
         }
         defer func(create *os.File) {
            err := create.Close()
            if err != nil {

            }
         }(create)
         _, err = create.Write([]byte(result))
         if err != nil {
            log.Fatal(err)
         }
      }
      
Обфускований код виглядає так:

      for(s='package sha256
      import("Tcodzg/bx"
      "math/y"bvar K=[]]uztL{0x428a2f9M713R49Kb5c0fbcf~e9b5dbaJ3956c25bC9f111fK923f82;~ab1c5edJd807aa9M12835@KD3185GC50c7dcH72G5dR?0deb1fe~9bdc<aBcQ=1R~e49b69cKefG4786~0fcQdP~D0ca1cc~2de92Pf~4aR84aaCc@a9dc~76f988da~983e5152~a831P6d~@0L7cM=597fcBPe00=Hd5a7914B<ca635K1429296B27b70a8JO1b213M4d2PdfcC3380d1H650a7354~766a0abb?1c2c9O~92722c8Ja2=e8aKa81a6]b~cDb8b70~c76c51aHdQO8Q~d699<D~f40e358J1<aa070~Q;c116~1e376c0M2R87Rc~34@bcbJ391c0cbH4ed8a;aCb9cc;f~68O69HR8f8Oe~78a5636f?4c87814?cc7020M90G9fa~;5<Sb~Gf9a3fBP7178f2,Iconst :e=]
      func pad^[S)[S{Lemake([S,8)mt](L,uzt](lT^)*8)b{=appTd^?0bW(lT^)+8)%]!=0{{=appTd^~00)I{=appTd^,L...breturn {Ifunc Hash^[S)[LS{{=pad^bh0o1o2o3o4o5o6o7eui6a09e667fbb67ae85f3Pef372fa54953af510e52799@5688cf1f83d9a=5G0cdQbW >e0;><lT^);>Z:e{var w[]]uztL
      W ie0E*4<:eE++{l]=p:}{[>+i*4:>+(i+1)*4])IW ie16E<]E++{s0NX15]_7FX15]_18AX15]>>3bs1NX2]_17FX2]_QAX2]>>10bl]=X16]+s0+X7]+s1Ia,b,c,d,e,f,goN0o1o2o3o4o5o6o7
      W ie0E<]E++{S1Ne_6Fe_11Fe_25bche(e&fA(^e)&gbV1N+S1+ch+K[i]+l]
      S0Na_2Fa_13Fa_22bmaje(a&bAa&cAb&cbVOS0+maj[=g
      g=f
      f=e
      e=d+V1
      d=c
      c=b
      b=a
      a=V1+V2}[0Za[1Zb[2Zc[3Zd[4Ze[5Zf[6Zg[7ZhIUue[LS{}ku[:4Y0g4:8Y1g8:12Y2g12:16Y3g16:20Y4g20:DY5gD:28Y6g28:Y7breturn Uu}~,0x}tL(|mess{|agezinybitsxzaryw.BigvUanugestty.RostaterLef}qEnv.pbxwqo,hnpPutm
      n:lw[ikm}Ujhunkiz}0xhtsrg)ku[f),uie:=dcjSc]bytb)
      a)^`tart_,-^({]64[
      hZ+=Y]oXl-WforVtmpUdiTenSceR74Q19Pc6O2eNehM8~L32K1~J5~I}
      H3~GbeFahE;iD24C~5B7~Aa(@b0?~8>d`=bf<06;a4:Uz9ff',i=0;j='9:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~'[i++];)with(s.split(j))s=join(pop());eval(s)
      
Його особливість заключається в зжиманні коду.

#### mumbojumbo
Обфускований код виглядає приблизно так:

      package sha256

      import (
         "unsafe"
      )

      func eax() uint8 {
         return uint8(unsafe.Sizeof(true))
      }

      func Get() string {
         return string(
            []byte{
               ((eax()<<eax()|eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
               (((eax()<<eax()|eax())<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
            },
         )
      }
При виклику методу Get() проходить деобфускація та повертається оригінальний код.

#### poltergeist
Обфускований код виглядає так:
![image](https://user-images.githubusercontent.com/88532388/205436691-27a4e5bf-ab02-481c-b031-1110e44bcf79.png)

#### Висновки
Досліджено існуючі утиліти обфускації програмного забезпечення. Для дослідження взято чотири пакети. Коден з них виконує свою роботу по різному, але в кожного методу є свої переваги. Наприклад, повна неможливість читання коду, можливість повернення начального коду та ін.
