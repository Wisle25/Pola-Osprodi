#include <stdio.h>

int main()
{
    int n;
    scanf("%d", &n);

    // Jika n adalah genap, ubah n ke bentuk ganjil di bawahnya
    if (n % 2 == 0)
        n -= 1;

    // Aku buatnya dengan konsep setengah...jadi setengah atas dan setengah bawah
    int half = n / 2;

    for (int i = 0; i < half; i++)
    {
        for (int j = 0; j < i; j++)
        {
            printf("  ");
        }

        for (int j = i; j < half; j++)
        {
            printf("* ");
        }

        for (int j = i; j <= half; j++)
        {
            printf("* ");
        }
        printf("\n");
    }

    for (int i = 1; i <= half + 1; i++)
    {
        for (int j = i; j <= half; j++)
            printf("  ");

        for (int j = 1; j < i; j++)
        {
            printf("* ");
        }

        for (int j = 1; j <= i; j++)
        {
            printf("* ");
        }
        printf("\n");
    }

    return 0;
}