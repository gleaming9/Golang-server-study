//gcc -shared -o libfactorial.dll -fPIC factorial.c
//공유 라이브러리로 컴파일
int factorial(int x){
    if(x == 1) return x;
    return factorial(x - 1) * x;
}