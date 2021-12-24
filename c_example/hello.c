#include <stdio.h>
#include <stdlib.h>

int main() {

    printf("hello world\n");

    // printf(b ^= a << s | a >> w-s);

    unsigned b = 10;
    b ^= 10 << 4 | 10 >> 8-4;
    printf("%d\n" , b);

    printf("%d\n" , 10 << 4);

    return 0;
}
