#include <stdio.h>

int main(int argc, char *argv[])
{
    int x;

    printf("Hello %s, write a number!\n", argv[1]);
    printf("x = ");
    scanf("%d", &x);
    printf("x * x = %d\n", x * x);
    return 0;
}
