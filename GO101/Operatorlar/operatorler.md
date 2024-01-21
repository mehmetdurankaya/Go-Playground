Aritmetik Operatörler:
+: Toplama
-: Çıkarma
*: Çarpma
/: Bölme
%: Mod (kalan)

a := 10
b := 5

sum := a + b    // 15
difference := a - b  // 5
product := a * b // 50
quotient := a / b // 2
remainder := a % b // 0

Karşılaştırma Operatörleri:
==: Eşitse
!=: Eşit değilse
<: Küçükse
>: Büyükse
<=: Küçük veya eşitse
>=: Büyük veya eşitse

x := 10
y := 5

isEqual := x == y  // false
isNotEqual := x != y  // true
isLess := x < y  // false
isGreater := x > y  // true
isLessOrEqual := x <= y  // false
isGreaterOrEqual := x >= y  // true

Mantıksal Operatörler:
&&: Mantıksal "ve" (and)
||: Mantıksal "veya" (or)
!: Mantıksal "değil" (not)

isTrue := true
isFalse := false

logicalAnd := isTrue && isFalse  // false
logicalOr := isTrue || isFalse   // true
logicalNot := !isTrue   // false


Golang (Go) programlama dilinde birçok operatör bulunmaktadır. İşte temel operatör kategorileri ve örnekler:

Aritmetik Operatörler:
+: Toplama
-: Çıkarma
*: Çarpma
/: Bölme
%: Mod (kalan)
go
Copy code
a := 10
b := 5

sum := a + b    // 15
difference := a - b  // 5
product := a * b // 50
quotient := a / b // 2
remainder := a % b // 0
Karşılaştırma Operatörleri:
==: Eşitse
!=: Eşit değilse
<: Küçükse
>: Büyükse
<=: Küçük veya eşitse
>=: Büyük veya eşitse
go
Copy code
x := 10
y := 5

isEqual := x == y  // false
isNotEqual := x != y  // true
isLess := x < y  // false
isGreater := x > y  // true
isLessOrEqual := x <= y  // false
isGreaterOrEqual := x >= y  // true
Mantıksal Operatörler:
&&: Mantıksal "ve" (and)
||: Mantıksal "veya" (or)
!: Mantıksal "değil" (not)
go
Copy code
isTrue := true
isFalse := false

logicalAnd := isTrue && isFalse  // false
logicalOr := isTrue || isFalse   // true
logicalNot := !isTrue   // false
Bit Manipülasyon Operatörleri:
&: Bitwise "ve"
|: Bitwise "veya"
^: Bitwise "XOR" (müstakil veya)
<<: Sol shift
>>: Sağ shift
&^: Bit clear (ve değil)

a := 5 // 101
b := 3 // 011

bitwiseAnd := a & b   // 001 (1)
bitwiseOr := a | b    // 111 (7)
bitwiseXOR := a ^ b   // 110 (6)
leftShift := a << 1   // 1010 (10)
rightShift := a >> 1  // 10 (2)
bitClear := a &^ b    // 100 (4)

