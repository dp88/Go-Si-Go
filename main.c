#include <stdio.h>
#include "hello.h"

int main (int argc, char *argv[]) {
    printf("Hello from C!\n");
    HelloGo();
}

void Finale() {
    printf("Hello from C from Go!\n");
}
