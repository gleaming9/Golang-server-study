//gcc -shared -o libfibonacci.dll -fPIC fibonacci.c
//공유 라이브러리로 컴파일
int fibonacci(int n){
    if(n <= 0) return 0;
    if(n == 1) return 1;
    return fibonacci(n - 1) + fibonacci(n - 2);
}