#include <stdlib.h>
#include<unistd.h>

const int KB = 1000;
const int MB = KB * 1000;
const int GB = MB * 1000;

int main(void)
{
   // allocate 100KB
  int n = 100;
  int* x[n];
  for (int i = 0; i < n; i++) {
    x[i] = malloc(KB);
  }

  // allocate 1MB
  int* mb = malloc(MB);

  // allocate 100KB
  n = 100;
  int* y[n];
  for (int i = 0; i < n; i++) {
     y[i] = malloc(KB);
  }

  sleep(7200);

 return 0;
}
